package pigpigo

import (
	"github.com/bpoetzschke/hvac-go/log"
)

func NewPiGpiGo(gpioPin uint) PiGpiGo {
	gpiGo := piGpiGo{
		gpioPin: gpioPin,
		wrapper: NewPiGpioWrapper(),
	}

	return &gpiGo
}

type piGpiGo struct {
	gpioPin uint
	wrapper PiGpioWrapper
}

func (gpiGo *piGpiGo) Setup() error {
	gpiGo.disableSigHandler()
	return gpiGo.initGpio()
}

func (gpiGo *piGpiGo) Shutdown() {
	gpiGo.terminate()
}

func (gpiGo *piGpiGo) disableSigHandler() {
	log.Debug("Disabling pigpio signal handler")
	cfg := gpiGo.wrapper.GpioCfgGetInternals()
	cfg |= piCfgNoSigHandler
	gpiGo.wrapper.GpioCfgSetInternals(cfg)
}

func (gpiGo *piGpiGo) initGpio() error {
	res := gpiGo.wrapper.GpioInitialise()
	if res < 0 {
		log.Errorf("Failed to initialise gpio, return value: %d", res)
		return ErrorInitFailed
	}

	res = gpiGo.wrapper.GpioSetMode(23, 1)
	if res != 0 {
		log.Errorf("Failed to set gpio mode, return value: %d", res)
		return ErrorSetGpioFailed
	}

	return nil
}

func (gpiGo *piGpiGo) terminate() {
	log.Debug("Shutdown pigpio")
	gpiGo.wrapper.GpioTerminate()
}
