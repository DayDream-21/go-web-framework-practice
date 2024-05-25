package main

import (
	"github.com/DayDream-21/big-web-framework/src/app/views/profile"
	"log/slog"
)

func main() {
	app := New()

	app.Get("/profile", HandleUserProfile)

	slog.Info("starting server on port :3000")
	err := app.Start(":3000")
	if err != nil {
		slog.Error("error starting server", "err", err)
	}
}

func HandleUserProfile(c *Context) error {
	return c.Render(profile.Index())
}
