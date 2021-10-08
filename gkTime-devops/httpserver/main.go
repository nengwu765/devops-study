package main

import (
	"flag"
	"github.com/golang/glog"
	"gkTime-devops-httpserver/frame"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// 自定义预设的 flag 参数， 直观便于理解
var (
	logLevel    string
	logToStderr bool
)

func init() {
	// frame.LogInfo 的类型是 glog.Level，直接string强制装换获取不到值
	flag.StringVar(&logLevel, "logLevel", strconv.Itoa(int(frame.LogInfo)), "log level")
	flag.BoolVar(&logToStderr, "logToStderr", true, "log output stderr")
}

func main() {

	flag.Parse()
	// glog init 预设的flag参数
	_ = flag.Set("v", logLevel)
	_ = flag.Set("logtostderr", strconv.FormatBool(logToStderr))
	defer glog.Flush()

	core := frame.NewCore()
	// health check
	core.Router.GET("/healthz", healthz)
	// get headers
	core.Router.GET("/getHeader", getHeader)
	// get os version
	core.Router.GET("/osVersion", osVersion)

	if err := core.Run(":5555"); err != nil {
		panic("server panic")
	}
}

func healthz(c *frame.Context) {
	c.String(http.StatusOK, "health check ok!")
}

func getHeader(c *frame.Context) {
	headers := c.Request.Header
	for key, values := range headers {
		c.SetHeader(key, strings.Join(values, ","))
	}
	var result = map[string]interface{}{
		"ok": "get header ok",
	}
	c.JSON(http.StatusOK, result)
}

func osVersion(c *frame.Context) {
	version := os.Getenv("VERSION")
	if version == "" {
		c.String(http.StatusExpectationFailed, "no VERSION environment variable, it is empty")
	} else {
		c.SetHeader("VERSION", version)
		c.FormatString(http.StatusOK, "VERSION environment is %s\n", version)
	}
}
