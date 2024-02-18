package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func getWeiboHot() {
	//https://weibo.com/ajax/side/hotSearch
	url := "https://weibo.com/ajax/side/hotSearch"
	client := http.Client{}

	// 发送GET请求
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	// 检查响应状态码是否成功（200表示成功）
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Request failed with status:", resp.Status)
	} else {
		// 读取并打印响应体内容
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
		} else {
			fmt.Println("Response Body:", string(bodyBytes))
		}
	}
}

var (
	GetWeiboHot = getWeiboHot
)
