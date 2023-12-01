package main

func funfunc() {
	println(funMax(1, 2))
}

func funMax(num1, num2 int) int {
	var res int
	if num1 > num2 {
		res = num1
	} else {
		res = num2
	}
	return res
}
