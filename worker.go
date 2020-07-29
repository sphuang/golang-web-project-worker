package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RootFunc(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"Happy": "Yes"})
}

func LilyFunc(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"Lily": "被信博加加"})
}

func ZooFunc(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"Day": "2020/8/2", "Weather": "nice"})
}

func AppleFunc(c *gin.Context) {
	// c.JSON(http.StatusOK, gin.H{"Say": "喜歡媽媽蛋跟爸爸蛋"})
	c.JSON(http.StatusOK, gin.H{"Say": appleWords})
}

type AppleBody struct {
	Message string `json:"message"`
}

var appleWords []string

func ApplePost(c *gin.Context) {
	var body AppleBody
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"messgae": err.Error()})
		return
	}
	appleWords = append(appleWords, body.Message)
	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}

func main() {
	router := gin.Default() // with Logger() and Recover() middleware
	router.GET("/", RootFunc)
	router.GET("/lily", LilyFunc)
	router.GET("/zoo", ZooFunc)
	router.GET("/apple", AppleFunc)
	router.POST("/apple", ApplePost)
	router.Run()
}
