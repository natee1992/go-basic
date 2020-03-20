package basic

import (
	"fmt"
	"math"
)

var (
	aa = 3
	bb = 4
	cc = 5
)

// 定义初始值
func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInitValue() {
	var a, b = 3, 4
	var s = "abc"
	fmt.Println(a, b, s)
}

//强制类型转换
func triangle() {
	var a, b = 3, 4
	fmt.Println(calcTriangle(a, b))
}

func calcTriangle(a, b int) int {
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	return c
}

//常量
func consts() {
	const filename = "abc.txt"
	const a, b = 3, 4
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

//枚举
func enums() {
	const (
		cpp = iota
		_
		python
		golang
		javascript
	)
	// b,kb,mb,gb,tb,pb
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(cpp, javascript, python, golang)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	fmt.Println("Hello world")
	variableZeroValue()
	variableInitValue()
	fmt.Println(aa, bb, cc)
	triangle()
	consts()
	enums()
}
