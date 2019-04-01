package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Req struct {
	User   string `json:"user"`
	Passwd string `json:"passwd"`
}

func main() {
	//get()
	//post()
	get2()
}

func get2() {

	client := &http.Client{}
	url := "http://mi.gdt.qq.com/api/v3"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("http NewRequest err:", err)
		return
	}

	// 设置Headers
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// 设置Cookies
	cookie := http.Cookie{Name: "sessionid", Value: "LSIE89SFLKGHHASLC9EETFBVNOPOXNM"}
	req.AddCookie(&cookie)

	//设置get请求行参数
	query := req.URL.Query()
	query.Add("api_version", "3.0")

	req.URL.RawQuery = query.Encode()
	log.Println("url:", req.URL.String())

	resp, err := client.Do(req)
	if err != nil {
		log.Println("client Do err:", err)
		return
	}

	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("read body err:", err)
		return
	}

	log.Println("content:", string(content))

}

func get() {
	url := "http://mi.gdt.qq.com/api/v3?api_version=3.0&pos=%7B%22id%22%3A1050604134524585%2C%22width%22%3A300%2C%22height%22%3A250%2C%22support_full_screen_interstitial%22%3Atrue%2C%22ad_count%22%3A1%7D&media=%7B%22app_id%22%3A%221104241296%22%2C%22app_bundle_id%22%3A%22com.test.android%22%7D&device=%7B%22os%22%3A%22android%22%2C%22os_version%22%3A%226.0.1%22%2C%22model%22%3A%22MI%204LTE%22%2C%22manufacturer%22%3A%22Xiaomi%22%2C%22device_type%22%3A1%2C%22screen_width%22%3A360%2C%22screen_height%22%3A640%2C%22dpi%22%3A480%2C%22orientation%22%3A0%2C%22imei%22%3A%22123456789009876%22%2C%22android_id%22%3A%229774d56d682e549c%22%2C%22android_ad_id%22%3A%22d725f723-86ea-466a-883d-5be2ca568241%22%7D&network=%7B%22connect_type%22%3A1%2C%22carrier%22%3A3%7D&geo=%7B%22lat%22%3A39925625%2C%22lng%22%3A116333285%2C%22coord_time%22%3A1473416991123%7D"
	resp, err := http.Get(url)
	if err != nil {
		log.Println("http Get err:", err)
		return
	}

	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("read body err:", err)
		return
	}

	log.Println("content:", string(content))
}

func post() {
	url := "http://127.0.0.1:8080/bar"
	contentType := "application/json;charset=utf-8"
	req := Req{User: "bluesgao", Passwd: "99999"}

	b, err := json.Marshal(req)
	if err != nil {
		log.Println("json format err:", err)
		return
	}

	body := bytes.NewBuffer(b)
	resp, err := http.Post(url, contentType, body)
	if err != nil {
		log.Println("http Post err:", err)
		return
	}

	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("read body err:", err)
		return
	}

	log.Println("content:", string(content))

}
