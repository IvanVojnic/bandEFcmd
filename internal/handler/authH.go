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
			"user": user,
		}).Errorf("Bind json, %s", errBind)
		return echo.NewHTTPError(http.StatusInternalServerError, "data not correct")
	}
	tokens, err := h.authS.SignIn(c.Request().Context(), &user)
	if err != nil {
		logrus.Errorf("SIGN IN USER request, %s", err)
		return echo.NewHTTPError(http.StatusUnauthorized, "wrong data")
	}
	return c.JSON(http.StatusOK, &tokens)
}

func (h *Handler) RefreshToken(c echo.Context) error {
	var tokens models.Tokens
	errBind := c.Bind(&tokens)
	if errBind != nil {
		logrus.WithFields(logrus.Fields{
			"tokens": tokens,
		}).Errorf("Refresh tokens request, %s", errBind)
		return echo.NewHTTPError(http.StatusInternalServerError, "cannot bind")
	}
	userID, errRT := utils.ParseToken(tokens.RefreshToken)
	if errRT != nil {
		logrus.WithFields(logrus.Fields{
			"userID": userID,
		}).Errorf("get user ID from token, %s", errRT)
		return echo.NewHTTPError(http.StatusUnauthorized, "login please")
	}
	newAt, errAT := utils.GenerateToken(userID, utils.TokenATDuration)
	if errAT != nil {
		logrus.WithFields(logrus.Fields{
			"new Access token": newAt,
		}).Errorf("generate access token, %s", errRT)
		return echo.NewHTTPError(http.StatusInternalServerError, "cannot bind")
	}
	return c.JSON(http.StatusOK, models.Tokens{AccessToken: newAt, RefreshToken: tokens.RefreshToken})
}
