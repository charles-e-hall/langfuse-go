package transmission

type Sender interface {
	SendAll() error
	Start() error
}

type DefaultSender struct {
	SendQueue chan Trace
	Started   bool
}

func (s *DefaultSender) SendAll() error {
	return nil
}

func (s *DefaultSender) Start() error {
	s.Started = true
	return nil
}
