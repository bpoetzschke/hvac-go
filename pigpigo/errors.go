package pigpigo

import "fmt"

var (
	ErrorInitFailed    = fmt.Errorf("failed to init gpio interface")
	ErrorSetGpioFailed = fmt.Errorf("failed to set gpio mode")
)
