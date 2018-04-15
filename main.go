package main

import (
	"fmt"
	"time"

	"github.com/gabrie30/kubieous/checks"
)

func main() {

	fmt.Println("starting kubieous...")

	// How often should the checks be performed
	podCheckTimer := time.NewTicker(99999 * time.Second)
	nodeCheckTimer := time.NewTicker(99999 * time.Second)

	// checks.PodEventStream()
	go checks.HPA()

	// Montoring Loop
	for {
		select {
		case _ = <-podCheckTimer.C:
			go checks.Pods()
		case _ = <-nodeCheckTimer.C:
			go checks.Nodes()
		}
	}

	fmt.Println("stopping kubieous")
}
