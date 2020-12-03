package irprotocol

type IRProtocol interface {
	ProcessCode(gpioPin uint32, irCode []IRCode) error
}
