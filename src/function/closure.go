package main

import "fmt"

func main()  {
	for i:=0;i<3 ;i++  {
		//闭包对捕获的外部变量是以引用的方式访问
		//这种方式会造成一些隐含的问题，使用时需要注意
		defer func(i int) {
			fmt.Println(i)
		}(i)
	}
}

func say(str string){
	fmt.Println(str)
}

