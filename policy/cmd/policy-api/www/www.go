package www

import (
	"github.com/julienschmidt/httprouter"
	"github.com/zucchinidev/go-ddd-skeleton/policy/cmd/policy-api/www/httpHandlers"
	"github.com/zucchinidev/go-ddd-skeleton/policy/cmd/policy-api/www/statusHandler"
	"github.com/zucchinidev/go-ddd-skeleton/policy/internal/policy"
	"github.com/zucchinidev/go-ddd-skeleton/policy/internal/user"
	"github.com/zucchinidev/go-ddd-skeleton/policy/shared/ping"
	"net/http"
)

type Conf struct {
	Addr    string
	Version string
}

func Server(c Conf, pings []ping.Pinger, policiesRepository policy.PolicyRepository, userRepository user.UserRepository) *http.Server {
	router := httprouter.New()
	router.GET("/status", statusHandler.Status(pings, c.Version))
	router.GET("/policies", httpHandlers.FetchPolicies(policiesRepository))
	router.GET("/user/:user_id", httpHandlers.FetchUserByID(userRepository))
	return &http.Server{Addr: c.Addr, Handler: router}
}
