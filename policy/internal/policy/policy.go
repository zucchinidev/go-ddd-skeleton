package policy

import (
	"github.com/zucchinidev/go-ddd-skeleton/policy/internal/user"
)

// domain

type Policy struct {
	Identifier int
	User       *user.User
}

type PolicyStatus int

const (
	OpenPolicy   PolicyStatus = 100
	ClosedPolicy PolicyStatus = 200
)

type PolicyRepository interface {
	OpenPolicies() ([]Policy, error)
}
