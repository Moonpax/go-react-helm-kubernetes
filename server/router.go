package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const spacex_api = "https://api.spacexdata.com/v4/launches/"

// Init gin router
func SetupRouter() *gin.Engine {

	// set gin router in Release mode
	//gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	// Enable CORS
	r.Use(cors.Default())

	// Custom Middleware
	r.Use(Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	// handlers
	r.GET("/", requestHandler)

	return r
}

func requestHandler(c *gin.Context) {

	body, err, status := HTTPGetRequest(spacex_api)
	checkError(err)

	var res interface{}
	err = json.Unmarshal(body, &res)
	checkError(err)

	c.JSON(status, res)

	// pin pong test
	// c.JSON(http.StatusOK, gin.H{
	// 	"message": "pong",
	// })
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// Set example variable
		c.Set("moonpax", "555")

		// before request

		c.Next()

		// after request

		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}
