package users

type Service interface {
	GetAllUsers() ([]User, error)
	GetOneUser()
	SaveUser(Name, LastName, Email string, Age int, Height float64, Active bool, DateCreate string) (User, error)
	UpdateUser(Id int, Name, LastName, Email string, Age int, Height float64, Active bool, DateCreate string) (User, error)
	DeleteUser(Id int) error
	UpdatePatch(Id int, LastName string, Age int) (User, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAllUsers() ([]User, error) {
	users, err := s.repository.GetAllUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *service) GetOneUser() {
}

func (s *service) SaveUser(Name, LastName, Email string, Age int, Height float64, Active bool, DateCreate string) (User, error) {
	lastId, err := s.repository.LastId()
	if err != nil {
		return User{}, err
	}
	lastId++
	user, err := s.repository.SaveUser(lastId, Name, LastName, Email, Age, Height, Active, DateCreate)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (s *service) UpdateUser(Id int, Name, LastName, Email string, Age int, Height float64, Active bool, DateCreate string) (User, error) {
	return s.repository.UpdateUser(Id, Name, LastName, Email, Age, Height, Active, DateCreate)
}

func (s *service) DeleteUser(Id int) error {
	return s.repository.DeleteUser(Id)
}

func (s *service) UpdatePatch(Id int, LastName string, Age int) (User, error) {
	return s.repository.UpdatePatch(Id, LastName, Age)
}
