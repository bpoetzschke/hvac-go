package pigpio

/*
#include "pigpio.h"
#cgo LDFLAGS: -lpigpio
*/
import "C"

func NewPiGpio(gpioPin uint) PiGpio {
	pg := piGpio{
		gpioPin: gpioPin,
	}
	pg.Setup()

	return pg
}

type piGpio struct {
	gpioPin uint
}

func (pg *piGpio) Setup() {
	pg.disableSigHandler()
}

func (pg *piGpio) disableSigHandler() {
	cfg := C.gpioCfgGetInternals()
	cfg |= C.PI_CFG_NOSIGHANDLER
	C.gpioCfgSetInternals(cfg)
}

//import "C"
//import (
//	"fmt"
//	"os"
//)
//
//func main() {
//	cfg := C.gpioCfgGetInternals()
//	fmt.Printf("Config: %d\n", cfg)
//	cfg |= C.PI_CFG_NOSIGHANDLER // (1<<10)
//	fmt.Printf("New config: %d\n", cfg)
//	C.gpioCfgSetInternals(cfg)
//
//	initRes := C.gpioInitialise()
//	if initRes == C.PI_INIT_FAILED {
//		fmt.Printf("Init failed: %d\n", initRes)
//		os.Exit(1)
//	}
//
//	fmt.Printf("Init status: %d\n", initRes)
//	//res = C.gpioSetMode(C.uint(23),C.uint(1))
//	//fmt.Printf("set mode status: %d\n", res)
//	//clear := C.gpioWaveClear()
//	//fmt.Printf("Clear: %d\n", clear)
//	//if clear != 0 {
//	//	C.gpioTerminate()
//	//	os.Exit(1)
//	//}
//
//	C.gpioTerminate()
//	//
//	//foo := C.gpioPulse_t{}
//	//fmt.Printf("%#v\n", foo)
//	//
//	//fmt.Printf("%#v\n", C.PI_INIT_FAILED)
//}
