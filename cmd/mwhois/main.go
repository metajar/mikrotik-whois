package main

import (
	"flag"
	"fmt"
	"github.com/metajar/mikrotik-whois/internal/api"
	"github.com/metajar/mikrotik-whois/internal/config"
	"log"
)

func main() {
	ipPtr := flag.String("ip", "", "ip to perform the lookup on.")
	flag.Parse()

	M := api.New(&config.MikrotikConfig{
		Address:  "192.168.88.1",
		Username: "api",
		Password: "api",
		Port:     "8728",
	})
	err := M.Connect()
	if err != nil {
		fmt.Println(err)
	}
	if *ipPtr == "" {
		log.Fatal("need an ip with --ip")
	}
	a, err := M.GetDHCPHost(*ipPtr)
	if err != nil {
		fmt.Println(err)
		return
	}
	a.Render()

}
