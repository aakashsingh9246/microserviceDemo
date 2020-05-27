package main

import (
	"github.com/gocql/gocql"
	"flag"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/go-kit/kit/log"

	"github.com/go-kit/kit/log/level"

	"net/http"
	"os"
	"os/signal"
	"syscall"

	"./account"
)

var Session *gocql.Session

func main() {
	var httpAddr = flag.String("http", ":8080", "http listen address")
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "account",
			"time:", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")


	var err error
 
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "mydb"
	Session, err = cluster.CreateSession()
	if err != nil {
		level.Error(logger).Log("exit", err)
			os.Exit(-1)
	}

	flag.Parse()
	var srv account.Service
	{
		repository := account.NewRepo(Session, logger)

		srv = account.NewService(repository, logger)
	}

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	endpoints := account.MakeEndpoints(srv)

	go func() {
		fmt.Println("listening on port", *httpAddr)
		handler := account.NewHTTPServer(endpoints)
		errs <- http.ListenAndServe(*httpAddr, handler)
	}()

	level.Error(logger).Log("exit", <-errs)
}