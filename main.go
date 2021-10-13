package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLFiles("templates/index.html")
	router.Static("/assets", "./assets")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	router.POST("/contactme", func(c *gin.Context) {
		name := c.PostForm("name")
		subject := c.PostForm("subject")
		replyto := c.PostForm("_replyto")
		message := c.PostForm("message")

		res := fmt.Sprintf("name : %s, subject : %s replyto : %s message : %s \n", name, subject, replyto, message)
		file, err := os.OpenFile("responses.txt", os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}
		if _, err := file.WriteString(res); err != nil {
			log.Println(err)
		}
		defer file.Close()

		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	router.Run()
}
