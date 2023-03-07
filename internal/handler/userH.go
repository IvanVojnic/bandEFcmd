// Package handler users comm handlers
package handler

import (
	"net/http"

	"cmdMS/models"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

// AcceptFriendsRequest used to accept friends request
func (h *Handler) AcceptFriendsRequest(ctx echo.Context) error { // nolint:dupl, gocritic
	var reqBody models.Friends
	err := ctx.Bind(&reqBody)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"user sender": reqBody,
		}).Errorf("Bind json %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "data not correct")
	}
	userID := ctx.Get("user_id").(uuid.UUID)
	err = h.userS.AcceptFriendsRequest(ctx.Request().Context(), reqBody.UserSender, userID)
	if err != nil {
		logrus.Errorf("Accept request, %s", err)
		return echo.NewHTTPError(http.StatusBadRequest, "accept failed")
	}
	return ctx.String(http.StatusOK, "request accepted")
}

// DeclineFriendsRequest used to decline friends request
func (h *Handler) DeclineFriendsRequest(ctx echo.Context) error { // nolint:dupl, gocritic
	var reqBody models.Friends
	err := ctx.Bind(&reqBody)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"user sender": reqBody,
		}).Errorf("Bind json %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "data not correct")
	}
	userID := ctx.Get("user_id").(uuid.UUID)
	err = h.userS.DeclineFriendsRequest(ctx.Request().Context(), reqBody.UserSender, userID)
	if err != nil {
		logrus.Errorf("Decline request, %s", err)
		return echo.NewHTTPError(http.StatusBadRequest, "decline failed")
	}
	return ctx.String(http.StatusOK, "request decline")
}

// GetFriends used to get friends
func (h *Handler) GetFriends(ctx echo.Context) error { // nolint:dupl, gocritic
	userID := ctx.Get("user_id").(uuid.UUID)
	friends, err := h.userS.GetFriends(ctx.Request().Context(), userID)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"friends": friends,
		}).Errorf("Get friends request, %s", err)
		return echo.NewHTTPError(http.StatusBadRequest, "get friends failed")
	}
	return ctx.JSON(http.StatusOK, friends)
}

// SendFriendsRequest used to send friends request
func (h *Handler) SendFriendsRequest(ctx echo.Context) error { // nolint:dupl, gocritic
	var reqBody models.Friends
	errBind := ctx.Bind(&reqBody)
	if errBind != nil {
		logrus.WithFields(logrus.Fields{
			"reqBody": reqBody,
		}).Errorf("Bind json, %s", errBind)
		return echo.NewHTTPError(http.StatusInternalServerError, "data not correct")
	}
	userID := ctx.Get("user_id").(uuid.UUID)
	err := h.userS.SendFriendsRequest(ctx.Request().Context(), userID, reqBody.UserReceiver)
	if err != nil {
		logrus.Errorf("send friends request, %s", err)
		return echo.NewHTTPError(http.StatusBadRequest, "smth went wrong")
	}
	return ctx.String(http.StatusOK, "request sent")
}

// FindUser used to find user
func (h *Handler) FindUser(ctx echo.Context) error {
	var reqBody models.UserFind
	errBind := ctx.Bind(&reqBody)
	if errBind != nil {
		logrus.Errorf("send friends request, %s", errBind)
		return echo.NewHTTPError(http.StatusBadRequest, "wrong data")
	}
	user, err := h.userS.FindUser(ctx.Request().Context(), reqBody.UserEmail)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"user": user,
		}).Errorf("GET find user, %s", err)
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	return ctx.JSON(http.StatusOK, user)
}

// GetRequest used to get requests
func (h *Handler) GetRequest(ctx echo.Context) error { // nolint:dupl, gocritic
	userID := ctx.Get("user_id").(uuid.UUID)
	users, err := h.userS.GetRequest(ctx.Request().Context(), userID)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"users": users,
		}).Errorf("GET users request, %s", err)
		return echo.NewHTTPError(http.StatusBadRequest, "cannot get users")
	}
	return ctx.JSON(http.StatusOK, users)
}
