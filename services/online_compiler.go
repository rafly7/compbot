package services

import (
	"bytes"
	"compbot/models"
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
	"go.mau.fi/whatsmeow"
)

func OnlineCompilerValidationService(rdb *redis.Client, infoSender string, cb func() (whatsmeow.SendResponse, error)) {
	// log.Println(res.ID)
	res, err := cb()
	if err != nil {
		log.Print(err)
		return
	}
	val, err := rdb.Get(context.Background(), infoSender).Result()
	if err != redis.Nil { // IF the key is exist and not error
		log.Println("The key is exist")
		if err != nil {
			return
		}
		val, err := rdb.Del(context.Background(), infoSender).Result()
		if err != nil {
			log.Print(err)
			return
		}
		if val == 1 {
			_, err := rdb.SetNX(context.Background(), infoSender, res.ID, 5*time.Minute).Result()
			if err != nil {
				log.Print(err)
				return
			}
		}
	} else if err == redis.Nil && val == "" {
		_, err := rdb.SetNX(context.Background(), infoSender, res.ID, 5*time.Minute).Result()
		if err != nil {
			log.Print(err)
			return
		}
		log.Println("The key is not exist")
	}
}

// func OnlineCompilerCodeResponse() {

// }

func RunCode(body map[string]interface{}) (string, error) {
	const url = "https://api.jdoodle.com/v1/execute"

	jsonValue, _ := json.Marshal(body)

	r, err := http.Post(url, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		return "", err
	}

	defer r.Body.Close()

	onlineCompiler := &models.OnlineCompiler{}
	err = json.NewDecoder(r.Body).Decode(onlineCompiler)

	if err != nil {
		return "", err
	}

	if r.StatusCode == http.StatusOK {
		return onlineCompiler.Output, nil
	}

	return "", errors.New("something went wrong")
}
