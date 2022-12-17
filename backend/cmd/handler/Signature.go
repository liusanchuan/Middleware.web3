package handler

import (
	"backend/cmd/models"
	"backend/cmd/utils"
	"encoding/hex"
	"fmt"
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
	randomStr := "0x" + hex.EncodeToString(random)

	randomStr = "0x" + hex.EncodeToString([]byte{
		38, 169, 229, 89, 164, 1, 181, 53,
		73, 218, 175, 232, 133, 246, 195, 247,
		90, 62, 195, 54, 14, 19, 235, 196,
		38, 245, 57, 139, 66, 2, 100, 27,
	})

	hash := utils.HashApplySyncMint(params.Address1, params.Address2, strconv.Itoa(params.Id), randomStr)
	fmt.Println("hash= ", hash)
	sig := utils.SigMessage(hash, "")

	c.JSON(http.StatusOK, gin.H{
		"id":     params.Id,
		"random": randomStr,
		"sig":    sig,
	})
}

func GetSignatureCrossTransfer(c *gin.Context) {

	params := models.CrossTransfer{}
	err := c.BindJSON(&params)

	if err != nil {
		c.JSON(http.StatusBadRequest, "Request should have signature and publicAddress")
		return
	}
	random := utils.RandBytes(32)
	randomStr := "0x" + hex.EncodeToString(random)
	hash := utils.HashApplyCrossTransfer(params.Sender, params.Receiver, strconv.Itoa(params.Id), strconv.Itoa(params.ReceiveChainId), randomStr)
	fmt.Println("hash= ", hash)
	sig := utils.SigMessage(hash, "")

	c.JSON(http.StatusOK, gin.H{
		"id":     params.Id,
		"random": random,
		"sig":    sig,
	})
}

func GetSignatureCrossReceive(c *gin.Context) {
	params := models.CrossReceive{}
	err := c.BindJSON(&params)

	if err != nil {
		c.JSON(http.StatusBadRequest, "Request should have signature and publicAddress")
		return
	}
	random := utils.RandBytes(32)
	randomStr := "0x" + hex.EncodeToString(random)
	hash := utils.HashApplyCrossReceive(params.Sender, params.Receiver, strconv.Itoa(params.Id), strconv.Itoa(params.SenderChainId), randomStr)
	fmt.Println("hash= ", hash)
	sig := utils.SigMessage(hash, "")

	c.JSON(http.StatusOK, gin.H{
		"id":     params.Id,
		"random": random,
		"sig":    sig,
	})
}
