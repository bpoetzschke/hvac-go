package irsender

type IRSender interface {
	Setup() error
	Shutdown()
}
