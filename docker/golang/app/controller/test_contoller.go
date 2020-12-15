package controller

import (
	"time"

	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context) {
	var today = time.Now().Format("2006-01-02")
	var hello = "Hello, Woeld!"
	c.JSON(200, gin.H{"today": today, "hello": hello})

}
