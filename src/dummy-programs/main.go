package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: ip-addr\n")
	}

	name := os.Args[1]
	addr := net.ParseIP(name)
	if addr == nil {
		fmt.Println("invalid address")
	} else {
		fmt.Println("The adress is ", addr.String())
	}
}
