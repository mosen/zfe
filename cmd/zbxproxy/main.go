package main

import (
	"flag"
	"fmt"
	"github.com/mosen/zfe/pkg/client"
	"os"
	"os/signal"
	"syscall"
)

var _ = flag.String("server-active", "localhost:10051", "The Zabbix Server to send Active Checks to")

func main() {
	flag.Parse()

	proxy, err := client.NewProxy("localhost:10051")
	if err != nil {
		fmt.Println("Unable to create Proxy")
		os.Exit(1)
	}

	errs := make(chan error, 2)

	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	proxy.Start(errs)
}
