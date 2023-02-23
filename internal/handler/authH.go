package handler

import (
	"cmdMS/internal/utils"
	"cmdMS/models"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
)

type response struct {
	User *models.User `json:"user"`
}

func (h *Handler) SignUp(c echo.Context) error {
	user := models.User{}
	errBind := c.Bind(&user)
	if errBind != nil {
		logrus.WithFields(logrus.Fields{
			"Error Bind json while creating user": errBind,
			"user":                                user,
		}).Info("Bind json")
		return echo.NewHTTPError(http.StatusInternalServerError, "data not correct")
	}
	err := h.authS.SignUp(c.Request().Context(), &user)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"Error create user": err,
		}).Info("CREATE USER request")
		return echo.NewHTTPError(http.StatusBadRequest, "user creating failed")
	}
	return c.String(http.StatusOK, "user created")
}

func (h *Handler) SignIn(c echo.Context) error {
	user := models.User{}
	errBind := c.Bind(&user)
	if errBind != nil {
		logrus.WithFields(logrus.Fields{
			"Error Bind json while creating user": errBind,
			"user":                                user,
		}).Info("Bind json")
		return echo.NewHTTPError(http.StatusInternalServerError, "data not correct")
	}
	tokens, err := h.authS.SignIn(c.Request().Context(), &user)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"Error sign in user": err,
		}).Info("SIGN IN USER request")
		return echo.NewHTTPError(http.StatusUnauthorized, "wrong data")
	}
	return c.JSON(http.StatusOK, &tokens)
}

func (h *Handler) GetUserAuth(c echo.Context) error {
	userID := c.Get("user_id").(uuid.UUID)
	user, err := h.authS.GetUser(c.Request().Context(), userID)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"Error get user": err,
			"user":           user,
		}).Info("GET USER request")
		return echo.NewHTTPError(http.StatusBadRequest, "sign up please")
	}
	return c.JSON(http.StatusOK, response{&user})
}

func (h *Handler) RefreshToken(c echo.Context) error {
	var tokens Tokens
	errBind := c.Bind(&tokens)
	if errBind != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "cannot bind")
	}
	checkRT, errRT := utils.ParseToken(tokens.RefreshToken)
	if checkRT {
		if errRT != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "bad refresh token")
		}
		id, errGetID := utils.ExtractIDFromToken(tokens.RefreshToken)
		if errGetID != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "smth went wrong")
		}
		newAt, errAT := utils.GenerateToken(id, utils.TokenATDuretion)
		if errAT != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "cannot bind")
		}
		return c.JSON(http.StatusOK, Tokens{AccessToken: newAt, RefreshToken: tokens.RefreshToken})
	}
	return echo.NewHTTPError(http.StatusUnauthorized, "login please")
}
