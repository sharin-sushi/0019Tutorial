package controller

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"net/http"

	"github.com/gin-gonic/gin"
)

//  https://github.com/sharin-sushi/0015Laboratoryn にて

// /aes
func aesHome(c *gin.Context) { //"/"
	c.HTML(http.StatusOK, "test4_aes.html", nil)
}

var text string
var key string

func toAes(c *gin.Context) { //"/to_aes"
	text := c.PostForm("user_id")
	key := c.PostForm("key")

	enxryptedPassword, _, _ := Encrypt(text, key)
	c.HTML(http.StatusOK, "test3.html", gin.H{"thTw": enxryptedPassword})
}
func deAes(c *gin.Context) { //"/de_aes"
	text := c.PostForm("user_id")
	key := c.PostForm("key")
	iv := c.PostForm("iv")
	decryptedPassword, _ := Decrypt(text, key, iv)
	c.HTML(http.StatusOK, "test3.html", gin.H{"thTw": decryptedPassword})
}

//////////////////////////////////////////////////////

// 暗号化
func GenerateIV() ([]byte, error) {
	iv := make([]byte, aes.BlockSize)
	if _, err := rand.Read(iv); err != nil {
		return nil, err
	}
	return iv, nil
}

func Pkcs7Pad(data []byte) []byte {
	length := aes.BlockSize - (len(data) % aes.BlockSize)
	trailing := bytes.Repeat([]byte{byte(length)}, length)
	return append(data, trailing...)
}

func Encrypt(dataString, keyString string) (iv []byte, encrypted []byte, err error) {
	keyByte, _ := hex.DecodeString(keyString)
	data, _ := hex.DecodeString(dataString)

	iv, err = GenerateIV()
	if err != nil {
		return nil, nil, err
	}
	block, err := aes.NewCipher(keyByte)
	if err != nil {
		return nil, nil, err
	}
	padded := Pkcs7Pad(data)
	encrypted = make([]byte, len(padded))
	cbcEncrypter := cipher.NewCBCEncrypter(block, iv)
	cbcEncrypter.CryptBlocks(encrypted, padded)
	return iv, encrypted, nil
}

// 複合化
func Pkcs7Unpad(data []byte) []byte {
	dataLength := len(data)
	padLength := int(data[dataLength-1])
	return data[:dataLength-padLength]
}

func Decrypt(data, keyString, iv string) ([]byte, error) {
	byteKey, _ := hex.DecodeString(key)
	byteIv, _ := hex.DecodeString(iv)
	block, err := aes.NewCipher(byteKey)
	if err != nil {
		return nil, err
	}
	decrypted := make([]byte, len(data))
	dataByte := //いい感じのbyteに変換する
	cbcDecrypter := cipher.NewCBCDecrypter(block, byteIv)
	cbcDecrypter.CryptBlocks(decrypted, data)
	return Pkcs7Unpad(decrypted), nil
}
