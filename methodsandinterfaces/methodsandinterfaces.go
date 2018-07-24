package methodsandinterfaces

import (
	"fmt"
	"io"
	"math"
	"os"
	"strings"
	"time"

	"github.com/Go-zh/tour/reader"
)

//DoTest 111
func DoTest() {
	rot13Exercies()
}

type rot13Reader struct {
	r io.Reader
}

func rot13Exercies() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

func (r rot13Reader) Read(b []byte) (int, error) {
	l, e := r.r.Read(b)
	if e != nil {
		return l, e
	}
	for i, v := range b {
		switch {
		case v < 65:
		case v < 78: //A-M
			b[i] = v + 13
		case v < 91: //N-Z
			b[i] = v - 13
		case v < 97:
		case v < 110: //a-m
			b[i] = v + 13
		case v < 123: //n-z
			b[i] = v - 13
		}
	}
	return l, e
}

func readerExercies() {
	r := myReader{}
	reader.Validate(r)
}

func (r myReader) Read(b []byte) (int, error) {
	c := byte('A')
	for i := range b {
		// v = c
		b[i] = c
	}
	return len(b), nil
}

type myReader struct {
}

func readerTest() {
	r := strings.NewReader("Hello, World!")
	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v, err = %v, b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		// fmt.Println("z is:", z)
		if z*z == x {
			return z, nil
		}

	}
	return z, nil
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number:%v", float64(e))
}

func errorTest2() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

func (e myError) Error() string {
	return fmt.Sprintf("at %v, %v", e.time, e.msg)
}

func errorTest() {
	err := run()
	fmt.Println(err)
}

func run() myError {
	return myError{"can't accesss.", time.Now()}
}

type myError struct {
	msg  string
	time time.Time
}

func (add IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", add[0], add[1], add[2], add[3])
}

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.

func stringerTest2() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

func stringerTest() {
	j := person{"Jack", 18}
	r := person{"Rose", 16}
	fmt.Println(j, r)
}

//实现了fmt中Stringer接口
func (p person) String() string {
	return fmt.Sprintf("%v (%v Years old)", p.name, p.old)
}

type person struct {
	name string
	old  int
}

func interfaceTest5() {
	do("hello")
	do(52)
	do(true)
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v.\n", v, v*2)
	case string:
		fmt.Printf("string bytes is:%v\n", len(v))
	default:
		fmt.Printf("I don't known about the type ! %T\n", v)
	}
}

func interfaceTest4() {
	var i interface{} = "hello"
	s := i.(string)
	fmt.Println(s)
	s, ok := i.(string)
	fmt.Println(ok, s)
	f, ok := i.(float64)
	fmt.Println(ok, f)
	// f = i.(float64)
	// fmt.Println(s)
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
