package pigpigo

/*
#include "pigpio.h"
#cgo LDFLAGS: -lpigpio
*/
import "C"
import (
	"github.com/bpoetzschke/hvac-go/log"
)

const (
	piCfgNoSigHandler = 1 << 10
)

func GpioCfgGetInternals() uint32 {
	res := C.gpioCfgGetInternals()
	goRes := uint32(res)
	log.Debugf("GpioCfgGetInternals: %d", goRes)

	return goRes
}

func GpioCfgSetInternals(cfg uint32) int {
	res := C.gpioCfgSetInternals(C.uint32_t(cfg))
	goRes := int(res)
	log.Debugf("GpioCfgSetInternals: %d", goRes)

	return goRes
}

func GpioInitialise() int {
	res := C.gpioInitialise()
	goRes := int(res)
	log.Debugf("GpioInitialise: %d", goRes)

	return goRes
}

func GpioSetMode(gpio uint, mode uint) int {
	res := C.gpioSetMode(C.unsigned(gpio), C.unsigned(mode))
	goRes := int(res)
	log.Debugf("GpioSetMode: %d", goRes)

	return goRes
}

func GpioTerminate() {
	C.gpioTerminate()
}

func GpioWaveGetMaxPulses() int {
	res := C.gpioWaveGetMaxPulses()
	goRes := int(res)
	log.Debugf("GpioWaveGetMaxPulses: %d", goRes)

	return goRes
}
