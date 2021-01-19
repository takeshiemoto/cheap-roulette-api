package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Ping struct {
	Status int `json:"status"`
}

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		s := new(Ping)
		s.Status = http.StatusOK
		return c.JSON(http.StatusOK, s)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
