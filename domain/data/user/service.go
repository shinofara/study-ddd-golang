package user

// User メッセージを操作する為に必要な、Repositoryなどを管理
type Service struct {
	userRepo Repository
}

func NewService(userRepo Repository) *Service {
	return &Service{
		userRepo: userRepo,
	}
}

func (m *Service) Register(name string) (*User, error) {
	u := &User{
		Name: name,
	}

	return m.userRepo.Add(u)
}
