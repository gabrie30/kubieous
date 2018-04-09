package main

import (
	"fmt"
	"time"

	"github.com/gabrie30/kubieous/checks"
)

func main() {

	fmt.Println("starting kubieous...")

	// How often should the checks be performed
	podCheckTimer := time.NewTicker(5 * time.Second)
	nodeCheckTimer := time.NewTicker(10 * time.Second)

	// Montoring Loop
	for {
		select {
		case _ = <-podCheckTimer.C:
			go checks.Pods()
		case _ = <-nodeCheckTimer.C:
			go checks.Nodes()
		}
	}
}
