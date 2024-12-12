package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"
    "github.com/joho/godotenv"
	"github.com/go-redis/redis/v8"
	"os"
	"errors"
)

const emailQueueList = "jobs"
const emailTempQueueList = "jobs-temp"
var client *redis.Client

func Client() *redis.Client {
	fmt.Println(GetValue("REDIS_URL"))
	url := GetValue("REDIS_URL")
    opts, err := redis.ParseURL(url)
    if err != nil {
        panic(err)
    }
    return redis.NewClient(opts)
}

func GetValue(key string) string {
	fmt.Println(os.Getenv("GO_ENV"))
	env := os.Getenv("GO_ENV")
    // load the .env file
	fmt.Printf("The env value is %s \n", env)

	if os.Getenv("GO_ENV") != "PRODUCTION" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalf("Error loading .env file!!\n")
		}
	}

    // return the value based on a given key
	return os.Getenv(key)
}

func main() {
	client := Client()
	_, err := client.Ping(context.Background()).Result()

	if err != nil {
		log.Fatal("ping failed. could not connect", err)
	}
	fmt.Println("reliable consumer ready")
	for {

		val, err := client.BLMove(context.Background(), emailQueueList, emailTempQueueList, "RIGHT", "LEFT", 2*time.Second).Result()

		if err != nil {
			if errors.Is(err, redis.Nil) {
				continue
			}
			log.Println("blmove issue", err)
		}

		job := val
		var jobInfo JobInfo

		err = json.Unmarshal([]byte(job), &jobInfo)
		if err != nil {
			log.Fatal("job info unmarshal issue issue", err)
		}

		fmt.Println("received job", jobInfo.JobId)
		fmt.Println("sending email", jobInfo.Email.Message, "to", jobInfo.Email.To)
	
		go func() {
			err = client.LRem(context.Background(), "jobs-temp", 0, job).Err()
			if err != nil {
				log.Fatal("lrem failed for", job, "error", err)
			}
			log.Println("removed job from temporary list", job)
		}()
	}

}

type Email struct {
	To      string `json:"to"`
	Message string `json:"message"`
}

type JobInfo struct {
	Email Email  `json:"email"`
	JobId string `json:"id"`
}