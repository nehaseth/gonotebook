package main

import (
	"context"
	"fmt"
	"time"
	
)

//func main() {
//
//	// We'll iterate over 2 values in the `queue` channel.
//	queue := make(chan string, 2)
//	queue <- "one"
//	queue <- "two"
//
//	var wg = &sync.WaitGroup{}
//	wg.Add(1)
//	go func() {
//		defer wg.Done()
//		for elem := range queue {
//			fmt.Println(elem)
//		}
//		fmt.Println("exiting go routine")
//	}()
//	fmt.Println("waiting1")
//	queue <- "three"
//	queue <- "four"
//	time.Sleep(5 * time.Second)
//	queue <- "five"
//	queue <- "five"
//	queue <- "five"
//	queue <- "five"
//	close(queue)
//	fmt.Println("closed now")
//	wg.Wait()
//}


func op(ctx context.Context, cancelFunc context.CancelFunc) {
	childContext, childCancel := context.WithCancel(ctx)
	defer childCancel()
	cancelFunc()
	select {
	case <-time.After(500 * time.Millisecond):
		fmt.Println(" op done ")
	case <-childContext.Done():
		fmt.Println("halted op")
	}
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	
	_, childCanc := context.WithCancel(ctx)
	
	fmt.Println("cancelling parent canc")
	cancel()
	time.Sleep(1 * time.Second)
	
	fmt.Println("cancelling child canc")
	childCanc()
}