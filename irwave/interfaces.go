package irwave

import "github.com/bpoetzschke/hvac-go/pigpigo"

type IRWave interface {
	One(gpio uint32, duration uint32)
	Zero(gpio uint32, duration uint32)

	GetPulses() []pigpigo.Pulse
}
