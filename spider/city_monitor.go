package spider

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/sclevine/agouti"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

var urlCity = `https://map.zq12369.com/#/layer=terrain/item=pm10/overlay=none/orthographic=114.403381,36.859845,8`

func FetchHTMLPage(uri string) ([]byte, error) {
	req, _ := http.NewRequest("GET", uri, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.96 Safari/537.36 Edg/88.0.705.50")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(body))
	return body, err
}

func GetCityData() error {
	text, err := FetchHTMLPage(urlCity)
	if err != nil {
		return err
	}
	// 获取表单token
	r := strings.NewReader(string(text))
	doc, err := goquery.NewDocumentFromReader(r)
	log.Println(doc)

	return nil
}

var UserAgentList = [...]string{
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/535.1 (KHTML, like Gecko) " +
		"Chrome/14.0.835.163 Safari/535.1",
	"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; WOW64; Trident/5.0; " +
		"SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; Media Center PC 6.0; InfoPath.3; .NET4.0C; .NET4.0E)",
	"Mozilla/5.0 (Windows NT 6.1) AppleWebKit/535.1 (KHTML, like Gecko) Chrome/13.0.782.41 Safari/535.1 " +
		"QQBrowser/6.9.11079.201",
}

func getRandUA() string {

	rand.Seed(time.Now().UnixNano())
	rd := rand.Intn(len(UserAgentList) - 1)
	return UserAgentList[rd]
}

func GetDataByJS() error {
	dr := getChromeDriver()

	// driver启动超时时间
	dr.Timeout = time.Second * 20

	// 启动phantomjs
	if err := dr.Start(); err != nil {
		fmt.Printf("failed to start driver, err:%s", err.Error())
		return err
	}

	defer dr.Stop()
	// 打开一个新的页面
	page, err := dr.NewPage()
	if err != nil {
		fmt.Printf("failed to open page, err:%s", err.Error())
		return err
	}

	defer page.CloseWindow()

	// 网页加载超时，单位毫秒
	page.SetPageLoad(10000)

	err = page.Navigate(urlCity)
	if err != nil {
		fmt.Printf("failed to navigate, err:%s ", err.Error())
		return err
	}

	html, err := page.HTML()
	if err != nil {
		fmt.Printf("failed to get html, err:%s ", err.Error())
		return err
	}
	fmt.Println(html)
	//截图
	page.Screenshot("./pm10.png")
	return nil
}

func getPhantomDriver() *agouti.WebDriver {
	capabilities := agouti.NewCapabilities()
	//使用随机的ua头
	capabilities["phantomjs.page.settings.userAgent"] = getRandUA()
	//是否加载图片
	capabilities["phantomjs.page.settings.loadImages"] = false
	capabilities["phantomjs.page.settings.resourceTimeout"] = 30
	capabilities["phantomjs.page.settings.disk-cache"] = false
	//phantomjs 不支持gzip
	capabilities["phantomjs.page.customHeaders.accept-encoding"] = "deflate, br"
	capabilities["phantomjs.page.settings.clearMemoryCaches"] = true

	capabilitiesOption := agouti.Desired(capabilities)

	dr := agouti.PhantomJS(capabilitiesOption)
	return dr
}

func getChromeDriver() *agouti.WebDriver {
	optionList := []string{
		"--headless",
		"--window-size=1000,900",
		"--incognito", //隐身模式
		"--blink-settings=imagesEnabled=true",
		"--no-default-browser-check",
		"--ignore-ssl-errors=true",
		"--ssl-protocol=any",
		"--no-sandbox",
		"--disable-breakpad",
		"--disable-gpu",
		"--disable-logging",
		"--no-zygote",
		"--allow-running-insecure-content",
	}

	userAgent := "--user-agent=" + getRandUA()
	optionList = append(optionList, userAgent)

	dr := agouti.ChromeDriver(
		agouti.ChromeOptions("args", optionList),
		agouti.Desired(
			agouti.Capabilities{
				"loggingPrefs": map[string]string{
					"performance": "ALL",
				},
				"acceptSslCerts":      true,
				"acceptInsecureCerts": true,
			},
		),
	)
	return dr
}
