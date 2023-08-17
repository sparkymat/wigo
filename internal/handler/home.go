package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sparkymat/wigo/internal/view"
)

func Home(_ ConfigService, _ UserService) echo.HandlerFunc {
	return func(c echo.Context) error {
		csrfToken := getCSRFToken(c)
		if csrfToken == "" {
			log.Print("error: csrf token not found")

			//nolint:wrapcheck
			return c.String(http.StatusInternalServerError, "server error")
		}

		pageHTML := view.Home()
		htmlString := view.BasicLayout("wigo", csrfToken, pageHTML)

		//nolint:wrapcheck
		return c.HTMLBlob(http.StatusOK, []byte(htmlString))
	}
}
