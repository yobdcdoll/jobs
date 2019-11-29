package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yobdc/jobs/models"
	"log"
	"net/http"
)

func main() {
	log.Println("hello")
	server := gin.Default()
	server.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	t1 := models.NewTask("task1", "i am task1", "")
	t2 := models.NewTask("task2", "i am task2", "")
	t3 := models.NewTask("task3", "i am task3", "")
	t4 := models.NewTask("task4", "i am task4", "")
	t1.AddChild(t2)
	t1.AddChild(t3)
	t1.AddChild(t4)
	t2.AddChild(t4)
	t3.AddChild(t4)
	ti1 := t1.NewInstance()
	ti1.Start()

	server.Run(":8080")
}
