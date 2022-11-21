package riak

import (
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"
)

func statsHandler(w http.ResponseWriter, r *http.Request) {
	// 打开json文件
	jsonFile, err := os.Open("./riak/banner.json")

	if err != nil {
		log.Errorf("[riak] open file error:%v", err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	log.Debugf("[riak] banner info:%v", string(byteValue))

	w.WriteHeader(200)
	_, err = w.Write(byteValue)
	if err != nil {
		log.Errorf("[riak] write response failed, err:%v", err)
	}
}

func Serve() {
	http.HandleFunc("/stats", statsHandler)

	log.Fatal(http.ListenAndServe(":8098", nil))
}
