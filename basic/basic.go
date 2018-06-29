package basic

import (
	"fmt"
	"math/rand"
)

//DoTest ...
func DoTest() {
	b1()
}

func b1() {
	fmt.Println("My favorite number is", rand.Intn(100))
}
