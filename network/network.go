// Copyright © 2020-2021 The EVEN Solutions Developers Team

package network

import (
	"net"
	"strings"
)

// UpInterfaces returns a list of the system's up network interfaces,
// optionally additional flags can be specified.
func UpInterfaces(fl ...net.Flags) (l []net.Interface) {
	var flags net.Flags
	for _, f := range fl {
		flags |= f
	}

	list, _ := net.Interfaces()
	for _, i := range list {
		switch {
		case i.Flags&net.FlagUp == 0:
			continue // skip - is not up interface
		case i.Flags&net.FlagLoopback != 0:
			continue // skip - is loopback interface
		case flags != 0 && i.Flags&flags == 0:
			continue // skip interface by additional flags
		default:
			l = append(l, i) // append interface to result list
		}
	}

	return l
}

// IPAddresses returns a list ip addresses
// of the system's up network interfaces.
func IPAddresses() (l []net.IP) {
	for _, i := range UpInterfaces() {
		addrs, _ := i.Addrs()
		for _, addr := range addrs {
			ip := strings.Split(addr.String(), "/")[0] // trim ip mask
			l = append(l, net.ParseIP(ip))
		}
	}

	return l
}
