package main

import "fmt"

func myfun(a int) int {
	fmt.Println(a)
	return a + 1
}

func main() {
	r := myfun(2)
	fmt.Println(r)
}
