package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func echoServer(c net.Conn) {
	for {
		//buf := make([]byte, 512)
		//nr, err := c.Read(buf)
		//if err != nil {
		//	return
		//}
		//
		//data := buf[0:1024]
		//println("Server got:", string(data))
		println("Writing data:", "testetstesetset")
		_, err := c.Write([]byte("testetstesetset"))
		if err != nil {
			log.Fatal("Writing client error: ", err)
		}
	}
}

func main() {
	log.Println("Starting echo server")
	_ = syscall.Unlink("/tmp/go.sock")
	ln, err := net.Listen("unix", "/tmp/go.sock")
	if err != nil {
		log.Fatal("Listen error: ", err)
	}
	
	defer ln.Close()
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt, syscall.SIGTERM)
	go func(ln net.Listener, c chan os.Signal) {
		sig := <-c
		log.Printf("Caught signal %s: shutting down.", sig)
		os.Exit(0)
	}(ln, sigc)
	
	for {
		fd, err := ln.Accept()
		if err != nil {
			log.Fatal("Accept error: ", err)
		}
		defer fd.Close()
		data := "test1\ntest2\ntest3\n"
		println("Writing data:", data)
		_, err = fd.Write([]byte(data))
		if err != nil {
			log.Fatal("Writing client error: ", err)
		}
	}
}