package utils

import (
	"errors"
	"net"
)

var ipCache = make(map[string]net.IP, 0)

// IP 获取本机ip,由于有多网卡的情况,不保证一定准确,传入同一网段的ip以提高可靠性.
// 在使用全局代理的情况下可能会有较大的错误.
func IP(neighbor net.IP) (ip net.IP, defaultError error) {
	if ans, ok := ipCache[neighbor.String()]; ok {
		return ans, nil
	}
	if neighbor == nil {
		neighbor = net.IPv4(8, 8, 8, 8)
	}
	defaultError = errors.New("network error")
	// 仅产生一个udp报文,但不实际地发送
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{IP: neighbor, Port: 65535})
	if err != nil {
		return nil, defaultError
	}
	defer func() { _ = socket.Close() }()
	localAddr, ok := socket.LocalAddr().(*net.UDPAddr)
	if !ok {
		return nil, defaultError
	}
	return localAddr.IP, err
}
