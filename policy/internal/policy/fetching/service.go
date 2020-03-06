package fetching

import (
	"github.com/zucchinidev/go-ddd-skeleton/policy/internal/policy"
	"github.com/zucchinidev/go-ddd-skeleton/policy/shared/errors"
)

type Service interface {
	Invoke() ([]policy.Policy, error)
}

type service struct {
	repository policy.PolicyRepository
}

func NewService(repository policy.PolicyRepository) Service {
	return &service{repository: repository}
}

func (s *service) Invoke() ([]policy.Policy, error) {
	pp, err := s.repository.OpenPolicies()

	if err != nil {
		if errors.IsPolicySearchError(err) {
			return nil, errors.WrapFetchingPoliciesUseCase(err, "error fetching open policies")
		}
		return nil, err
	}
	return pp, nil
}
