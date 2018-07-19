package more

import (
	"fmt"
	"strings"

	"golang.org/x/tour/pic"
	"golang.org/x/tour/wc"
)

//DoTest DoTest
func DoTest() {
	froballTest()
}

func fibonacci() func() int {
	x, y := 0, 1
	return func() int {
		x, y = y, x+y
		return y
	}
}

func froballTest() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

//WordCount 11
func WordCount(input string) map[string]int {
	aMap := make(map[string]int)
	// l := strings.Split(input, " ")
	l := strings.Fields(input)
	for _, v := range l {
		aMap[v]++

	}
	return aMap
}

func mapTest2() {
	wc.Test(WordCount)
}

var m map[string]Vertex

func mapTest() {
	m = make(map[string](Vertex))
	m["test"] = Vertex{11, 12}
	// fmt.Println(m["test"])
	delete(m, "new")
	v, ok := m["test"]
	fmt.Println(ok)
	fmt.Println(v)
}

func getpic(x, y int) [][]uint8 {
	result := make([][]uint8, y)
	for i := 0; i < y; i++ {
		result[i] = make([]uint8, x)
	}
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			switch {
			case j%15 == 0:
				result[i][j] = 240
			case j%3 == 0:
				result[i][j] = 120
			case j%5 == 0:
				result[i][j] = 150
			default:
				result[i][j] = 100
			}
		}
	}
	return result
}

func showPic() {
	pic.Show(getpic)
}

func rangeSlice2() {
	for _, v := range pow {
		fmt.Println("range,", v, " ")
	}
}

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func rangeSlice() {
	for i, v := range pow {
		fmt.Println("range,", i, v)
	}
}

/**
*切片的长度就是它所包含的元素个数。
*切片的容量是从它的第一个元素开始数，到其底层数组元素末尾的个数。
 */
func sliceTest1() {
	primes := []int{2, 3, 5, 7, 11, 13}
	printSliceLength(primes[1:3])
	printSliceLength(primes[:3])
	printSliceLength(primes[1:])
}

func printSliceLength(a []int) {
	fmt.Println("length is:", len(a), "  cap is:", cap(a))
}

func sliceDefault() {
	primes := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes[1:3])
	fmt.Println(primes[:3])
	fmt.Println(primes[1:])
}

func arraysTest2() {
	primes := []int{2, 3, 5, 7, 11, 13}
	s := primes[1:4]
	s[2] = 999
	fmt.Println(s)
	fmt.Println(primes)
}

func arraysTest1() {
	var a [2]string
	a[0] = "hello"
	a[1] = "world"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{1, 3, 5, 7, 9}
	fmt.Println(primes)

}

var (
	v1 = Vertex{11, 22}
	v2 = Vertex{x: 15}
	v3 = Vertex{}
	p  = &Vertex{33, 44}
)

func structTest4() {
	fmt.Println(v1, v2, v3, p)
}

func structTest3() {
	v := Vertex{11, 22}
	p := &v
	p.x = 12345
	fmt.Println(v)
}

func structTest2() {
	v := Vertex{10, 15}
	v.x = 11
	fmt.Println(v.x)
}

func structTest() {
	fmt.Println(Vertex{50, 100})
}

//Vertex Vertex
type Vertex struct {
	x, y int
}

func pointers() {
	i, j := 520, 2701
	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)
	p = &j
	*p = *p / 37
	fmt.Println(j)
}
