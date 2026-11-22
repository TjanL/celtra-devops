package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

type AccountData struct {
	AccountID string `json:"accountId"`
	Timestamp int64  `json:"timestamp"`
	Data      string `json:"data"`
}

func ready(c *gin.Context) {
	c.String(200, "Ready")
}

func getAccountData(c *gin.Context) {
	accountId := c.Param("accountId")
	now := time.Now().UTC()
	data := AccountData{accountId, now.Unix(), "Celtra DevOps"}
	c.JSON(200, data)
}

func newRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/ready", ready)
	router.GET("/:accountId/data", getAccountData)
	return router
}

func main() {
	router := newRouter()
	if err := router.Run(); err != nil {
		log.Fatal(err)
	}
	router.Run()
}
