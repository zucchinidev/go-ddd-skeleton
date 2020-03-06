package statusHandler

import (
	"github.com/julienschmidt/httprouter"
	"github.com/zucchinidev/go-ddd-skeleton/policy/cmd/policy-api/www/engine"
	"github.com/zucchinidev/go-ddd-skeleton/policy/shared/ping"
	"net/http"
	"strings"
)

type statusResp struct {
	Status  string `json:"status"`
	Msg     string `json:"msg"`
	Version string `json:"version"`
}

func Status(pings []ping.Pinger, version string) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		var errs []string
		for _, ping := range pings {
			if err := ping.Ping(); err != nil {
				errs = append(errs, err.Error())
			}
		}
		if len(errs) > 0 {
			engine.Respond(w, r, http.StatusInternalServerError, statusResp{Status: "DOWN", Msg: strings.Join(errs, " - "), Version: version})
		} else {
			engine.Respond(w, r, http.StatusOK, statusResp{Status: "UP", Version: version})
		}
	}
}
