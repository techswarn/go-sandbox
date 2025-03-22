// In this example we'll look at how to implement
// a _worker pool_ using goroutines and channels.

package main

import (
	"fmt"
//	"time"
	"runtime"
	"log"

	"os"
	"strings"
	"github.com/joho/godotenv"
	"github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/credentials"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)



// Here's the worker, of which we'll run several
// concurrent instances. These workers will receive
// work on the `jobs` channel and send the corresponding
// results on `results`. We'll sleep a second per job to
// simulate an expensive task.
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		res, _ := UploadImage("sample.jpg")
		log.Printf("file uploaded to, %s\n", res)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {

	// In order to use our pool of workers we need to send
	// them work and collect their results. We make 2
	// channels for this.
	const numJobs = 1000
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// This starts up 3 workers, initially blocked
	// because there are no jobs yet.
	threads := runtime.NumCPU()
	fmt.Println(threads)
	for w := 1; w <= 100; w++ {
	//	log.Printf("CPU count: %d \n", runtime.NumCPU())
	//	log.Printf("GO routine count: %d \n", runtime.NumGoroutine())
		go worker(w, jobs, results)
	}

	// Here we send 5 `jobs` and then `close` that
	// channel to indicate that's all the work we have.
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// Finally we collect all the results of the work.
	// This also ensures that the worker goroutines have
	// finished. An alternative way to wait for multiple
	// goroutines is to use a [WaitGroup](waitgroups).
	for a := 1; a <= numJobs; a++ {
		<-results
	}
}





// GetValue returns configuration value based on a given key from the .env file
func GetValue(key string) string {
	//fmt.Println(os.Getenv("GO_ENV"))
	//env := os.Getenv("GO_ENV")
    // load the .env file

	if os.Getenv("GO_ENV") != "PRODUCTION" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalf("Error loading .env file!!\n")
		}
	}

    // return the value based on a given key
	return os.Getenv(key)
}

func UploadImage(filename string) (string , error) {

	var name string = strings.TrimSuffix(filename, ".jpg") 
//	fmt.Println(name)
    key := GetValue("SPACES_KEY")
    secret := GetValue("SPACES_SECRET")
	endpoint := GetValue("SPACES_ENDPOINT")
	// fmt.Println(endpoint)
	// fmt.Println(key)
	// fmt.Println(secret)

    s3Config := &aws.Config{
        Credentials: credentials.NewStaticCredentials(key, secret, ""),
        Endpoint:    aws.String(endpoint),
        Region:      aws.String("nyc3"),
        S3ForcePathStyle: aws.Bool(false), // // Configures to use subdomain/virtual calling format. Depending on your version, alternatively use o.UsePathStyle = false
    }

    newSession := session.New(s3Config)
    s3Client := s3.New(newSession)
	fmt.Printf("%#v", s3Client)
	uploader := s3manager.NewUploader(newSession)

	pwd, _ := os.Getwd()

    filepath := pwd + "/" + filename
    
	f, err := os.Open(filepath)
	defer f.Close()
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	// Upload the file to S3.
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("backend"),
		Key:    aws.String(name),
		Body:   f,
	})
	if err != nil {	
		fmt.Printf("error while uploading photos: %v \n", err)
		return "", err
	}
	//fmt.Printf("Result: %#v", result)
	fmt.Printf("file uploaded to, %s\n", aws.StringValue(&result.Location))
	
	return result.Location, nil
}