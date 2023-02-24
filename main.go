package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"

	"cmdMS/internal/handler"
	"cmdMS/internal/repository"
	"cmdMS/internal/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"

	pr "github.com/IvanVojnic/bandEFuser/proto"
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
	conn, err := grpc.Dial(os.Getenv("PORT"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logrus.Fatalf("error while conecting to user ms, %s", err)
	}

	client := pr.NewUserClient(conn)
	userRepo := repository.NewUserMS(client)
	userAuthServ = service.NewAuthService(userRepo)
	userCommServ = service.NewUserCommSrv(userRepo)
	profileHandlers := handler.NewHandler(userAuthServ, userCommServ)
	profileHandlers.InitRoutes(e)
}
