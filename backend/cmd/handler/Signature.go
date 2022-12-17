package handler

import (
	"backend/cmd/models"
	"backend/cmd/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetSignatureSyncMint(c *gin.Context) {
	params := models.SyncMint{}
	err := c.BindJSON(&params)

	if err != nil {
		c.JSON(http.StatusBadRequest, "Request should have signature and publicAddress")
		return
	}
	random := utils.RandBytes(32)

	sig := utils.HashApplySyncMint(params.Address1, params.Address2, strconv.Itoa(params.Id), random)

}

func GetSignatureCrossTransfer(c *gin.Context) {

	params := models.CrossTransfer{}
	err := c.BindJSON(&params)

	if err != nil {
		c.JSON(http.StatusBadRequest, "Request should have signature and publicAddress")
		return
	}
	random := utils.RandBytes(32)
	sig := utils.HashApplyCrossTransfer(params.Sender, params.Receiver, strconv.Itoa(params.Id), strconv.Itoa(params.ReceiveChainId), random)
}

func GetSignatureCrossReceive(c *gin.Context) {
	params := models.CrossReceive{}
	err := c.BindJSON(&params)

	if err != nil {
		c.JSON(http.StatusBadRequest, "Request should have signature and publicAddress")
		return
	}
	random := utils.RandBytes(32)
	sig := utils.HashApplyCrossReceive(params.Sender, params.Receiver, strconv.Itoa(params.Id), strconv.Itoa(params.SenderChainId), random)
}
