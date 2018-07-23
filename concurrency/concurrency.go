package concurrency

import (
	"fmt"
	"time"
)

//DoTest concurrency
func DoTest() {
	bufferedChannels()
}

func bufferedChannels() {
	c := make(chan int, 2)
	c <- 2
	c <- 1
	fmt.Println(<-c)
	fmt.Println(<-c)
}

func channels() {
	ints := []int{7, 2, 8, -9, 4, 0}
	channel := make(chan int)
	go sum(ints[:len(ints)/2], channel)
	go sum(ints[len(ints)/2:], channel)
	x, y := <-channel, <-channel
	fmt.Println(x, y, x+y)

}

//默认情况下，发送和接收操作在另一端准备好之前都会阻塞。这使得 Go 程可以在没有显式的锁或竞态变量的情况下进行同步。
func sum(ints []int, c chan int) {
	sum := 0
	for _, v := range ints {
		sum += v
		fmt.Printf("sum += %v\n", v)
	}
	c <- sum
}

func goroutine() {
	go say("world")
	say("hello")
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}
