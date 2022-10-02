package main

import (
	"fmt"
	"github.com/jasonlvhit/gocron"
)

func SayHello() {
	fmt.Println("Hello World")
}

func main() {
	s := gocron.NewScheduler()
	s.Every(1).Second().Do(SayHello)
	<-s.Start()
}
