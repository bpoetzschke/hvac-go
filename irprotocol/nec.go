package irprotocol

import (
	"github.com/bpoetzschke/hvac-go/log"
	"github.com/bpoetzschke/hvac-go/wavegenerator"
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

		waveGenerator: wavegenerator.NewWaveGenerator(),
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

	waveGenerator wavegenerator.WaveGenerator
}

func (nec necIRProtocol) Zero() {
	log.Debugf("NEC: Zero")
	nec.waveGenerator.One(nec.zeroPulseDuration)
	nec.waveGenerator.Zero(nec.zeroGapDuration)
}

func (nec necIRProtocol) One() {
	log.Debugf("NEC: One")
	nec.waveGenerator.One(nec.onePulseDuration)
	nec.waveGenerator.Zero(nec.oneGapDuration)
}

func (nec necIRProtocol) ProcessCode(irCodes []IRCode) {
	for _, irCode := range irCodes {
		if irCode == IRCodeZero {

		} else if irCode == IRCodeOne {

		} else {
			log.Warnf("Unknown ir code: %s", irCode)
		}
	}
}
