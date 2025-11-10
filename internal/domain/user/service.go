package user

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo}
}

func (s *Service) ListUsers() ([]User, error) {
	return s.repo.GetAll()
}

func (s *Service) CreateUser(u *User) error {
	return s.repo.Create(u)
}
