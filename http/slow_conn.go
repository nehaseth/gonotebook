package main
//
//import (
//	"fmt"
//	"io"
//	"log"
//	"net"
//	"net/http"
//	"os"
//	"time"
//)
//
//func main() {
//	c := http.Client{
//		Transport: &http.Transport{
//			Proxy:               http.ProxyFromEnvironment,
//			Dial:                slowDial,
//			TLSHandshakeTimeout: 10 * time.Second,
//		},
//		Timeout: 5 * time.Second,
//	}
//	res, err := c.Get("http://www.golang.org")
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer res.Body.Close()
//	io.Copy(os.Stdout, res.Body)
//}
//
//// slow connection
//type slowConn struct {
//	net.Conn
//	sr slowReader
//}
//
//func slowDial(network, addr string) (net.Conn, error) {
//	conn, err := net.Dial(network, addr)
//	if err != nil {
//		return nil, err
//	}
//	return slowConn{conn, slowReader{conn}}, nil
//}
//
//func (conn slowConn) Read(data []byte) (int, error) {
//	return conn.sr.Read(data)
//}
//
//// slow reader
//type slowReader struct{ r io.Reader }
//
//func (r slowReader) Read(data []byte) (int, error) {
//	// wait for 250 ms before reading a single byte.
//	time.Sleep(10 * time.Millisecond)
//	n, err := r.r.Read(data[:1])
//	if n > 0 {
//		fmt.Printf("%s", data[:1])
//	}
//	return n, err
//}
