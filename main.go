// package main used to start program
package main

import (
	"os"

	"cmdMS/internal/handler"
	"cmdMS/internal/repository"
	"cmdMS/internal/service"

	prRoom "github.com/IvanVojnic/bandEFroom/proto"
	prUser "github.com/IvanVojnic/bandEFuser/proto"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	e := echo.New()
	logger := logrus.New()
	logger.Out = os.Stdout
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, values middleware.RequestLoggerValues) error {
			logrus.WithFields(logrus.Fields{
				"URI":    values.URI,
				"status": values.Status,
			}).Info("request")
			return nil
		},
	}))

	var userAuthServ *service.AuthService
	var userCommServ *service.UserCommSrv
	var roomServ *service.RoomInviteService
	connUserMS, err := grpc.Dial("app:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logrus.Fatalf("error while conecting to user ms, %s", err)
	}

	connRoomMS, err := grpc.Dial("app:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logrus.Fatalf("error while conecting to user ms, %s", err)
	}

	clientUserComm := prUser.NewUserCommClient(connUserMS)
	clientUserAuth := prUser.NewUserAuthClient(connUserMS)

	clientRoom := prRoom.NewRoomClient(connRoomMS)
	clientInvite := prRoom.NewInviteClient(connRoomMS)

	userRepo := repository.NewUserMS(clientUserComm, clientUserAuth)
	roomRepo := repository.NewRoomMS(clientRoom, clientInvite)

	userAuthServ = service.NewAuthService(userRepo)
	userCommServ = service.NewUserCommSrv(userRepo)
	roomServ = service.NewRoomInviteService(roomRepo)

	profileHandlers := handler.NewHandler(userAuthServ, userCommServ, roomServ)
	profileHandlers.InitRoutes(e)
}
