package handler

import (
	"backend/cmd/db"
	"backend/cmd/models"
	"backend/cmd/utils"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"time"
)

const (
	DatabaseName   = "Web3"
	UserCollection = "User"
)

func MetamaskLoginAPI(c *gin.Context) {
	params := models.MetamaskLoginParameter{}
	err := c.BindJSON(&params)

	if err != nil {
		c.JSON(http.StatusBadRequest, "Request should have signature and publicAddress")
		return
	}

	isValid := utils.VerifySig(params.PublicAddress, params.Signature, []byte("I am signing my one-time nonce: "+params.Nonce))
	fmt.Printf("is valid: %v", isValid)
	if !isValid {
		c.JSON(http.StatusUnauthorized, "login failed")
		return
	}

	var user models.User
	MongoUserCollection := db.Client.Database(DatabaseName).Collection(UserCollection)
	err = MongoUserCollection.FindOne(context.TODO(), bson.M{"public_address": params.PublicAddress}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "mongo parse failed")
		return
	}
	//db.Where("public_address = ?", params.PublicAddress).Limit(1).Find(&user)

	if user.Id.Hex() > 0 {
		if user.Nonce != params.Nonce {
			c.JSON(http.StatusUnauthorized, "login failed")
			return
		}

		user.Nonce = utils.RandStringRunes(6)
		db.Save(&user)

		c.JSON(http.StatusOK, MetamaskLoginResponseData{
			AccessToken: utils.GenerateToken(utils.CustomClaims{
				PublicAddress: user.PublicAddress,
				RegisteredClaims: jwt.RegisteredClaims{
					ExpiresAt: jwt.NewNumericDate(time.Unix(time.Now().Unix()+60*60, 0)),
					Issuer:    "middleware.web3",
				},
			}),
		})
		return
	}

	newUser := models.User{PublicAddress: params.PublicAddress, Nonce: utils.RandStringRunes(8)}
	err = db.Create(&newUser).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, "login failed")
		return
	}

	c.JSON(http.StatusOK, MetamaskLoginResponseData{
		AccessToken: utils.GenerateToken(utils.CustomClaims{
			Uid:           newUser.ID,
			PublicAddress: user.PublicAddress,
		}),
	})
}
