package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

//位置
type Pos struct {
	Id      int64 `json:"id"`
	Width   int   `json:"width"`
	Height  int   `json:"height"`
	AdCount int   `json:"ad_count"`
}

//媒体
type Media struct {
	AppId       string `json:"app_id"`
	AppBundleId string `json:"app_bundle_id"`
}

//设备
type Device struct {
	Os           string `json:"os"`
	OsVersion    string `json:"os_version"`
	Model        string `json:"model"`
	Manufacturer string `json:"manufacturer"`
	DeviceType   int    `json:"device_type"`
	Idfa         string `json:"idfa"`
	Imei         string `json:"imei"`
	ImeiMd5      string `json:"imei_md5"`
	AndroidId    string `json:"android_id"`
	AndroidIdMd5 string `json:"android_id_md5"`
}

//网络
type NetWork struct {
	ConnectType int `json:"connect_type"`
	Carrier     int `json:"carrier"`
}

func main() {
	get()
}

func setReqUrlParams(req *http.Request) {
	query := req.URL.Query()
	query.Add("api_version", "3.0")

	pos := Pos{
		Id:      1050604134524585,
		Width:   300,
		Height:  250,
		AdCount: 1}
	posb, poserr := json.Marshal(pos)
	if poserr != nil {
		log.Println("pos json format err:", poserr)
		return
	}
	log.Println("pos json:", string(posb))

	query.Add("pos", string(posb))

	media := Media{
		AppId:       "1104241296",
		AppBundleId: "com.test.android"}
	mediab, mediaerr := json.Marshal(media)
	if mediaerr != nil {
		log.Println("media json format err:", mediaerr)
		return
	}
	log.Println("media json:", string(mediab))
	query.Add("media", string(mediab))

	device := Device{
		Os:           "android",
		OsVersion:    "6.0.1",
		Model:        "MI 4LTE",
		Manufacturer: "Xiaomi",
		DeviceType:   1,
		Imei:         "123456789009876",
		AndroidId:    "9774d56d682e549c"}

	deviceb, deviceerr := json.Marshal(device)
	if deviceerr != nil {
		log.Println("device json format err:", deviceerr)
		return
	}
	log.Println("device json:", string(deviceb))
	query.Add("device", string(deviceb))

	network := NetWork{
		ConnectType: 1,
		Carrier:     3}

	networkb, networkerr := json.Marshal(network)
	if networkerr != nil {
		log.Println("network json format err:", networkerr)
		return
	}
	log.Println("network json:", string(networkb))
	query.Add("network", string(networkb))

	//get url 编码
	req.URL.RawQuery = query.Encode()
	log.Println("url:", req.URL.String())

}

func get() {

	client := &http.Client{}
	url := "http://mi.gdt.qq.com/api/v3"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("http NewRequest err:", err)
		return
	}

	// 设置Headers
	//req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// 设置Cookies
	//cookie := http.Cookie{Name: "sessionid", Value: "LSIE89SFLKGHHASLC9EETFBVNOPOXNM"}
	//req.AddCookie(&cookie)

	//设置get请求行参数
	setReqUrlParams(req)

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
