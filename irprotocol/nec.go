package irprotocol

import (
	"github.com/bpoetzschke/hvac-go/irwave"
	"github.com/bpoetzschke/hvac-go/log"
)

const (
	frequency             = 38_000
	dutyCycle             = 0.33
	leadingPulseDuration  = 9_000
	leadingGapDuration    = 4_500
	onePulseDuration      = 562
	oneGapDuration        = 1_686
	zeroPulseDuration     = 562
	zeroGapDuration       = 562
	trailingPulseDuration = 562
	trailingGapDuration   = 0
)

func NewNECIRProtocol() IRProtocol {
	return &necIRProtocol{
		frequency:             frequency,
		dutyCycle:             dutyCycle,
		leadingPulseDuration:  leadingPulseDuration,
		leadingGapDuration:    leadingGapDuration,
		onePulseDuration:      onePulseDuration,
		oneGapDuration:        oneGapDuration,
		zeroPulseDuration:     zeroPulseDuration,
		zeroGapDuration:       zeroGapDuration,
		trailingPulseDuration: trailingPulseDuration,
		trailingGapDuration:   trailingGapDuration,
	}
}

type necIRProtocol struct {
	frequency             uint32
	dutyCycle             float32
	leadingPulseDuration  uint32
	leadingGapDuration    uint32
	onePulseDuration      uint32
	oneGapDuration        uint32
	zeroPulseDuration     uint32
	zeroGapDuration       uint32
	trailingPulseDuration uint32
	trailingGapDuration   uint32
}

func (nec necIRProtocol) ProcessCode(gpioPin uint32, irCodes []IRCode) {
	wave := irwave.NewWave()
	// TODO: THis check can be done in the AGC method itself
	if nec.leadingPulseDuration > 0 || nec.leadingGapDuration > 0 {
		nec.addAGCPulse(gpioPin, wave)
	}
	for _, irCode := range irCodes {
		if irCode == IRCodeZero {
			nec.zero(gpioPin, wave)
		} else if irCode == IRCodeOne {
			nec.one(gpioPin, wave)
		} else {
			log.Warnf("Unknown ir code: %s", irCode)
		}
	}
	// TODO: do check in trailing pulse method
	if nec.trailingPulseDuration > 0 {
		nec.addTrailingPulse(gpioPin, wave)
	}
}

func (nec necIRProtocol) addAGCPulse(gpioPin uint32, wave irwave.IRWave) {
	log.Debug("NEC: Add AGC burst pulse")
	wave.One(gpioPin, nec.leadingPulseDuration)
	wave.Zero(gpioPin, nec.leadingGapDuration)
}

func (nec necIRProtocol) addTrailingPulse(gpioPin uint32, wave irwave.IRWave) {
	log.Debug("NEC: Add trailing pulse")
	wave.One(gpioPin, nec.trailingPulseDuration)
	if nec.trailingGapDuration > 0 {
		wave.Zero(gpioPin, nec.trailingGapDuration)
	}
}

func (nec necIRProtocol) zero(gpioPin uint32, wave irwave.IRWave) {
	log.Debugf("NEC: Zero")
	wave.One(gpioPin, nec.zeroPulseDuration)
	wave.Zero(gpioPin, nec.zeroGapDuration)
}

func (nec necIRProtocol) one(gpioPin uint32, wave irwave.IRWave) {
	log.Debugf("NEC: One")
	wave.One(gpioPin, nec.onePulseDuration)
	wave.Zero(gpioPin, nec.oneGapDuration)
}
