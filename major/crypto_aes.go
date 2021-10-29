package major

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

func padding(src []byte, blocksize int) []byte {
	padnum := blocksize - len(src)%blocksize
	pad := bytes.Repeat([]byte{byte(padnum)}, padnum)
	return append(src, pad...)
}

func unpadding(src []byte) []byte {
	defer func() {
		if r := recover(); r != nil {
			// log.Debug("Aes解密失败", 0, log.String("src", string(src)))
		}
	}()
	n := len(src)
	unpadnum := int(src[n-1])
	return src[:n-unpadnum]
}

func AesEncrypt(s string, k string) string {
	if len(s) == 0 || len(k) != 16 {
		return ""
	}
	key := []byte(k)
	src := []byte(s)
	block, _ := aes.NewCipher(key)
	src = padding(src, block.BlockSize())
	blockmode := cipher.NewCBCEncrypter(block, key)
	blockmode.CryptBlocks(src, src)
	return Base64URLEncode(src)
}

func AesDecrypt(s string, k string) string {
	if len(s) == 0 || len(k) != 16 {
		return ""
	}
	key := []byte(k)
	src := Base64URLDecode(s)
	block, _ := aes.NewCipher(key)
	blockmode := cipher.NewCBCDecrypter(block, key)
	blockmode.CryptBlocks(src, src)
	src = unpadding(src)
	return string(src)
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func AesEncrypt1(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = PKCS5Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func AesDecrypt1(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	return origData, nil
}

// url base64 - 正向
func Base64URLEncode(input interface{}) string {
	var dataByte []byte
	if v, b := input.(string); b {
		dataByte = []byte(v)
	} else if v, b := input.([]byte); b {
		dataByte = v
	} else {
		return ""
	}
	if dataByte == nil || len(dataByte) == 0 {
		return ""
	}
	return base64.URLEncoding.EncodeToString(dataByte)
}

// url base64 - 逆向
func Base64URLDecode(input interface{}) []byte {
	dataStr := ""
	if v, b := input.(string); b {
		dataStr = v
	} else if v, b := input.([]byte); b {
		dataStr = string(v)
	} else {
		return nil
	}
	if len(dataStr) == 0 {
		return nil
	}
	if r, err := base64.URLEncoding.DecodeString(dataStr); err != nil {
		return nil
	} else {
		return r
	}
}
