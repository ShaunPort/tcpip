package tcpip

import (
	"encoding/binary"
	"net"
)

/**
 * 获取IP的版本
 */
func Version(pkg []byte) int {
	return int(pkg[0]) >> 4
}

/**
 *	从字符数组中解析ipv4 start初始位置下标索引 end是最后位置下标索引+1
 *  后续需要完善判断pkg长度是否大于20
 */
func ParseIPv4(pkg []byte, start int, end int) (net.IP, error) {
	v := Version(pkg)
	if v != 4 && v != 6 {
		return nil, ErrNotIP
	}
	if v != 4 {
		return nil, ErrNotIPv4
	}
	return pkg[start:end], nil
}

/**
 * 从pkg报文中获取ipv4源地址
 */
func ParseIPv4Src(pkg []byte) (net.IP, error) {
	return ParseIPv4(pkg, 12, 16)
}

/**
 * 从pkg报文中获取ipv4目标地址
 */
func ParseIPv4Dst(pkg []byte) (net.IP, error) {
	return ParseIPv4(pkg, 16, 20)
}

/**
 *	把点分十进制ip地址写入字符数组pkg start初始位置下标索引 end是最后位置下标索引+1
 */
func AdjustIPv4(pkg []byte, v4 string, start int, end int) error {
	v := Version(pkg)
	if v != 4 && v != 6 {
		return ErrNotIP
	}
	if v != 4 {
		return ErrNotIPv4
	}
	copy(pkg[start:end], net.ParseIP(v4).To4())
	return nil
}

/**
 *	写入IPv4源地址到pkg报文中
 */
func AdjustIPv4Src(pkg []byte, v4 string) error {
	return AdjustIPv4(pkg, v4, 12, 16)
}

/**
 *	写入IPv4目标地址到pkg报文中
 */
func AdjustIPv4Dst(pkg []byte, v4 string) error {
	return AdjustIPv4(pkg, v4, 16, 20)
}

// Uint32toIP uint32转IP 不用考虑字节序，已经处理了
func Uint32toIP(uint32Value uint32) net.IP {
	ipBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(ipBytes, uint32Value)
	return net.IP(ipBytes)
}

// IPToUint32 IP转uint32 不用考虑字节序，已经处理了
func IPToUint32(ip net.IP) uint32 {
	ip = ip.To4()
	if ip == nil {
		return 0
	}
	return binary.BigEndian.Uint32(ip)
}

// SumIP 根据偏移量计算IP地址
func SumIP(ip net.IP, offset int) net.IP {
	// 将IP转换为32位整数
	ipInt := IPToUint32(ip)
	// 计算偏移后的IP地址
	offsetIPInt := ipInt + uint32(offset)
	// 将偏移后的IP转换为4字节表示
	offsetIP := Uint32toIP(offsetIPInt)
	return offsetIP
}

// DiffIP 计算两个IP地址的差
func DiffIP(s, e net.IP) uint32 {
	start := IPToUint32(s)
	end := IPToUint32(e)
	if start > end {
		return start - end
	}
	return end - start
}

// CountIPNum 通过起始地址、结束地址计算地址总数
func CountIPNum(startIP, endIP net.IP) uint32 {
	dif := DiffIP(startIP, endIP)
	totalNum := dif + 1
	return totalNum
}

// CountHostWithIPNet 通过掩码计算得到CIDR地址范围内可拥有的主机数（去除广播地址和组播地址）
func CountHostWithIPNet(ipNet *net.IPNet) uint32 {
	return CountIPNumWithIPNet(ipNet) - 2
}

// CountIPNumWithIPNet 通过掩码计算得到CIDR地址范围内可拥有的地址总数
func CountIPNumWithIPNet(ipNet *net.IPNet) uint32 {
	return CountIPNumWithMask(&ipNet.Mask)
}

// CountIPNumWithMask 通过掩码计算得到CIDR地址范围内地址总数
func CountIPNumWithMask(mask *net.IPMask) uint32 {
	prefixLen, bits := mask.Size()
	total := 1 << (uint64(bits) - uint64(prefixLen))
	return uint32(total)
}

// ObtainRangeSubnet 得到一段子网的IP地址范围
func ObtainRangeSubnet(ipNet *net.IPNet) (startIP, endIP net.IP) {
	startIP = ipNet.IP
	endIP = make(net.IP, len(startIP))
	copy(endIP, startIP)
	for i := range startIP {
		endIP[i] |= ^ipNet.Mask[i]
	}
	return startIP, endIP
}
