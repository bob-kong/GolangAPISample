package main

import "github.com/gin-gonic/gin"

type HTTPError struct {
	Success bool     `json:"Success" example:"false"`
	Errors  []string `json:"Errors,omitempty" example:"Name is not vaild,OTP is not correct"`
}

func NewError(ctx *gin.Context, status int, Errors []string) {
	er := HTTPError{
		Success: status == 200,
		Errors:  Errors,
	}
	ctx.JSON(status, er)
}
