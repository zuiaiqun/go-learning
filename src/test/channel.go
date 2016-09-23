package test

import (
	"fmt"
	"time"
)

func count(ch chan int) {
	ch <- 1
	fmt.Println("Counting")
}

func TestChannel() {
	chs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go count(chs[i])
	}
	for _, ch := range chs {
		<-ch
	}
	time.Sleep(time.Second)
}
