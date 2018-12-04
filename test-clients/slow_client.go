package main

import (
	"log"
	"flag"
	"time"
	"net"
	"net/http"
	"sync/atomic"
	"io/ioutil"
	"io"
)

var ops uint32 = 0
func sendRequest(req string) {

	tr := &http.Transport{
		Dial:                slowDial,
		TLSHandshakeTimeout: 10 * time.Second,
		DisableKeepAlives: false,
		MaxIdleConnsPerHost: 1,
		//DisableKeepAlives:   !(*keepAlive),
		//MaxIdleConnsPerHost: *concurrency,
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get(req)
	if err != nil {
		log.Println(err)
	} else {
		_, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			log.Println(err)
		}
		atomic.AddUint32(&ops, 1)
	}
}

var (
	slowRead *int
	defReq = "http://10.1.1.0/fast"
)

func main() {
	concurrency := flag.Int("concurrency", 1000, "Concurent connection to the server")
	maxQPS := flag.Int("maxQPS", 100, "Maximum QPS the client will generate")
	req := flag.String("req", defReq, "Request url")
	keepAlive := flag.Bool("keepAlive", true, "Whether to keep connection alive, if enabled keep alive for 5 min")
	slowRead = flag.Int("slowRead", 50, "Slow read per byte time in milliseconds")

	flag.Parse()

	log.Printf("Current concurrency: %d\n", *concurrency)
	log.Printf("Current max QPS: %d\n", *maxQPS)
	log.Printf("KeepAlive: %v\n", *keepAlive)
	log.Printf("Url: %s\n", *req)
	log.Printf("SlowRead: %d\n", *slowRead)

	fin := make(chan bool)
	bucket := make(chan bool, *maxQPS)

	go func() {
		// QPS calc
		for {
			currentOps := ops;
			time.Sleep(time.Second)
			var qps = (int32)(ops - currentOps);
			log.Printf("QPS: %d\n", qps)
		}
	}()

	go func() {
		for {
			for i := 0; i < *maxQPS; i++ {
				select {
				case bucket <- true:
				default:
				}
			}
			time.Sleep(time.Second)
		}
	}()

	for i := 0; i < *concurrency; i++ {
		go func() {
			for {
				<- bucket
				sendRequest(*req)
			}
		}()
	}
	// make it never end
	<- fin
}

// slow connection
type slowConn struct {
	net.Conn
	sr slowReader
}

func slowDial(network, addr string) (net.Conn, error) {
	conn, err := (&net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 3000 * time.Second,
	}).Dial(network, addr)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	//conn.(*net.TCPConn).SetLinger(0)
	return slowConn{conn, slowReader{conn}}, nil
}

func (conn slowConn) Read(data []byte) (int, error) {
	return conn.sr.Read(data)
}

// slow reader
type slowReader struct{ r io.Reader }

func (r slowReader) Read(data []byte) (int, error) {
	// wait for x ms before reading a single byte.
	duration := time.Duration(int64(*slowRead)) * time.Millisecond
	time.Sleep(duration)
	n, err := r.r.Read(data[:1])
	return n, err
}

//Ref https://github.com/jiaz/simpleloadclient
//https://stackoverflow.com/questions/3757289/tcp-option-so-linger-zero-when-its-required