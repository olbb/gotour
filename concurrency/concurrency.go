package concurrency

import (
	"fmt"
	"time"

	"golang.org/x/tour/tree"
)

//DoTest concurrency
func DoTest() {
	treeExercies()
}

func treeExercies() {
	t1, t2 := tree.New(1), tree.New(1)
	fmt.Println(Same(t1, t2))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}

// Walk 步进 tree t 将所有的值从 tree 发送到 channel ch。
func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	fmt.Println(t, t.Value)
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}

}

func walkTree(t *tree.Tree, ch chan int) {
	Walk(t, ch)
	close(ch)
}

// Same 检测树 t1 和 t2 是否含有相同的值。
func Same(t1, t2 *tree.Tree) bool {
	// fmt.Println(t1, t2)
	ch1, ch2 := make(chan int), make(chan int)
	go walkTree(t1, ch1)
	go walkTree(t2, ch2)
	for i := range ch1 {
		j := <-ch2
		fmt.Println(i, j)
		if i != j {
			return false
		}
	}
	return true
}

func selectTest2() {

	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick")
		case <-boom:
			fmt.Println("Boom!")
			return
		default:
			fmt.Println("...")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

//select 语句使一个 Go 程可以等待多个通信操作。
//select 会阻塞到某个分支可以继续执行为止，这时就会执行该分支。当多个分支都准备好时会随机选择一个执行。
func selectTest() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci2(c, quit)
}

func fibonacci2(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y

		case <-quit:
			fmt.Println("quit.")
			return
		}
	}

}

func rangeAndClose() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}

// 循环 for i := range c 会不断从信道接收值，直到它被关闭。
// 注意： 只有发送者才能关闭信道，而接收者不能。向一个已经关闭的信道发送数据会引发程序恐慌（panic）。
// 还要注意： 信道与文件不同，通常情况下无需关闭它们。只有在必须告诉接收者不再有值需要发送的时候才有必要关闭，例如终止一个 range 循环。
func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
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
