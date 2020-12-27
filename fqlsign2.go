package main

//import (
//	"bytes"
//	"crypto/aes"
//	"crypto/cipher"
//	"crypto/des"
//	"crypto/sha1"
//	"encoding/base64"
//	"fmt"
//	//_ "github.com/caddyserver/caddy/v2/modules/caddyhttp"
//	// plug in Caddy modules here
//	//_ "github.com/caddyserver/caddy/v2/modules/standard"
//)
//
//func main() {
//	text := "a=1&c=2"
//	//ces, err := AesEcbpkCes([]byte(text), []byte("93be37a09b1ff3da2d8b86d129db9b35"))
//	//fmt.Println(ces)
//	//fmt.Println(err)
//
//	prng, err2 := AESSHA1PRNG( []byte("93be37a09b1ff3da2d8b86d129db9b35"), 128)
//
//	ecb := EncryptDES_ECB([]byte(text), prng)
//	fmt.Println(prng)
//	fmt.Println(err2)
//	fmt.Println(ecb)
//
//
//}
//
//
//func AesEcbpkCes(plaintext, key []byte)  (string,error){
//
//	prng, err2 := AESSHA1PRNG(key, 128)
//	fmt.Println(prng)
//	fmt.Println(err2)
//
//	//block,err := aes.NewCipher(key)
//	block,err := aes.NewCipher(prng)
//	if err!=nil {
//		fmt.Println(err)
//	}
//	blockSize := block.BlockSize()
//	plaintext = ZeroPadding(plaintext,blockSize)
//	ciphertext := make([]byte,len(plaintext))
//
//	iv := []byte("000000000000000000000000000000000000000000000000")[:blockSize]
//
//	//mode := cipher.NewCBCEncrypter(block,iv)
//	mode := cipher.NewCBCEncrypter(block,iv)
//	mode.CryptBlocks(ciphertext,plaintext)
//
//	return base64.StdEncoding.EncodeToString(ciphertext),nil
//}
//
//
//
//func ZeroPadding(ciphertext []byte, blockSize int) []byte  {
//	padding := blockSize - len(ciphertext)%blockSize
//	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
//	return append(ciphertext,padtext...)
//}
//
//
//// 模拟 Java 接口 generatorKey()
//// 目前 encryptLength 仅支持 128bit 长度
//// 因为 SHA1() 返回的长度固定为 20byte 160bit
//// 所以 encryptLength 超过这个长度，就无法生成了
//// 因为不知道 java 中 AES SHA1PRNG 的生成逻辑
//func AESSHA1PRNG(keyBytes []byte,  encryptLength int) ([]byte, error) {
//	hashs := SHA1(SHA1(keyBytes))
//	maxLen := len(hashs)
//	realLen := encryptLength/8
//	if realLen > maxLen {
//		return nil, nil
//	}
//
//	return hashs[0:realLen], nil
//}
//
//func SHA1(data []byte) []byte {
//	h := sha1.New()
//	h.Write(data)
//	return h.Sum(nil)
//}
//
//
//
//func EncryptDES_ECB(src, key []byte) string {
//	//data := []byte(src)
//	//keyByte := []byte(key)
//	block, err := des.NewCipher(key)
//	if err != nil {
//		panic(err)
//	}
//	bs := block.BlockSize()
//	//对明文数据进行补码
//	data := PKCS5Padding(src, bs)
//	if len(data)%bs != 0 {
//		panic("Need a multiple of the blocksize")
//	}
//	out := make([]byte, len(data))
//	dst := out
//	for len(data) > 0 {
//		//对明文按照blocksize进行分块加密
//		//必要时可以使用go关键字进行并行加密
//		block.Encrypt(dst, data[:bs])
//		data = data[bs:]
//		dst = dst[bs:]
//	}
//	return fmt.Sprintf("%X", out)
//}
//
//
//func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
//	padding := blockSize - len(ciphertext)%blockSize
//	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
//	return append(ciphertext, padtext...)
//}