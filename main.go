package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/charlesbourget/WeatherFetcher/config"
	"github.com/charlesbourget/WeatherFetcher/server"
)

func main() {
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*environment)
	server.Init()
}
