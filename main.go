package main

import (
	"cheap-roulette-api/domain"
	"net/http"
	"os"

	"github.com/labstack/echo/v4/middleware"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"

	"github.com/labstack/echo/v4"
)

var (
	db  *gorm.DB
	err error
)

func main() {

	// Database
	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&domain.User{})

	// Server
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	// Routing
	e.GET("/users", allUsers(db))
	e.POST("/users", newUser(db))
	e.DELETE("/users/:id", deleteUser(db))

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}

func allUsers(db *gorm.DB) func(ctx echo.Context) error {
	return func(ctx echo.Context) error {
		var users []domain.User
		db.Find(&users)
		return ctx.JSON(http.StatusOK, users)
	}
}

func newUser(db *gorm.DB) func(ctx echo.Context) error {
	return func(ctx echo.Context) error {
		u := new(domain.User)
		if err := ctx.Bind(u); err != nil {
			return err
		}
		db.Create(&domain.User{
			Name: u.Name,
		})
		return ctx.String(http.StatusOK, u.Name+" user successfully created")
	}
}

func deleteUser(db *gorm.DB) func(ctx echo.Context) error {
	return func(ctx echo.Context) error {
		id := ctx.Param("id")
		var user domain.User
		db.Delete(&user, id)
		return ctx.String(http.StatusOK, "user successfully deleted")
	}
}
