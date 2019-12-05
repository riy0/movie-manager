package api

import (
	"github.com/labstack/echo"
	"github.com/valyala/fasthttp"
)

func FetchPopularVideos() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(fasthttp.StatusOK, "Good")
	}
}
