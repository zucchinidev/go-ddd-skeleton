package httpHandlers

import (
	"github.com/julienschmidt/httprouter"
	"github.com/zucchinidev/go-ddd-skeleton/policy/cmd/policy-api/www/engine"
	"github.com/zucchinidev/go-ddd-skeleton/policy/internal/policy"
	"github.com/zucchinidev/go-ddd-skeleton/policy/internal/policy/fetching"
	"github.com/zucchinidev/go-ddd-skeleton/policy/shared/errors"
	"net/http"
)

type policyResponse struct {
	ID        int    `json:"id"`
	UserEmail string `json:"user_email"`
}

type fetchPoliciesResponse struct {
	Policies   []policyResponse `json:"policies"`
	StatusText string           `json:"status_text"`
}

func FetchPolicies(repository policy.PolicyRepository) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		svc := fetching.NewService(repository)
		pp, err := svc.Invoke()

		if errors.IsPolicyNotFoundError(err) {
			data := &fetchPoliciesResponse{
				Policies:   []policyResponse{},
				StatusText: http.StatusText(http.StatusNotFound),
			}
			engine.Respond(w, r, http.StatusNotFound, data)
			return
		}

		if err != nil {
			engine.Respond(w, r, http.StatusInternalServerError, err)
			return
		}

		engine.Respond(w, r, http.StatusOK, &fetchPoliciesResponse{
			Policies: mapPolicies(pp, func(p policy.Policy) policyResponse {
				return policyResponse{
					ID:        p.Identifier,
					UserEmail: p.User.Email,
				}
			}),
			StatusText: http.StatusText(http.StatusOK),
		})
		return
	}
}

func mapPolicies(pp []policy.Policy, fn func(p policy.Policy) policyResponse) []policyResponse {
	var newPP = []policyResponse{}
	for _, it := range pp {
		newPP = append(newPP, fn(it))
	}
	return newPP
}
