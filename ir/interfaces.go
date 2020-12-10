package ir

import "github.com/bpoetzschke/hvac-go/pigpigo"

type Wave interface {
	One(gpio uint32, duration uint32)
	Zero(gpio uint32, duration uint32)

	GetPulses() []pigpigo.Pulse
}

type Sender interface {
	Transmit(w Wave)
}
