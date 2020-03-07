package errors

import "github.com/pkg/errors"

type policyNotFound struct {
	error
}

func WrapPolicyNotFound(err error, format string, args ...interface{}) error {
	return &policyNotFound{errors.Wrapf(err, format, args...)}
}

func NewPolicyNotFound(format string, args ...interface{}) error {
	return &policyNotFound{errors.Errorf(format, args...)}
}

func IsPolicyNotFoundError(err error) bool {
	err = errors.Cause(err)
	_, ok := err.(*policyNotFound)
	return ok
}

type policySearch struct {
	error
}

func WrapPolicySearch(err error, format string, args ...interface{}) error {
	return &policySearch{errors.Wrapf(err, format, args...)}
}

func NewPolicySearch(format string, args ...interface{}) error {
	return &policySearch{errors.Errorf(format, args...)}
}

func IsPolicySearchError(err error) bool {
	err = errors.Cause(err)
	_, ok := err.(*policySearch)
	return ok
}

type fetchingPoliciesUseCase struct {
	error
}

func WrapFetchingPoliciesUseCase(err error, format string, args ...interface{}) error {
	return &fetchingPoliciesUseCase{errors.Wrapf(err, format, args...)}
}

func NewFetchingPoliciesUseCase(format string, args ...interface{}) error {
	return &fetchingPoliciesUseCase{errors.Errorf(format, args...)}
}

func IsFetchingPoliciesUseCaseError(err error) bool {
	err = errors.Cause(err)
	_, ok := err.(*fetchingPoliciesUseCase)
	return ok
}

type userSearch struct {
	error
}

func WrapUserSearch(err error, format string, args ...interface{}) error {
	return &userSearch{errors.Wrapf(err, format, args...)}
}

func NewUserSearch(format string, args ...interface{}) error {
	return &userSearch{errors.Errorf(format, args...)}
}

func IsUserSearchError(err error) bool {
	err = errors.Cause(err)
	_, ok := err.(*userSearch)
	return ok
}

type userNotFound struct {
	error
}

func WrapUserNotFound(err error, format string, args ...interface{}) error {
	return &userNotFound{errors.Wrapf(err, format, args...)}
}

func NewUserNotFound(format string, args ...interface{}) error {
	return &userNotFound{errors.Errorf(format, args...)}
}

func IsUserNotFoundError(err error) bool {
	err = errors.Cause(err)
	_, ok := err.(*userNotFound)
	return ok
}

type updatePolicy struct {
	error
}

func WrapUpdatePolicy(err error, format string, args ...interface{}) error {
	return &updatePolicy{errors.Wrapf(err, format, args...)}
}

func NewUpdatePolicy(format string, args ...interface{}) error {
	return &updatePolicy{errors.Errorf(format, args...)}
}

func IsUpdatePolicyError(err error) bool {
	err = errors.Cause(err)
	_, ok := err.(*updatePolicy)
	return ok
}

type blockUser struct {
	error
}

func WrapBlockUser(err error, format string, args ...interface{}) error {
	return &blockUser{errors.Wrapf(err, format, args...)}
}

func NewBlockUser(format string, args ...interface{}) error {
	return &blockUser{errors.Errorf(format, args...)}
}

func IsBlockUserError(err error) bool {
	err = errors.Cause(err)
	_, ok := err.(*blockUser)
	return ok
}
