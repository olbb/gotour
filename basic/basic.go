package basic

import (
	"fmt"
	"math"
	"math/rand"
)

//DoTest ...
func DoTest() {
	b9()
}

var iInt, jInt = 1, 2

func b9() {
	j := 5
	var c, java, python = true, false, "test"
	fmt.Println(iInt, jInt, c, java, python, j)
}

var c, python, java bool

func b8() {
	var aInt int
	fmt.Println(aInt, c, python, java)
}

func b7() {
	fmt.Println(b7test(9))
}

func b7test(input int) (x, y int) {
	x = input / 3
	y = input * 199
	return
}

func b6() {
	fmt.Println(swap("hello", "world"))
}

func swap(x, y string) (string, string) {
	return y, x
}

func b4() {
	fmt.Println(add(15, 20))
}

func add(x, y int) int {

	return x + y
}

func b3() {
	fmt.Println(math.Pi)
}

func b2() {
	fmt.Printf("Now you havd %g problems.", math.Sqrt(7))
}

func b1() {
	fmt.Println("My favorite number is", rand.Intn(100))
}
