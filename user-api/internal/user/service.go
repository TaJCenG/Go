package user

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(u User) error {
	if err := Validate(u); err != nil {
		return err
	}
	return s.repo.Create(u)
}

func (s *Service) GetAll() ([]User, error) {
	return s.repo.GetAll()
}

func (s *Service) GetByID(id int) (User, error) {
	return s.repo.GetByID(id)
}

func (s *Service) Update(id int, u User) error {
	if err := Validate(u); err != nil {
		return err
	}
	return s.repo.Update(id, u)
}

func (s *Service) Delete(id int) error {
	return s.repo.Delete(id)
}
