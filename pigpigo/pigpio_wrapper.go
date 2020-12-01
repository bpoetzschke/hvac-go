package pigpigo

/*
#include "pigpio.h"
#cgo LDFLAGS: -lpigpio
*/
import "C"
import "fmt"

const (
	piCfgNoSigHandler = 1 << 10
)

var (
	ErrorFailedTypeConversion = fmt.Errorf("failed to convert c type to go type")
)

type PiGpioWrapper interface {
	GpioCfgGetInternals() uint32
	GpioCfgSetInternals(cfg uint32) int
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
	res := C.gpioCfgGetInternals()

	return uint32(res)
}

func (wrapper piGpioWrapper) GpioCfgSetInternals(cfg uint32) int {
	res := C.gpioCfgSetInternals(C.uint32_t(cfg))

	return int(res)
}

func (wrapper piGpioWrapper) GpioInitialise() int {
	res := C.gpioInitialise()

	return int(res)
}

func (wrapper piGpioWrapper) GpioSetMode(gpio uint, mode uint) int {
	res := C.gpioSetMode(C.unsigned(gpio), C.unsigned(mode))

	return int(res)
}

func (wrapper piGpioWrapper) GpioTerminate() {
	C.gpioTerminate()
}
