package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var messages []string

type MessageRequest struct {
	Message string `json:"message"`
}

type MessageResponse struct {
	Messages []string `json:"messages"`
}

func handler(c *gin.Context) {
	c.String(http.StatusOK, "Hello World")
}

func handlerGetMessages(c *gin.Context) {
	c.JSON(http.StatusOK, MessageResponse{
		Messages: messages,
	})
}

func handlerCreateMessage(c *gin.Context) {
	var requestBody MessageRequest
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if requestBody.Message == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Empty message"})
	}

	messages = append(messages, requestBody.Message)

	c.JSON(http.StatusCreated, gin.H{"message": "Message created successfully"})
}

func main() {
	port := "8000"
	gin.DisableConsoleColor()
	r := gin.Default()
	r.GET("/", handler)
	r.GET("/message", handlerGetMessages)
	r.POST("/message", handlerCreateMessage)
	log.Printf("Server is running on http://0.0.0.0:%s\n", port)
	err := r.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalln("issue with server", err)
	}
}
