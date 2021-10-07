package main

import (
	"fmt"
	"github.com/tebeka/selenium"
	"os"
	"time"
)

//const (
//	chromeDriverPath = "/home/qiannw/config/chrome/chromedriver"
//	port             = 8080
//)

func main() {
	// Start a WebDriver server instance
	opts := []selenium.ServiceOption{
		//selenium.StartFrameBuffer(),
		selenium.Output(os.Stderr), // Output debug information to STDERR.
	}
	selenium.SetDebug(true)
	service, err := selenium.NewChromeDriverService(chromeDriverPath, port, opts...)
	if err != nil {
		panic(err) // panic is used only as an example and is not otherwise recommended.
	}
	defer service.Stop()

	// Connect to the WebDriver instance running locally.
	caps := selenium.Capabilities{"browserName": "chrome"}
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		panic(err)
	}
	defer wd.Quit()


	if err:= wd.Get("https://account.xiaomi.com/fe/service/login/password?_locale=zh_CN"); err != nil {
		panic(err)
	}
	// 等待页面加载完成，避免后续查找页面元素报错
	time.Sleep(3 * time.Second)

	// 填充用户名和密码
	account, err := wd.FindElement(selenium.ByName, "account")
	if err != nil {
		panic(err)
	}
	if err = account.SendKeys("xxx"); err != nil {
		panic(err)
	}

	password, err := wd.FindElement(selenium.ByName, "password")
	if err != nil {
		panic(err)
	}
	if err = password.SendKeys("xxx"); err != nil {
		panic(err)
	}

	btn, err := wd.FindElement(selenium.ByCSSSelector, ".mi-button")
	if err != nil {
		panic(err)
	}
	if err := btn.Click(); err != nil {
		panic(err)
	}

	// 根据${appID}，直接打开应用详情，避免页面元素多次点击查找（对应${appID}，研发中台进行维护，可能有多个APP）
	if err:= wd.Get("https://dev.mi.com/distribute/app/2882303761518817554/create?namespaceValue=null&statusType=1&action=update"); err != nil {
		panic(err)
	}



	// 附件上传 -- 待继续调研
	//apk, err := os.Open("/home/qiannw/tmp/android-package/天天天气-3.0.3-12-host-release-105-20210906_1832.apk")




	time.Sleep(10 * time.Minute)

}



