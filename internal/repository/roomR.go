package repository

import (
	"cmdMS/models"
	"context"
	"fmt"
	pr "github.com/IvanVojnic/bandEFroom/proto"
	"github.com/google/uuid"
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
func (r *RoomMS) GetRooms(ctx context.Context, userID uuid.UUID) ([]models.Room, error) {
	var rooms []models.Room
	res, errGRPC := r.client.GetRooms(ctx, &pr.GetRoomsRequest{UserID: userID})
	if errGRPC != nil {
		return rooms, fmt.Errorf("error while sign up, %s", errGRPC)
	}
	return rooms, nil
}

func (r *RoomMS) GetUsersRoom(ctx context.Context, roomID uuid.UUID) ([]uuid.UUID, error) {
	var usersID []uuid.UUID

	return usersID, nil
}

// SendInvite used to send request to be a friends
func (r *RoomMS) SendInvite(ctx context.Context, users []models.User, roomID uuid.UUID, creatorID uuid.UUID) error {

	return nil
}

// AcceptInvite used to accept invite to the room
func (r *RoomMS) AcceptInvite(ctx context.Context, userID uuid.UUID, roomID uuid.UUID, status int) error {

	return nil
}

// DeclineInvite used to accept invite to the room
func (r *RoomMS) DeclineInvite(ctx context.Context, userID uuid.UUID, roomID uuid.UUID, status int) error {

	return nil
}