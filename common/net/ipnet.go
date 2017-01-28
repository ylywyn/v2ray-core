package net

import (
	"net"
)

var (
	onesCount = make(map[byte]byte)
)

type IPNet struct {
	cache map[uint32]byte
}

func NewIPNet() *IPNet {
	return &IPNet{
		cache: make(map[uint32]byte, 1024),
	}
}

func ipToUint32(ip net.IP) uint32 {
	value := uint32(0)
	for _, b := range []byte(ip) {
		value <<= 8
		value += uint32(b)
	}
	return value
}

func ipMaskToByte(mask net.IPMask) byte {
	value := byte(0)
	for _, b := range []byte(mask) {
		value += onesCount[b]
	}
	return value
}

func (v *IPNet) Add(ipNet *net.IPNet) {
	ipv4 := ipNet.IP.To4()
	if ipv4 == nil {
		// For now, we don't support IPv6
		return
	}
	mask := ipMaskToByte(ipNet.Mask)
	v.AddIP(ipv4, mask)
}

func (v *IPNet) AddIP(ip []byte, mask byte) {
	k := ipToUint32(ip)
	existing, found := v.cache[k]
	if !found || existing > mask {
		v.cache[k] = mask
	}
}

func (v *IPNet) Contains(ip net.IP) bool {
	ipv4 := ip.To4()
	if ipv4 == nil {
		return false
	}
	originalValue := ipToUint32(ipv4)

	if entry, found := v.cache[originalValue]; found {
		if entry == 32 {
			return true
		}
	}

	mask := uint32(0)
	for maskbit := byte(1); maskbit <= 32; maskbit++ {
		mask += 1 << uint32(32-maskbit)

		maskedValue := originalValue & mask
		if entry, found := v.cache[maskedValue]; found {
			if entry == maskbit {
				return true
			}
		}
	}
	return false
}

func (v *IPNet) IsEmpty() bool {
	return len(v.cache) == 0
}

func init() {
	value := byte(0)
	for mask := byte(1); mask <= 8; mask++ {
		value += 1 << byte(8-mask)
		onesCount[value] = mask
	}
}
