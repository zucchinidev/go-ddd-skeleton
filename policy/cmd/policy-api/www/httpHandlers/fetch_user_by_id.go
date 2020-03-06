package httpHandlers

import (
	nativeErrors "errors"
	"github.com/julienschmidt/httprouter"
	"github.com/zucchinidev/go-ddd-skeleton/policy/cmd/policy-api/www/engine"
	"github.com/zucchinidev/go-ddd-skeleton/policy/internal/user"
	"github.com/zucchinidev/go-ddd-skeleton/policy/internal/user/fetchingByID"
	"github.com/zucchinidev/go-ddd-skeleton/policy/shared/errors"
	"net/http"
	"strconv"
)

type userResponse struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}

func FetchUserByID(repository user.UserRepository) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		userID := params.ByName("user_id")
		id, err := strconv.Atoi(userID)
		if userID == "" || err != nil {
			engine.Respond(w, r, http.StatusBadRequest, nativeErrors.New("invalid user_id param"))
			return
		}
		u, err := fetchingByID.NewService(repository).Invoke(id)

		if errors.IsUserNotFoundError(err) {
			engine.Respond(w, r, http.StatusNotFound, nil)
			return
		}

		if err != nil {
			engine.Respond(w, r, http.StatusInternalServerError, err)
			return
		}

		engine.Respond(w, r, http.StatusOK, &userResponse{
			ID:    u.ID,
			Email: u.Email,
		})
		return
	}
}
