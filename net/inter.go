package main

import (
	"fmt"
	"net"
	"os"
)

func myip() {
	os.Stdout.WriteString("myip:\n")
	
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Errorf("error: %v\n", err.Error())
		return
	}
	
	for _, a := range addrs {
		ip := net.ParseIP(a.String())
		fmt.Printf("addr: %v loopback=%v\n", a, ip.IsLoopback())
	}
	
	fmt.Println()
}

func myip2() {
	os.Stdout.WriteString("myip2:\n")
	
	tt, err := net.Interfaces()
	if err != nil {
		fmt.Errorf("error: %v\n", err.Error())
		return
	}
	for _, t := range tt {
		aa, err := t.Addrs()
		if err != nil {
			fmt.Errorf("error: %v\n", err.Error())
			continue
		}
		for _, a := range aa {
			ip := net.ParseIP(a.String())
			fmt.Printf("%v addr: %v loopback=%v\n", t.Name, a, ip.IsLoopback())
		}
	}
	
	fmt.Println()
}

func main() {
	fmt.Println("myip -- begin")
	myip()
	myip2()
	fmt.Println("myip -- end")
}