package main

import (
	"fmt"
)

func greatestOf(l ...int) int {
	var g int
	for i, n := range l {
		if n > g || i == 0 {
			g = n
		}
	}
	return g
}

func main() {
	g := greatestOf(5, 8, 2, 7, 1, 6, 9, 7)
	fmt.Println(g)

}
