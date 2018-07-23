package flowcontrol

import (
	"fmt"
	"runtime"
	"time"
)

//DoTest ...
func DoTest() {
	// fmt.Println("DoTest")
	sqrtTest()
}

//defer后进先出
func defer2() {
	fmt.Println("counting")
	for i := 0; i < 20; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}

func defer1() {
	defer fmt.Println("World!")
	fmt.Println("Hello")
}

//z -= (z*z - x) / (2*z)
func sqrt(f float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - f) / (2 * z)
		// fmt.Println("z is:", z)
		if z*z == f {
			return z
		}

	}
	return z
}

func sqrtTest() {
	// c := complex(1, 2)
	// fmt.Println(c)
	fmt.Println("result is :", sqrt(100))
}

func switchTest3() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning")
	case t.Hour() < 18:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good night")
	}
}

func switchTest2() {
	today := time.Now().Weekday()
	switch time.Saturday {
	case today:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("Tommorrow")
	case today + 2:
		fmt.Println("Two days later")
	default:
		fmt.Println("Too far away")
	}
}

func switchTest() {

	var result string
	switch os := runtime.GOOS; os {
	case "linux":
		result = "Linux"
	default:
		result = "Unknow"
	}
	fmt.Println("OS is:", result)
}

func forLoop() {
	sum := 0
	for i := 0; i < 100; i++ {
		sum += i
	}
	fmt.Println("Result is:", sum)
}

func forWhile() {
	sum := 0
	for sum < 100 {
		sum++
	}
	fmt.Println("Result is:", sum)
}

func ifTest() {
	fmt.Println("if Test:", ifFunc(25))
}

func ifFunc(a int) (x int) {
	if newA := a * 666; newA%5 == 0 {
		x = newA
		return
	}
	x = -1
	return
}
