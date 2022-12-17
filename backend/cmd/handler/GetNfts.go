package handler

import (
	"backend/cmd/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"io"
	"net/http"
)

const (
	GoerliTestApi  = "https://eth-goerli.g.alchemy.com/v2/oZ0Vutjl7Rk4s_D8TgSQz_Q0xbqtByN9"
	PolygonTestApi = "https://polygon-mumbai.g.alchemy.com/v2/geFREsW3WWYg0QrKYneByJoHyrDxmFDF"

	params = "/getNFTs/?owner="
)

func GetNFTS(c *gin.Context) {
	owner, _ := c.GetQuery("owner")

	if owner == "" {
		c.JSON(http.StatusBadRequest, "bad request")
		return
	}
	var nftResp models.NFTResponse
	//get eth testnetwork nfts
	resp, err := http.Get(GoerliTestApi + params + owner)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, "requet GoerliTestApi1 error")
		return
	}
	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	//fmt.Println(body)
	//fmt.Println(len(body))

	var GoerliTestResp models.NFTSingleResponse
	err = json.Unmarshal(body, &GoerliTestResp)
	if err != nil {
		c.JSON(http.StatusBadRequest, "requet GoerliTestApi2 error")
		return
	}
	nftResp.ETHTest = GoerliTestResp

	resp, err = http.Get(PolygonTestApi + params + owner)
	if err != nil {
		c.JSON(http.StatusBadRequest, "requet PolygonTestApi1 error")
		return
	}
	body, err = io.ReadAll(resp.Body)
	defer resp.Body.Close()
	var PolygonTestResp models.NFTSingleResponse
	err = json.Unmarshal(body, &PolygonTestResp)
	if err != nil {
		c.JSON(http.StatusBadRequest, "requet PolygonTestApi2 error")
		return
	}
	nftResp.PolygonTest = PolygonTestResp

	c.JSON(http.StatusOK, gin.H{
		"data": nftResp,
	})
}
