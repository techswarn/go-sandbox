package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

type photo struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Post string `json:"post"`
}

var photos = []photo{
	{ID:"1", Title:"Golden hour", Post:"Captured during a fine evening at malpe"},
	{ID:"2", Title:"Moyo", Post:"Picture of suraj's cat"},
	{ID:"3", Title:"Boats", Post:"Boats entering malpe harbour"},
}

func main(){
	router := gin.Default()
	router.GET("/photos", getPhotos)
	router.POST("/photos", postPhotos)
	router.Run("localhost:8080")
}

//getPhotos responds with the list of all Photos as JSON.
func getPhotos(c *gin.Context){
	c.IndentedJSON(http.StatusOK, photos)
}
// postPhotos adds an Photos from JSON received in the request body.
func postPhotos(c *gin.Context) {
	var newPhoto photo

	// Call BindJSON to bind the received JSON to
    // newAlbum.
	if err := c.BindJSON(&newPhoto); err != nil {
		return
	}

	//Add new album to slice
	photos = append(photos, newPhoto)
	c.IndentedJSON(http.StatusCreated, newPhoto)
}