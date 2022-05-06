package server

import (
	"fmt"
	"github.com/charlesbourget/WeatherFetcher/config"
	"github.com/gin-contrib/cors"
	"log"
)

func Init() {
	configuration := config.GetConfig()
	r := NewRouter()
	err := r.SetTrustedProxies(nil)
	if err != nil {
		log.Fatalln("error on starting gin server", err)
	}
	r.Use(cors.Default())
	err = r.Run(fmt.Sprint(":", configuration.GetString("server.port")))
	if err != nil {
		log.Fatal("error on starting gin server", err)
	}
}
