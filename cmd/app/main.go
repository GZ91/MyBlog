package main

import (
	"github.com/GZ91/MyBlog/internal/app"
	"github.com/GZ91/MyBlog/internal/app/config"
	"github.com/GZ91/MyBlog/internal/app/initializing"
	"go.uber.org/zap"
)

func main() {
	envs, err := initializing.ReadEnv()
	if err != nil {
		panic(err)
	}
	config := config.New(envs.ConnectionStringDB, envs.MainPath)
	app := app.New("info", config)
	err = app.Start()
	if err != nil {
		app.Logger.Error("error started", zap.Error(err))
	}
}
