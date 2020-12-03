package irwave

import "github.com/bpoetzschke/hvac-go/pigpigo"

func NewWave() IRWave {
	return &wave{}
}

type wave struct {
	pulses []pigpigo.Pulse
}

func (w *wave) Zero(gpio uint32, duration uint32) {

}

func (w *wave) One(gpio uint32, duration uint32) {

}
