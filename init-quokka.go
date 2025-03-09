package main

import (
	"log"
	"os"

	quokka "github.com/ali-ahadi1105/Quokka"
)

func initApplication() *application {

	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	quo := &quokka.Quokka{}
	err = quo.New(path)

	if err != nil {
		log.Fatal(err)
	}

	quo.AppName = "myApp"
	quo.InfoLog.Println("Debug is set to " ,quo.Debug)
	quo.Debug = true

	app := &application{
		App: quo,
	}

	return app
}
