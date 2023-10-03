package model

import (
	"fmt"
	"gorm.io/gorm"
	"server/app/dbhandle"
)

type User struct {
	gorm.Model
	Firstname string `gorm:"unique" json:"firstname"`
	Lastname  string `gorm:"unique" json:"lastname"`
	Email     string `gorm:"unique" json:"email"`
	City      string `json:"city"`
	Age       int    `json:"age"`
	Status    bool   `json:"status"`
}

func CreateSchema() {
	db, err := dbhandle.Connect()
	if err != nil{
		panic("Could not connect to db")
	}
	db.AutoMigrate(&User{})
	res := db.Create(&User{Firstname: "Swarn", Lastname: "suvarna",
		Email: "s.swarn619@gmail.com", City: "Mangalore", Age: 30, Status: true})
	fmt.Println(res)
}
