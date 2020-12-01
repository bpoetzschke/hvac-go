package main

import (
	"time"

	"github.com/bpoetzschke/hvac-go/pigpigo"
)

func main() {
	gpiGo := pigpigo.NewPiGpiGo(23)
	gpiGo.Setup()
	defer gpiGo.Shutdown()
	<-time.After(10 * time.Second)
}
