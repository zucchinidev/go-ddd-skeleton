package conf

import (
	"github.com/joeshaw/envdecode"
	"github.com/zucchinidev/go-ddd-skeleton/policy/shared/adapters/logger"
	"sync"
)

type Conf struct {
	SQLDBUser string `env:"GO_DDD_SKELETON_SQL_DB_USER,required"`
	SQLDBPass string `env:"GO_DDD_SKELETON_SQL_DB_PASS,required"`
	SQLDBName string `env:"GO_DDD_SKELETON_SQL_DB_NAME,required"`
	SQLDBHost string `env:"GO_DDD_SKELETON_SQL_DB_HOST,required"`
	Addr      string `env:"GO_DDD_SKELETON_HTTP_SERVER_LISTEN_ADDRESS,required"`
}

var config *Conf
var once sync.Once

func init() {
	config = &Conf{}
}

func C() *Conf {
	once.Do(func() {
		if err := envdecode.Decode(config); err != nil {
			logger.New().FatalError(err)
		}
	})
	return config
}
