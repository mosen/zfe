package main

import (
	"flag"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/jmoiron/sqlx"
	"github.com/mosen/zfe/pkg/hosts"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var (
		httpAddr = flag.String("http.addr", ":8080", "HTTP listen address")
		pgsqlUri = flag.String("pgsql.uri", "postgres://zabbix:zabbix@localhost/zabbix?sslmode=disable", "PostgreSQL Database URI")
	)
	flag.Parse()

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	db, err := sqlx.Open("postgres", *pgsqlUri)
	if err != nil {
		fmt.Errorf("Cannot open database connection: %s", err)
	}

	hostsRepo := hosts.NewHostsRepository(db)
	hostsSvc := hosts.NewHostService(hostsRepo)
	templatesRepo := hosts.NewTemplatesRepository(db)
	templateSvc := hosts.NewTemplateService(templatesRepo)
	hostsHandler := hosts.MakeHandler(hostsSvc, templateSvc, logger)

	mux := http.NewServeMux()
	mux.Handle("/hosts/v1/", hostsHandler)
	mux.Handle("/templates/v1/", hostsHandler)

	http.Handle("/", mux)

	errs := make(chan error, 2)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		logger.Log("transport", "http", "address", *httpAddr, "msg", "listening")
		errs <- http.ListenAndServe(*httpAddr, nil)
	}()

	logger.Log("terminated", <-errs)

}
