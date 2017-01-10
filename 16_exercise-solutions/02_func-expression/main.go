package main

import "fmt"

func main() {
	half := func(n int) (float64, bool) {
		return float64(n) / 2, n%2 == 0
	}
	h, even := half(5)
	fmt.Println(h, even)
}
