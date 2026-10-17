package cli

import (
	"SPRk_Space/internal/config"
	"SPRk_Space/internal/logger"
	"bufio"
	"context"
	"os"
	"strconv"
	"strings"
)

type Cli struct {
	ctx    context.Context
	config *config.Cli
	logger logger.Logger

	Menu *Menu
}

func NewCli(ctx context.Context, cfg *config.Cli, log logger.Logger) (*Cli, error) {
	cli := Cli{
		ctx:    ctx,
		config: cfg,
		logger: log,
		Menu:   newMenu(),
	}
	return &cli, nil
}

func (c Cli) Run() {
	sc := bufio.NewScanner(os.Stdin)
	for !c.Menu.exit {
		select {
		case <-c.ctx.Done():
			return
		default:
		}
		c.Menu.show()
		if !sc.Scan() {
			return
		}
		n, err := strconv.Atoi(strings.TrimSpace(sc.Text()))
		if err != nil {
			continue
		}
		c.Menu.Select(n)
	}
}

func (c Cli) Close() {
	c.logger.Debug_info("Cli Close")
}
