package irwave

import (
	"math"

	"github.com/bpoetzschke/hvac-go/log"
	"github.com/bpoetzschke/hvac-go/pigpigo"
)

func NewWave(frequency uint32, dutyCycle float64) IRWave {
	return &wave{
		frequency: frequency,
		dutyCycle: dutyCycle,
	}
}

type wave struct {
	frequency uint32
	dutyCycle float64

	pulses []pigpigo.Pulse
}

func (w *wave) GetPulses() []pigpigo.Pulse {
	return w.pulses
}

func (w *wave) Zero(gpioPin uint32, duration uint32) {
	log.Debugf("SPACE\t%s", duration)
	w.addPulse(0, 1<<gpioPin, duration)
}

func (w *wave) One(gpioPin uint32, duration uint32) {
	log.Debugf("MARK\t%s", duration)
	periodTime := 1_000_000.0 / float64(w.frequency)
	onDuration := int(math.Round(periodTime * w.dutyCycle))
	offDuration := int(math.Round(periodTime * (1.0 - w.dutyCycle)))
	totalPeriods := int(math.Round(float64(duration) / periodTime))
	totalPulses := totalPeriods * 2

	for i := 0; i < totalPulses; i++ {
		if i%2 == 0 {
			w.addPulse(1<<gpioPin, 0, uint32(onDuration))
		} else {
			w.addPulse(0, 1<<gpioPin, uint32(offDuration))
		}
	}
}

func (w *wave) addPulse(gpioOn uint32, gpioOff uint32, usDelay uint32) {
	w.pulses = append(w.pulses, pigpigo.Pulse{
		GpioOn:  gpioOn,
		GpioOff: gpioOff,
		UsDelay: usDelay,
	})
}
