package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

var himawarLatestJsonUrl = "https://himawari8-dl.nict.go.jp/himawari8/img/FULL_24h/latest.json"

func GetImageId() string {
	response, _ := http.Get(himawarLatestJsonUrl)
	defer response.Body.Close()
	responseData, _ := ioutil.ReadAll(response.Body)
	var m map[string]string
	json.Unmarshal(responseData, &m)

	fmt.Println("未能解析到最后的时间: url=", himawarLatestJsonUrl)
	date := m["date"]
	if date == "" {
		fmt.Println("未能解析到最后的时间: url=", himawarLatestJsonUrl)
		return ""
	}
	imgId := date[11:]
	imgId = strings.ReplaceAll(imgId, ":", "")
	fmt.Println(imgId)
	return imgId
}
