package main

import "fmt"

type Fn func(x, y int) int

func (fn Fn) Chain(f Fn) Fn {
	return func(x, y int) int {
		fmt.Println(fn(x, y))
		return f(x, y)
	}
}

func add(x, y int) int {
	fmt.Printf("%d + %d = ", x, y)
	return x + y
}
func minus(x, y int) int {
	fmt.Printf("%d - %d = ", x, y)
	return x - y
}
func mul(x, y int) int {
	fmt.Printf("%d * %d = ", x, y)
	return x * y
}

func main() {
	var result = Fn(add).Chain(Fn(minus)).Chain(Fn(mul))(3, 5)
	fmt.Println(result)
}
