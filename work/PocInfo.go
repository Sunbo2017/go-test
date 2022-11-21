package main

import (
	"crypto/tls"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var (
	token = ""
	agent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Goby/2.1.0 Chrome/96.0.4664.110 Electron/16.0.6 Safari/537.36"
	//total = 4887/100 + 1
	url = ""
)

type pocResp struct {
	StatusCode int    `json:"statusCode"`
	Messages   string `json:"messages"`
	Data       Data   `json:"data"`
}

type Data struct {
	CurrPage   int   `json:"currPage"`
	PageSize   int   `json:"pageSize"`
	TotalCount int   `json:"totalCount"`
	TotalPage  int   `json:"totalPage"`
	List       []Poc `json:"list"`
}

type Poc struct {
	AuditDate   string    `json:"auditDate"`
	Cats        string    `json:"cats"`
	CreateDate  time.Time `json:"createDate"`
	FileName    string    `json:"fileName"`
	FofaQuery   string    `json:"fofaQuery"`
	FofaRecords int64     `json:"fofaRecords"`
	HasExp      bool      `json:"hasExp"`
	Homepage    string    `json:"homepage"`
	Ids         []string  `json:"ids"`
	Is0day      bool      `json:"is0Day"`
	Level       int       `json:"level"`
	Name        string    `json:"name"`
	ParentCats  string    `json:"parentCats"`
	PocId       string    `json:"pocId"`
	Product     string    `json:"product"`
	Tags        []string  `json:"tags"`
	UpdateDate  time.Time `json:"updateDate"`
	VulIds      []string  `json:"vulIds"`
}

func GetPocList(pageNum int) string {

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	paramMap := map[string]int{
		"page":  pageNum,
		"limit": 20,
	}
	// json.Marshal
	reqParam, err := json.Marshal(&paramMap)
	if err != nil {
		log.Error("Marshal RequestParam fail, err:%v", err)
		panic(err)
	}

	// 准备: HTTP请求
	reqBody := strings.NewReader(string(reqParam))
	reqest, _ := http.NewRequest("POST", url, reqBody)

	//增加header选项
	reqest.Header.Add("User-Agent", agent)
	reqest.Header.Add("Authorization", token)

	// DO: HTTP请求
	httpRsp, err := client.Do(reqest)
	if err != nil {
		log.Error("do http fail, url: %s, reqBody: %s, err:%v", url, reqBody, err)
		panic(err)
	}
	defer httpRsp.Body.Close()

	// Read: HTTP结果
	rspBody, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		log.Error("ReadAll failed, url: %s, reqBody: %s, err: %v", url, reqBody, err)
		panic(err)
	}

	// unmarshal: 解析HTTP返回的结果
	// 		body: {"Result":{"RequestId":"12131","HasError":true,"ResponseItems":{"ErrorMsg":"错误信息"}}}
	result := pocResp{}
	if err = json.Unmarshal(rspBody, &result); err != nil {
		log.Error("Unmarshal fail, err:%v", err)
		panic(err)
	}

	log.Println(result)

	//data := result.Data

	return ""
}

func main() {

	GetPocList(1)

}
