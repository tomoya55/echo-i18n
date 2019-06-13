package middleware

import (
	"github.com/BurntSushi/toml"
	"github.com/labstack/echo"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var bundle *i18n.Bundle

func init() {
	bundle = i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	bundle.MustLoadMessageFile("active.en.toml")
	bundle.MustLoadMessageFile("active.es.toml")
}

func I18n() echo.MiddlewareFunc {
	return func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			accept := c.Request().Header.Get("Accept-Language")
			localizer := i18n.NewLocalizer(bundle, accept)
			c.Set("localizer", localizer)
			return handlerFunc(c)
		}
	}
}
