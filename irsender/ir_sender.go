package irsender

/*
#include "pigpio.h"
#cgo LDFLAGS: -lpigpio
*/
import "C"

import (
	"github.com/bpoetzschke/hvac-go/log"
)

func NewIrSender(gpioPin uint) IRSender {
	ir := irSender{
		gpioPin: gpioPin,
	}

	return &ir
}

type irSender struct {
	gpioPin uint
}

func (ir *irSender) Setup() error {
	ir.disableSigHandler()
	return ir.initGpio()
}

func (ir *irSender) Shutdown() {
	ir.terminate()
}

func (ir *irSender) disableSigHandler() {
	cfg := C.gpioCfgGetInternals()
	cfg |= C.PI_CFG_NOSIGHANDLER
	C.gpioCfgSetInternals(cfg)
}

func (ir *irSender) initGpio() error {
	res := C.gpioInitialise()
	if res < 0 {
		log.Errorf("Failed to initialise gpio, return value: %d", res)
		return ErrorInitFailed
	}

	res = C.gpioSetMode(C.uint(ir.gpioPin), C.uint(1))
	if res != 0 {
		log.Errorf("Failed to set gpio mode, return value: %d", res)
		return ErrorSetGpioFailed
	}

	return nil
}

func (ir *irSender) terminate() {
	C.gpioTerminate()
}
