package main
//
//import (
//	"time"
//	"fmt"
//	"net/http"
//	log "github.com/Sirupsen/logrus"
//	"os"
//	"sync"
//	"net"
//	"io"
//	"flag"
//)
//
//var (
//	httpClient *http.Client
//	dialTimeout = 4 * time.Second
//	requestTimeout = 2 * time.Minute
//	idleConnectionCount = 40000
//	endpoint = "10.47.11.0"
//
//	slowRead *int
//	requestCount *int
//	apiPath *string
//)
//
//func init() {
//	httpClient = initHttpClient(dialTimeout, requestTimeout, idleConnectionCount)
//}
//
//func main() {
//	slowRead = flag.Int("byteReadTime", 50, "an int")
//	requestCount = flag.Int("requestCount", 2000, "an int")
//	apiPath = flag.String("apiPath", "/fast", " a string")
//
//	flag.Parse()
//	fmt.Println("using values ", *slowRead, *requestCount, *apiPath)
//	createLog()
//	startLoad()
//}
//
//func startLoad() {
//
//	url := fmt.Sprintf("http://%v%v", endpoint, *apiPath)
//
//	for {
//		wg := &sync.WaitGroup{}
//		//log.Info("Sending new batch of requests to ", url)
//		//for i:=0; i< *requestCount; i++ {
//		//	wg.Add(1)
//		//	go func(i int) {
//		//		defer wg.Done()
//		//		_, err := httpClient.Get(url)
//		//		if err != nil {
//		//			log.Error("Errored request ", i, " with ", err.Error())
//		//			return
//		//		}
//		//		//resp.Body.Close()
//		//		log.Info("Finished GET request ", i)
//		//	}(i)
//		//}
//
//		for j:=0; j < 1000; j++ {
//			for i:=0; i< *requestCount; i++ {
//				wg.Add(1)
//				go func(i int) {
//					defer wg.Done()
//					_, err := httpClient.Get(url)
//					if err != nil {
//						log.Error("Errored request ", i, " with ", err.Error())
//						return
//					}
//					//resp.Body.Close()
//					log.Info("Finished GET request ", i)
//				}(i)
//			}
//			time.Sleep(1 * time.Second)
//		}
//		wg.Wait()
//	}
//}
//
//func createLog(){
//	if f, err := os.OpenFile("/tmp/run_load.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666); err == nil {
//		log.SetOutput(f)
//	} else {
//		fmt.Printf("Failed to open log file: %s",  err.Error())
//		os.Exit(1)
//	}
//}
//
//func initHttpClient(dialTimeout, requestTimeout time.Duration, idleConnsPerHost int) *http.Client {
//	transport := &http.Transport{
//		//Dial: (&net.Dialer{
//		//	Timeout: dialTimeout,
//		//}).Dial,
//		Dial: slowDial,
//		DisableCompression:  true,
//		DisableKeepAlives:   true,
//		//MaxIdleConns: 0,
//		IdleConnTimeout: 90 * time.Second,
//		MaxIdleConnsPerHost: idleConnsPerHost,
//	}
//
//	return &http.Client{
//		Timeout:   requestTimeout,
//		Transport: transport,
//	}
//}
//
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
//	// wait for x ms before reading a single byte.
//	duration := time.Duration(int64(*slowRead)) * time.Millisecond
//	time.Sleep(duration)
//	n, err := r.r.Read(data[:1])
//	if n > 0 {
//		//fmt.Printf("%s", data[:1])
//	}
//	return n, err
//}
