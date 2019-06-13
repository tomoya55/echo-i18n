package main

import (
	"errors"
	"net/http"

	"github.com/nicksnyder/go-i18n/v2/i18n"

	miw "github.com/tomoya55/echo-i18n/middleware"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func t(c echo.Context, msg *i18n.Message) (string, error) {
	lz, ok := c.Get("localizer").(*i18n.Localizer)
	if ok {
		return lz.LocalizeMessage(msg)
	}
	return "", errors.New("cannot find localizer")
}

func hello(c echo.Context) error {
	return c.JSON(http.StatusOK, struct{ Message string }{Message: "Hello, World"})
}

func helloi18n(c echo.Context) error {
	msg := &i18n.Message{
		ID:    "hello.message",
		Other: "Hello, world",
	}
	tr, _ := t(c, msg)
	return c.JSON(http.StatusOK, struct{ Message string }{Message: tr})
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(miw.I18n())

	// Routes
	e.GET("/hello", hello)
	e.GET("/helloi18n", helloi18n)

	// Start server
	e.Logger.Fatal(e.Start(":3333"))
}
