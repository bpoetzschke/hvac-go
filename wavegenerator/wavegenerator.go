package wavegenerator

import "github.com/bpoetzschke/hvac-go/pigpigo"

// TODO: Refactor; maybe create and return a wave instead of wavegenerator
func NewWaveGenerator() WaveGenerator {
	wg := waveGenerator{
		addedPulses: 0,
	}
	wg.pulses = make([]pigpigo.Pulse, pigpigo.GpioWaveGetMaxPulses())

	return &wg
}

type waveGenerator struct {
	pulses      []pigpigo.Pulse
	addedPulses int
}

func (wg *waveGenerator) Zero(duration uint32) {

}

func (wg *waveGenerator) One(duration uint32) {

}
