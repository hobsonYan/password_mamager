package main

import (
	"password_manager/src/common/route"
	"password_manager/src/handler"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.New()

	r.Use(handler.Recover)

	r.Use(handler.JwtVerify)

	r.Use(handler.Cors())

	r = route.PathRoute(r)

	r.Run(":6991")
}
