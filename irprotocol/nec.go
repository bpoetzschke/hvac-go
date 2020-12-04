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

var defaultConfig = NECConfig{
	LeadingPulse: NECPulseConfig{
		PulseDuration: leadingPulseDuration,
		GapDuration:   leadingGapDuration,
	},
	OnePulse: NECPulseConfig{
		PulseDuration: onePulseDuration,
		GapDuration:   oneGapDuration,
	},
	ZeroPulse: NECPulseConfig{
		PulseDuration: zeroPulseDuration,
		GapDuration:   zeroGapDuration,
	},
	TrailingPulse: NECPulseConfig{
		PulseDuration: trailingPulseDuration,
		GapDuration:   trailingGapDuration,
	},
}

type NECPulseConfig struct {
	PulseDuration uint32
	GapDuration   uint32
}

type NECConfig struct {
	LeadingPulse  NECPulseConfig
	OnePulse      NECPulseConfig
	ZeroPulse     NECPulseConfig
	TrailingPulse NECPulseConfig
}

func NewNECIRProtocol(config *NECConfig) IRProtocol {
	cfgToApply := defaultConfig
	if config != nil {
		cfgToApply = *config
	}
	return &necIRProtocol{
		frequency: frequency,
		dutyCycle: dutyCycle,
		config:    cfgToApply,
	}
}

type necIRProtocol struct {
	frequency uint32
	dutyCycle float64
	config    NECConfig
}

func (nec necIRProtocol) ProcessCode(gpioPin uint32, irCodes []IRCode) error {
	wave := irwave.NewWave(nec.frequency, nec.dutyCycle)
	if nec.config.LeadingPulse.PulseDuration > 0 || nec.config.LeadingPulse.GapDuration > 0 {
		nec.addAGCPulse(gpioPin, wave)
	}
	for _, irCode := range irCodes {
		if irCode == IRCodeZero {
			nec.zero(gpioPin, wave)
		} else if irCode == IRCodeOne {
			nec.one(gpioPin, wave)
		} else {
			log.Errorf("Unknown ir code: %s", irCode)
			return ErrUnknownIRCode
		}
	}
	if nec.config.TrailingPulse.PulseDuration > 0 {
		nec.addTrailingPulse(gpioPin, wave)
	}
	return nil
}

func (nec necIRProtocol) addAGCPulse(gpioPin uint32, wave irwave.IRWave) {
	log.Debug("NEC: Add AGC burst pulse")
	wave.One(gpioPin, nec.config.LeadingPulse.PulseDuration)
	wave.Zero(gpioPin, nec.config.LeadingPulse.GapDuration)
}

func (nec necIRProtocol) addTrailingPulse(gpioPin uint32, wave irwave.IRWave) {
	log.Debug("NEC: Add trailing pulse")
	wave.One(gpioPin, nec.config.TrailingPulse.PulseDuration)
	if nec.config.TrailingPulse.GapDuration > 0 {
		wave.Zero(gpioPin, nec.config.TrailingPulse.GapDuration)
	}
}

func (nec necIRProtocol) zero(gpioPin uint32, wave irwave.IRWave) {
	log.Debugf("NEC: Zero")
	wave.One(gpioPin, nec.config.ZeroPulse.PulseDuration)
	wave.Zero(gpioPin, nec.config.ZeroPulse.GapDuration)
}

func (nec necIRProtocol) one(gpioPin uint32, wave irwave.IRWave) {
	log.Debugf("NEC: One")
	wave.One(gpioPin, nec.config.OnePulse.PulseDuration)
	wave.Zero(gpioPin, nec.config.OnePulse.GapDuration)
}
