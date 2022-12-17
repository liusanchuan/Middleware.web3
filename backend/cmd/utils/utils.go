package utils

import (
	"crypto/ecdsa"
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

func RandBytes(n int) []byte {
	token := make([]byte, n)
	rand.Read(token)
	fmt.Println(token)
	return token
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

	// !! this private key is for test only
	// do not use in Product env
	//Public Key: 0x0408476593202ad2182e3c01624de2770f70a987b834634986632b442530d6634be835697ef8caae26af4f2c7b95f5b9ce96b4120df1f6a6b562922cee314f9bf7
	//Address: 0xa36461eD0cd2d34c20C63cB52Dc3190Aa3b57529

	hexPrivateKey = "0x21b565205d29ff8e3c0c50a0664911b137dd2b476987cbe370e9ab70a1f4a4c6"
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

func GenerateKey() {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	fmt.Println("SAVE BUT DO NOT SHARE THIS (Private Key):", hexutil.Encode(privateKeyBytes))

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println("Public Key:", hexutil.Encode(publicKeyBytes))

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println("Address:", address)
}
