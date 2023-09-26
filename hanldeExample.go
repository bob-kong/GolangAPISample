package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandlerExample(c *gin.Context) {
	var req Member
	if err := c.BindJSON(&req); err != nil {
		NewError(c, http.StatusBadRequest, []string{err.Error()})
		return
	}

	data := req.Example()

	c.JSON(http.StatusOK, gin.H{
		`Success`: true,
		`Message`: data,
	})
}

func HandlerGetExample(c *gin.Context) {
	fmt.Println("asd")
}
