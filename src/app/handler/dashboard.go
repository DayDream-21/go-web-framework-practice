package handler

import (
	slick "github.com/DayDream-21/big-web-framework"
	"github.com/DayDream-21/big-web-framework/src/app/view/dashboard"
)

func HandleDashboardIndex(c *slick.Context) error {
	return c.Render(dashboard.ShowIndex(dashboard.Show()))
}
