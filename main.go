package main

import quokka "github.com/ali-ahadi1105/Quokka"

type application struct {
	App *quokka.Quokka
}

func main() {
	quo := initApplication()
	quo.App.ListenAndServe()
}
