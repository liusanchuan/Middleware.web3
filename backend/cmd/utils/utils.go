package utils

import (
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
	"log"
	"math/rand"
	"regexp"
	"time"
)

var letterRunes = []rune("1234567890abcdefghijklmnopqrstuvwxyz")
var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))
var letterBytes = []byte("1234567890abcdef")

func RandBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[seededRand.Intn(len(letterBytes))]
	}
	return string(b)
}

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[seededRand.Intn(len(letterRunes))]
	}
	return string(b)
}

func SignHash(data []byte) []byte {
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(data), data)
	return crypto.Keccak256([]byte(msg))
}

func VerifySig(from, sigHex string, msg []byte) bool {
	return VerifySig1(from, sigHex, msg) || VerifySig2(from, sigHex, msg)
}

func VerifySig1(from, sigHex string, msg []byte) bool {
	fromAddr := common.HexToAddress(from)
	sig := hexutil.MustDecode(sigHex)

	if sig[64] != 27 && sig[64] != 28 {
		return false
	}
	sig[64] -= 27

	pubKey, err := crypto.SigToPub(SignHash(msg), sig)
	if err != nil {
		return false
	}

	recoveredAddr := crypto.PubkeyToAddress(*pubKey)

	return fromAddr == recoveredAddr
}

func VerifySig2(from, sigHex string, msg []byte) bool {
	fromAddr := common.HexToAddress(from)

	sig := hexutil.MustDecode(sigHex)

	pubKey, err := crypto.SigToPub(SignHash(msg), sig)
	if err != nil {
		return false
	}

	recoveredAddr := crypto.PubkeyToAddress(*pubKey)

	return fromAddr == recoveredAddr
}

func IsAddressValid(address string) bool {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	return re.MatchString(address)
}

func HashApplySyncMint(addr1, addr2, id, random string) string {
	// "0x935F7770265D0797B621c49A5215849c333Cc3ce"
	// "100000000000000000",
	// "0x4e03657aea45a94fc7d47ba826c8d667c0d1e6e33a64a036ec44f58fa12d6c45"
	types := []string{"address", "address", "uint256", "bytes32"}
	values := []interface{}{
		addr1,
		addr2,
		id,
		random,
	}

	hash := solsha3.SoliditySHA3(types, values)

	fmt.Println(hex.EncodeToString(hash))
	return hex.EncodeToString(hash)
}
func HashApplyCrossTransfer(addr1, addr2, id, receiveChainId, random string) string {
	types := []string{"address", "address", "uint256", "uint256", "bytes32"}
	values := []interface{}{
		addr1,
		addr2,
		id,
		receiveChainId,
		random,
	}

	hash := solsha3.SoliditySHA3(types, values)

	fmt.Println(hex.EncodeToString(hash))
	return hex.EncodeToString(hash)
}

func HashApplyCrossReceive(addr1, addr2, id, senderChainId, random string) string {
	types := []string{"address", "address", "uint256", "uint256", "bytes32"}
	values := []interface{}{
		addr1,
		addr2,
		id,
		senderChainId,
		random,
	}

	hash := solsha3.SoliditySHA3(types, values)

	fmt.Println(hex.EncodeToString(hash))
	return hexutil.Encode(hash)
}

func SigMessage(dataToSign, hexPrivateKey string) string {
	//hexPrivateKey := "0xae78c8b502571dba876742437f8bc78b689cf8518356c0921393d89caaf284ce"
	//dataToSign := "bou"

	privateKey, err := crypto.HexToECDSA(hexPrivateKey[2:])
	if err != nil {
		log.Fatal(err)
	}

	// keccak256 hash of the data
	dataBytes := []byte(dataToSign)
	hashData := crypto.Keccak256Hash(dataBytes)

	signatureBytes, err := crypto.Sign(hashData.Bytes(), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	signature := hexutil.Encode(signatureBytes)

	fmt.Println(signature) // 0xc83d417a3b99535e784a72af0d9772c019c776aa0dfe4313c001a5548f6cf254477f5334c30da59531bb521278edc98f1959009253dda4ee9f63fe5562ead5aa00
	return signature
}
