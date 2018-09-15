package response

import "github.com/labstack/echo"

func Success(c echo.Context, result interface{}) error {
	return c.JSON(200, result)
}

func ErrInternalServerError(c echo.Context, result interface{}) error {
	return c.JSON(500, result)
}
