package templates

const Database = `
package database

import (
	"os"
	"time"
	"github.com/jinzhu/gorm"
	"github.com/go-redis/redis/v7"
	log "github.com/sirupsen/logrus"
)

//Context MySQL Database Context Pool
var Context *gorm.DB

//Redis Client
var Redis *redis.Client

func NewDatabase(connectionString string) {
	var err error

	if connectionString == "" {
		log.Println("[database] mock data non connect database if you want to connect database please setting db user")
		return
	}

	Context, err = gorm.Open("mysql", connectionString)
	if err != nil {
		log.Panic("[database] connect to database failed", err)
		return
	}

	err = Context.DB().Ping()
	if err != nil {
		log.Panic("[database] connect to database failed", err)
		return
	}

	Context.DB().SetMaxOpenConns(10)
	Context.DB().SetMaxIdleConns(5)
	Context.DB().SetConnMaxLifetime(time.Second * 60)
}

//NewRedisDatabase 連線到Redis
func NewRedisDatabase() {
	Redis = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDRESS"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})
	pong, err := Redis.Ping().Result()
	if err != nil {
		log.Panic("[database] connect to redis failed ", err)
	} else {
		log.Println("[database] connect to redis success ", pong)
	}
}

func CloseDatabase() {
	if Context != nil{
		_ = Context.Close()
	}
}

//CloseRedis 關閉Redis
func CloseRedis() {
	if Redis != nil {
		_ = Redis.Close()
	}
}`

