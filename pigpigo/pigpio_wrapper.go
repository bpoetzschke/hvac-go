package pigpigo

/*
#include "pigpio.h"
#cgo LDFLAGS: -lpigpio
*/
import "C"
import (
	"fmt"

	"github.com/bpoetzschke/hvac-go/log"
)

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
	goRes := uint32(res)
	log.Debugf("GpioCfgGetInternals: %d", goRes)

	return goRes
}

func (wrapper piGpioWrapper) GpioCfgSetInternals(cfg uint32) int {
	res := C.gpioCfgSetInternals(C.uint32_t(cfg))
	goRes := int(res)
	fmt.Printf("FOO: %d\n", goRes)
	log.Debugf("GpioCfgSetInternals: %d", goRes)

	return goRes
}

func (wrapper piGpioWrapper) GpioInitialise() int {
	res := C.gpioInitialise()
	goRes := int(res)
	log.Debugf("GpioInitialise: %d", goRes)

	return goRes
}

func (wrapper piGpioWrapper) GpioSetMode(gpio uint, mode uint) int {
	res := C.gpioSetMode(C.unsigned(gpio), C.unsigned(mode))
	goRes := int(res)
	log.Debugf("GpioSetMode: %d", goRes)

	return goRes
}

func (wrapper piGpioWrapper) GpioTerminate() {
	C.gpioTerminate()
}
