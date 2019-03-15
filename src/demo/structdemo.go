package main

import "log"

type user struct {
	name  string
	age   int
	phone []string
	data  map[string]string
}

func main() {
	user := user{name: "gx", age: 20, phone: []string{"1890000000", "18899999999"}, data: map[string]string{"aa": "asdfasdf", "bb": "aaaa"}}
	log.Printf("user:%+v \n", user)
	log.Printf("user.data.aa:%s \n", user.data["aa"])

}
