package main

import (
	"log"
	"time"
)

func main() {
	tc := time.After(time.Second * 3)
	log.Printf("tc type=%T \n", tc)
	log.Printf("tc=%s \n", <-tc)
}
