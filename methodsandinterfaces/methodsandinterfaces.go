package methodsandinterfaces

import (
	"fmt"
	"math"
)

//DoTest 111
func DoTest() {
	interfaceTest3()
}

func interfaceTest3() {
	var i interface{}
	describe(i)
	i = 52
	describe(i)
	i = "Hello"
	describe(i)
}

func interfaceTest2() {
	var shower iShow
	str := aStr{"测试数据"}
	shower = &str
	shower.show()
	describe(shower)
	shower = aFloat{1990}
	shower.show()
	describe(shower)

	var shower2 iShow
	var str2 *aStr
	shower2 = str2
	shower2.show()
	describe(shower2)
}

func describe(i interface{}) {
	fmt.Printf("value is:%v, type is:%T\n", i, i)
}

type iShow interface {
	show()
}

type aStr struct {
	s string
}

func (str *aStr) show() {
	if str == nil {
		fmt.Println("<wtf nil>")
		return
	}
	fmt.Println(str.s)
}

type aFloat struct {
	f float64
}

func (f aFloat) show() {
	fmt.Println(f.f)
}

func interfaceTest() {
	var abser IAbs
	v := Vertex{3, 4}
	f := MyFloat(-15)
	abser = &v
	fmt.Println(abser.NewAbs())
	abser = f

	fmt.Println(abser.NewAbs())

}

//NewAbs 1
func (f MyFloat) NewAbs() float64 {
	if f > 0 {
		return float64(f)
	}
	return float64(-f)
}

//NewAbs 1
func (v *Vertex) NewAbs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//IAbs 1
type IAbs interface {
	NewAbs() float64
	// Test()
}

func testMethods3() {
	v := Vertex{3, 4}
	v.scale(5)
	fmt.Println(v)
	scale(&v, 5)
	fmt.Println(v)
}

func scale(v *Vertex, scale float64) {
	v.X = v.X * scale
	v.Y = v.Y * scale
}

//传指针才能改变该接收者的值
func (v *Vertex) scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func testMethods2() {
	f := MyFloat(-1)
	fmt.Println(f.abs())
}

//MyFloat 1
type MyFloat float64

func (f MyFloat) abs() float64 {
	if f > 0 {
		return float64(f)
	}
	return float64(-f)

}

//方法只是个带接收者参数的函数。
func testMethods() {
	fmt.Println(Vertex{15, 20}.Abs())
	fmt.Println(Abs(Vertex{15, 20}))
}

//Abs 111
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//Abs 11
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Vertex 111
//Go 没有类。不过你可以为结构体类型定义方法。
//方法就是一类带特殊的 接收者 参数的函数。
type Vertex struct {
	X, Y float64
}
