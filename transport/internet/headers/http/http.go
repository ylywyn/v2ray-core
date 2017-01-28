package http

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net"
	"net/http"
	"strings"
	"time"

	"v2ray.com/core/common"
	"v2ray.com/core/common/buf"
	"v2ray.com/core/common/serial"
)

const (
	CRLF   = "\r\n"
	ENDING = CRLF + CRLF

	// max length of HTTP header. Safety precaution for DDoS attack.
	maxHeaderLength = 8192
)

var (
	ErrHeaderToLong = errors.New("Header too long.")
	writeCRLF       = serial.WriteString(CRLF)
)

type Reader interface {
	Read(io.Reader) (*buf.Buffer, error)
}

type Writer interface {
	Write(io.Writer) error
}

type NoOpReader struct{}

func (v *NoOpReader) Read(io.Reader) (*buf.Buffer, error) {
	return nil, nil
}

type NoOpWriter struct{}

func (v *NoOpWriter) Write(io.Writer) error {
	return nil
}

type HeaderReader struct {
}

func (*HeaderReader) Read(reader io.Reader) (*buf.Buffer, error) {
	buffer := buf.NewSmall()
	totalBytes := 0
	endingDetected := false
	for totalBytes < maxHeaderLength {
		err := buffer.AppendSupplier(buf.ReadFrom(reader))
		if err != nil {
			return nil, err
		}
		if n := bytes.Index(buffer.Bytes(), []byte(ENDING)); n != -1 {
			buffer.SliceFrom(n + len(ENDING))
			endingDetected = true
			break
		}
		if buffer.Len() >= len(ENDING) {
			totalBytes += buffer.Len() - len(ENDING)
			leftover := buffer.BytesFrom(-len(ENDING))
			buffer.Reset(func(b []byte) (int, error) {
				return copy(b, leftover), nil
			})
		}
	}
	if buffer.IsEmpty() {
		buffer.Release()
		return nil, nil
	}
	if !endingDetected {
		buffer.Release()
		return nil, ErrHeaderToLong
	}
	return buffer, nil
}

type HeaderWriter struct {
	header *buf.Buffer
}

func NewHeaderWriter(header *buf.Buffer) *HeaderWriter {
	return &HeaderWriter{
		header: header,
	}
}

func (v *HeaderWriter) Write(writer io.Writer) error {
	if v.header == nil {
		return nil
	}
	_, err := writer.Write(v.header.Bytes())
	v.header.Release()
	v.header = nil
	return err
}

type HttpConn struct {
	net.Conn

	readBuffer    *buf.Buffer
	oneTimeReader Reader
	oneTimeWriter Writer
	errorWriter   Writer
}

func NewHttpConn(conn net.Conn, reader Reader, writer Writer, errorWriter Writer) *HttpConn {
	return &HttpConn{
		Conn:          conn,
		oneTimeReader: reader,
		oneTimeWriter: writer,
		errorWriter:   errorWriter,
	}
}

func (v *HttpConn) Read(b []byte) (int, error) {
	if v.oneTimeReader != nil {
		buffer, err := v.oneTimeReader.Read(v.Conn)
		if err != nil {
			return 0, err
		}
		v.readBuffer = buffer
		v.oneTimeReader = nil
	}

	if v.readBuffer.Len() > 0 {
		nBytes, err := v.readBuffer.Read(b)
		if nBytes == v.readBuffer.Len() {
			v.readBuffer.Release()
			v.readBuffer = nil
		}
		return nBytes, err
	}

	return v.Conn.Read(b)
}

func (v *HttpConn) Write(b []byte) (int, error) {
	if v.oneTimeWriter != nil {
		err := v.oneTimeWriter.Write(v.Conn)
		v.oneTimeWriter = nil
		if err != nil {
			return 0, err
		}
	}

	return v.Conn.Write(b)
}

// Close implements net.Conn.Close().
func (v *HttpConn) Close() error {
	if v.oneTimeWriter != nil && v.errorWriter != nil {
		// Connection is being closed but header wasn't sent. This means the client request
		// is probably not valid. Sending back a server error header in this case.
		v.errorWriter.Write(v.Conn)
	}

	return v.Conn.Close()
}

func formResponseHeader(config *ResponseConfig) *HeaderWriter {
	header := buf.NewSmall()
	header.AppendSupplier(serial.WriteString(strings.Join([]string{config.GetFullVersion(), config.GetStatusValue().Code, config.GetStatusValue().Reason}, " ")))
	header.AppendSupplier(writeCRLF)

	headers := config.PickHeaders()
	for _, h := range headers {
		header.AppendSupplier(serial.WriteString(h))
		header.AppendSupplier(writeCRLF)
	}
	if !config.HasHeader("Date") {
		header.AppendSupplier(serial.WriteString("Date: "))
		header.AppendSupplier(serial.WriteString(time.Now().Format(http.TimeFormat)))
		header.AppendSupplier(writeCRLF)
	}
	header.AppendSupplier(writeCRLF)
	return &HeaderWriter{
		header: header,
	}
}

type HttpAuthenticator struct {
	config *Config
}

func (v HttpAuthenticator) GetClientWriter() *HeaderWriter {
	header := buf.NewSmall()
	config := v.config.Request
	header.AppendSupplier(serial.WriteString(strings.Join([]string{config.GetMethodValue(), config.PickUri(), config.GetFullVersion()}, " ")))
	header.AppendSupplier(writeCRLF)

	headers := config.PickHeaders()
	for _, h := range headers {
		header.AppendSupplier(serial.WriteString(h))
		header.AppendSupplier(writeCRLF)
	}
	header.AppendSupplier(writeCRLF)
	return &HeaderWriter{
		header: header,
	}
}

func (v HttpAuthenticator) GetServerWriter() *HeaderWriter {
	return formResponseHeader(v.config.Response)
}

func (v HttpAuthenticator) Client(conn net.Conn) net.Conn {
	if v.config.Request == nil && v.config.Response == nil {
		return conn
	}
	var reader Reader = new(NoOpReader)
	if v.config.Request != nil {
		reader = new(HeaderReader)
	}

	var writer Writer = new(NoOpWriter)
	if v.config.Response != nil {
		writer = v.GetClientWriter()
	}
	return NewHttpConn(conn, reader, writer, new(NoOpWriter))
}

func (v HttpAuthenticator) Server(conn net.Conn) net.Conn {
	if v.config.Request == nil && v.config.Response == nil {
		return conn
	}
	return NewHttpConn(conn, new(HeaderReader), v.GetServerWriter(), formResponseHeader(&ResponseConfig{
		Version: &Version{
			Value: "1.1",
		},
		Status: &Status{
			Code:   "500",
			Reason: "Internal Server Error",
		},
		Header: []*Header{
			{
				Name:  "Connection",
				Value: []string{"close"},
			},
			{
				Name:  "Cache-Control",
				Value: []string{"private"},
			},
			{
				Name:  "Content-Length",
				Value: []string{"0"},
			},
		},
	}))
}

func NewHttpAuthenticator(ctx context.Context, config *Config) (HttpAuthenticator, error) {
	return HttpAuthenticator{
		config: config,
	}, nil
}

func init() {
	common.Must(common.RegisterConfig((*Config)(nil), func(ctx context.Context, config interface{}) (interface{}, error) {
		return NewHttpAuthenticator(ctx, config.(*Config))
	}))
}
