package main

import (
	"time"
	"fmt"
	"net/http"
	log "github.com/Sirupsen/logrus"
	"os"
	"sync"
	"net"
	"io"
	"flag"
	"crypto/tls"
)

var (
	slowRead *int
	requestCount, sleepBW *int
	apiPath *string
)

func main() {
	slowRead = flag.Int("byteReadTime", 0, "wait for x ms before reading a single byte")
	requestCount = flag.Int("requestCount", 100, "Concurrent requests")
	apiPath = flag.String("apiPath", "https://10.0.0.1/large", " Complete endpoint ")
	sleepBW = flag.Int("sleepInBw", 10, "Sleep before next request batch of concurrent requests")

	flag.Parse()
	fmt.Printf("Requests per second: %d\n", *requestCount)
	fmt.Printf("Sleep before next request batch: %d%s\n", *sleepBW, "ms")
	fmt.Printf("Endpoint: %v\n", *apiPath)
	//fmt.Printf("Byte read wait(slow read) : %d%s\n", *slowRead, "ms")

	createLog()
	startLoad()
}

func startLoad() {

	wg := &sync.WaitGroup{}
	for {
		log.Info("Sending new batch of requests to ", *apiPath)
		for i:=0; i< *requestCount; i++ {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				httpClient := initHttpClient()
				_, err := httpClient.Get(*apiPath)
				/* NOT closing client + reading response body */
				if err != nil {
					log.Error("Errored request ", i, " with ", err.Error())
					return
				}
			}(i)
			time.Sleep(time.Duration(int64(*sleepBW)) * time.Millisecond)
		}
	}
	wg.Wait()
}

func createLog(){
	if f, err := os.OpenFile("/tmp/run_load.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666); err == nil {
		log.SetOutput(f)
	} else {
		fmt.Printf("Failed to open log file: %s",  err.Error())
		os.Exit(1)
	}
}

func initHttpClient() *http.Client {
	transport := &http.Transport{
		Dial: (&net.Dialer{
		 KeepAlive: 300 * time.Second,
		 Timeout: 30 * time.Second,
		}).Dial,
		//Dial: slowDial,
		DisableCompression:  true,
		DisableKeepAlives:   false,
		IdleConnTimeout:  30 * time.Second,
		MaxIdleConnsPerHost: 1,
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},	//Ignore ssl errors - Https VIP with no domain mapping
	}
	return &http.Client{
		Timeout:   1 * time.Minute,	//Request timeout
		Transport: transport,
	}
}

/* slow connection */
type slowConn struct {
	net.Conn
	sr slowReader
}

func slowDial(network, addr string) (net.Conn, error) {
	conn, err := (&net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 300 * time.Second,
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

/* slow reader */
type slowReader struct{ r io.Reader }

/* sleep for x ms before reading a single byte. */
func (r slowReader) Read(data []byte) (int, error) {
	duration := time.Duration(int64(*slowRead)) * time.Millisecond
	time.Sleep(duration)
	return r.r.Read(data[:1])
}
