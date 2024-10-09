package main

import (
	"fmt"
	"time"
)

func ping(b chan string, r int) {

	if r < 10 {
		b <- "ping"
	} else {
		close(b)
	}
}
func pong(b chan string, r int) {

	if r < 10 {
		b <- "pong"
	} else {
		close(b)
	}
}

func main() {
	x := make(chan string)

	for i := 0; i < 10; i++ {

		if i%2 == 0 {
			go ping(x, i)
		} else {
			go pong(x, i)
		}

		time.Sleep(time.Millisecond * 500)
		p := <-x
		fmt.Println(p)
	}
}
