package main

import (
	"fmt"
	"time"
)

func funIf() {
	/* 定义局部变量 */
	var a int = 10

	/* 使用 if 语句判断布尔表达式 */
	if a < 20 {
		/* 如果条件为 true 则执行以下语句 */
		fmt.Printf("a 小于 20\n")
	}
	fmt.Printf("a 的值为 : %d\n", a)
}

func funSelect() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}

func funSelect2() {
	// 定义两个通道
	ch1 := make(chan string)
	ch2 := make(chan string)

	// 启动两个 goroutine，分别从两个通道中获取数据
	go func() {
		for {
			ch1 <- "from 1"
		}
	}()
	go func() {
		for {
			ch2 <- "from 2"
		}
	}()

	// 使用 select 语句非阻塞地从两个通道中获取数据
	for {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		default:
			// 如果两个通道都没有可用的数据，则执行这里的语句
			fmt.Println("no message received")
		}
	}
}
