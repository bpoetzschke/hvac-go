package pigpigo

import (
	"github.com/bpoetzschke/hvac-go/log"
)

func NewPiGpiGo(gpioPin GpioPin) PiGpiGo {
	gpiGo := piGpiGo{
		gpioPin:       gpioPin,
		pigpioWrapper: NewPiGpioWrapper(),
	}

	return &gpiGo
}

type piGpiGo struct {
	gpioPin       GpioPin
	pigpioWrapper PiGpioWrapper
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
	cfg := gpiGo.pigpioWrapper.GpioCfgGetInternals()
	cfg |= piCfgNoSigHandler
	gpiGo.pigpioWrapper.GpioCfgSetInternals(cfg)
}

func (gpiGo *piGpiGo) initGpio() error {
	res := gpiGo.pigpioWrapper.GpioInitialise()
	if res < 0 {
		log.Errorf("Failed to initialise gpio, return value: %d", res)
		return ErrorInitFailed
	}

	res = gpiGo.pigpioWrapper.GpioSetMode(uint32(gpiGo.gpioPin), 1)
	if res != 0 {
		log.Errorf("Failed to set gpio mode, return value: %d", res)
		return ErrorSetGpioFailed
	}

	return nil
}

func (gpiGo *piGpiGo) terminate() {
	log.Debug("Shutdown pigpio")
	gpiGo.pigpioWrapper.GpioTerminate()
}
