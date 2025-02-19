package cfg

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Server struct {
	Host string
	Port string
}

type DataBase struct {
	User string
	Pass string
	Host string
	Port string
	Name string
}

type Kafka struct {
	Server  string
	Port    string
	GroupId string
	Retry   string
}

type Config struct {
	Server
	DataBase
	Kafka
}

// <-- CONSTRUCTOR --> //

func GetConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error in load .env : ", err.Error())
	}

	return &Config{
		Server: Server{
			Host: os.Getenv("SERVER_HOST"),
			Port: os.Getenv("SERVER_PORT"),
		},
		DataBase: DataBase{
			User: os.Getenv("DB_USER"),
			Pass: os.Getenv("DB_PASS"),
			Host: os.Getenv("DB_HOST"),
			Port: os.Getenv("DB_PORT"),
			Name: os.Getenv("DB_NAME"),
		},
		Kafka: Kafka{
			Server:  os.Getenv("KAFKA_SERVER"),
			Port:    os.Getenv("KAFKA_PORT"),
			GroupId: os.Getenv("KAFKA_GROUP_ID"),
			Retry:   os.Getenv("RETRY_TIME"),
		},
	}

}
