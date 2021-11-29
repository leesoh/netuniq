package iputils

import "net"

func IsIP(target string) bool {
	return net.ParseIP(target) != nil
}

func IsCIDR(target string) bool {
	_, ipnet, _ := net.ParseCIDR(target)
	return ipnet != nil
}
