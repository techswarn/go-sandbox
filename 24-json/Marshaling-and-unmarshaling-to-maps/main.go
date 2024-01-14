package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	b,err := os.ReadFile("data.json")
	if err != nil {
		panic(err)
	}
	var data map[string]interface{}

	json.Unmarshal([]byte(b), &data)

	_, ok := data["name"]
	if !ok {
		fmt.Println("json does not contain key 'name'")
	}

	fmt.Println("Name :", data["name"],
	"\nCountry :", data["country"],
	"\nCity :", data["city"],
	"\nType :", data["type"])
}