package hvac

const (
	PowerOff = PowerMode(0b00000000) // 0x00      0000 0000        0
	PowerOn  = PowerMode(0b00100000) // 0x20      0010 0000       32

	ClimateModeHot  = ClimateMode(0b00000000) // 0x00      0000 0000        0
	ClimateModeCold = ClimateMode(0b00000110) // 0x06      0000 0110        6
	ClimateModeDry  = ClimateMode(0b00000010) // 0x02      0000 0010        2
	ClimateModeAuto = ClimateMode(0b00000000) // 0x00      0000 0000        0

	ISeeOff = ISeeMode(0b00000000) // 0x00      0000 0000        0
	ISeeOn  = ISeeMode(0b01000000) // 0x40      0100 0000       64

	PowerfulOff = PowerfulMode(0b00000000) // 0x00      0000 0000        0
	PowerfulOn  = PowerfulMode(0b00001000) // 0x08      0000 1000        8

	VaneHorizontalNotSet      = VaneHorizontalMode(0b00000000) // 0x00      0000 0000        0
	VaneHorizontalLeft        = VaneHorizontalMode(0b00010000) // 0x10      0001 0000       16
	VaneHorizontalMiddleLeft  = VaneHorizontalMode(0b00100000) // 0x20      0010 0000       32
	VaneHorizontalMiddle      = VaneHorizontalMode(0b00110000) // 0x30      0011 0000       48
	VaneHorizontalMiddleRight = VaneHorizontalMode(0b01000000) // 0x40      0100 0000       64
	VaneHorizontalRight       = VaneHorizontalMode(0b01010000) // 0x50      0101 0000       80
	VaneHorizontalSwing       = VaneHorizontalMode(0b11000000) // 0xC0      1100 0000      192

	FanModeSpeed1 = FanMode(0b00000001) // 0x01      0000 0001        1
	FanModeSpeed2 = FanMode(0b00000010) // 0x02      0000 0010        2
	FanModeSpeed3 = FanMode(0b00000011) // 0x03      0000 0011        3
	FanModeAuto   = FanMode(0b10000000) // 0x80      1000 0000      128

	VaneVerticalAuto         = VaneVerticalMode(0b01000000) // 0x40      0100 0000       64
	VaneVerticalTop          = VaneVerticalMode(0b01001000) // 0x48      0100 1000       72
	VaneVerticalMiddleTop    = VaneVerticalMode(0b01010000) // 0x50      0101 0000       80
	VaneVerticalMiddle       = VaneVerticalMode(0b01011000) // 0x58      0101 1000       88
	VaneVerticalMiddleBottom = VaneVerticalMode(0b01100000) // 0x60      0110 0000       96
	VaneVerticalBottom       = VaneVerticalMode(0b01101000) // 0x68      0110 1000      104
	VaneVerticalSwing        = VaneVerticalMode(0b01111000) // 0x78      0111 1000      120

	TimeControlNoTimeControl = TimeControlMode(0b00000000) // 0x00      0000 0000        0
	TimeControlControlStart  = TimeControlMode(0b00000101) // 0x05      0000 0101        5
	TimeControlControlEnd    = TimeControlMode(0b00000011) // 0x03      0000 0011        3
	TimeControlControlBoth   = TimeControlMode(0b00000111) // 0x07      0000 0111        7

	AreaModeNotSet = AreaMode(0b00000000) // 0x00      0000 0000        0
	AreaModeLeft   = AreaMode(0b01000000) // 0x40      0100 0000       64
	AreaModeRight  = AreaMode(0b11000000) // 0xC0      1100 0000      192
	AreaModeFull   = AreaMode(0b10000000) // 0x80      1000 0000      128

	temperatureMin = Temperature(16)
	temperatureMax = Temperature(31)
)
