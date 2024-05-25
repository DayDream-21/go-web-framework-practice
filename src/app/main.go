package main

import (
	"github.com/DayDream-21/big-web-framework"
	"github.com/DayDream-21/big-web-framework/src/app/handler"
	"log/slog"
)

func main() {
	app := slick.New()

	app.Get("/profile", handler.HandleProfileIndex)
	app.Get("/dashboard", handler.HandleDashboardIndex)

	err := app.Start(":3000")
	if err != nil {
		slog.Error("error starting server", "err", err)
	}
}
