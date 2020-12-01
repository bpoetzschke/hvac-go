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
	GpioCfgSetInternals(cfg uint32) (int, error)
	GpioInitialise() (int, error)
	GpioSetMode(gpio uint, mode uint) (int, error)
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

func (wrapper piGpioWrapper) GpioCfgSetInternals(cfg uint32) (int, error) {
	res := C.gpioCfgSetInternals(C.uint32_t(cfg))
	goRes, ok := res.(int)
	if ok {
		return goRes, nil
	}

	return -1, ErrorFailedTypeConversion
}

func (wrapper piGpioWrapper) GpioInitialise() (int, error) {
	res := C.gpioInitialise()
	goRes, ok := res.(int)
	if ok {
		return goRes, nil
	}

	return -1, ErrorFailedTypeConversion
}

func (wrapper piGpioWrapper) GpioSetMode(gpio uint, mode uint) (int, error) {
	res := C.gpioSetMode(C.unsigned(gpio), C.unsigned(mode))
	goRes, ok := res.(int)
	if ok {
		return goRes, nil
	}

	return -1, ErrorFailedTypeConversion
}

func (wrapper piGpioWrapper) GpioTerminate() {
	C.gpioTerminate()
}
