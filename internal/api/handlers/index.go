package handlers

import (
	"bytes"
	"io"
	"net/http"
	"text/template"

	"github.com/GZ91/MyBlog/internal/models"
	"go.uber.org/zap"
)

func (h *Handlers) Index(w http.ResponseWriter, r *http.Request) {
	UserAuthorized, err := h.authorized(r)
	if err != nil {
		h.logger.Error("error authorized", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !UserAuthorized {
		w.Header().Add("Location", "/login")
		w.WriteHeader(http.StatusTemporaryRedirect)
		return
	}
	tmpl, err := template.ParseFiles(h.config.GetMainPath() + "/source/template.html")
	if err != nil {
		h.logger.Error("error main page", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	artList, err := h.NodeService.GetArts()
	if err != nil {
		h.logger.Error("error get arts", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	tmplt := &models.PageArt{
		Content:    "",
		AlterLabel: "",
		ArtList:    artList,
	}

	tmpl2, err := template.ParseFiles(h.config.GetMainPath() + "/source/index.html")
	if err != nil {
		h.logger.Error("error main page", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	buf := bytes.Buffer{}
	tmpl2.Execute(&buf, tmplt)
	dataBuf, err := io.ReadAll(&buf)
	if err != nil {
		h.logger.Error("error read buf", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	tmplt.Content = string(dataBuf)

	w.Header().Set("Content-Type", "text/html")
	tmpl.Execute(w, tmplt)
}

func (h *Handlers) authorized(r *http.Request) (bool, error) {
	var UserID string
	var userIDCTX models.CtxString = "userID"

	// Извлекаем идентификатор пользователя (UserID) из контекста запроса
	UserIDVal := r.Context().Value(userIDCTX)
	if UserIDVal != nil {
		UserID = UserIDVal.(string)
	}

	ok, err := h.NodeService.Authorized(UserID)
	if err != nil {
		h.logger.Error("error authorized", zap.Error(err))
		return false, err
	}
	return ok, nil
}
