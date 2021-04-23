package main

import (
	"fmt"
	"runtime"
)

// 包内部变量
// := 只能用在函数内部
var aa = 3

var (
	bb = 4
	ss = "kkk"
)

func variableInitValues() {
	var a, b int = 6, 8
	var s string = "abc"
	fmt.Println(a, s, b)
}

func variableInitValuesTypeDeduction() {
	var a, b, c = 6, 8, true
	var s = "abc"
	fmt.Println(a, s, b, c)
}

func variableInitValuesTypeDeduction2() {
	a, b, c := 6, 8, true
	var s = "abc"
	fmt.Println(a, s, b, c)
}

func main() {
	fmt.Println("hello, there")
	fmt.Println(runtime.GOARCH)

	variableInitValuesTypeDeduction2()
}
