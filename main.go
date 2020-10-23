package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("Server started")

	register router
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/api/test/post", MyPostHandler).Methods("POST")

    start server listening
    err := http.ListenAndServe(":8080", router)
    if err != nil {
        log.Fatalln("ListenAndServe err:", err)
	}
	
	var lo sync.Mutex
	var nums int
    total := 0
    for i := 1; i <= 10; i++ {
        nums += i
        lo.Lock()
        go func() {
            total += i
            lo.Unlock()
        }()
	}
	log.Printf("sum:%d", nums)
    log.Printf("total:%d", total)

    log.Println("Server end")
}

type Asd struct {
	Title      string `schema:"title,required"`  // must be supplied
	Category   []string `schema:"category,required"`  // must be supplied
	Cancomment string `schema:"comment"`
	Content    string `schema:"content"`
	TotalWords int64 `schema:"-"` //this field is never set
}

func MyPostHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("server start")
	err := r.ParseForm()
	if err != nil {
		fmt.Println("解析表单数据失败!")
	}

	
	fmt.Println(r.Form)
	cate := r.Form["category"]
	fmt.Println(cate)

	var decoder = schema.NewDecoder()
	var asd Asd
	// 将请求参数直接转为struct
	err = decoder.Decode(&asd, r.Form)
	if err != nil {
		fmt.Println("解码表单数据失败!")
		fmt.Println(err)
	}
	fmt.Println(asd)
	fmt.Println("category---", asd.Category)

	for k, v := range asd.Category {
		fmt.Println(k, v)
	}
	var res = map[string]interface{}{"result":"success", "code":200, "data":asd}
	response, _ := json.Marshal(res)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
  }