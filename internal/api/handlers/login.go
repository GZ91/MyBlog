package handlers

import (
	"io"
	"net/http"
	"net/url"
	"os"
	"text/template"

	"github.com/GZ91/MyBlog/internal/models"
	"go.uber.org/zap"
)

func (h *Handlers) Login(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles(h.config.GetMainPath() + "/source/template.html")
	if err != nil {
		h.logger.Error("error main page", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	file, err := os.Open(h.config.GetMainPath() + "/source/login.html")
	if err != nil {
		h.logger.Error("error open index file", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	data, err := io.ReadAll(file)
	if err != nil {
		h.logger.Error("error read index file", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	dataLine := string(data)
	if r.Method == http.MethodPost {
		dataLine = "<p> Неверный логин или пароль<p>" + dataLine
	}
	tmplt := &models.Page{
		Content:    dataLine,
		AlterLabel: "",
	}
	w.Header().Set("Content-Type", "text/html")
	tmpl.Execute(w, tmplt)

}

func (h *Handlers) LoginPost(w http.ResponseWriter, r *http.Request) {

	data, err := io.ReadAll(r.Body)
	if err != nil {
		h.logger.Error("error read body", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var UserID string
	var userIDCTX models.CtxString = "userID"
	UserIDVal := r.Context().Value(userIDCTX)
	if UserIDVal != nil {
		UserID = UserIDVal.(string)
	}

	ok, err := h.authorization(data, UserID)
	if err != nil {
		h.logger.Error("error authorization", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !ok {
		w.Header().Add("Location", "/login")
		w.WriteHeader(http.StatusTemporaryRedirect)
		return
	}
	w.Header().Add("Location", "/")
	w.WriteHeader(http.StatusSeeOther)

}

func (h *Handlers) authorization(data []byte, userID string) (bool, error) {
	var login, password string

	dataLine := string(data)
	parsedQuery, err := url.ParseQuery(dataLine)
	if err != nil {
		return false, err
	}

	login = parsedQuery.Get("login")
	password = parsedQuery.Get("password")

	return h.NodeService.Login(login, password, userID)
}
