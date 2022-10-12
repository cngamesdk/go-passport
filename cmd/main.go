package main

import (
	"github.com/cngamesdk/go-gin/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.New()

	g.Use(middleware.Cors())
}
