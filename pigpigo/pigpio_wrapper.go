// +build pigpio

package pigpigo

/*
#include "pigpio.h"
#cgo LDFLAGS: -lpigpio
*/
import "C"
import (
	"github.com/bpoetzschke/hvac-go/log"
)

func NewPiGpioWrapper() PiGpioWrapper {
	return &pigpioWrapper{}
}

type pigpioWrapper struct {
}

func (_ pigpioWrapper) GpioCfgGetInternals() uint32 {
	res := C.gpioCfgGetInternals()
	goRes := uint32(res)
	log.Debugf("GpioCfgGetInternals: %d", goRes)

	return goRes
}

func (_ pigpioWrapper) GpioCfgSetInternals(cfg uint32) int {
	res := C.gpioCfgSetInternals(C.uint32_t(cfg))
	goRes := int(res)
	log.Debugf("GpioCfgSetInternals(cfg=%d): %d", cfg, goRes)

	return goRes
}

func (_ pigpioWrapper) GpioInitialise() int {
	res := C.gpioInitialise()
	goRes := int(res)
	log.Debugf("GpioInitialise: %d", goRes)

	return goRes
}

func (_ pigpioWrapper) GpioSetMode(gpioPin uint32, mode uint32) int {
	res := C.gpioSetMode(C.unsigned(gpio), C.unsigned(mode))
	goRes := int(res)
	log.Debugf("GpioSetMode(gpio=%d, mode=%d): %d", gpioPin, mode, goRes)

	return goRes
}

func (_ pigpioWrapper) GpioTerminate() {
	log.Debug("GpioTerminate")
	C.gpioTerminate()
}

func (_ pigpioWrapper) GpioWaveGetMaxPulses() int {
	res := C.gpioWaveGetMaxPulses()
	goRes := int(res)
	log.Debugf("GpioWaveGetMaxPulses: %d", goRes)

	return goRes
}
