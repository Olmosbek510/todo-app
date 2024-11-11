package main

import (
	"github.com/Olmosbek510/todo-app"
	"github.com/Olmosbek510/todo-app/pkg/handler"
	"github.com/Olmosbek510/todo-app/pkg/repository"
	"github.com/Olmosbek510/todo-app/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	dbConfig := viper.Sub("db")
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     dbConfig.GetString("host"),
		Port:     dbConfig.GetString("port"),
		Username: dbConfig.GetString("username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   dbConfig.GetString("name"),
		SSLMode:  dbConfig.GetString("ssl.mode"),
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running http server: %s",
			err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
