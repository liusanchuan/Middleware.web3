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
	r.GET("v1/getNFTs", handler.GetNFTS)
	r.POST("v1/signature/syncMint", handler.GetSignatureSyncMint)
	r.POST("v1/signature/crossTransfer", handler.GetSignatureCrossTransfer)
	r.POST("v1/signature/crossReceive", handler.GetSignatureCrossReceive)
	//r.POST("")

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	//randomStr := "0x" + hex.EncodeToString([]byte{
	//	38, 169, 229, 89, 164, 1, 181, 53,
	//	73, 218, 175, 232, 133, 246, 195, 247,
	//	90, 62, 195, 54, 14, 19, 235, 196,
	//	38, 245, 57, 139, 66, 2, 100, 27,
	//})
	//fmt.Println(randomStr)
	//res := utils.HashApplySyncMint("0x70997970C51812dc3A010C7d01b50e0d17dc79C8", "0x70997970C51812dc3A010C7d01b50e0d17dc79C8", "0", randomStr)
	//fmt.Println(res)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
