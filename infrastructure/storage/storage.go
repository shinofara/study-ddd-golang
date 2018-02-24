package storage

type Storage interface {
	Write() error
	Read() error
}

type Local struct{}

func NewLocal() *Local {
	return &Local{}
}

func (l *Local) Write() error {
	return nil
}
func (l *Local) Read() error {
	return nil
}
