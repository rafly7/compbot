package services

import (
	"bytes"
	"compbot/models"
	"compbot/utils"
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
	res, err := cb()
	if err != nil {
		utils.Recover(err)
		return
	}
	val, err := rdb.Get(context.Background(), infoSender).Result()
	if err != redis.Nil { // IF the key is exist and not error
		log.Println("The key is exist")
		if err != nil {
			utils.Recover(err)
			return
		}
		val, err := rdb.Del(context.Background(), infoSender).Result()
		if err != nil {
			utils.Recover(err)
			return
		}
		if val == 1 {
			_, err := rdb.SetNX(context.Background(), infoSender, res.ID, 5*time.Minute).Result()
			if err != nil {
				utils.Recover(err)
				return
			}
		}
	} else if err == redis.Nil && val == "" {
		_, err := rdb.SetNX(context.Background(), infoSender, res.ID, 5*time.Minute).Result()
		if err != nil {
			utils.Recover(err)
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
		utils.Recover(err)
		return "", err
	}

	defer r.Body.Close()

	onlineCompiler := &models.OnlineCompiler{}
	err = json.NewDecoder(r.Body).Decode(onlineCompiler)

	if err != nil {
		utils.Recover(err)
		return "", err
	}

	if r.StatusCode == http.StatusOK {
		return onlineCompiler.Output, nil
	}

	return "", errors.New("something went wrong")
}
