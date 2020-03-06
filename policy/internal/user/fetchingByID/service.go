package fetchingByID

import (
	"github.com/zucchinidev/go-ddd-skeleton/policy/internal/user"
)

type Service interface {
	Invoke(userID int) (*user.User, error)
}

type service struct {
	repository user.UserRepository
}

func NewService(repository user.UserRepository) Service {
	return &service{repository: repository}
}

func (s *service) Invoke(userID int) (*user.User, error) {
	return s.repository.ByID(userID)
}
