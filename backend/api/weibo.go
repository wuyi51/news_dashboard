package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type WeiboHot struct {
	HotWords   []string `json:"hot_words"`
	HotCount   []int    `json:"hot_count"`
	BodyString string   `json:"body_string"`
}

func getWeiboHot() WeiboHot {
	//https://weibo.com/ajax/side/hotSearch
	url := "https://weibo.com/ajax/side/hotSearch"
	client := http.Client{}

	// 发送GET请求
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer resp.Body.Close()

	// 检查响应状态码是否成功（200表示成功）
	bodyString := ""
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Request failed with status:", resp.Status)
	} else {
		// 读取并打印响应体内容
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
		} else {
			fmt.Println("Response Body:", string(bodyBytes))
			bodyString = string(bodyBytes)
		}
	}
	return WeiboHot{
		HotWords:   []string{"热词1", "热词2", "热词3", "热词4", "热词5", "热词6", "热词7", "热词8", "热词9", "热词10"},
		HotCount:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		BodyString: bodyString,
	}
}

var (
	GetWeiboHot = getWeiboHot
)
