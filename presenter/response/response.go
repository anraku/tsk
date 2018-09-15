package response

import (
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func Success(c echo.Context, result interface{}) error {
	log.Info(result)
	return c.JSON(200, result)
}

func ErrInternalServerError(c echo.Context, result interface{}) error {
	log.Error(result)
	return c.JSON(500, result)
}
