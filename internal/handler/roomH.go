// Package handler room invite handlers
package handler

import (
	"net/http"

	"cmdMS/models"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

// SendInvite used to send invite
func (h *Handler) SendInvite(ctx echo.Context) error { // nolint:dupl, gocritic
	var reqBody models.SendInvite
	err := ctx.Bind(&reqBody)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"body": reqBody,
		}).Errorf("Bind json %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "data not correct")
	}
	userID := ctx.Get("user_id").(uuid.UUID)
	err = h.roomS.SendInvite(ctx.Request().Context(), userID, reqBody.UsersID, reqBody.Place, reqBody.Date)
	if err != nil {
		logrus.Errorf("send invite, %s", err)
		return echo.NewHTTPError(http.StatusBadRequest, "sending invite failed")
	}
	return ctx.String(http.StatusOK, "invite sent")
}

// AcceptInvite used to accept invite
func (h *Handler) AcceptInvite(ctx echo.Context) error { // nolint:dupl, gocritic
	var reqBody models.Invite
	err := ctx.Bind(&reqBody)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"body": reqBody,
		}).Errorf("Bind json %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "data not correct")
	}
	userID := ctx.Get("user_id").(uuid.UUID)
	err = h.roomS.AcceptInvite(ctx.Request().Context(), userID, reqBody.RoomID)
	if err != nil {
		logrus.Errorf("accept invite, %s", err)
		return echo.NewHTTPError(http.StatusBadRequest, "accepting invite failed")
	}
	return ctx.String(http.StatusOK, "invite accepted")
}

// DeclineInvite used to decline invite
func (h *Handler) DeclineInvite(ctx echo.Context) error { // nolint:dupl, gocritic
	var reqBody models.Invite
	err := ctx.Bind(&reqBody)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"body": reqBody,
		}).Errorf("Bind json %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "data not correct")
	}
	userID := ctx.Get("user_id").(uuid.UUID)
	err = h.roomS.AcceptInvite(ctx.Request().Context(), userID, reqBody.RoomID)
	if err != nil {
		logrus.Errorf("decline invite, %s", err)
		return echo.NewHTTPError(http.StatusBadRequest, "decline invite failed")
	}
	return ctx.String(http.StatusOK, "invite declined")
}

// GetRooms used to get rooms
func (h *Handler) GetRooms(ctx echo.Context) error {
	userID := ctx.Get("user_id").(uuid.UUID)
	rooms, err := h.roomS.GetRooms(ctx.Request().Context(), userID)
	if err != nil {
		logrus.Errorf("getting room, %s", err)
		return echo.NewHTTPError(http.StatusBadRequest, "getting room failed")
	}
	return ctx.JSON(http.StatusOK, &rooms)
}

// GetRoomUsers used to get users from current room
func (h *Handler) GetRoomUsers(ctx echo.Context) error {
	var reqBody models.Invite
	err := ctx.Bind(&reqBody)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"body": reqBody,
		}).Errorf("Bind json %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "data not correct")
	}
	users, err := h.roomS.GetRoomUsers(ctx.Request().Context(), reqBody.RoomID)
	if err != nil {
		logrus.Errorf("getting room, %s", err)
		return echo.NewHTTPError(http.StatusBadRequest, "getting room failed")
	}
	return ctx.JSON(http.StatusOK, &users)
}
