package main

import (
	"cheap-roulette-api/domain"
	"net/http"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"

	"github.com/labstack/echo/v4"
)

type Ping struct {
	Status int `json:"status"`
}

var (
	db  *gorm.DB
	err error
	dsn = "host=localhost user=roulette dbname=roulette password=roulette sslmode=disable"
)

func main() {
	dbInit()

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		s := new(Ping)
		s.Status = http.StatusOK
		return c.JSON(http.StatusOK, s)
	})

	e.Logger.Fatal(e.Start(":1323"))
}

func dbInit() {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(domain.User{})
}
