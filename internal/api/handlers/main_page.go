package handlers

import (
	"net/http"
	"text/template"

	"go.uber.org/zap"
)

type mainPage struct {
	Title      string
	Grid       string
	AlterLabel string
}

func (h *Handlers) MainPage(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("../../source/template.html")
	if err != nil {
		h.logger.Error("error main page", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
	}
	tmplt := &mainPage{
		Title:      "My Blog",
		Grid:       "",
		AlterLabel: "",
	}
	w.Header().Set("Content-Type", "text/html")
	tmpl.Execute(w, tmplt)

}
