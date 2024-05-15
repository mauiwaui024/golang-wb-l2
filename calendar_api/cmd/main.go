package main

import (
	calendarapi "calendar_api"
	"calendar_api/pkg/handler"
	"calendar_api/pkg/repository"
	"calendar_api/pkg/service"
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	InitConfig()
	repo := repository.NewRepository()
	services := service.NewService(repo)
	handler := handler.NewHandler(services)
	srv := new(calendarapi.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
			logrus.Fatal(err)
		}
	}()
	logrus.Printf("Server started on the port %s", viper.GetString("port"))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	logrus.Print("APP shutting down")
	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occurred on server shutting down: %s", err.Error())
	}

}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
