package main

//go:generate go run github.com/valyala/quicktemplate/qtc -dir=internal/view

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/sparkymat/wigo/internal/config"
	"github.com/sparkymat/wigo/internal/database"
	"github.com/sparkymat/wigo/internal/dbx"
	"github.com/sparkymat/wigo/internal/route"
	"github.com/sparkymat/wigo/internal/service/user"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	dbDriver, err := database.New(cfg.DatabaseURL())
	if err != nil {
		log.Error(err)
		panic(err)
	}

	if err = dbDriver.AutoMigrate(); err != nil {
		log.Error(err)
		panic(err)
	}

	// Initialize web server
	db := dbx.New(dbDriver.DB())

	userService := user.New(db)

	e := echo.New()
	route.Setup(e, cfg, userService)

	e.Logger.Panic(e.Start(":8080"))
}
