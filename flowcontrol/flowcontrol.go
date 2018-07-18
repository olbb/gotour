package flowcontrol

import (
	"fmt"
	"runtime"
	"time"
)

//DoTest ...
func DoTest() {
	// fmt.Println("DoTest")
	switchTest2()
}

func switchTest2() {
	today := time.Now().Weekday()
	switch time.Friday {
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
