package main

import "fmt"

func hello() {
	fmt.Print("hello")
}

func exclamation() {
	fmt.Print("!\n")
}

func space() {
	fmt.Print(" ")
}

func world() {
	fmt.Print("world")
}

func main() {
	hello()
	defer exclamation()
	defer world()
	space()
}
