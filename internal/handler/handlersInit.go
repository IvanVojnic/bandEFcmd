// Package handler handlers init
package handler

import (
	"cmdMS/internal/middleware"
	"context"
	"time"

	"cmdMS/models"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

// UserComm service consists of methods of user actions
type UserComm interface {
	GetFriends(ctx context.Context, userID uuid.UUID) ([]*models.User, error)
	SendFriendsRequest(ctx context.Context, userSender, userReceiver uuid.UUID) error
	AcceptFriendsRequest(ctx context.Context, userSenderID, userReceiverID uuid.UUID) error
	DeclineFriendsRequest(ctx context.Context, userSenderID, userID uuid.UUID) error
	FindUser(ctx context.Context, userEmail string) (models.User, error)
	GetRequest(ctx context.Context, userID uuid.UUID) ([]*models.User, error)
}

// Authorization service consists of methods for user
type Authorization interface {
	SignUp(ctx context.Context, user *models.User) error
	SignIn(ctx context.Context, user *models.User) (models.Tokens, error)
	UpdateRefreshToken(context.Context, string, uuid.UUID) error
}

// RoomInvite service consists of methods for rooms invites
type RoomInvite interface {
	SendInvite(ctx context.Context, userCreatorID uuid.UUID, usersID []*uuid.UUID, place string, date time.Time) error
	AcceptInvite(ctx context.Context, userID, roomID uuid.UUID) error
	DeclineInvite(ctx context.Context, userID, roomID uuid.UUID) error
	GetRooms(ctx context.Context, user uuid.UUID) ([]*models.Room, error)
	GetRoomUsers(ctx context.Context, roomID uuid.UUID) ([]*models.User, error)
}

// Handler object consists of handlers
type Handler struct {
	authS Authorization
	userS UserComm
	roomS RoomInvite
}

// NewHandler used to init Handler obj
func NewHandler(authS Authorization, userS UserComm, roomS RoomInvite) *Handler {
	return &Handler{authS: authS, userS: userS, roomS: roomS}
}

// InitRoutes used to init routes.txt
func (h *Handler) InitRoutes(router *echo.Echo) *echo.Echo {
	router.Use(middleware.JwtAuthMiddleware())
	rAuth := router.Group("/auth")
	rAuth.POST("/refreshToken", h.RefreshToken)
	rAuth.POST("/createUser", h.SignUp)
	rAuth.POST("/signIn", h.SignIn)
	rUserComm := router.Group("/userComm")
	rUserComm.Use(echoMiddleware.Logger())
	rUserComm.POST("/getFriends", h.GetFriends)
	rUserComm.GET("/getFriendsRequest", h.SendFriendsRequest)
	rUserComm.POST("/acceptFriendsRequest", h.AcceptFriendsRequest)
	rUserComm.POST("/declineFriendsRequest", h.DeclineFriendsRequest)
	rUserComm.GET("/findFriend", h.FindUser)
	rUserComm.GET("/sendRequest", h.GetRequest)
	rRoom := router.Group("/room")
	rRoom.POST("/sendInvite", h.SendInvite)
	rRoom.POST("/acceptInvite", h.AcceptInvite)
	rRoom.POST("/declineInvite", h.DeclineInvite)
	rRoom.POST("/getRooms", h.GetRooms)
	rRoom.POST("/getRoomUsers", h.GetRoomUsers)
	router.Logger.Fatal(router.Start(":40000"))
	return router
}

/*func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}*/
