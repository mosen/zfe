package main

import "flag"

var serverActive = flag.String("server-active", "localhost:10051", "The Zabbix Server to send Active Checks to")

func main() {
	flag.Parse()

}
