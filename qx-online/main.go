package main

import "log"

func main() {
	InitConfig()

	aes := GlobalConfig.Aes
	iv := aes.Token[:16]
	result, err := EncryptAES(aes.Token, aes.Id, iv)
	if err != nil {
		log.Fatalf("Encrypt AES failed,err:%v", err)
	}
	log.Printf("AES:%s", result)

	md5Conf := GlobalConfig.MD5
	md5Str := GetMD5Encode(md5Conf.Text)
	log.Printf("MD5:%s", md5Str)
}
