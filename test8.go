package main

import "fmt"

func funAddress() {
	var a int = 10

	fmt.Printf("变量的地址: %x\n", &a)
}

// 要取指针的值要用*p，如 b = *p，直接输出p是地址
func funPoint() {
	var a int = 20
	var p *int
	p = &a
	fmt.Printf("a 变量的地址是: %x\n", &a)
	/* 指针变量的存储地址 */
	fmt.Printf("p 变量储存的指针地址: %x\n", p)
	/* 使用指针访问值 */
	fmt.Printf("*p 变量的值: %d\n", *p)
	var b int = *p
	fmt.Println(b)
}

func funPointEmpty() {
	var ptr *int
	fmt.Printf("ptr 为 : %x\n", ptr)
}
