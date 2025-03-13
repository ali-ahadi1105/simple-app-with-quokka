package handlers

import (
	"net/http"

	quokka "github.com/ali-ahadi1105/Quokka"
)

type Handlers struct {
	App *quokka.Quokka
}

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	err := h.App.Render.Page(w, r, "home", nil, nil)
	if err != nil {
		h.App.ErrorLog.Println("error in generating view: ", err)
	}
}
