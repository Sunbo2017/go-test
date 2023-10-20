package main

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

var GlobalConfig Config

type Config struct {
	Aes AesConf `yaml:"aes"`
	MD5 MD5Conf `yaml:"md5"`
}

type AesConf struct {
	Id    string `yaml:"id"`
	Token string `yaml:"token"`
}

type MD5Conf struct {
	Text string `yaml:"text"`
}

func InitConfig() {
	dataBytes, err := os.ReadFile("./config.yaml")
	if err != nil {
		log.Fatalf("读取文件失败：%v", err)
		return
	}
	err = yaml.Unmarshal(dataBytes, &GlobalConfig)
	if err != nil {
		log.Fatalf("read config file to struct err: %s", err)
		return
	}

	log.Println("init config successful")
}
