package test

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"

	"gopkg.in/yaml.v2"
)

func TestBase64(t *testing.T) {
	input := []byte("hello world")

	// 演示base64编码
	encodeString := base64.StdEncoding.EncodeToString(input)
	fmt.Println(encodeString)

	// 对上面的编码结果进行base64解码
	decodeBytes, err := base64.StdEncoding.DecodeString(encodeString)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(decodeBytes))

	fmt.Println("-----------------------------------")

	str := "EQKPyQEAAQEBAQABAQEAAAAAAAAAAGoAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAPbuuwcHAQAAAAAAAAAAAAAAAAAAAAAAAQEBAQABAQEAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAFAGU9EBAAAAAAAAAAAAAAAAAAAAAAAAAQEBAQABAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAALzebnkOAAAAAAAAAAAAAAAAAAAAAAAAHwAyMDE2MjkAAFDDAABJDAAAwCcJAEkMAAADAIEIAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=="
	decodeBytes, err = base64.StdEncoding.DecodeString(str)
	fmt.Println(decodeBytes)
	if err != nil {
		log.Fatalln(err)
	}
	// fmt.Println(string(decodeBytes))
	smartBytes := decodeBytes[2:362]
	fmt.Println(smartBytes)
	var attrBytes []byte
	attrMap := make(map[byte][]byte)
	for i, v := range smartBytes {
		attrBytes = append(attrBytes, v)
		if (i+1)%12 == 0 {
			fmt.Println(attrBytes)
			attrMap[attrBytes[0]] = attrBytes
			attrBytes = []byte{}
		}
	}
	fmt.Println(attrMap)
}

func TestStringConv(t *testing.T) {
	vendorSpecific := "1,0,5,50,0,100,100,0,0,0,0,0,0,0,9,50,0,100,100,188,18,0,0,0,0,0,12,50,0,100,100,66,3,0,0,0,0,0,165,50,0,100,100,131,5,0,0,0,0,0,166,50,0,100,100,7,0,0,0,0,0,0,167,50,0,100,100,0,0,0,0,0,0,0,168,50,0,100,100,21,0,0,0,0,0,0,169,50,0,100,100,172,0,0,0,0,0,0,170,50,0,100,100,0,0,0,0,0,0,0,171,50,0,100,100,0,0,0,0,0,0,0,172,50,0,100,100,0,0,0,0,0,0,0,173,50,0,100,100,7,0,0,0,0,0,0,174,50,0,100,100,6,0,0,0,0,0,0,184,50,0,100,100,0,0,0,0,0,0,0,187,50,0,100,100,9,0,0,0,0,0,0,188,50,0,100,100,0,0,0,0,0,0,0,194,34,0,55,56,45,0,18,0,56,0,0,199,50,0,100,100,0,0,0,0,0,0,0,230,50,0,100,100,3,4,40,1,3,4,0,232,51,0,100,100,100,0,0,0,0,0,4,233,50,0,100,100,212,6,0,0,0,0,0,234,50,0,100,100,4,42,0,0,0,0,0,241,48,0,100,100,233,14,0,0,0,0,0,242,48,0,100,100,107,8,0,0,0,0,0,244,50,0,0,100,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0"
	vendorBytes := strings.Split(vendorSpecific, ",")
	smartBytes := vendorBytes[2:362]
	fmt.Println(smartBytes)
	var attrBytes []byte
	attrMap := make(map[byte][]byte)
	for i, v := range smartBytes {
		val, _ := strconv.Atoi(v)
		attrBytes = append(attrBytes, byte(val))
		if (i+1)%12 == 0 {
			fmt.Println(attrBytes)
			attrMap[attrBytes[0]] = attrBytes
			attrBytes = []byte{}
		}
	}
	delete(attrMap, 0)
	fmt.Println(len(attrMap))
	fmt.Println(attrMap)
}

func TestAttrYaml(t *testing.T) {
	var attrMap map[byte]string

	f, err := os.Open("attribute.yaml")
	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()
	doc := yaml.NewDecoder(f)

	if err := doc.Decode(&attrMap); err != nil {
		fmt.Println(err)
	}

	fmt.Println(attrMap)
}
