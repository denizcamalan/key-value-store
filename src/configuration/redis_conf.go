package configuration

import (
	"log"

	"github.com/go-redis/redis"
)

func NewDatabase() *redis.Client{
	client := redis.NewClient(&redis.Options{
		// Addr: "redis:6379",
		Addr: "localhost:6379",
	   	Password: "",
	   	DB: 0,
	})

	if err := client.Ping().Err(); err != nil {
		log.Println(err)
		return nil
	}

	
	return client
 }
