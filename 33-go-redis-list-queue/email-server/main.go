package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
    "github.com/joho/godotenv"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"os"
)

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

func init() {

	client = Client()

	err := client.Ping(context.Background()).Err()
	if err != nil {
		log.Fatalf("failed to connect to redis. error message - %v", err)
	}

	log.Println("successfully connected to redis")
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

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "application is ready")
	}).Methods(http.MethodGet)

	r.HandleFunc("/", send).Methods(http.MethodPost)

	log.Println("started HTTP server....")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func send(w http.ResponseWriter, req *http.Request) {
	var email Email
	err := json.NewDecoder(req.Body).Decode(&email)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("recieved email request", email)

	jobID := strconv.Itoa(rand.Intn(1000) + 1)

	jobInfo := JobInfo{Email: email, JobId: jobID}
	job, err := json.Marshal(jobInfo)
	if err != nil {
		log.Fatal(err)
	}

	err = client.LPush(context.Background(), "jobs", job).Err()
	if err != nil {
		log.Fatal("lpush issue", err)
	}

	w.Header().Add("jobid", jobID)

}

type Email struct {
	To      string `json:"to"`
	Message string `json:"message"`
}

type JobInfo struct {
	Email Email  `json:"email"`
	JobId string `json:"id"`
}