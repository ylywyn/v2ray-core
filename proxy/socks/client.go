package socks

import (
	"context"

	"v2ray.com/core/common"
	"v2ray.com/core/common/buf"
	"v2ray.com/core/common/log"
	"v2ray.com/core/common/net"
	"v2ray.com/core/common/protocol"
	"v2ray.com/core/common/retry"
	"v2ray.com/core/common/signal"
	"v2ray.com/core/proxy"
	"v2ray.com/core/transport/internet"
	"v2ray.com/core/transport/ray"
)

type Client struct {
	serverPicker protocol.ServerPicker
}

func NewClient(ctx context.Context, config *ClientConfig) (*Client, error) {
	serverList := protocol.NewServerList()
	for _, rec := range config.Server {
		serverList.AddServer(protocol.NewServerSpecFromPB(*rec))
	}
	client := &Client{
		serverPicker: protocol.NewRoundRobinServerPicker(serverList),
	}

	return client, nil
}

func (c *Client) Process(ctx context.Context, ray ray.OutboundRay) error {
	destination := proxy.DestinationFromContext(ctx)

	var server *protocol.ServerSpec
	var conn internet.Connection

	dialer := proxy.DialerFromContext(ctx)
	err := retry.ExponentialBackoff(5, 100).On(func() error {
		server = c.serverPicker.PickServer()
		dest := server.Destination()
		rawConn, err := dialer.Dial(ctx, dest)
		if err != nil {
			return err
		}
		conn = rawConn

		return nil
	})

	if err != nil {
		log.Warning("Socks|Client: Failed to find an available destination.")
		return err
	}

	defer conn.Close()
	conn.SetReusable(false)

	request := &protocol.RequestHeader{
		Version: socks5Version,
		Command: protocol.RequestCommandTCP,
		Address: destination.Address,
		Port:    destination.Port,
	}
	if destination.Network == net.Network_UDP {
		request.Command = protocol.RequestCommandUDP
	}

	user := server.PickUser()
	if user != nil {
		request.User = user
	}

	udpRequest, err := ClientHandshake(request, conn, conn)
	if err != nil {
		log.Warning("Socks|Client: Failed to establish connection to server: ", err)
		return err
	}

	var requestFunc func() error
	var responseFunc func() error
	if request.Command == protocol.RequestCommandTCP {
		requestFunc = func() error {
			return buf.PipeUntilEOF(ray.OutboundInput(), buf.NewWriter(conn))
		}
		responseFunc = func() error {
			defer ray.OutboundOutput().Close()
			return buf.PipeUntilEOF(buf.NewReader(conn), ray.OutboundOutput())
		}
	} else if request.Command == protocol.RequestCommandUDP {
		udpConn, err := dialer.Dial(ctx, udpRequest.Destination())
		if err != nil {
			log.Info("Socks|Client: Failed to create UDP connection: ", err)
			return err
		}
		defer udpConn.Close()
		requestFunc = func() error {
			return buf.PipeUntilEOF(ray.OutboundInput(), &UDPWriter{request: request, writer: udpConn})
		}
		responseFunc = func() error {
			defer ray.OutboundOutput().Close()
			reader := &UDPReader{reader: net.NewTimeOutReader(16, udpConn)}
			return buf.PipeUntilEOF(reader, ray.OutboundOutput())
		}
	}

	requestDone := signal.ExecuteAsync(requestFunc)
	responseDone := signal.ExecuteAsync(responseFunc)
	if err := signal.ErrorOrFinish2(requestDone, responseDone); err != nil {
		log.Info("Socks|Client: Connection ends with ", err)
		ray.OutboundInput().CloseError()
		ray.OutboundOutput().CloseError()
		return err
	}

	return nil
}

func init() {
	common.Must(common.RegisterConfig((*ClientConfig)(nil), func(ctx context.Context, config interface{}) (interface{}, error) {
		return NewClient(ctx, config.(*ClientConfig))
	}))
}
