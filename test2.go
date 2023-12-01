package main

import "fmt"

// 语言变量
func funVar() {
	var str string = "testString"
	fmt.Println(str)

	var aInt, bInt int = 1, 2
	fmt.Println(aInt, bInt)
}

func funVar2() {
	var x, y int
	var ( // 这种因式分解关键字的写法一般用于声明全局变量
		a int
		b bool
	)

	var c, d int = 1, 2
	var e, f = 123, "hello"

	g, h := 123, "hello"
	fmt.Println(x, y, a, b, c, d, e, f, g, h)
}
