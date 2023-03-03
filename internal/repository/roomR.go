package repository

import (
	"cmdMS/models"
	"context"
	"fmt"
	pr "github.com/IvanVojnic/bandEFroom/proto"
	"github.com/google/uuid"
	"time"
)

// RoomMS has an internal grpc object
type RoomMS struct {
	client pr.RoomClient
}

// NewRoomMS used to init RoomAP
func NewRoomMS(client pr.RoomClient) *RoomMS {
	return &RoomMS{client: client}
}

// GetRooms used to get rooms where you had invited
func (r *RoomMS) GetRooms(ctx context.Context, userID uuid.UUID) (*[]models.Room, error) {
	var rooms *[]models.Room
	res, errGRPC := r.client.GetRooms(ctx, &pr.GetRoomsRequest{UserID: userID.String()})
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
		*rooms = append(*rooms, models.Room{ID: roomID, Place: room.Place, Date: date, UserCreatorID: userCreatorID})
	}
	if errGRPC != nil {
		return rooms, fmt.Errorf("error while getting rooms, %s", errGRPC)
	}
	return rooms, nil
}

func (r *RoomMS) GetRoomUsers(ctx context.Context, roomID uuid.UUID) (*[]models.User, error) {
	var users *[]models.User
	res, err := r.client.GetUsersRoom(ctx, &pr.GetUsersRoomRequest{RoomID: roomID.String()})
	for _, user := range res.Users {
		userID, err := uuid.Parse(user.ID)
		if err != nil {
			return users, fmt.Errorf("error while parsing room ID, %s", err)
		}
		*users = append(*users, models.User{ID: userID, Name: user.Name, Email: user.Email})
	}
	if err != nil {
		return users, fmt.Errorf("error while getting users for current room, %s", err)
	}
	return users, nil
}

// SendInvite used to send request to be a friends
func (r *RoomMS) SendInvite(ctx context.Context, userCreatorID uuid.UUID, usersID *[]uuid.UUID, place string, date time.Time) error {
	var users *[]models.User
	res, err := r.client.(ctx, &pr.GetUsersRoomRequest{RoomID: roomID.String()})
	for _, user := range res.Users {
		userID, err := uuid.Parse(user.ID)
		if err != nil {
			return users, fmt.Errorf("error while parsing room ID, %s", err)
		}
		*users = append(*users, models.User{ID: userID, Name: user.Name, Email: user.Email})
	}
	if err != nil {
		return users, fmt.Errorf("error while getting users for current room, %s", err)
	}
	return users, nil
	return nil
}

// AcceptInvite used to accept invite to the room
func (r *RoomMS) AcceptInvite(ctx context.Context, userID uuid.UUID, roomID uuid.UUID) error {

	return nil
}

// DeclineInvite used to accept invite to the room
func (r *RoomMS) DeclineInvite(ctx context.Context, userID uuid.UUID, roomID uuid.UUID) error {

	return nil
}
