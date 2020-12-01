package main

import (
	"fmt"
	"os"
	"time"

	"github.com/bpoetzschke/hvac-go/pigpigo"
)

func main() {
	gpiGo := pigpigo.NewPiGpiGo(23)
	err := gpiGo.Setup()
	if err != nil {
		fmt.Printf("Failed to setup gpio. %s", err)
		os.Exit(1)
	}
	defer gpiGo.Shutdown()
	<-time.After(10 * time.Second)
}
