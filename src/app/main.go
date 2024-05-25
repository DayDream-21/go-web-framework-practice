package main

import (
	"github.com/DayDream-21/big-web-framework"
	"github.com/DayDream-21/big-web-framework/src/app/model"
	"github.com/DayDream-21/big-web-framework/src/app/views/profile"
	"log/slog"
)

func main() {
	app := slick.New()

	app.Get("/profile", HandleUserProfile)

	slog.Info("starting server on port :3000")
	err := app.Start(":3000")
	if err != nil {
		slog.Error("error starting server", "err", err)
	}
}

func HandleUserProfile(c *slick.Context) error {
	user := model.User{
		FirstName: "Slava",
		LastName:  "Mashkov",
		Email:     "mashkov.vd@gmail.com",
	}

	return c.Render(profile.Index(user))
}
