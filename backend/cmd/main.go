package main

import (
	"backend/cmd/db"
	"backend/cmd/handler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"net/http"
)

func init() {
	db.InitCollection()
}

func main() {
	r := gin.Default()
	corsCfg := cors.DefaultConfig()
	corsCfg.AllowAllOrigins = true
	corsCfg.AddAllowHeaders("Authorization")
	r.Use(cors.New(corsCfg))

	r.GET("/v1/user/nonce", handler.GetLoginNonceAPI)
	r.POST("/v1/user/auth", handler.MetamaskLoginAPI)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
