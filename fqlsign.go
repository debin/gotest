package main

//func main() {
//	text := "a=1&c=2"
//	ces, err := AesEcbpkCes([]byte(text), []byte("93be37a09b1ff3da2d8b86d129db9b35"))
//	fmt.Println(ces)
//	fmt.Println(err)
//
//}
//
//
//
//func AesEcbpkCes(plaintext, key []byte)  (string,error){
//	block,err := aes.NewCipher(key)
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