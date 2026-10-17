package app

import (
	"SPRk_Space/internal/config"
	"SPRk_Space/internal/controllers"
	"SPRk_Space/internal/logger"
	"SPRk_Space/internal/services"
	"SPRk_Space/internal/storage"
	"SPRk_Space/internal/utils"
	"context"
	"fmt"
)

type App struct {
	ctx         context.Context
	logger      logger.Logger
	config      *config.Config
	storage     *storage.Storage
	services    *services.Services
	controllers *controllers.Controllers
}

func (a App) Init() (app *App, err error) {
	err = utils.LoadEnvFile(".env")
	if err != nil {
		fmt.Print("error: app: Init: LoadFromEnv: ", err)
		panic("error: app: Init: LoadFromEnv: ")
	}
	a.config, err = config.NewConfig()
	if err != nil {
		fmt.Print("Error: app: Init: Config: NewConfig: ", err)
		fmt.Print("Info: Default config accepted ")
	}
	a.logger, err = logger.NewLogger(&a.config.Logger)
	if err != nil {
		panic(fmt.Sprintf("error: app: Init: LoadFromEnv: %w", err))
	}
	a.storage, err = storage.NewStorage(a.ctx, &a.config.Storage, a.logger)
	if err != nil {
		a.logger.Fatal(err.Error())
	}
	a.logger.Info("Storage init complete")

	a.controllers, err = controllers.NewControllers(a.ctx, &a.config.Controllers, a.logger)
	if err != nil {
		a.logger.Fatal("Error: App: Init: NewControllers:  %w", err)
	}

	a.services = services.NewServices(a.ctx, &a.config.Services, a.logger, a.storage)
	return &a, nil
}

func (a App) Run() error {
	a.controllers.Run()

	//a.services.Run()
	return nil
}

func (a *App) Close() {
	a.logger.Close()
}
