package routes

import (
	"github.com/labstack/echo/v4"
)

func Post_post(c echo.Context) error {
	return c.String(200, "post new user")
}
