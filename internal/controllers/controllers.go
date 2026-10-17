package controllers

import (
	"SPRk_Space/internal/config"
	"SPRk_Space/internal/controllers/cli"
	"SPRk_Space/internal/logger"
	"context"
	"errors"
	"fmt"
)

type Controller interface {
	Run() error
	Close()
}

type Controllers struct {
	ctx    context.Context
	config *config.Controllers
	logger logger.Logger

	cli      Controller
	telegram Controller
	HTTP     Controller
}

func NewControllers(ctx context.Context, cfg *config.Controllers, log logger.Logger) (*Controllers, error) {
	controllers := Controllers{
		ctx:    ctx,
		config: cfg,
		logger: log,
	}
	cli, err := cli.NewCli(ctx, &cfg.Cli, log)
	if err != nil {
		return &Controllers{}, fmt.Errorf("controllers: newControllers: newCli: %w", err)
	}
	/*
		http, err := cli.NewHttp(ctx, &cfg.HTTP, log)
		if err != nil {
			return Controllers{}, fmt.Errorf("controllers: newControllers: NewHttp: %w", err)
		}
		telegram, err := cli.NewTelegram(ctx, &cfg.Telegram, log)
		if err != nil {
			return Controllers{}, fmt.Errorf("controllers: newControllers: NewTelegram: %w", err)
		}
		controllers.controller = append(controllers.controller, http)
		controllers.controller = append(controllers.controller, telegram)
	*/
	controllers.cli = cli
	return &controllers, nil
}

func (c Controllers) Run() error {
	err := c.cli.Run()
	if err != nil && !errors.Is(err, context.Canceled) {
		return fmt.Errorf("Controllers: Run: Cli: %w", err)
	}
	return nil
}
