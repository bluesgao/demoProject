package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type LoginInfo struct {
	User   string `json:"user"`
	Passwd string `json:"passwd"`
}

func main() {
	http.HandleFunc("/", welcome)
	http.HandleFunc("/login", login)
	fmt.Println("服务启动...")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("服务端错误 err:", err)
	}
	fmt.Println("服务关闭...")
}

func login(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("login request:", request)

	user := request.FormValue("user")
	passwd := request.FormValue("passwd")

	logininfo := LoginInfo{User: user, Passwd: passwd}

	fmt.Printf("logininfo:%+v \n", logininfo)

	if len(user) == 0 || len(passwd) == 0 {
		fmt.Fprintf(writer, "%s", "用户名和密码不能为空")
	} else {
		b, _ := json.Marshal(logininfo)
		fmt.Fprintf(writer, "欢迎%s", b)
	}
}

func welcome(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "%s", "首页")
}
