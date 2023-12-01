package main

import "fmt"

func funFor() {
	sum := 0
	for i := 0; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func funFor2() {
	strs := []string{"google", "runoob"}
	for i, s := range strs {
		fmt.Println(i, s)
	}

	nums := []int{1, 2, 3, 5}
	for i, x := range nums {
		fmt.Printf("第 %d 位是 %d \n", i, x)
	}
}

func funcFor3() {
	map1 := make(map[int]float32)
	map1[1] = 1.0
	map1[2] = 2.0
	map1[3] = 3.0
	map1[4] = 4.0

	for key, value := range map1 {
		fmt.Printf("keys is %d value is %f \n", key, value)
	}
}
