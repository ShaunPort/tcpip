package tcpip_test

import (
	"net"
	"tcpip"

	"testing"
)

func TestInt32ToIP(t *testing.T) {
	expectedIP := net.ParseIP("192.168.1.2")
	var i = 3232235778
	resultIP := tcpip.Uint32toIP(uint32(i))
	if !resultIP.Equal(expectedIP) {
		t.Errorf("Expected %s, but got %s", expectedIP, resultIP)
	}
}

func TestIPToInt32(t *testing.T) {
	var expectedValueInt32 = 3232235778
	ip := net.ParseIP("192.168.1.2")
	resultIPInt32 := tcpip.IPToUint32(ip)
	if resultIPInt32 != uint32(expectedValueInt32) {
		t.Errorf("Expected %d, but got %d", uint32(expectedValueInt32), resultIPInt32)
	}
}
