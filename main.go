package main

import (
	"myApp/handlers"

	quokka "github.com/ali-ahadi1105/Quokka"
)

type application struct {
	App *quokka.Quokka
	Handlers *handlers.Handlers
}

func main() {
	quo := initApplication()
	quo.App.ListenAndServe()
}
