package cli

import (
	"SPRk_Space/internal/config"
	"SPRk_Space/internal/logger"
	"context"
)

type Cli struct {
	ctx    context.Context
	confgi *config.Cli
	logger logger.Logger

	Menu menu
}

type menu struct {
}

func NewCli(ctx context.Context, cfg *config.Cli, log logger.Logger) (Cli, error) {
	cli := Cli{
		ctx:    ctx,
		confgi: cfg,
		logger: log,
	}
	return cli, nil
}

func (c Cli) Run() {

}

func (c Cli) Close() {
	c.logger.Debug_info("Cli Close")
}
