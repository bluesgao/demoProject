package main

import "log"

func main() {
	var dict = map[string]string{
		"123": "abc",
		"234": "dbc",
	}

	for k, v := range dict {
		log.Printf("key:%s,value:%s \n", k, v)
	}

	value := dict["1111"]
	log.Printf(value)

}
