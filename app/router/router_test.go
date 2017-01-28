package router_test

import (
	"context"
	"testing"

	"v2ray.com/core/app"
	"v2ray.com/core/app/dispatcher"
	_ "v2ray.com/core/app/dispatcher/impl"
	"v2ray.com/core/app/dns"
	_ "v2ray.com/core/app/dns/server"
	"v2ray.com/core/app/proxyman"
	_ "v2ray.com/core/app/proxyman/outbound"
	. "v2ray.com/core/app/router"
	"v2ray.com/core/common/net"
	"v2ray.com/core/proxy"
	"v2ray.com/core/testing/assert"
)

func TestSimpleRouter(t *testing.T) {
	assert := assert.On(t)

	config := &Config{
		Rule: []*RoutingRule{
			{
				Tag: "test",
				NetworkList: &net.NetworkList{
					Network: []net.Network{net.Network_TCP},
				},
			},
		},
	}

	space := app.NewSpace()
	ctx := app.ContextWithSpace(context.Background(), space)
	assert.Error(app.AddApplicationToSpace(ctx, new(dns.Config))).IsNil()
	assert.Error(app.AddApplicationToSpace(ctx, new(dispatcher.Config))).IsNil()
	assert.Error(app.AddApplicationToSpace(ctx, new(proxyman.OutboundConfig))).IsNil()
	assert.Error(app.AddApplicationToSpace(ctx, config)).IsNil()
	assert.Error(space.Initialize()).IsNil()

	r := FromSpace(space)

	ctx = proxy.ContextWithDestination(ctx, net.TCPDestination(net.DomainAddress("v2ray.com"), 80))
	tag, err := r.TakeDetour(ctx)
	assert.Error(err).IsNil()
	assert.String(tag).Equals("test")
}
