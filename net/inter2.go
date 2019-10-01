package main

import (
	"fmt"
	"net"
	"strings"
)

func findIPAddress() string {
	if interfaces, err := net.Interfaces(); err == nil {
		for _, interfac := range interfaces {
		
			if interfac.HardwareAddr.String() != "" && strings.Contains(strings.ToLower(interfac.Flags.String()), "up") {
				if addrs, err := interfac.Addrs(); err == nil {
					for _, addr := range addrs {
						println("gs ",interfac.Flags.String(), addr.String())
						//if addr.Network() == "ip+net" {
						//	pr := strings.Split(addr.String(), "/")
						//	if len(pr) == 2 && len(strings.Split(pr[0], ".")) == 4 {
						//		return pr[0]
						//	}
						//}
					}
				}
			}
		}
	}
	return ""
}

type MachineDetails struct {
	MacAddr   string
	Ipv4Local string
	Ipv6Local string
}


func GetMachineDetails(interfaceName string) (machineInfo MachineDetails, err error) {
	var itf *net.Interface
	var addrs []net.Addr
	
	if itf, err = net.InterfaceByName(interfaceName); err != nil {
		return
	}
	if addrs, err = itf.Addrs(); err != nil {
		return
	}
	
	if itf.HardwareAddr != nil {
		machineInfo.MacAddr = itf.HardwareAddr.String()
	}
	
	for _, addr := range addrs {
		switch v := addr.(type) {
		case *net.IPNet:
			if !v.IP.IsLoopback() {
				if v.IP.To4() != nil { /* Check if IP is IPV4 */
					machineInfo.Ipv4Local = v.IP.String()
				} else if len(v.IP) == net.IPv6len {  /* Check if IP is IPV6 */
					machineInfo.Ipv6Local = v.IP.String()
				}
			}
		}
	}
	return
}


func main() {
	m, err := GetMachineDetails("awdl0")
	fmt.Println(err, m)
	//println(findIPAddress())
}