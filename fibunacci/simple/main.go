package main

import "fmt"

const N = 10

func main() {

	// with loop
	f1, f2 := 0, 1
	for i := 0; i < N; i++ {
		fmt.Printf("%v ", f1)
		sum := f1 + f2
		f1 = f2
		f2 = sum
	}

	fmt.Println()

	// recursive
	for i := 0; i < N; i++ {
		fmt.Printf("%v ", fib(i))
	}

}

func fib(f int) int {
	if f == 0 || f == 1 {
		return f
	}
	return fib(f-2) + fib(f-1)
}
