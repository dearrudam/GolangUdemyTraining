package main

import "fmt"

func half(n int) (float64, bool) {
	return float64(n) / 2, n%2 == 0
}

func main() {
	fmt.Println(h, even)
	h, even := half(5)
}
