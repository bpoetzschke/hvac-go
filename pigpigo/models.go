package pigpigo

type Pulse struct {
	GpioOn  uint32
	GpioOff uint32
	UsDelay uint32
}

type GpioPin uint32
