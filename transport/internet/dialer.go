package internet

import (
	"context"
	"net"

	"v2ray.com/core/common/errors"
	v2net "v2ray.com/core/common/net"
)

type Dialer func(ctx context.Context, dest v2net.Destination) (Connection, error)

var (
	transportDialerCache = make(map[TransportProtocol]Dialer)
)

func RegisterTransportDialer(protocol TransportProtocol, dialer Dialer) error {
	if _, found := transportDialerCache[protocol]; found {
		return errors.New("Internet|Dialer: ", protocol, " dialer already registered.")
	}
	transportDialerCache[protocol] = dialer
	return nil
}

func Dial(ctx context.Context, dest v2net.Destination) (Connection, error) {
	if dest.Network == v2net.Network_TCP {
		streamSettings, _ := StreamSettingsFromContext(ctx)
		protocol := streamSettings.GetEffectiveProtocol()
		transportSettings, err := streamSettings.GetEffectiveTransportSettings()
		if err != nil {
			return nil, err
		}
		ctx = ContextWithTransportSettings(ctx, transportSettings)
		if streamSettings != nil && streamSettings.HasSecuritySettings() {
			securitySettings, err := streamSettings.GetEffectiveSecuritySettings()
			if err != nil {
				return nil, err
			}
			ctx = ContextWithSecuritySettings(ctx, securitySettings)
		}
		dialer := transportDialerCache[protocol]
		if dialer == nil {
			return nil, errors.New("Internet|Dialer: ", protocol, " dialer not registered.")
		}
		return dialer(ctx, dest)
	}

	udpDialer := transportDialerCache[TransportProtocol_UDP]
	if udpDialer == nil {
		return nil, errors.New("Internet|Dialer: UDP dialer not registered.")
	}
	return udpDialer(ctx, dest)
}

// DialSystem calls system dialer to create a network connection.
func DialSystem(src v2net.Address, dest v2net.Destination) (net.Conn, error) {
	return effectiveSystemDialer.Dial(src, dest)
}
