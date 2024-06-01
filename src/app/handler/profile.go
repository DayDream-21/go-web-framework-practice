package handler

import (
	slick "github.com/DayDream-21/big-web-framework"
	"github.com/DayDream-21/big-web-framework/src/app/model"
	"github.com/DayDream-21/big-web-framework/src/app/view/profile"
)

func HandleProfileIndex(c *slick.Context) error {
	users := []model.User{
		{
			FirstName: "Slava",
			LastName:  "Mashkov",
			Email:     "mashkov.vd@gmail.com",
		},
		{
			FirstName: "Alex",
			LastName:  "Starkov",
			Email:     "starkov.am@gmail.com",
		},
	}

	return c.Render(profile.ShowIndex(profile.Show(users)))
}
