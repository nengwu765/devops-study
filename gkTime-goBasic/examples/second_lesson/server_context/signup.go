package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/signup", SignUpWithoutContext)

	if err := http.ListenAndServe("0.0.0.0:5555", nil); err != nil {
		panic(err)
	}}

type signUpReq struct {
	Email             string `json:"email"`
	Password          string `json:"password"`
	ConfirmedPassword string `json:"confirmed_password"`
}

type commonResponse struct {
	BizCode int         `json:"biz_code"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func SignUpWithoutContext(w http.ResponseWriter, r *http.Request) {
	req := &signUpReq{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "read body failed: %v", err)
		return
	}

	if err = json.Unmarshal(body, req); err != nil {
		fmt.Fprintf(w, "deserialized failed: %v", err)
		return
	}

	// return 一个虚拟的 user id 标识注册成功了
	fmt.Fprintf(w, "%d", 1)
	//fmt.Printf("server err %v \n", err)
}
