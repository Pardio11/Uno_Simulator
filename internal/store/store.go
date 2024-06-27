package store

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
	"unoSimulator/internal/game"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

var (
	ErrNotFound = errors.New("not found")
)

type redisHandler struct {
	client *redis.Client
}

func NewRedisHandler() *redisHandler{
	godotenv.Load("../../.env")
	redisAddress := os.Getenv("REDIS_ADDRESS")
	redisPassword := os.Getenv("REDIS_PASSWORD")
	fmt.Println("Redis Address:", redisAddress)
    fmt.Println("Redis Password:", redisPassword)
	
	client := redis.NewClient(&redis.Options{
		Addr: redisAddress,
		Password: redisPassword,
		DB:0,
	})
	rh := redisHandler{client}
	ping,err :=	rh.client.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	fmt.Println(ping)
	return &rh
}

func (rh redisHandler) CreateEditGame(game game.Game) error{
	key:=game.Id
	g,err:= json.Marshal(game)
	
	if err != nil {
		return err
	}
	err = rh.client.Set(context.Background(), key,g,20*time.Minute).Err()
	if err != nil {
		fmt.Printf("Fail to Set Pair Value Key:%v, Value:%v\n",key,g)
		return err
	}
	return nil
}

func (rh redisHandler) GameExist(id string) bool{
	return rh.client.Exists(context.Background(),id).Val()==1
}

func (rh redisHandler) GetGame(id string) (game.Game,error){
	g:=game.Game{}
	val, err := rh.client.Get(context.Background(),id).Result()
	if err != nil {
		fmt.Println("Failed to GET", err)
		return g,ErrNotFound
	}
	if err := json.Unmarshal([]byte(val), &g); err != nil{
		return g, errors.New("UnMarshal Failed")
	}
	return g, nil
}
