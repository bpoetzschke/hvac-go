// +build !pigpio

package pigpigo

import "C"
import (
	"math"

	"github.com/bpoetzschke/hvac-go/log"
)

func NewPiGpioWrapper() PiGpioWrapper {
	return &pigpioWrapperFake{}
}

type pigpioWrapperFake struct {
}

func (_ pigpioWrapperFake) GpioCfgGetInternals() uint32 {
	log.Debug("GpioCfgGetInternals: fake - 0")

	return 0
}

func (_ pigpioWrapperFake) GpioCfgSetInternals(cfg uint32) int {
	log.Debugf("GpioCfgSetInternals(cfg=%d): fake - 0", cfg)

	return 0
}

func (_ pigpioWrapperFake) GpioInitialise() int {
	log.Debug("GpioInitialise: fake - 1")

	return 1
}

func (_ pigpioWrapperFake) GpioSetMode(gpioPin uint32, mode uint32) int {
	log.Debugf("GpioSetMode(gpioPin=%d, mode-%d): fake - 0", gpioPin, mode)

	return 0
}

func (_ pigpioWrapperFake) GpioTerminate() {
	log.Debug("GpioTerminate: fake")
}

func (_ pigpioWrapperFake) GpioWaveGetMaxPulses() int {
	log.Debugf("GpioWaveGetMaxPulses: fake - %d", math.MaxInt32)

	return math.MaxInt32
}
