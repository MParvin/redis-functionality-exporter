package tools

import (
	"net"
)

// ValidateIPV4 will validate ipv4 address
func ValidateIPV4(ip string) bool {
	return net.ParseIP(ip).To4() != nil
}

// ValidateIPV6 will validate ipv6 address
func ValidateIPV6(ip string) bool {
	return net.ParseIP(ip).To16() != nil
}

// ValidateDomain will validate domain name
func ValidateDomain(domain string) bool {
	return net.ParseIP(domain) == nil
}

func ValidateIP(ip string) bool {
	return ValidateIPV4(ip) || ValidateIPV6(ip) || ValidateDomain(ip)
}
