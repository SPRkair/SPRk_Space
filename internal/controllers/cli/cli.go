package cli

import (
	"SPRk_Space/internal/config"
	"SPRk_Space/internal/logger"
	"SPRk_Space/internal/services"
	"bufio"
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Cli struct {
	ctx    context.Context
	config *config.Cli
	logger logger.Logger

	sc  *bufio.Scanner
	svc *services.Services

	Menu *Menu
	h    *handlers
	con  *console
}

func NewCli(ctx context.Context, cfg *config.Cli, log logger.Logger, svc *services.Services) (*Cli, error) {
	cli := Cli{
		ctx:    ctx,
		config: cfg,
		logger: log,
		svc:    svc,
		sc:     bufio.NewScanner(os.Stdin),
	}
	cli.con = newConsole(cli.sc)
	cli.h = newHandlers(ctx, svc, cli.con, log)
	cli.Menu = newMenu(cli.h)
	return &cli, nil
}

func (c *Cli) Run() error {
	for !c.Menu.exit {
		select {
		case <-c.ctx.Done():
			return c.ctx.Err()
		default:
		}
		c.Menu.show()
		if !c.sc.Scan() {
			break
		}
		n, err := strconv.Atoi(strings.TrimSpace(c.sc.Text()))
		if err != nil {
			continue
		}
		c.Menu.Select(n)
	}
	if err := c.sc.Err(); err != nil {
		return fmt.Errorf("Cli: read input: %w", err)
	}
	return nil
}

func (c *Cli) Close() {
	c.logger.Debug_info("Cli Close")
}
