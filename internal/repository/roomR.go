// Package repository define methods for rooms and invites using RoomMS
package repository

import (
	"cmdMS/models"
	"context"
	"fmt"
	"time"

	pr "github.com/IvanVojnic/bandEFroom/proto"

	"github.com/google/uuid"
)

// RoomMS has an internal grpc object
type RoomMS struct {
	clientRoom   pr.RoomClient
	clientInvite pr.InviteClient
}

// NewRoomMS used to init RoomAP
func NewRoomMS(clientRoom pr.RoomClient, clientInvite pr.InviteClient) *RoomMS {
	return &RoomMS{clientRoom: clientRoom, clientInvite: clientInvite}
}

// GetRooms used to get rooms where you had invited
func (r *RoomMS) GetRooms(ctx context.Context, userID uuid.UUID) ([]*models.Room, error) {
	rooms := make([]*models.Room, 0)
	res, errGRPC := r.clientRoom.GetRooms(ctx, &pr.GetRoomsRequest{UserID: userID.String()})
	for _, room := range res.Rooms {
		roomID, errRoomID := uuid.Parse(room.RoomID)
		if errRoomID != nil {
			return rooms, fmt.Errorf("error while parsing room ID, %s", errRoomID)
		}
		date, errDateParse := time.Parse("2006-01-02 15:04:05", room.Date)
		if errDateParse != nil {
			return rooms, fmt.Errorf("error while parsing date, %s", errDateParse)
		}
		userCreatorID, errUserID := uuid.Parse(room.UserCreatorId)
		if errUserID != nil {
			return rooms, fmt.Errorf("error while parsing user ID, %s", errUserID)
		}
		rooms = append(rooms, &models.Room{ID: roomID, Place: room.Place, Date: date, UserCreatorID: userCreatorID})
	}
	if errGRPC != nil {
		return rooms, fmt.Errorf("error while getting rooms, %s", errGRPC)
	}
	return rooms, nil
}

// GetRoomUsers used to get users from current room
func (r *RoomMS) GetRoomUsers(ctx context.Context, roomID uuid.UUID) ([]*models.User, error) {
	users := make([]*models.User, 0)
	res, err := r.clientRoom.GetUsersRoom(ctx, &pr.GetUsersRoomRequest{RoomID: roomID.String()})
	for _, user := range res.Users {
		userID, errParse := uuid.Parse(user.ID)
		if errParse != nil {
			return users, fmt.Errorf("error while parsing room ID, %s", errParse)
		}
		users = append(users, &models.User{ID: userID, Name: user.Name, Email: user.Email})
	}
	if err != nil {
		return users, fmt.Errorf("error while getting users for current room, %s", err)
	}
	return users, nil
}

// SendInvite used to send request to be a friends
func (r *RoomMS) SendInvite(ctx context.Context, userCreatorID uuid.UUID, usersID []*uuid.UUID, place string, date time.Time) error {
	users := make([]string, 0)
	for _, user := range usersID {
		userID := user.String()
		users = append(users, userID)
	}
	_, err := r.clientInvite.SendInvite(ctx, &pr.SendInviteRequest{Date: date.String(), Place: place, UserCreatorID: userCreatorID.String(), UsersID: users})
	if err != nil {
		return fmt.Errorf("error while sending invite to users for current room, %s", err)
	}
	return nil
}

// AcceptInvite used to accept invite to the room
func (r *RoomMS) AcceptInvite(ctx context.Context, userID, roomID uuid.UUID) error { // nolint:dupl, gocritic
	_, err := r.clientInvite.AcceptInvite(ctx, &pr.AcceptInviteRequest{RoomID: roomID.String(), UserID: userID.String()})
	if err != nil {
		return fmt.Errorf("error while sending invite to users for current room, %s", err)
	}
	return nil
}

// DeclineInvite used to accept invite to the room
func (r *RoomMS) DeclineInvite(ctx context.Context, userID, roomID uuid.UUID) error { // nolint:dupl, gocritic
	_, err := r.clientInvite.DeclineInvite(ctx, &pr.DeclineInviteRequest{RoomID: roomID.String(), UserID: userID.String()})
	if err != nil {
		return fmt.Errorf("error while sending invite to users for current room, %s", err)
	}
	return nil
}
