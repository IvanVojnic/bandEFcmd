package main

import (
	"os"

	"cmdMS/internal/config"
	"cmdMS/internal/handler"
	"cmdMS/internal/repository"
	"cmdMS/internal/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
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
	cfg, err := config.NewConfig()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"Error":  err,
			"config": cfg,
		}).Fatal("failed to get config")
	}
	var profileServ *service.AuthService
	var userServ *service.UserCommSrv
	db, err := repository.NewPostgresDB(cfg)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"Error connection to database rep.NewPostgresDB()": err,
		}).Fatal("DB ERROR CONNECTION")
	}
	defer repository.ClosePool(db)
	profileRepo := repository.NewUserPostgres(db)
	userRepo := repository.NewUserCommPostgres(db)
	profileServ = service.NewAuthService(profileRepo)
	userServ = service.NewUserCommSrv(userRepo)
	profileHandlers := handler.NewHandler(profileServ, userServ)
	profileHandlers.InitRoutes(e)
}
