package pdfparser

import (
	"fmt"
	"os"
	"io/ioutil"
	"github.com/unidoc/unipdf/v3/extractor"
	pdf "github.com/unidoc/unipdf/v3/model" // not free,need to buy license
)

// PdfReader pdf struct
type PdfReader struct {
	PdfPath  string
	TextPath string
	AllPage  int
	CurPage  int
	Content  string
	File     *os.File
	Size	int64
}

// Load load pdf
func (p *PdfReader) Load() {
	fi, err := os.Stat(p.PdfPath)
	if err == nil {
		p.Size = fi.Size()/1024/1024
		fmt.Printf("the size of pdf file %s: %v M \n", p.PdfPath, p.Size)
	}
	
	// p.AllPage = file.NumPage()
	// fmt.Printf("the pdf file has total %v pages \n", p.AllPage)
}

// ReadCurrentPageContent read the content of current page
func (p *PdfReader) ReadCurrentPageContent() {
	file, err := os.Open(p.PdfPath)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	p.File = file
	pdfReader, err := pdf.NewPdfReader(file)
	if err != nil {
		panic(err)
	}
	page, err := pdfReader.GetPage(p.CurPage)
	//导出文本
	extract, err := extractor.New(page)
	if err != nil {
		panic(err)
	}
	text,err := extract.ExtractText()
	if err != nil{
		panic(err)
	}
	p.Content = text
	fmt.Println(p.Content)
}

// WriteContent write the content to text file
func (p *PdfReader) WriteContent() {
	ioutil.WriteFile(p.TextPath, []byte(p.Content), 0655);
	fmt.Println("the content has been written into " + p.TextPath)
}
