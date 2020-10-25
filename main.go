package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"

	// "github.com/gorilla/mux"
	pdfp "go-test/pdfparser"

	"github.com/gorilla/schema"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("Server started")
	page := flag.Int("page", 0, "the page number of pdf which you want to convert")
	pdfPath := flag.String("pdf", "", "the path of pdf file which you want to convert")
	txtPath := flag.String("txt", "", "the path of tet file which you want to save")

	flag.Parse()

	if *pdfPath == "" {
		panic("Please input the pdf file path!")
	}

	pdf := new(pdfp.PdfReader)
	pdf.PdfPath = *pdfPath
	pdf.TextPath = *txtPath
	pdf.CurPage = *page

	pdf.Load()

	if pdf.CurPage > 0 {
		pdf.ReadCurrentPageContent()
	}

	if pdf.TextPath != "" {
		pdf.WriteContent()
	}

	fmt.Println("pdf parse success")

	// register router
	// router := mux.NewRouter().StrictSlash(true)
	// router.HandleFunc("/api/test/post", MyPostHandler).Methods("POST")

	// // start server listening
	// err := http.ListenAndServe(":8080", router)
	// if err != nil {
	// 	log.Fatalln("ListenAndServe err:", err)
	// }

	// log.Println("Server end")
}

type Asd struct {
	Title      string   `schema:"title,required"`    // must be supplied
	Category   []string `schema:"category,required"` // must be supplied
	Cancomment string   `schema:"comment"`
	Content    string   `schema:"content"`
	TotalWords int64    `schema:"-"` //this field is never set
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
	var res = map[string]interface{}{"result": "success", "code": 200, "data": asd}
	response, _ := json.Marshal(res)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
