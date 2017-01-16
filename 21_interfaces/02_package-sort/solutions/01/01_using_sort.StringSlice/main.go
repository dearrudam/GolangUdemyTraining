package main

import (
	"fmt"
	"sort"
)

type people []string


func main() {
	studyGroup := people{"Zeno", "John", "Al", "Jenny"}
	fmt.Println(studyGroup)
	sort.Sort(sort.StringSlice(studyGroup))
	fmt.Println(studyGroup)
}

