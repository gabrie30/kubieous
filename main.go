package main

import (
	"fmt"
	"time"

	"github.com/gabrie30/kubieous/checks"
)

func main() {

	fmt.Println("starting kubieous...")
	fmt.Println("")

	// Montoring Loop
	for {

		checks.Pods()
		checks.Nodes()

		time.Sleep(5 * time.Second)
	}
}
