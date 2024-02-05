package main

import (
	"github.com/GZ91/MyBlog/internal/app"
	"go.uber.org/zap"
)

func main() {
	app := app.New("info")
	err := app.Start()
	if err != nil {
		app.Logger.Error("error started", zap.Error(err))
	}
}
