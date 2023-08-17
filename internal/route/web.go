package route

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sparkymat/wigo/internal/auth"
	"github.com/sparkymat/wigo/internal/handler"
)

func registerWebRoutes(app *echo.Group, cfg ConfigService, userService UserService) {
	webApp := app.Group("")

	webApp.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "form:csrf",
	}))

	webApp.GET("/login", handler.Login(cfg, userService))
	webApp.POST("/login", handler.DoLogin(cfg, userService))

	if !cfg.DisableRegistration() {
		webApp.GET("/register", handler.Register(cfg, userService))
		webApp.POST("/register", handler.DoRegister(cfg, userService))
	}

	authenticatedWebApp := webApp.Group("")

	if cfg.ReverseProxyAuthentication() {
		authenticatedWebApp.Use(auth.ProxyAuthMiddleware(cfg, userService))
	} else {
		authenticatedWebApp.Use(auth.Middleware(cfg, userService))
	}

	authenticatedWebApp.GET("/", handler.Home(cfg, userService))
}
