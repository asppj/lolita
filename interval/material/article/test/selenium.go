package main

import (
	"fmt"
	"os"
	"t-mk-opentrace/ext/log-driver/log"
	"time"

	"github.com/tebeka/selenium/chrome"

	"github.com/tebeka/selenium"
)

const (
	path    = `C:\Users\asppj\AppData\Local\Google\Chrome\Application\chromedriver.exe`
	port    = 18080
	timeOut = 10 * time.Second
)

var debug = true

// Browser 浏览器封装 headless
type Browser struct {
	s  *selenium.Service
	wb selenium.WebDriver
}

// NewBrowser 打开chrome
func NewBrowser() (b *Browser, err error) {
	selenium.SetDebug(debug)
	s, err := selenium.NewChromeDriverService(path, port)
	if err != nil {
		return
	}
	caps := selenium.Capabilities{"browserName": "chrome"}
	chromeCaps := chrome.Capabilities{
		Path: "",
		Args: []string{
			// "--headless", // 设置Chrome无头模式，在linux下运行，需要设置这个参数，否则会报错
			// "--no-sandbox",
			"--user-agent=Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Safari/537.36", // 模拟user-agent，防反爬
		},
	}
	caps.AddChrome(chromeCaps)
	// 调起chrome浏览器
	urlPrefix := fmt.Sprintf("http://localhost:%d/wd/hub", port)
	wb, err := selenium.NewRemote(caps, urlPrefix)
	if err != nil {
		fmt.Println("connect to the webDriver faild", err.Error())
		return
	}
	b = &Browser{s: s, wb: wb}
	return
}

// Get 打开链接
func (b *Browser) Get(url string) error {
	if err := b.wb.Get(url); err != nil {
		return err
	}
	if err := b._until(func(wd selenium.WebDriver) (b bool, e error) {
		t, err2 := wd.Title()
		if err2 != nil {
			return false, err2
		}
		log.Info(t)
		return true, nil
	}); err != nil {
		return err
	}
	return nil
}

// Stop 停止chrome headless
func (b *Browser) Stop() error {
	if err := b.wb.Close(); err != nil {
		return err
	}
	return b.s.Stop()
}
func (b *Browser) _until(condition selenium.Condition) error {
	if err := b.wb.WaitWithTimeout(condition, timeOut); err != nil {
		return nil
	}
	return nil
}
func (b *Browser) savePageSource() error {
	body, err := b.wb.PageSource()
	if err != nil {
		return err
	}
	f, err := os.Create("./body.html")
	if err != nil {
		return err
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Error(err)
		}
	}()
	_, err = f.WriteString(body)
	return err

}

// Example 例子
func Example() {
	b, err := NewBrowser()
	if err != nil {
		log.Error(err)
		return
	}
	url := `https://mp.weixin.qq.com/s?__biz=MjM5MzYwNDAxNg==&amp;mid=2247484001&amp;idx=3&amp;sn=7f05d5618d99446d75252e3cad5132a6&amp;chksm=a695cabe91e243a8a89fb134df88399dc700288304aebd75d945c1c64141a4da7caadb46ee51
`
	if err = b.Get(url); err != nil {
		log.Error(err)
		return
	}
	if err = b.savePageSource(); err != nil {
		log.Error(err)
		return
	}
	defer func() {
		if err2 := b.Stop(); err2 != nil {
			log.Error(err2)
		}
	}()

}
