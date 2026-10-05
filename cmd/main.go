package main

import (
	app "SPRk_Space/internal"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	a, err := app.App{}.Init()
	if err != nil {
		fmt.Println("Error: main: App: Init")
		panic("Error: main: App.Init:")
	}

	// Обработка сигналов для корректного завершения приложения
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		a.Close()
		os.Exit(0)
	}()

	a.Run()
	a.Close()
}
