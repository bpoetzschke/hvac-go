// +build !pigpio

package pigpigo

import "C"
import (
	"math"

	"github.com/bpoetzschke/hvac-go/log"
)

func GpioCfgGetInternals() uint32 {
	log.Debug("GpioCfgGetInternals: fake - 0")

	return 0
}

func GpioCfgSetInternals(_ uint32) int {
	log.Debug("GpioCfgSetInternals: fake - 0")

	return 0
}

func GpioInitialise() int {
	log.Debug("GpioInitialise: fake - 1")

	return 1
}

func GpioSetMode(_ uint, _ uint) int {
	log.Debugf("GpioSetMode: fake - 0")

	return 0
}

func GpioTerminate() {
	log.Debug("GpioTerminate: fake")
}

func GpioWaveGetMaxPulses() int {
	log.Debugf("GpioWaveGetMaxPulses: fake - %d", math.MaxInt32)

	return math.MaxInt32
}
