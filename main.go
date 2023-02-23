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

	"github.com/IvanVojnic/bandEFuser/proto"
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

	var profileServ *service.AuthService
	var userServ *service.UserCommSrv
	conn, err := grpc.Dial(":8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logrus.Fatalf("error while conecting to user ms, %s", err)
	}
	c := pr.NewMs1Client(conn)
	profileRepo := repository.NewUserPostgres(db)
	userRepo := repository.NewUserCommPostgres(db)
	profileServ = service.NewAuthService(profileRepo)
	userServ = service.NewUserCommSrv(userRepo)
	profileHandlers := handler.NewHandler(profileServ, userServ)
	profileHandlers.InitRoutes(e)
}
