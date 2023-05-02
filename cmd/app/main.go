package main

import (
	"context"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/skorolevskiy/wallet-backend-go/internal/config"
	"github.com/skorolevskiy/wallet-backend-go/internal/repository"
	"github.com/skorolevskiy/wallet-backend-go/internal/server"
	"github.com/skorolevskiy/wallet-backend-go/internal/service"
	transport "github.com/skorolevskiy/wallet-backend-go/internal/transport/rest"
	"github.com/skorolevskiy/wallet-backend-go/pkg/database"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
)

//	@title			Wallet Simple API
//	@version		1.0
//	@description	API server for Wallet Application

//	@host		localhost:8000
//	@BasePath	/

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := config.InitConfig(); err != nil {
		logrus.Fatalf("error initialization configs: %s", err.Error())
	}
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variable: %s", err.Error())
	}
	db, err := database.NewPostgresDB(database.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.ports"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Username: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASS"),
	})
	if err != nil {
		logrus.Fatalf("failed to install DB: %s", err.Error())
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handler := transport.NewHandler(services)
	srv := new(server.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while runing http server: %s", err.Error())
		}
	}()

	logrus.Print("Wallet App Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("Wallet Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}
	if err = db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}
}
