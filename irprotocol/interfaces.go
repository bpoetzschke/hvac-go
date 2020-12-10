package irprotocol

import "github.com/bpoetzschke/hvac-go/ir"

type Protocol interface {
	ProcessCode(gpioPin uint32, irCode []IRCode) (ir.Wave, error)
}
