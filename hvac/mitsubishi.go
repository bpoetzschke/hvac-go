package hvac

import "github.com/bpoetzschke/hvac-go/irprotocol"

type Mitsubishi interface {
}

func NewMitsubishi() Mitsubishi {
	return &mitsubishi{
		irProtocol: irprotocol.NewNECIRProtocol(),
	}
}

type mitsubishi struct {
	irProtocol irprotocol.IRProtocol
}
