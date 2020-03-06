package main

import (
	"context"
	"fmt"
	mysqlPolicy "github.com/zucchinidev/go-ddd-skeleton/policy/internal/policy/store/mysql"
	mysqlUser "github.com/zucchinidev/go-ddd-skeleton/policy/internal/user/store/mysql"
	"github.com/zucchinidev/go-ddd-skeleton/policy/shared/ping"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/zucchinidev/go-ddd-skeleton/policy/cmd/policy-api/www"
	"github.com/zucchinidev/go-ddd-skeleton/policy/shared/adapters/conf"
	"github.com/zucchinidev/go-ddd-skeleton/policy/shared/adapters/logger"
	"github.com/zucchinidev/go-ddd-skeleton/policy/shared/adapters/store"
)

var Version string

var rootCmd = &cobra.Command{
	Use:   "policy-api",
	Short: "policy-api is a very fast service to execute sql queries",
	Run: func(cmd *cobra.Command, args []string) {
		var (
			sqlManager *store.Storage
			err        error
			c          = conf.C()
			l          = logger.New()
		)
		if sqlManager, err = initSQLManager(c); err != nil {
			l.SQLConnectionError(err)
			return
		}
		defer sqlManager.Close()
		policiesRepository := mysqlPolicy.NewPolicyRepository(sqlManager)
		userRepository := mysqlUser.NewUserRepository(sqlManager)

		shutdown := make(chan struct{}, 1)

		wwwConf := www.Conf{Addr: c.Addr, Version: Version}
		pings := []ping.Pinger{sqlManager}
		httpServer := www.Server(wwwConf, pings, policiesRepository, userRepository)
		kill := terminate(httpServer, []io.Closer{sqlManager})
		go func() {
			l.HTTPServerInitialization(c.Addr)
			if err := httpServer.ListenAndServe(); err != http.ErrServerClosed {
				l.HTTPServerError(err)
				kill(l)
			}
		}()
		go interruptSignal(l, shutdown)
		for {
			select {
			case <-shutdown:
				kill(l)
				return
			}
		}
	},
}

// Execute assembles the app commands necessaries to up the applications
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func terminate(svc *http.Server, closers []io.Closer) func(l *logger.Standard) {
	return func(l *logger.Standard) {
		if err := svc.Shutdown(context.Background()); err != nil {
			l.HTTPServerShutdownError(err)
		}
		for _, closer := range closers {
			err := closer.Close()
			if err != nil {
				l.ShowCloserError(err)
			}
		}
	}
}

func interruptSignal(l *logger.Standard, shutdown chan struct{}) {
	signals := make(chan os.Signal, 1)
	// sigterm signal sent from kubernetes, interrupt signal sent from terminal
	signal.Notify(signals, syscall.SIGTERM, os.Interrupt)
	<-signals
	l.ReceivedInterruptSignal()
	shutdown <- struct{}{}
}

func initSQLManager(c *conf.Conf) (*store.Storage, error) {
	sqlManager := store.NewSqlManager(store.Cfg{Addr: c.SQLDBHost, DBName: c.SQLDBName, User: c.SQLDBUser, Passwd: c.SQLDBPass})
	if err := sqlManager.Connect(); err != nil {
		return nil, err
	}
	return sqlManager, nil
}
