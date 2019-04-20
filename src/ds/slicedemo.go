package main

import "log"

func main()  {
	arr:=[]int{1,2,3,4,5,6}
	s1:=arr[0:0]
	log.Printf("len=%d, cap=%d, %v \n",len(s1), cap(s1), s1)
}
