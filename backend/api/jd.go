package api

import (
	"github.com/gocolly/colly/v2"
)

func StockRemind() {
	c := colly.NewCollector()

	//更新时间
	c.OnHTML("body > div.wrap.clearfix.wea_info > div.left > div.wea_weather.clearfix > strong", func(e *colly.HTMLElement) {
		content := e.Text
		println(content)
	})
	//存储时间
	c.Visit("https://tianqi.moji.com/weather/china/hunan/kaifu-district")
	//将天气信息转为json格式
}
