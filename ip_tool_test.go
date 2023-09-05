package tcpip_test

import (
	"net"

	"testing"

	"github.com/ShaunPort/tcpip"
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

func TestSumIP(t *testing.T) {
	type args struct {
		ip     net.IP
		offset int
	}
	tests := []struct {
		name    string
		args    args
		want    net.IP
		wantErr bool
	}{
		// Add test cases.
		{
			name: "test1",
			args: args{
				ip:     net.ParseIP("192.168.0.1"),
				offset: 1,
			},
			want:    net.ParseIP("192.168.0.2"),
			wantErr: false,
		},
		{
			name: "test2",
			args: args{
				ip:     net.ParseIP("192.168.0.255"),
				offset: 1,
			},
			want:    net.ParseIP("192.168.1.0"),
			wantErr: false,
		},
		{
			name: "test3",
			args: args{
				ip:     net.ParseIP("255.255.255.255"),
				offset: 1,
			},
			want:    net.ParseIP("0.0.0.0"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tcpip.SumIP(tt.args.ip, tt.args.offset)
			var err error
			if (err != nil) != tt.wantErr {
				t.Errorf("tcpip.SumIP() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// reflect.DeepEqual(got, tt.want)
			if !got.Equal(tt.want) {
				t.Errorf("tcpip.SumIP() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDiffIP(t *testing.T) {
	type args struct {
		ip1 net.IP
		ip2 net.IP
	}
	tests := []struct {
		name    string
		args    args
		want    uint32
		wantErr bool
	}{
		// Add test cases.
		{
			name: "test1",
			args: args{
				ip1: net.ParseIP("192.168.0.1"),
				ip2: net.ParseIP("192.168.0.1"),
			},
			want:    0,
			wantErr: false,
		},
		{
			name: "test2",
			args: args{
				ip1: net.ParseIP("192.168.0.2"),
				ip2: net.ParseIP("192.168.0.1"),
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "test3",
			args: args{
				ip1: net.ParseIP("192.168.0.255"),
				ip2: net.ParseIP("192.168.0.1"),
			},
			want:    254,
			wantErr: false,
		},
		{
			name: "test4",
			args: args{
				ip1: net.ParseIP("0.0.0.255"),
				ip2: net.ParseIP("0.0.0.0"),
			},
			want:    255,
			wantErr: false,
		},
		{
			name: "test5",
			args: args{
				ip1: net.ParseIP("192.168.0.1"),
				ip2: net.ParseIP("192.168.0.255"),
			},
			want:    254,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tcpip.DiffIP(tt.args.ip1, tt.args.ip2)
			var err error
			if (err != nil) != tt.wantErr {
				t.Errorf("tcpip.SumIP() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// reflect.DeepEqual(got, tt.want)
			if got != tt.want {
				t.Errorf("tcpip.SumIP() got = %v, want %v", got, tt.want)
			}
		})
	}
}
