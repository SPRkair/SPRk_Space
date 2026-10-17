package cli

import (
	"SPRk_Space/internal/logger"
	"SPRk_Space/internal/services"
	"context"
	"fmt"
)

type handlers struct {
	ctx    context.Context
	logger logger.Logger
	svc    *services.Services
	c      *console
}

func newHandlers(ctx context.Context, svc *services.Services, c *console, log logger.Logger) *handlers {
	return &handlers{ctx: ctx, svc: svc, c: c, logger: log}
}

func (h handlers) ShowUsersList() {
	users, err := h.svc.FakeService.ShowUsersList()
	if err != nil {
		h.logger.Error("%v", err)
	}
	fmt.Println("\nUser list:")
	for _, u := range users {
		fmt.Printf("%d - %s - %s\n", u.Id, u.Name, u.Email)
	}
}

func (h handlers) GetUserById() {
	id, ok := h.c.ReadInt64("Enter user ID: ")
	if !ok {
		return
	}
	user, err := h.svc.FakeService.GetUserById(id)
	if err != nil {
		h.logger.Error("%v", err)
	}
	fmt.Printf("\nUser %d: %s - %s\n", id, user.Name, user.Email)
}

func (h handlers) UpdateUser() {
	id, ok := h.c.ReadInt64("Enter user ID: ")
	if !ok {
		return
	}
	user, err := h.svc.FakeService.GetUserById(id)
	if err != nil {
		h.logger.Error("Cli: handlers: UpdateUser: %d: %v", id, err)
		return
	}

	updated, ok := editStruct(*h.c, fmt.Sprintf("Edit user %d", id), &user)
	if !ok {
		return
	}
	err = h.svc.FakeService.Users.UpdateUser(h.ctx, *updated)
	if err != nil {
		h.logger.Error("Cli: handlers: SaveUser: %d: %v", id, err)
		return
	}
	fmt.Println("\nSaved.")
	h.logger.Info("Saved!")
}
