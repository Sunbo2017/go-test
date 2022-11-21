package pcworx

import (
	"bufio"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net"
	"os"
)

var (
	//模拟设备session id
	sid        byte = 66
	initComms       = string([]byte{1, 1, 0, 26, 0, 0, 0, 0, 120, 128, 0, 3, 0, 12, 73, 66, 69, 84, 72, 48, 49, 78, 48, 95, 77, 0})
	initComms2      = string([]byte{1, 5, 0, 22, 0, 1, 0, 0, 120, 128, 0, sid, 0, 0, 0, 6, 0, 4, 2, 149, 0, 0})
	reqInfo         = string([]byte{1, 6, 0, 14, 0, 2, 0, 0, 0, 0, 0, sid, 4, 0})
)

func Serve() {
	listen, err := net.Listen("tcp", ":1962")
	if err != nil {
		log.Fatalf("[pcworx] listen failed, err: %v", err)
		return
	}
	for {
		conn, err := listen.Accept() // 监听客户端的连接请求
		if err != nil {
			log.Errorf("[pcworx] Accept failed, err: %v", err)
			continue
		}
		go process(conn) // 启动一个goroutine来处理客户端的连接请求
	}
}

func process(conn net.Conn) {
	defer conn.Close() // 关闭连接

	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:]) // 读取数据
		if err != nil {
			log.Errorf("[pcworx] read from client failed, err: %v", err)
			break
		}
		log.Debugf("[pcworx] receive bytes:%v", buf)
		recvStr := string(buf[:n])
		log.Infof("[pcworx] receive data from client：%v", recvStr)

		//第一次交互
		if recvStr == initComms {
			//lua string.sub(response, 18, 18) 下标从1开始
			resBytes := []byte{129, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, sid, 0, 0}
			log.Debugf("[pcworx] first:%v", string(resBytes))
			_, err := conn.Write(resBytes) // 发送数据
			if err != nil {
				log.Errorf("[pcworx] send data to client failed, err: %v", err)
			}
			continue
		}

		//第二次交互，无用，但必须有结果返回
		if recvStr == initComms2 {
			resBytes := []byte{129, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, sid, 0}
			log.Debugf("[pcworx] second:%v", string(resBytes))
			_, err := conn.Write(resBytes) // 发送数据
			if err != nil {
				log.Errorf("[pcworx] send data to client failed, err: %v", err)
			}
			continue
		}

		//最后一次交互，发送banner
		if recvStr == reqInfo {
			result := makeResult()
			log.Debugf("[pcworx] result:%v", string(result))
			conn.Write(result) // 发送数据
			if err != nil {
				log.Errorf("[pcworx] send data to client failed, err: %v", err)
			}
			break
		}
	}
}

func makeResult() []byte {
	resArr := make([]byte, 165)
	resArr[0] = 129
	b := readBanner("./pcworx/banner.json")
	bannerMap := make(map[string]string)
	err := json.Unmarshal(b, &bannerMap)
	if err != nil {
		log.Errorf("[pcworx] Unmarshal banner failed, err: %v", err)
	}
	plcType := bannerMap["plc_type"]
	modelNum := bannerMap["model_num"]
	version := bannerMap["firmware_version"]
	date := bannerMap["firmware_date"]
	time := bannerMap["firmware_time"]

	//lua index start from 1
	joinResult(resArr, plcType, 31-1)
	joinResult(resArr, modelNum, 153-1)
	joinResult(resArr, version, 67-1)
	joinResult(resArr, date, 80-1)
	joinResult(resArr, time, 92-1)

	return resArr
}

func joinResult(arr []byte, str string, pos int) {
	for i := pos; i < pos+len(str); i++ {
		arr[i] = str[i-pos]
	}
	arr[pos+len(str)] = 0
}

func readBanner(path string) []byte {
	// 打开json文件
	jsonFile, err := os.Open(path)
	defer jsonFile.Close()

	if err != nil {
		log.Fatalf("[pcworx] open file error:%v", err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	log.Debugf("[pcworx] banner info:%v", string(byteValue))

	return byteValue
}
