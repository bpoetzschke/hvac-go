package hvac

import (
	"math"

	"github.com/bpoetzschke/hvac-go/ir"
	"github.com/bpoetzschke/hvac-go/irprotocol"
	"github.com/bpoetzschke/hvac-go/pigpigo"
)

type Mitsubishi interface {
	Off()
}

func NewMitsubishi(gpioPin pigpigo.GpioPin) Mitsubishi {
	return &mitsubishi{
		irProtocol: irprotocol.NewNECIRProtocol(&irprotocol.NECConfig{
			LeadingPulse: irprotocol.NECPulseConfig{
				PulseDuration: 3400,
				GapDuration:   1750,
			},
			OnePulse:      irprotocol.NECPulseConfig{},
			ZeroPulse:     irprotocol.NECPulseConfig{},
			TrailingPulse: irprotocol.NECPulseConfig{},
		}),
	}
}

type mitsubishi struct {
	irProtocol         irprotocol.Protocol
	irSender           ir.Sender
	powerMode          PowerMode
	climateMode        ClimateMode
	iSeeMode           ISeeMode
	powerfulMode       PowerfulMode
	vaneHorizontalMode VaneHorizontalMode
	fanMode            FanMode
	vaneVerticalMode   VaneVerticalMode
	timeControlMode    TimeControlMode
	areaMode           AreaMode
	temperature        Temperature
}

func (m *mitsubishi) Off() {
}

func (m *mitsubishi) sendCommand() {

}

func (m *mitsubishi) getAdjustedTemperature() Temperature {
	return Temperature(
		math.Max(
			float64(temperatureMin),
			math.Min(
				float64(temperatureMax),
				float64(m.temperature),
			),
		) - 16,
	)
}

func (m *mitsubishi) assembleCommand() []uint32 {
	data := []uint32{
		// Fixed header frame
		0x23, 0xCB, 0x26, 0x01, 0x00,
		uint32(m.powerMode),
		uint32(m.climateMode) | uint32(m.iSeeMode),
		uint32(m.getAdjustedTemperature()),
		m.climateMode.toClimateMode2() | uint32(m.vaneHorizontalMode),
	}
}
