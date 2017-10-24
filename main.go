package main

import (
	"brochat/lib"
	"flag"
	"os"
)

func main() {
	var isHost bool

	flag.BoolVar(&isHost, "listen", false, "We're gonna listen bro")
	flag.Parse()

	if isHost {
		connIP := os.Args[2]
		lib.RunHost(connIP)
	} else {
		connIP := os.Args[1]
		lib.RunGuest(connIP)
	}
}
