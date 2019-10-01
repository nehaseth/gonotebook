package main

import (
	"time"
	"fmt"
	"net/http"
	"io/ioutil"

	//"os"
	//"strings"
)

func main() {
	go funcCall2()
	time.Sleep(10 *time.Millisecond)
	//f, _ := os.Open("")

	funcCall(nil)
	//funcCall(&http.Request{})
	//funcCall(&http.Request{Body: ioutil.NopCloser(strings.NewReader(""))})
	time.Sleep(1* time.Second)
	fmt.Println("testing sleep on panic")
}

func funcCall(r *http.Request) {
	defer func() {
		time.Sleep(1 * time.Millisecond)
		ioutil.ReadAll(r.Body)
		r.Body.Close()
	}()
	panic("TEST")
}

func funcCall2() {
	ticker := time.NewTicker(1 * time.Millisecond)
	for {
		select {
		case <-ticker.C :
			fmt.Println("testing goroutine on panic")
		}
	}
}