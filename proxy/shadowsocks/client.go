package shadowsocks

import (
	"context"

	"v2ray.com/core/common"
	"v2ray.com/core/common/buf"
	"v2ray.com/core/common/bufio"
	"v2ray.com/core/common/log"
	"v2ray.com/core/common/net"
	"v2ray.com/core/common/protocol"
	"v2ray.com/core/common/retry"
	"v2ray.com/core/common/signal"
	"v2ray.com/core/proxy"
	"v2ray.com/core/transport/internet"
	"v2ray.com/core/transport/ray"
)

// Client is a inbound handler for Shadowsocks protocol
type Client struct {
	serverPicker protocol.ServerPicker
}

// NewClient create a new Shadowsocks client.
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

// Process implements OutboundHandler.Process().
func (v *Client) Process(ctx context.Context, outboundRay ray.OutboundRay) error {
	destination := proxy.DestinationFromContext(ctx)
	network := destination.Network

	var server *protocol.ServerSpec
	var conn internet.Connection

	dialer := proxy.DialerFromContext(ctx)
	err := retry.ExponentialBackoff(5, 100).On(func() error {
		server = v.serverPicker.PickServer()
		dest := server.Destination()
		dest.Network = network
		rawConn, err := dialer.Dial(ctx, dest)
		if err != nil {
			return err
		}
		conn = rawConn

		return nil
	})
	if err != nil {
		log.Warning("Shadowsocks|Client: Failed to find an available destination:", err)
		return err
	}
	log.Info("Shadowsocks|Client: Tunneling request to ", destination, " via ", server.Destination())

	conn.SetReusable(false)

	request := &protocol.RequestHeader{
		Version: Version,
		Address: destination.Address,
		Port:    destination.Port,
	}
	if destination.Network == net.Network_TCP {
		request.Command = protocol.RequestCommandTCP
	} else {
		request.Command = protocol.RequestCommandUDP
	}

	user := server.PickUser()
	rawAccount, err := user.GetTypedAccount()
	if err != nil {
		log.Warning("Shadowsocks|Client: Failed to get a valid user account: ", err)
		return err
	}
	account := rawAccount.(*ShadowsocksAccount)
	request.User = user

	if account.OneTimeAuth == Account_Auto || account.OneTimeAuth == Account_Enabled {
		request.Option |= RequestOptionOneTimeAuth
	}

	if request.Command == protocol.RequestCommandTCP {
		bufferedWriter := bufio.NewWriter(conn)
		bodyWriter, err := WriteTCPRequest(request, bufferedWriter)
		if err != nil {
			log.Info("Shadowsocks|Client: Failed to write request: ", err)
			return err
		}

		bufferedWriter.SetBuffered(false)

		requestDone := signal.ExecuteAsync(func() error {
			if err := buf.PipeUntilEOF(outboundRay.OutboundInput(), bodyWriter); err != nil {
				return err
			}
			return nil
		})

		responseDone := signal.ExecuteAsync(func() error {
			defer outboundRay.OutboundOutput().Close()

			responseReader, err := ReadTCPResponse(user, conn)
			if err != nil {
				return err
			}

			if err := buf.PipeUntilEOF(responseReader, outboundRay.OutboundOutput()); err != nil {
				return err
			}

			return nil
		})

		if err := signal.ErrorOrFinish2(requestDone, responseDone); err != nil {
			log.Info("Shadowsocks|Client: Connection ends with ", err)
			outboundRay.OutboundInput().CloseError()
			outboundRay.OutboundOutput().CloseError()
			return err
		}

		return nil
	}

	if request.Command == protocol.RequestCommandUDP {

		writer := &UDPWriter{
			Writer:  conn,
			Request: request,
		}

		requestDone := signal.ExecuteAsync(func() error {
			if err := buf.PipeUntilEOF(outboundRay.OutboundInput(), writer); err != nil {
				log.Info("Shadowsocks|Client: Failed to transport all UDP request: ", err)
				return err
			}
			return nil
		})

		timedReader := net.NewTimeOutReader(16, conn)

		responseDone := signal.ExecuteAsync(func() error {
			defer outboundRay.OutboundOutput().Close()

			reader := &UDPReader{
				Reader: timedReader,
				User:   user,
			}

			if err := buf.PipeUntilEOF(reader, outboundRay.OutboundOutput()); err != nil {
				log.Info("Shadowsocks|Client: Failed to transport all UDP response: ", err)
				return err
			}
			return nil
		})

		if err := signal.ErrorOrFinish2(requestDone, responseDone); err != nil {
			log.Info("Shadowsocks|Client: Connection ends with ", err)
			outboundRay.OutboundInput().CloseError()
			outboundRay.OutboundOutput().CloseError()
			return err
		}

		return nil
	}

	return nil
}

func init() {
	common.Must(common.RegisterConfig((*ClientConfig)(nil), func(ctx context.Context, config interface{}) (interface{}, error) {
		return NewClient(ctx, config.(*ClientConfig))
	}))
}
