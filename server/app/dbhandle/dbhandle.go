package dbhandle

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func init() {
	fmt.Println("INIT function called")
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error occured while loading env")
	}

}

func Connect() (*gorm.DB, error) {

	dbuser := os.Getenv("dbuser")
	dbpassword := os.Getenv("dbpassword")
	dbhost := os.Getenv("dbhost")
	dbport := os.Getenv("dbport")
	dbname := os.Getenv("dbname")

	dsn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	fmt.Println("Db connected")
	sqlDB, err := db.DB() 
	if err != nil {
		fmt.Println("dot db returned error")
	}
    sqlDB.Ping()
	// db.AutoMigrate(&User{})
	// db.Create(&User{Firstname: "Swarn", Lastname: "suvarna",
	// 	Email: "s.swarn619@gmail.com", City: "Mangalore", Age: 30, Status: true})
	return db, err
}
