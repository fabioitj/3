package main

import (
	"fmt"
	"time"
)

func pingpong(ping chan<- bool, pong <-chan bool) {
	for {
		<-pong
		fmt.Println("ping")
		time.Sleep(time.Second)
		ping <- true
	}
}

func main() {
	pingChan := make(chan bool)
	pongChan := make(chan bool)

	go pingpong(pingChan, pongChan)
	pongChan <- true

	for {
		<-pingChan
		fmt.Println("pong")
		time.Sleep(time.Second)
		pongChan <- true
	}
}