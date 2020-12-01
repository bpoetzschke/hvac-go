package pigpigo

/*
#include "pigpio.h"
#cgo LDFLAGS: -lpigpio
*/
import "C"

const (
	piCfgNoSigHandler = 1 << 10
)

type PiGpioWrapper interface {
	GpioCfgGetInternals() uint32
	GpioCfgSetInternals(uint322 uint32) int
	GpioInitialise() int
	GpioSetMode(gpio uint, mode uint) int
	GpioTerminate()
}

func NewPiGpioWrapper() PiGpioWrapper {
	return &piGpioWrapper{}
}

type piGpioWrapper struct {
}

func (wrapper piGpioWrapper) GpioCfgGetInternals() uint32 {
	return C.gpioCfgGetInternals()
}

func (wrapper piGpioWrapper) GpioCfgSetInternals(cfg uint32) int {
	return C.gpioCfgSetInternals(cfg)
}

func (wrapper piGpioWrapper) GpioInitialise() int {
	return C.gpioInitialise()
}

func (wrapper piGpioWrapper) GpioSetMode(gpio uint, mode uint) int {
	return C.gpioSetMode(gpio, mode)
}

func (wrapper piGpioWrapper) GpioTerminate() {
	C.gpioTerminate()
}
