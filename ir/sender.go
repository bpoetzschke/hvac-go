package ir

func NewSender() Sender {
	return &irSender{}
}

type irSender struct {
}

func (s *irSender) Transmit(w Wave) {

}
