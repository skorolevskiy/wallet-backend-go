package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/skorolevskiy/wallet-backend-go/internal/config"
	"github.com/skorolevskiy/wallet-backend-go/internal/repository"
	"github.com/skorolevskiy/wallet-backend-go/internal/repository/postgres"
	"github.com/skorolevskiy/wallet-backend-go/internal/server"
	"github.com/skorolevskiy/wallet-backend-go/internal/service"
	transport "github.com/skorolevskiy/wallet-backend-go/internal/transport/rest"
	"github.com/spf13/viper"
	"log"
	"os"
)

func main() {
	if err := config.InitConfig(); err != nil {
		log.Fatalf("error initialization configs: %s", err.Error())
	}
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variable: %s", err.Error())
	}
	db, err := postgres.NewPostgresDB(postgres.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.ports"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Username: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASS"),
	})
	if err != nil {
		log.Fatalf("failed to install DB: %s", err.Error())
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handler := transport.NewHandler(services)
	srv := new(server.Server)
	if err := srv.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
		log.Fatalf("error occured while runing http server: %s", err.Error())
	}
}
