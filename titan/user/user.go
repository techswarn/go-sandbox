package user

import (
	"fmt"
	"net/http"
)

func UserHanlder(w http.ResponseWriter, r *http.Request){
	fmt.Println("Create user")
	w.Write([]byte("Created user!\n"))
}