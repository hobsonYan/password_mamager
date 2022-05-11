package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Recover(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("panic: %v\n", r)

			c.JSON(http.StatusOK, gin.H{
				"code": 500,
				"msg":  error2String(r),
				"data": nil,
			})
			c.Abort()
		}
	}()
	c.Next()
}

func error2String(r interface{}) string {
	switch v := r.(type) {
	case error:
		return v.Error()
	default:
		return r.(string)
	}
}
