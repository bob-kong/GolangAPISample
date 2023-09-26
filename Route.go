package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// a
func RouterEngine() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "OPTIONS", "DELETE"},
		AllowHeaders:    []string{"Origin", "x-api-key", "Authorization", "Referer", `Accept`, `Content-Type`},
		ExposeHeaders:   []string{"Content-Length"},
		MaxAge:          1 * time.Minute,
	}))

	r.GET(`/`, HandlerGetExample)

	Example := r.Group(`Example`)
	{
		Example.POST(`/post`, HandlerExample)
	}

	return r
}
