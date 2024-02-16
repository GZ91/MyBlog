package handlers

import (
	"net/http"
	"text/template"

	"go.uber.org/zap"
)

func (h *Handlers) Course(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(h.config.GetMainPath() + "/source/course.html")
	if err != nil {
		h.logger.Error("error course page", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	tmpl.Execute(w, struct{}{})
}
