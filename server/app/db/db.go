package db

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"github.com/go-sql-driver/mysql"
	"log"
	"database/sql"
)

func init(){
	fmt.Println("INIT function called")
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error occured while loading env")
	}
}

type Database struct {
	SqlDb *sql.DB
 }

func Connect()(*sql.DB, error){
	dbuser := os.Getenv("dbuser")
	dbpassword := os.Getenv("dbpassword")
	dbhost := os.Getenv("dbhost")
	dbport := os.Getenv("dbport")
	dbname := os.Getenv("dbname")
	dbAdrss := dbhost + ":" + dbport

	cfg := mysql.Config{
        User:   dbuser,
        Passwd: dbpassword,
        Net:    "tcp",
        Addr:   dbAdrss,
        DBName: dbname,
    }

	var err error
    db, err := sql.Open("mysql", cfg.FormatDSN())
    if err != nil {
        log.Fatal(err)
    }

    pingErr := db.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
    }
    fmt.Println("Connected!")

	return db , err
}

func (pointerToDb *Database) Query() {
	fmt.Println("database queried")
}