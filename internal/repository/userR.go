package repository

import (
	"cmdMS/internal/handler"
	"cmdMS/models"
	"context"
	"fmt"
	pr "github.com/IvanVojnic/bandEFuser/proto"
	"github.com/google/uuid"
)

// UserMS has an internal grpc object
type UserMS struct {
	clientComm pr.UserCommClient
	clientAuth pr.UserAuthClient
}

// NewUserMS used to init UsesAP
func NewUserMS(clientComm pr.UserCommClient, clientAuth pr.UserAuthClient) *UserMS {
	return &UserMS{clientComm: clientComm, clientAuth: clientAuth}
}

// SignUp used to create user
func (r *UserMS) SignUp(ctx context.Context, user *models.User) error {
	res, errGRPC := r.clientAuth.SignUp(ctx, &pr.SignUpRequest{Password: user.Password, Name: user.Name, Email: user.Email})
	if errGRPC != nil {
		return fmt.Errorf("error while sign up, %s", errGRPC)
	}
	if res.IsCreated {
		return nil
	}
	return fmt.Errorf("cannot create user")
}

// SignInUser used to sign in user
func (r *UserMS) SignIn(ctx context.Context, user *models.User) (handler.Tokens, error) {
	res, errGRPC := r.clientAuth.SignIn(ctx, &pr.SignInRequest{Name: user.Name, Password: user.Password})
	if errGRPC != nil {
		return handler.Tokens{}, fmt.Errorf("error while sign up, %s", errGRPC)
	}

	return handler.Tokens{RefreshToken: res.Rt, AccessToken: res.At}, nil
}

// UpdateRefreshToken used to update rt
func (r *UserMS) UpdateRefreshToken(ctx context.Context, rt string, id uuid.UUID) error {

	return nil
}

// GetFriends used to send friends
func (r *UserMS) GetFriends(ctx context.Context, userID uuid.UUID) ([]models.User, error) {
	users := make([]models.User, 0)
	res, errGRPC := r.clientComm.GetFriends(ctx, &pr.GetFriendsRequest{UserID: userID.String()})
	if errGRPC != nil {
		return users, fmt.Errorf("error while getting friends, %s", errGRPC)
	}
	for i := 0; i < len(res.Friends); i++ {
		friendID, errParse := uuid.Parse(res.Friends[i].ID)
		if errParse != nil {
			return users, fmt.Errorf("error while getting friends, %s", errParse)
		}
		user := models.User{ID: friendID, Name: res.Friends[i].Name, Email: res.Friends[i].Email}
		users = append(users, user)
	}
	return users, nil
}

// SendFriendsRequest used to send requests for user
func (r *UserMS) SendFriendsRequest(ctx context.Context, userSender uuid.UUID, userReceiver uuid.UUID) error {
	_, errGRPC := r.clientComm.SendFriendRequest(ctx, &pr.SendFriendRequestReq{UserID: userSender.String(), ReceiverID: userReceiver.String()})
	if errGRPC != nil {
		return fmt.Errorf("error while send request to be a friend, %s", errGRPC)
	}
	return nil
}

// AcceptFriendsRequest used to accept request
func (r *UserMS) AcceptFriendsRequest(ctx context.Context, userSenderID uuid.UUID, userID uuid.UUID) error {
	_, errGRPC := r.clientComm.AcceptFriendsRequest(ctx, &pr.AcceptFriendsRequestReq{UserSenderID: userSenderID.String(), UserID: userID.String()})
	if errGRPC != nil {
		return fmt.Errorf("error while accepting request to be a friend, %s", errGRPC)
	}
	return nil
}

// DeclineFriendsRequest used to accept request
func (r *UserMS) DeclineFriendsRequest(ctx context.Context, userSenderID uuid.UUID, userID uuid.UUID) error {
	_, errGRPC := r.clientComm.DeclineFriendsRequest(ctx, &pr.DeclineFriendsRequestReq{UserSenderID: userSenderID.String(), UserID: userID.String()})
	if errGRPC != nil {
		return fmt.Errorf("error while decline request to be a friend, %s", errGRPC)
	}
	return nil
}

// FindUser used to find user by email
func (r *UserMS) FindUser(ctx context.Context, userEmail string) (models.User, error) {
	var user models.User
	res, err := r.clientComm.FindUser(ctx, &pr.FindUserRequest{UserEmail: userEmail})
	if err != nil {
		return user, fmt.Errorf("error while sign up, %s", err)
	}
	userID, err := uuid.Parse(res.Friend.ID)
	if err != nil {
		return user, fmt.Errorf("error while getting friends, %s", err)
	}
	user = models.User{ID: userID, Name: res.Friend.Name, Email: res.Friend.Email}
	return user, nil
}

// GetRequest used to send request to be a friends
func (r *UserMS) GetRequest(ctx context.Context, userID uuid.UUID) ([]models.User, error) {
	users := make([]models.User, 0)
	res, errGRPC := r.clientComm.GetRequest(ctx, &pr.GetRequestReq{UserID: userID.String()})
	if errGRPC != nil {
		return users, fmt.Errorf("error while getting requests to be a friend, %s", errGRPC)
	}
	for i := 0; i < len(res.Users); i++ {
		friendID, errParse := uuid.Parse(res.Users[i].ID)
		if errParse != nil {
			return users, fmt.Errorf("error while getting requests to be a friend, %s", errParse)
		}
		user := models.User{ID: friendID, Name: res.Users[i].Name, Email: res.Users[i].Email}
		users = append(users, user)
	}
	return users, nil
}
