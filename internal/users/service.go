package users

type Service interface {
	Create(user *User) (*User, error)
	Update(user *User) (*User, error)
	Get(user *User) (*User, error)
	List() ([]User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) Create(user *User) (*User, error) {
	return s.repository.Insert(user)
}

func (s *service) Update(user *User) (*User, error) {
	return s.repository.Update(user)
}

func (s *service) Get(user *User) (*User, error) {
	return s.repository.Get(user)
}

func (s *service) List() ([]User, error) {
	return s.repository.List()
}
