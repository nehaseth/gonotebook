package main

import (
	"fmt"
	"net"
)

func localAddresses() {
	ifaces, err := net.Interfaces()
	if err != nil {
		fmt.Print(fmt.Errorf("localAddresses: %+v\n", err.Error()))
		return
	}
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			fmt.Print(fmt.Errorf("localAddresses: %+v\n", err.Error()))
			continue
		}
		for _, a := range addrs {
			fmt.Println(a)
			switch v := a.(type) {
			case *net.IPAddr:
				fmt.Printf("ip addr type %v : %s (%s)\n", i.Name, v, v.IP.DefaultMask())
			
			case *net.IPNet:
				fmt.Printf("ip net type %v : %s [%v/   %v]\n", i.Name, v, v.IP, v.Mask)
			}
			
		}
	}
}

func main() {




}