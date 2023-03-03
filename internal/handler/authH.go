package handler

import (
	"cmdMS/internal/utils"
	"cmdMS/models"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (h *Handler) SignUp(c echo.Context) error {
	user := models.User{}
	errBind := c.Bind(&user)
	if errBind != nil {
		logrus.WithFields(logrus.Fields{
			"Error Bind json while creating user": errBind,
			"user":                                user,
		}).Errorf("Bind json, %s", errBind)
		return echo.NewHTTPError(http.StatusInternalServerError, "data not correct")
	}
	err := h.authS.SignUp(c.Request().Context(), &user)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"Error create user": err,
		}).Errorf("CREATE USER request, %s", err)
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
		}).Errorf("Bind json, %s", errBind)
		return echo.NewHTTPError(http.StatusInternalServerError, "data not correct")
	}
	tokens, err := h.authS.SignIn(c.Request().Context(), &user)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"Error sign in user": err,
		}).Errorf("SIGN IN USER request, %s", err)
		return echo.NewHTTPError(http.StatusUnauthorized, "wrong data")
	}
	return c.JSON(http.StatusOK, &tokens)
}

func (h *Handler) RefreshToken(c echo.Context) error {
	var tokens Tokens
	errBind := c.Bind(&tokens)
	if errBind != nil {
		logrus.WithFields(logrus.Fields{
			"Error get tokens": errBind,
			"tokens":           tokens,
		}).Errorf("Refresh tokens request, %s", errBind)
		return echo.NewHTTPError(http.StatusInternalServerError, "cannot bind")
	}
	userID, errRT := utils.ParseToken(tokens.RefreshToken)
	if errRT != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "login please")
	}
	newAt, errAT := utils.GenerateToken(userID, utils.TokenATDuretion)
	if errAT != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "cannot bind")
	}
	return c.JSON(http.StatusOK, Tokens{AccessToken: newAt, RefreshToken: tokens.RefreshToken})
}
