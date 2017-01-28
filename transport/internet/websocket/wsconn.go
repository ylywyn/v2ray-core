package websocket

import (
	"bufio"
	"io"
	"net"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"v2ray.com/core/common/errors"
	"v2ray.com/core/common/log"
)

type wsconn struct {
	wsc         *websocket.Conn
	readBuffer  *bufio.Reader
	connClosing bool
	reusable    bool
	rlock       *sync.Mutex
	wlock       *sync.Mutex
	config      *Config
}

func (ws *wsconn) Read(b []byte) (n int, err error) {
	ws.rlock.Lock()
	n, err = ws.read(b)
	ws.rlock.Unlock()
	return n, err

}

func (ws *wsconn) read(b []byte) (n int, err error) {
	if ws.connClosing {
		return 0, io.EOF
	}

	n, err = ws.readNext(b)
	return n, err
}

func (ws *wsconn) getNewReadBuffer() error {
	_, r, err := ws.wsc.NextReader()
	if err != nil {
		log.Warning("WS transport: ws connection NewFrameReader return ", err)
		ws.connClosing = true
		ws.Close()
		return err
	}
	ws.readBuffer = bufio.NewReader(r)
	return nil
}

func (ws *wsconn) readNext(b []byte) (n int, err error) {
	if ws.readBuffer == nil {
		err = ws.getNewReadBuffer()
		if err != nil {
			return 0, err
		}
	}

	n, err = ws.readBuffer.Read(b)

	if err == nil {
		return n, err
	}

	if errors.Cause(err) == io.EOF {
		ws.readBuffer = nil
		if n == 0 {
			return ws.readNext(b)
		}
		return n, nil
	}
	return n, err

}

func (ws *wsconn) Write(b []byte) (n int, err error) {
	ws.wlock.Lock()
	if ws.connClosing {
		return 0, io.ErrClosedPipe
	}

	n, err = ws.write(b)
	ws.wlock.Unlock()
	return n, err
}

func (ws *wsconn) write(b []byte) (n int, err error) {
	wr, err := ws.wsc.NextWriter(websocket.BinaryMessage)
	if err != nil {
		log.Warning("WS transport: ws connection NewFrameReader return ", err)
		ws.connClosing = true
		ws.Close()
		return 0, err
	}
	n, err = wr.Write(b)
	if err != nil {
		return 0, err
	}
	err = wr.Close()
	if err != nil {
		return 0, err
	}
	return n, err
}

func (ws *wsconn) Close() error {
	ws.connClosing = true
	ws.wlock.Lock()
	ws.wsc.WriteControl(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""), time.Now().Add((time.Second * 5)))
	ws.wlock.Unlock()
	err := ws.wsc.Close()
	return err
}
func (ws *wsconn) LocalAddr() net.Addr {
	return ws.wsc.LocalAddr()
}
func (ws *wsconn) RemoteAddr() net.Addr {
	return ws.wsc.RemoteAddr()
}
func (ws *wsconn) SetDeadline(t time.Time) error {
	if err := ws.SetReadDeadline(t); err != nil {
		return err
	}
	return ws.SetWriteDeadline(t)
}
func (ws *wsconn) SetReadDeadline(t time.Time) error {
	return ws.wsc.SetReadDeadline(t)
}
func (ws *wsconn) SetWriteDeadline(t time.Time) error {
	return ws.wsc.SetWriteDeadline(t)
}

func (ws *wsconn) setup() {
	ws.connClosing = false

	/*
		https://godoc.org/github.com/gorilla/websocket#Conn.NextReader
		https://godoc.org/github.com/gorilla/websocket#Conn.NextWriter

		Both Read and write access are both exclusive.
		And in both case it will need a lock.

	*/
	ws.rlock = &sync.Mutex{}
	ws.wlock = &sync.Mutex{}

	ws.pingPong()
}

func (ws *wsconn) Reusable() bool {
	return ws.config.IsConnectionReuse() && ws.reusable && !ws.connClosing
}

func (ws *wsconn) SetReusable(reusable bool) {
	ws.reusable = reusable
}

func (ws *wsconn) pingPong() {
	pongRcv := make(chan int, 1)
	ws.wsc.SetPongHandler(func(data string) error {
		pongRcv <- 0
		return nil
	})

	go func() {
		for !ws.connClosing {
			ws.wlock.Lock()
			ws.wsc.WriteMessage(websocket.PingMessage, nil)
			ws.wlock.Unlock()
			tick := time.After(time.Second * 3)

			select {
			case <-pongRcv:
			case <-tick:
				if !ws.connClosing {
					log.Debug("WS:Closing as ping is not responded~" + ws.wsc.UnderlyingConn().LocalAddr().String() + "-" + ws.wsc.UnderlyingConn().RemoteAddr().String())
				}
				ws.Close()
			}
			<-time.After(time.Second * 27)
		}

		return
	}()

}
