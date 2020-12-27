/*
描述 :  golang  AES/ECB/PKCS5  加密解密
*/

package main

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"github.com/forgoer/openssl"
)

func main() {
	src := []byte("a=1&c=2")
	key := []byte("93be37a09b1ff3da2d8b86d129db9b35")

	//加密
	prng, err := AESSHA1PRNG(key, 128)
	fmt.Println(prng)
	fmt.Println(err)
	dst , _ := openssl.AesECBEncrypt(src, prng, openssl.PKCS5_PADDING)
	//fmt.Println(dst)
	fmt.Printf(base64.StdEncoding.EncodeToString(dst))  // iA2GV3g4A4tr9uxZr/Wegw==

	//解密
	src2 := "iA2GV3g4A4tr9uxZr/Wegw=="
	decodeString, err := base64.StdEncoding.DecodeString(src2)
	//fmt.Println(decodeString)
	result , _ := openssl.AesECBDecrypt(decodeString, prng, openssl.PKCS5_PADDING)
	fmt.Println("\n解密结果:")
	fmt.Println(string(result))

}

// 模拟 Java 接口 generatorKey()
// 目前 encryptLength 仅支持 128bit 长度
// 因为 SHA1() 返回的长度固定为 20byte 160bit
// 所以 encryptLength 超过这个长度，就无法生成了
// 因为不知道 java 中 AES SHA1PRNG 的生成逻辑
func AESSHA1PRNG(keyBytes []byte,  encryptLength int) ([]byte, error) {
	hashs := SHA1(SHA1(keyBytes))
	maxLen := len(hashs)
	realLen := encryptLength/8
	if realLen > maxLen {
		return nil, nil
	}

	return hashs[0:realLen], nil
}

func SHA1(data []byte) []byte {
	h := sha1.New()
	h.Write(data)
	return h.Sum(nil)
}
