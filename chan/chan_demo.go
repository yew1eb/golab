package main

import "fmt"

type service struct {
	Ch chan int
}

func main() {
	ch1 := make(chan int, 10)

	serv := service{ch1}
	ch1 <- 1
	ch1 <- 2

	fmt.Println("ch1 len:", len(ch1))
	fmt.Println("service len:", len(serv.Ch))

	<-serv.Ch

	fmt.Println("ch1 len:", len(ch1))
	fmt.Println("service len:", len(serv.Ch))
}
