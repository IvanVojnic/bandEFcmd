package handler

import (
	"cmdMS/internal/errorwrapper"
	"cmdMS/internal/utils"
	"cmdMS/models"
	"context"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"strings"
)

// UserComm service consists of methods of user actions
type UserComm interface {
	GetFriends(ctx context.Context, userID uuid.UUID) ([]models.User, error)
	SendFriendsRequest(ctx context.Context, userSender uuid.UUID, userReceiver uuid.UUID) error
	AcceptFriendsRequest(ctx context.Context, userSenderID uuid.UUID, userReceiverID uuid.UUID) error
	FindUser(ctx context.Context, userEmail string) (models.User, error)
	GetRequest(ctx context.Context, userID uuid.UUID) ([]models.User, error)
}

// Authorization service consists of methods fo user
type Authorization interface {
	SignUp(ctx context.Context, user *models.User) error
	GetUser(ctx context.Context, id uuid.UUID) (models.User, error)
	SignIn(ctx context.Context, user *models.User) (Tokens, error)
	UpdateRefreshToken(context.Context, string, uuid.UUID) error
}

type Handler struct {
	authS Authorization
	userS UserComm
}

// Tokens used to define at and rt
type Tokens struct {
	AccessToken  string `json:"access"`
	RefreshToken string `json:"refresh"`
}

func NewHandler(authS Authorization, users UserComm) *Handler {
	return &Handler{authS: authS, userS: users}
}

// InitRoutes used to init routes.txt
func (h *Handler) InitRoutes(router *echo.Echo) *echo.Echo {
	router.Use(jwtAuthMiddleware())
	rAuth := router.Group("/auth")
	rAuth.POST("/refreshToken", h.RefreshToken)
	rAuth.POST("/createUser", h.SignUp)
	rAuth.POST("/signIn", h.SignIn)
	rAuth.POST("/getUserAuth", h.GetUserAuth)
	rUserComm := router.Group("/userComm")
	rUserComm.Use(middleware.Logger())
	rUserComm.POST("/getFriends", h.GetFriends)
	rUserComm.GET("/getFriendsRequest", h.SendFriendsRequest)
	rUserComm.POST("/acceptFriendsRequest", h.AcceptFriendsRequest)
	rUserComm.GET("/findFriend", h.FindUser)
	rUserComm.GET("/sendRequest", h.GetRequest)
	router.Logger.Fatal(router.Start(":40000"))
	return router
}

func jwtAuthMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			if c.Path() == "/auth/createUser" || c.Path() == "/auth/signIn" || c.Path() == "/auth/refreshToken" {
				return next(c)
			}
			var tokens Tokens
			req := c.Request()
			headers := req.Header
			atHeader := headers.Get("Authorization")
			if atHeader == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "no access token in header")
			}
			atHeaderArr := strings.Split(atHeader, " ")
			tokens.AccessToken = atHeaderArr[1]
			authorized, errIsAuth := utils.IsAuthorized(tokens.AccessToken)
			if authorized {
				userID, errGetID := utils.ExtractIDFromToken(tokens.AccessToken)
				if errGetID != nil {
					return echo.NewHTTPError(http.StatusUnauthorized, errorwrapper.ErrorResponse{Message: errGetID.Error()})
				}
				c.Set("user_id", userID)
				return next(c)
			}
			return echo.NewHTTPError(http.StatusUnauthorized, errorwrapper.ErrorResponse{Message: errIsAuth.Error()})
		}
	}
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
