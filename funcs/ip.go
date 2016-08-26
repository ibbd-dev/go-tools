package funcs

import (
	"errors"
	"net"
)

// 将IPv4地址转为uint32的结果类型
// 注意：ParseIP的结果
func Ip2uint(ip_str string) (uint32, error) {
	ip := net.ParseIP(ip_str)
	if ip == nil {
		return 0, errors.New("ParseIP error")
	}

	return uint32(ip[12])<<24 | uint32(ip[13])<<16 | uint32(ip[14])<<8 | uint32(ip[15]), nil
}

// 获取本地IP地址
func GetInternalIp() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}
