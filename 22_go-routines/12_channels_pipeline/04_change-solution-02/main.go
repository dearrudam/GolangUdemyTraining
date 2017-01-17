package main

import (
	"fmt"
)

func main() {
	in := gen(10)
	for f := range fatorial(in) {
		fmt.Println(f)
	}
}

func gen(n int) <-chan int {
	out := make(chan int)
	go func() {
		for i := 1; i <= n; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func fatorial(in <-chan int) <-chan string {
	out := make(chan string)
	go func() {
		for n := range in {
			for f := range fat(n) {
				out <- fmt.Sprint("fatorial for ", n, " is ", f)
			}
		}
		close(out)
	}()
	return out
}

func fat(n int) <-chan int {
	out := make(chan int)
	go func() {
		total := 1
		for i := n; i > 0; i-- {
			total *= i
		}
		out <- total
		close(out)
	}()
	return out
}
