package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// baseSelect()
	gracefulShutdown()
}

func baseSelect() {
	// bufferChan := make(chan string, 2) //2
	// bufferChan <- "first"

	// select {
	// case bufferChan <- "second":
	// 	fmt.Println("write", <-bufferChan, <-bufferChan)
	// case str := <-bufferChan:
	// 	fmt.Println("read", str)
	// }

	// // fmt.Printf("Len: %v Cap: %v\n", len(bufferChan), cap(bufferChan))

	// unbufChan := make(chan int)

	// go func() {
	// 	time.Sleep(time.Second)
	// 	unbufChan <- 1
	// }()

	// select {
	// // case bufferChan <- "third":
	// // 	fmt.Println("Unblocking writing")
	// case val := <-unbufChan:
	// 	fmt.Println("blocking reading", val)
	// case <-time.After(time.Millisecond * 1500):
	// 	fmt.Println("time's up")
	// 	// default:
	// 	// 	fmt.Println("default case")
	// }

	resultChan := make(chan int)
	timer := time.After(time.Second) // time outside loop

	go func() {
		defer close(resultChan)

		for i := 0; i < 1000; i++ {
			select {
			case lol := <-timer:
				fmt.Println("time's up", lol)
				return
			default:
				time.Sleep(time.Nanosecond)
				resultChan <- i
			}
		}
	}()

	for v := range resultChan {
		fmt.Println(v)
	}
}

func gracefulShutdown() {
	sigChan := make(chan os.Signal, 1)

	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	timer := time.After(10 * time.Second)

	select {
	case <-timer:
		fmt.Println("time's up")
		return
	case sig := <-sigChan:
		fmt.Println("Stoped by signal", sig)
		return
	}
}
