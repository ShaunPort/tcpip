package tcpip

import "errors"

// 错误列表
var ErrNotIP = errors.New("Not IPv4")                             // 不是IP数据包
var ErrNotIPv4 = errors.New("Not IPv4 package")                   // 不是IPv4 无法解析
var ErrNotIPv6 = errors.New("Not IPv6 package")                   // 不是IPv4 无法解析
var ErrCountIPNum = errors.New("Start IP is greater than end IP") // 起始IP>结束IP
var ErrNullIP = errors.New("IP is null")                          // IP地址为空
