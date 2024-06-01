package handler

import (
	slick "github.com/SlavaMashkov/big-web-framework"
	"github.com/SlavaMashkov/big-web-framework/src/app/view/dashboard"
)

func HandleDashboardIndex(c *slick.Context) error {
	return c.Render(dashboard.ShowIndex(dashboard.Show()))
}
