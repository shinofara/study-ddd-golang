package locator

import "gitlab.com/shinofara/alpha/infrastructure/storage"

type Locator struct {
	sl *ServiceLocator
}

func New() *Locator {
	return &Locator{
		sl: &ServiceLocator{},
	}
}

func (l *Locator) SetStorage() error {
	l.sl.storage = storage.NewLocal()
	return nil
}

func (l *Locator) ServiceLocator() *ServiceLocator {
	return l.sl
}

type ServiceLocator struct {
	storage storage.Storage
}

func (sl *ServiceLocator) Storage() storage.Storage { return sl.storage }
