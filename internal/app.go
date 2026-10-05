package app

import (
	"SPRk_Space/internal/config"
	"SPRk_Space/internal/logger"
	"SPRk_Space/internal/utils"
	"fmt"
)

type App struct {
	logger   logger.Logger
	config   config.Config
	cache    services.Cache
	services services.Services
}

func (a App) Init() (app *App, err error) {
	err = utils.LoadFromEnv(".env")
	if err != nil {
		fmt.Print("error: app: Init: LoadFromEnv: ", err)
		panic("error: app: Init: LoadFromEnv: ")
	}
	a.logger = logger.NewLogger()
	a.config, err = config.NewConfig()
	if err != nil {
		a.logger.Println("Error", "app: Init: Config: NewConfig: ", err)
		a.logger.Println("Info", "Default config accepted ")
	}
	a.services = services.NewServices(&a.config, a.logger, a.cache)
	return &a, nil
}

func (a App) Run() error {
	a.services.Run()
	return nil
}

func (a *App) Close() {
	a.logger.Close()
}
