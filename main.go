package main

import (
	"fmt"
	"log"

	"github.com/numericals/DFS/p2p"
)

func main() {
	tr := p2p.NewTCPTransport(":3000")

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Hello world")

	select {}
}

// Test-NetConnection localhost -Port 3000
