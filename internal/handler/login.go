package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sparkymat/wigo/internal/auth"
	"github.com/sparkymat/wigo/internal/view"
	"golang.org/x/crypto/bcrypt"
)

func Login(cfg ConfigService, _ UserService) echo.HandlerFunc {
	return func(c echo.Context) error {
		return renderLoginPage(cfg, c, "", "")
	}
}

func DoLogin(cfg ConfigService, userService UserService) echo.HandlerFunc {
	return func(c echo.Context) error {
		username := c.FormValue("username")
		password := c.FormValue("password")

		user, err := userService.FetchUserByUsername(c.Request().Context(), username)
		if err != nil {
			log.Printf("failed to load user. err: %v", err)

			return renderLoginPage(cfg, c, username, "Authentication failed")
		}

		if bcrypt.CompareHashAndPassword([]byte(user.EncryptedPassword), []byte(password)) != nil {
			log.Printf("password does not match")

			return renderLoginPage(cfg, c, username, "Authentication failed")
		}

		err = auth.SaveUsernameToSession(cfg, c, user.Username)
		if err != nil {
			log.Printf("failed to save email to session. err: %v", err)

			return renderLoginPage(cfg, c, username, "Authentication failed")
		}

		//nolint:wrapcheck
		return c.Redirect(http.StatusSeeOther, "/")
	}
}

func renderLoginPage(cfg ConfigService, c echo.Context, email string, errorMessage string) error {
	csrfToken := getCSRFToken(c)
	if csrfToken == "" {
		log.Print("error: csrf token not found")

		//nolint:wrapcheck
		return c.String(http.StatusInternalServerError, "server error")
	}

	pageHTML := view.Login(cfg.DisableRegistration(), csrfToken, email, errorMessage)
	htmlString := view.Layout("wigo | login", csrfToken, pageHTML)

	//nolint:wrapcheck
	return c.HTMLBlob(http.StatusOK, []byte(htmlString))
}
