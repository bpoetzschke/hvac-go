package hvac

type PowerMode uint32
type ClimateMode uint32
type ISeeMode uint32
type PowerfulMode uint32
type VaneHorizontalMode uint32
type FanMode uint32
type VaneVerticalMode uint32
type TimeControlMode uint32
type AreaMode uint32
type Temperature uint32

func (cm ClimateMode) toClimateMode2() uint32 {
	if cm == ClimateModeAuto {
		return 0x00
	} else if cm == ClimateModeHot {
		return 0x00
	} else if cm == ClimateModeCold {
		return 0x06
	} else if cm == ClimateModeDry {
		return 0x02
	}

	return 0x00
}
