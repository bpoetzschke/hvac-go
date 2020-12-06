package pigpigo

const (
	piCfgNoSigHandler = 1 << 10
)

type PiGpioWrapper interface {
	GpioCfgGetInternals() uint32
	GpioCfgSetInternals(cfg uint32) int
	GpioInitialise() int
	GpioSetMode(gpioPin uint32, mode uint32) int
	GpioTerminate()
	GpioWaveGetMaxPulses() int
}
