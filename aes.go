package tool

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

// 长度必须为16, 24或者32
//var key = []byte("1234567812345678")

func AesEncrypt(plaintext, key []byte) (ciphertext []byte, e error) {
	block, e := aes.NewCipher(key) // 分组秘钥
	if e != nil {
		return
	}
	blockSize := block.BlockSize()                              // 获取秘钥块的长度
	plaintext = pkcs7Padding(plaintext, blockSize)              // 补全码
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize]) // 加密模式
	ciphertext = make([]byte, len(plaintext))                   // 创建数组
	blockMode.CryptBlocks(ciphertext, plaintext)                // 加密
	return
}

func AesDecrypt(ciphertext, key []byte) (plaintext []byte, e error) {
	block, e := aes.NewCipher(key) // 分组秘钥
	if e != nil {
		return
	}
	blockSize := block.BlockSize()                              // 获取秘钥块的长度
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize]) // 加密模式
	plaintext = make([]byte, len(ciphertext))                   // 创建数组
	blockMode.CryptBlocks(plaintext, ciphertext)                // 解密
	plaintext = pkcs7UnPadding(plaintext)                       // 去除补全码
	return
}

func pkcs7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func pkcs7UnPadding(plaintext []byte) []byte {
	length := len(plaintext)
	unpadding := int(plaintext[length-1])
	return plaintext[:(length - unpadding)]
}
