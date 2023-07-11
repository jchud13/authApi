package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type authResponse struct {
    Key     string  `json:"secretKey"`
}

func main() {
    router := gin.Default()
    router.GET("/secretKey/:file", getSecretKeyByFile)

    router.Run("localhost:8080")
}

func getSecretKeyByFile(c *gin.Context) {
	fmt.Println("here")
	retKey := &authResponse{}     
	fileBytes, _ := os.ReadFile("WeatherKey.json")   
	err := json.Unmarshal(fileBytes, retKey)
	fmt.Println(retKey)
	if err != nil {
		fmt.Println(err.Error())
	}
    c.IndentedJSON(http.StatusOK, retKey)
}