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
	err := gpiGo.disableSigHandler()
	if err != nil {
		return err
	}
	return gpiGo.initGpio()
}

func (gpiGo *piGpiGo) Shutdown() {
	gpiGo.terminate()
}

func (gpiGo *piGpiGo) disableSigHandler() error {
	log.Debug("Disabling pigpio signal handler")
	cfg, err := gpiGo.wrapper.GpioCfgGetInternals()
	if err != nil {
		log.Errorf("Failed to get gpio config internals. %s", err)
		return err
	}
	cfg |= piCfgNoSigHandler
	_, err = gpiGo.wrapper.GpioCfgSetInternals(cfg)
	if err != nil {
		log.Errorf("Failed to set gpio config internals. %s", err)
		return err
	}

	return nil
}

func (gpiGo *piGpiGo) initGpio() error {
	res, err := gpiGo.wrapper.GpioInitialise()
	if err != nil {
		log.Errorf("Failed to initialise gpio. %s", err)
		return err
	}
	if res < 0 {
		log.Errorf("Failed to initialise gpio, return value: %d", res)
		return ErrorInitFailed
	}

	res, err = gpiGo.wrapper.GpioSetMode(23, 1)
	if err != nil {
		log.Errorf("Failed to set gpio mode. %s", err)
		return err
	}
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
