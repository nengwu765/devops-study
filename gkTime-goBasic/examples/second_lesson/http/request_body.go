package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/body/once", readBodyOnce)
	http.HandleFunc("/body/isnil", readBodyIsNil)
	http.HandleFunc("/url/query", queryParams)
	http.HandleFunc("/header", header)
	http.HandleFunc("/wholeUrl", wholeUrl)
	http.HandleFunc("/form", form)

	if err := http.ListenAndServe("0.0.0.0:5555", nil); err != nil {
		panic(err)
	}
}

func form(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "before parse form %v\n", request.Form)
	err := request.ParseForm()
	if err != nil {
		fmt.Fprintf(writer, "parse form err %v\n", err)
	}
	fmt.Fprintf(writer, "after parse form %v\n", request.Form)
}

func wholeUrl(writer http.ResponseWriter, request *http.Request) {
	// request.URL有些参数填充不全面，需要实际验证
	data, _ := json.Marshal(request.URL)
	fmt.Fprintf(writer, string(data))
}

func header(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "header is %v\n", request.Header)
}

func queryParams(writer http.ResponseWriter, request *http.Request) {
	values := request.URL.Query()
	fmt.Fprintf(writer, "query is %v\n", values)
}

func readBodyIsNil(writer http.ResponseWriter, request *http.Request) {
	// 原生的GetBody返回是nil
	if request.GetBody != nil {
		fmt.Fprint(writer, "getbody data is not nil")
	} else {
		fmt.Fprint(writer, "getbody data is nil")
	}
}

func readBodyOnce(writer http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		fmt.Fprintf(writer, "read body failed: %v", err)
		return
	}
	fmt.Fprintf(writer, "read the data: %s \n", string(body))

	// 尝试再次读取，啥也读不到，但是也不会报错
	body, err = ioutil.ReadAll(request.Body)
	if err != nil {
		// 不会进来这里
		fmt.Fprintf(writer, "read the data one more time get error: %v", err)
	}
	fmt.Fprintf(writer, "read the data one more time: [%s] and read data length %d \n", string(body), len(body))
}

func home(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Hello go! this is home page")
}


