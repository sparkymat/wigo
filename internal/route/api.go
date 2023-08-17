package route

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sparkymat/wigo/internal/auth"
)

func registerAPIRoutes(app *echo.Group, cfg ConfigService, userService UserService) {
	apiGroup := app.Group("api")

	if cfg.ReverseProxyAuthentication() {
		apiGroup.Use(auth.ProxyAuthMiddleware(cfg, userService))
	} else {
		apiGroup.Use(auth.APIMiddleware(cfg, userService))
	}

	apiGroup.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "header:X-CSRF-Token",
	}))
}
