package main

import (
	"gkTime-devops-httpserver/frame"
	"net/http"
	"os"
	"strings"
)

func main() {
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
