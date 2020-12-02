package irprotocol

type IRProtocol interface {
	One()
	Zero()
	ProcessCode([]IRCode)
}
