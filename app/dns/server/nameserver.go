package server

import (
	"context"
	"net"
	"sync"
	"time"

	"github.com/miekg/dns"
	"v2ray.com/core/app/dispatcher"
	"v2ray.com/core/common/buf"
	"v2ray.com/core/common/dice"
	"v2ray.com/core/common/log"
	v2net "v2ray.com/core/common/net"
	"v2ray.com/core/transport/internet/udp"
)

const (
	DefaultTTL       = uint32(3600)
	CleanupInterval  = time.Second * 120
	CleanupThreshold = 512
)

var (
	pseudoDestination = v2net.UDPDestination(v2net.LocalHostIP, v2net.Port(53))
)

type ARecord struct {
	IPs    []net.IP
	Expire time.Time
}

type NameServer interface {
	QueryA(domain string) <-chan *ARecord
}

type PendingRequest struct {
	expire   time.Time
	response chan<- *ARecord
}

type UDPNameServer struct {
	sync.Mutex
	address     v2net.Destination
	requests    map[uint16]*PendingRequest
	udpServer   *udp.Dispatcher
	nextCleanup time.Time
}

func NewUDPNameServer(address v2net.Destination, dispatcher dispatcher.Interface) *UDPNameServer {
	s := &UDPNameServer{
		address:   address,
		requests:  make(map[uint16]*PendingRequest),
		udpServer: udp.NewDispatcher(dispatcher),
	}
	return s
}

// Private: Visible for testing.
func (v *UDPNameServer) Cleanup() {
	expiredRequests := make([]uint16, 0, 16)
	now := time.Now()
	v.Lock()
	for id, r := range v.requests {
		if r.expire.Before(now) {
			expiredRequests = append(expiredRequests, id)
			close(r.response)
		}
	}
	for _, id := range expiredRequests {
		delete(v.requests, id)
	}
	v.Unlock()
	expiredRequests = nil
}

// Private: Visible for testing.
func (v *UDPNameServer) AssignUnusedID(response chan<- *ARecord) uint16 {
	var id uint16
	v.Lock()
	if len(v.requests) > CleanupThreshold && v.nextCleanup.Before(time.Now()) {
		v.nextCleanup = time.Now().Add(CleanupInterval)
		go v.Cleanup()
	}

	for {
		id = uint16(dice.Roll(65536))
		if _, found := v.requests[id]; found {
			continue
		}
		log.Debug("DNS: Add pending request id ", id)
		v.requests[id] = &PendingRequest{
			expire:   time.Now().Add(time.Second * 8),
			response: response,
		}
		break
	}
	v.Unlock()
	return id
}

// Private: Visible for testing.
func (v *UDPNameServer) HandleResponse(payload *buf.Buffer) {
	msg := new(dns.Msg)
	err := msg.Unpack(payload.Bytes())
	if err != nil {
		log.Warning("DNS: Failed to parse DNS response: ", err)
		return
	}
	record := &ARecord{
		IPs: make([]net.IP, 0, 16),
	}
	id := msg.Id
	ttl := DefaultTTL
	log.Debug("DNS: Handling response for id ", id, " content: ", msg.String())

	v.Lock()
	request, found := v.requests[id]
	if !found {
		v.Unlock()
		return
	}
	delete(v.requests, id)
	v.Unlock()

	for _, rr := range msg.Answer {
		switch rr := rr.(type) {
		case *dns.A:
			record.IPs = append(record.IPs, rr.A)
			if rr.Hdr.Ttl < ttl {
				ttl = rr.Hdr.Ttl
			}
		case *dns.AAAA:
			record.IPs = append(record.IPs, rr.AAAA)
			if rr.Hdr.Ttl < ttl {
				ttl = rr.Hdr.Ttl
			}
		}
	}
	record.Expire = time.Now().Add(time.Second * time.Duration(ttl))

	request.response <- record
	close(request.response)
}

func (v *UDPNameServer) BuildQueryA(domain string, id uint16) *buf.Buffer {

	msg := new(dns.Msg)
	msg.Id = id
	msg.RecursionDesired = true
	msg.Question = []dns.Question{
		{
			Name:   dns.Fqdn(domain),
			Qtype:  dns.TypeA,
			Qclass: dns.ClassINET,
		}}

	buffer := buf.New()
	buffer.AppendSupplier(func(b []byte) (int, error) {
		writtenBuffer, err := msg.PackBuffer(b)
		return len(writtenBuffer), err
	})

	return buffer
}

func (v *UDPNameServer) QueryA(domain string) <-chan *ARecord {
	response := make(chan *ARecord, 1)
	id := v.AssignUnusedID(response)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*8)
	v.udpServer.Dispatch(ctx, v.address, v.BuildQueryA(domain, id), v.HandleResponse)

	go func() {
		for i := 0; i < 2; i++ {
			time.Sleep(time.Second)
			v.Lock()
			_, found := v.requests[id]
			v.Unlock()
			if found {
				v.udpServer.Dispatch(ctx, v.address, v.BuildQueryA(domain, id), v.HandleResponse)
			} else {
				break
			}
		}
		cancel()
	}()

	return response
}

type LocalNameServer struct {
}

func (v *LocalNameServer) QueryA(domain string) <-chan *ARecord {
	response := make(chan *ARecord, 1)

	go func() {
		defer close(response)

		ips, err := net.LookupIP(domain)
		if err != nil {
			log.Info("DNS: Failed to lookup IPs for domain ", domain)
			return
		}

		response <- &ARecord{
			IPs:    ips,
			Expire: time.Now().Add(time.Second * time.Duration(DefaultTTL)),
		}
	}()

	return response
}
