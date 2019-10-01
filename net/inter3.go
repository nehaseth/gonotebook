package main

import (
	"fmt"
	"net"
)

func GetNetworkInterface() {
	if interfaces, err := net.Interfaces(); err == nil {
		for _, iface := range interfaces {
			if iface.Flags&net.FlagUp == 0 { // interface down
				continue
			}
			if iface.Flags&net.FlagLoopback != 0 { // loopback interface
				continue
			}
			addrs, err := iface.Addrs()
			if err != nil {
				return //err
			}
			
			fmt.Println("tset ", iface)
			for _, addr := range addrs {
				var ip net.IP
				switch v := addr.(type) {
				case *net.IPNet:
					ip = v.IP
				case *net.IPAddr:
					ip = v.IP
				}
				if ip == nil || ip.IsLoopback() {
					continue
				}
				ip = ip.To4()
				if ip == nil {
					continue // not an ipv4 address
				}
				fmt.Println(ip.String())
				fmt.Println(ip.To16())
			}
		}
	}
}

func GetDetailsByInterface() {

}

func main() {
	GetNetworkInterface()
}