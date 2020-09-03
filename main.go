package main

import (
	"log"

	"github.com/manh737/go-mux/app"
	"github.com/manh737/go-mux/config"
)

// return the value of the key
func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	appConfig := config.NewConfig()
	app.ConfigAndRunApp(appConfig)
}