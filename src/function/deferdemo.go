package main

import (
	"fmt"
)

//defer是在return之前执行的
//函数返回的过程是这样的：先给返回值赋值，然后调用defer表达式，最后才是返回到调用函数中。
func main() {
	r := 1
	fmt.Println(f(r)) //0
	fmt.Println(r)    //1

}

func f(r int) int {
	defer func() {
		r++
	}()
	return 0
}
