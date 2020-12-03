package irprotocol

import "fmt"

type IRCode string

const (
	IRCodeZero = IRCode("0")
	IRCodeOne  = IRCode("1")
)

var (
	ErrUnknownIRCode = fmt.Errorf("unknown IR code")
)
