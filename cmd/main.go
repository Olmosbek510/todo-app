package main

import (
	"context"
	"fmt"
	"github.com/Olmosbek510/todo-app"
	"github.com/Olmosbek510/todo-app/pkg/handler"
	"github.com/Olmosbek510/todo-app/pkg/repository"
	"github.com/Olmosbek510/todo-app/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&SimpleFormatter{})
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

	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server: %s",
				err.Error())
		}
	}()

	logrus.Println("TodoApp Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Println("TodoApp Shutting Down")

	if err := srv.ShutDown(context.Background()); err != nil {
		logrus.Errorf("error occurred on server shutting down: %s", err.Error())
	}
	if err := db.Close(); err != nil {
		logrus.Errorf("error occurred on server database connection close: %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

type SimpleFormatter struct{}

func (f *SimpleFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	log := fmt.Sprintf("[%s] %s: %s\n", entry.Time.Format("2006-01-02 15:04:05"), entry.Level, entry.Message)
	return []byte(log), nil
}
