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
}

func NewCli(ctx context.Context, cfg *config.Cli, log logger.Logger, svc *services.Services) (*Cli, error) {
	cli := Cli{
		ctx:    ctx,
		config: cfg,
		logger: log,
		svc:    svc,
		sc:     bufio.NewScanner(os.Stdin),
	}
	cli.Menu = cli.newMenu()
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

func (c *Cli) ReadLine(prompt string) (string, bool) {
	fmt.Printf(prompt)
	if !c.sc.Scan() {
		return "", false
	}
	return strings.TrimSpace(c.sc.Text()), true
}

func (c *Cli) ReadInt64(prompt string) (int64, bool) {
	s, ok := c.ReadLine(prompt)
	if !ok {
		return 0, false
	}
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		fmt.Printf("\nInvalid number\n")
		return 0, false
	}
	return n, true
}
