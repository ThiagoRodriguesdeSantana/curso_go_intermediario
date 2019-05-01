package main

import (
	"fmt"
	"time"
)

var irc = make(chan string)
var sms = make(chan string)

func ponger(canal chan string) {
	for {
		canal <- "pong"
	}
}

func eai(canal chan string) {
	for {
		canal <- "e ai"
	}
}
func blz(canal chan string) {
	for {
		canal <- "blz"
	}
}
func pinger(canal chan string) {
	for {
		canal <- "ping"
	}
}
func impressora() {
	var msg string
	for {
		select {
		case msg = <-irc:
			fmt.Println(msg)
		case msg = <-sms:
			fmt.Println("zzz..zzz", msg)
		}
		time.Sleep(time.Second * 1)

	}
}

func main() {

	go pinger(irc)
	go ponger(irc)
	go eai(sms)
	go blz(sms)
	go impressora()

	var entrada string
	fmt.Scanln(&entrada)
}