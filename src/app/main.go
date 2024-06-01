package main

import (
	slick "github.com/SlavaMashkov/big-web-framework"
	"github.com/SlavaMashkov/big-web-framework/src/app/handler"
	"log/slog"
)

func main() {
	app := slick.New()

	app.Plug(WithAuth)
	app.Get("/profile", handler.HandleProfileIndex)
	app.Get("/dashboard", handler.HandleDashboardIndex)

	err := app.Start(":3000")
	if err != nil {
		slog.Error("error starting server", "err", err)
	}
}

func WithAuth(h slick.Handler) slick.Handler {
	return func(c *slick.Context) error {
		c.Set("email", "mail@gmail.com")

		slog.Info("hello from the middleware")

		return h(c)
	}
}
