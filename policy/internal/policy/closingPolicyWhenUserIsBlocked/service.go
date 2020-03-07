package closingPolicyWhenUserIsBlocked

import (
	"github.com/zucchinidev/go-ddd-skeleton/policy/internal/policy"
	"github.com/zucchinidev/go-ddd-skeleton/policy/shared/errors"
)

type Service interface {
	Invoke(userID int) error
}

type service struct {
	repository policy.PolicyRepository
}

func NewService(repository policy.PolicyRepository) Service {
	return &service{repository: repository}
}

func (s *service) Invoke(userID int) error {
	return s.repository.WithTransaction(func(tx policy.TransactionManager) error {
		if err := s.repository.CloseByUserID(tx, userID); err != nil {
			return errors.WrapUpdatePolicy(err, "error updating policies to user with id %d", userID)
		}

		if err := s.repository.BlockUser(tx, userID); err != nil {
			return errors.WrapBlockUser(err, "error blocking user with id %d", userID)
		}
		return nil
	})
}
