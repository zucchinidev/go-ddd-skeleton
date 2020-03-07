package httpHandlers

import (
	nativeErrors "errors"
	"github.com/julienschmidt/httprouter"
	"github.com/zucchinidev/go-ddd-skeleton/policy/cmd/policy-api/www/engine"
	"github.com/zucchinidev/go-ddd-skeleton/policy/internal/policy"
	"github.com/zucchinidev/go-ddd-skeleton/policy/internal/policy/closingPolicyWhenUserIsBlocked"
	"net/http"
	"strconv"
)

func ClosePolicyWhenUserIsBlocked(repository policy.PolicyRepository) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		userID := params.ByName("user_id")
		id, err := strconv.Atoi(userID)
		if userID == "" || err != nil {
			engine.Respond(w, r, http.StatusBadRequest, nativeErrors.New("invalid user_id param"))
			return
		}

		if err = closingPolicyWhenUserIsBlocked.NewService(repository).Invoke(id); err != nil {
			engine.Respond(w, r, http.StatusInternalServerError, err)
			return
		}

		engine.Respond(w, r, http.StatusNoContent, nil)
		return
	}
}
