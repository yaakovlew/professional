package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"proffesional"
	"proffesional/handler"
	"proffesional/repository"
	"proffesional/service"
	"syscall"
)

func main() {

	if err := initConfig(); err != nil {
		err.Error()
	}

	if err := godotenv.Load(); err != nil {
		err.Error()
	}
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		fmt.Println(err.Error())
		panic("Fatal to connect to DB")
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(proffesional.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			panic(err.Error())
		}
	}()
	fmt.Println("MephiSRW started")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	fmt.Println("MephiSRW shutting down")

	if err := srv.ShutDown(context.Background()); err != nil {
		panic(err.Error())
	}

	if err := db.Close(); err != nil {
		panic(err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
