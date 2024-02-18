package api

import (
	"encoding/json"
	"github.com/gocolly/colly/v2"
	"new_dashboard_data/utils"
	"strings"
	"time"
)

// https://dev.qweather.com/docs/finance/subscription/ 每日限制请求1000次
// https://github.com/qwd/LocationList/blob/master/China-City-List-latest.csv 城市列表
// 开福区 101250110
// key
// https://devapi.qweather.com/v7/weather/24h?{查询参数} 逐小时预报 24/72/168
// https://devapi.qweather.com/v7/weather/now?{查询参数} 实况天气
// https://devapi.qweather.com/v7/weather/7d?{查询参数} 7天预报  3/7/10/15/30
// 参数 key location
//https://a.hecdn.net/img/common/icon/202106d/101.png

type CurrentWeather struct {
	City        string        `json:"city"`
	Weather     string        `json:"weather"`
	Icon        string        `json:"icon"`
	Temperature string        `json:"temperature"`
	AirQuality  string        `json:"air_quality"`
	Wind        string        `json:"wind"`
	Humidity    string        `json:"humidity"`
	DaysWeather []daysWeather `json:"days_weather"`
	UpdateTime  string        `json:"update_time"`
	SaveAt      string        `json:"save_at"`
}

type daysWeather struct {
	Weather     string `json:"weather"`
	Icon        string `json:"icon"`
	Temperature string `json:"temperature"`
	AirQuality  string `json:"air_quality"`
	Wind        string `json:"wind"`
	WindLevel   string `json:"wind_level"`
}

func getWeather() CurrentWeather {
	// 检查./data/weather.json数据更新的时间 如果小于10分钟则直接返回
	data, err := utils.ReadFile("./data/weather.json")
	if err == nil {
		var weatherData CurrentWeather
		err := json.Unmarshal(data, &weatherData)
		if err == nil {
			layout := "2006-01-02 15:04:05"
			t, err := time.Parse(layout, weatherData.SaveAt)
			if err == nil {
				timestamp := t.Unix()
				if time.Now().Sub(time.Unix(timestamp, 0)).Minutes() < 10 {
					return weatherData
				}
			}

		}
	}
	currentWeather := CurrentWeather{
		City: "湖南省长沙市开福区",
	}
	c := colly.NewCollector()

	//实时天气情况
	c.OnHTML("body > div:nth-child(5) > div.left > div.forecast.clearfix > ul:nth-child(2) > li:nth-child(2)", func(e *colly.HTMLElement) {
		content := e.Text
		currentWeather.Weather = strings.TrimSpace(content)
	})
	//天气图标
	c.OnHTML("body > div.wrap.clearfix.wea_info > div.left > div.wea_weather.clearfix > span > img", func(e *colly.HTMLElement) {
		content := e.Attr("src")
		currentWeather.Icon = strings.TrimSpace(content)
	})
	//温度
	c.OnHTML("body > div.wrap.clearfix.wea_info > div.left > div.wea_weather.clearfix > em", func(e *colly.HTMLElement) {
		content := e.Text
		currentWeather.Temperature = strings.TrimSpace(content)
	})

	//空气质量
	c.OnHTML("body > div.wrap.clearfix.wea_info > div.left > div.wea_alert.clearfix > ul > li > a > em", func(e *colly.HTMLElement) {
		content := e.Text
		currentWeather.AirQuality = strings.TrimSpace(content)
	})
	//湿度
	c.OnHTML("body > div.wrap.clearfix.wea_info > div.left > div.wea_about.clearfix > span", func(e *colly.HTMLElement) {
		content := e.Text
		currentWeather.Humidity = strings.TrimSpace(content)
	})
	//风向
	c.OnHTML("body > div.wrap.clearfix.wea_info > div.left > div.wea_about.clearfix > em", func(e *colly.HTMLElement) {
		content := e.Text
		currentWeather.Wind = strings.TrimSpace(content)
	})
	//更新时间
	c.OnHTML("body > div.wrap.clearfix.wea_info > div.left > div.wea_weather.clearfix > strong", func(e *colly.HTMLElement) {
		content := e.Text
		currentWeather.UpdateTime = strings.TrimSpace(content)
	})
	//存储时间
	currentWeather.SaveAt = time.Now().Format("2006-01-02 15:04:05")

	//预报
	dayWeather := getDayWeather()

	currentWeather.DaysWeather = dayWeather
	c.Visit("https://tianqi.moji.com/weather/china/hunan/kaifu-district")
	//将天气信息转为json格式
	currentWeatherJson, _ := json.Marshal(currentWeather)
	_ = utils.WriteFile("./data/weather.json", currentWeatherJson)
	return currentWeather
}

func getDayWeather() []daysWeather {
	var dayWeather []daysWeather
	todayWeather := daysWeather{}
	tomorrowWeather := daysWeather{}
	theDayAfterTomorrowWeather := daysWeather{}
	c := colly.NewCollector()
	//今天
	c.OnHTML("body > div:nth-child(5) > div.left > div.forecast.clearfix > ul:nth-child(2) > li:nth-child(2)", func(e *colly.HTMLElement) {
		content := e.Text
		todayWeather.Weather = strings.TrimSpace(content)
	})
	c.OnHTML("body > div:nth-child(5) > div.left > div.forecast.clearfix > ul:nth-child(2) > li:nth-child(3)", func(e *colly.HTMLElement) {
		content := e.Text
		todayWeather.Temperature = strings.TrimSpace(content)
	})
	c.OnHTML("body > div:nth-child(5) > div.left > div.forecast.clearfix > ul:nth-child(2) > li:nth-child(4) > em", func(e *colly.HTMLElement) {
		content := e.Text
		todayWeather.Wind = strings.TrimSpace(content)
	})
	c.OnHTML("body > div:nth-child(5) > div.left > div.forecast.clearfix > ul:nth-child(2) > li:nth-child(2) > span > img", func(e *colly.HTMLElement) {
		content := e.Attr("src")
		todayWeather.Icon = strings.TrimSpace(content)
	})
	c.OnHTML("body > div:nth-child(5) > div.left > div.forecast.clearfix > ul:nth-child(2) > li:nth-child(5) > strong", func(e *colly.HTMLElement) {
		content := e.Text
		todayWeather.AirQuality = strings.TrimSpace(content)
	})
	c.OnHTML("body > div:nth-child(5) > div.left > div.forecast.clearfix > ul:nth-child(2) > li:nth-child(4) > b", func(e *colly.HTMLElement) {
		content := e.Text
		todayWeather.WindLevel = strings.TrimSpace(content)
	})
	//明天
	c.OnHTML("body > div:nth-child(5) > div.left > div.forecast.clearfix > ul:nth-child(3) > li:nth-child(2)", func(e *colly.HTMLElement) {
		content := e.Text
		tomorrowWeather.Weather = strings.TrimSpace(content)
	})
	c.OnHTML("body > div:nth-child(5) > div.left > div.forecast.clearfix > ul:nth-child(3) > li:nth-child(3)", func(e *colly.HTMLElement) {
		content := e.Text
		tomorrowWeather.Temperature = strings.TrimSpace(content)
	})
	c.OnHTML("body > div:nth-child(5) > div.left > div.forecast.clearfix > ul:nth-child(3) > li:nth-child(4) > em", func(e *colly.HTMLElement) {
		content := e.Text
		tomorrowWeather.Wind = strings.TrimSpace(content)
	})
	c.OnHTML("body > div:nth-child(5) > div.left > div.forecast.clearfix > ul:nth-child(3) > li:nth-child(2) > span > img", func(e *colly.HTMLElement) {
		content := e.Attr("src")
		tomorrowWeather.Icon = strings.TrimSpace(content)
	})
	c.OnHTML("body > div:nth-child(5) > div.left > div.forecast.clearfix > ul:nth-child(3) > li:nth-child(5) > strong", func(e *colly.HTMLElement) {
		content := e.Text
		tomorrowWeather.AirQuality = strings.TrimSpace(content)
	})
	c.OnHTML("body > div:nth-child(5) > div.left > div.forecast.clearfix > ul:nth-child(3) > li:nth-child(4) > b", func(e *colly.HTMLElement) {
		content := e.Text
		tomorrowWeather.WindLevel = strings.TrimSpace(content)
	})
	//后天
	c.OnHTML("body > div:nth-child(5) > div.left > div.forecast.clearfix > ul:nth-child(4) > li:nth-child(2)", func(e *colly.HTMLElement) {
		content := e.Text
		theDayAfterTomorrowWeather.Weather = strings.TrimSpace(content)
	})
	c.OnHTML("body > div:nth-child(5) > div.left > div.forecast.clearfix > ul:nth-child(4) > li:nth-child(3)", func(e *colly.HTMLElement) {
		content := e.Text
		theDayAfterTomorrowWeather.Temperature = strings.TrimSpace(content)
	})
	c.OnHTML("body > div:nth-child(5) > div.left > div.forecast.clearfix > ul:nth-child(4) > li:nth-child(4) > em", func(e *colly.HTMLElement) {
		content := e.Text
		theDayAfterTomorrowWeather.Wind = strings.TrimSpace(content)
	})
	c.OnHTML("body > div:nth-child(5) > div.left > div.forecast.clearfix > ul:nth-child(4) > li:nth-child(2) > span > img", func(e *colly.HTMLElement) {
		content := e.Attr("src")
		theDayAfterTomorrowWeather.Icon = strings.TrimSpace(content)
	})
	c.OnHTML("body > div:nth-child(5) > div.left > div.forecast.clearfix > ul:nth-child(4) > li:nth-child(5) > strong", func(e *colly.HTMLElement) {
		content := e.Text
		theDayAfterTomorrowWeather.AirQuality = strings.TrimSpace(content)
	})
	c.OnHTML("body > div:nth-child(5) > div.left > div.forecast.clearfix > ul:nth-child(4) > li:nth-child(4) > b", func(e *colly.HTMLElement) {
		content := e.Text
		theDayAfterTomorrowWeather.WindLevel = strings.TrimSpace(content)
	})
	c.Visit("https://tianqi.moji.com/weather/china/hunan/kaifu-district")
	dayWeather = append(dayWeather, todayWeather)
	dayWeather = append(dayWeather, tomorrowWeather)
	dayWeather = append(dayWeather, theDayAfterTomorrowWeather)
	//fmt.Println(dayWeather)

	return dayWeather
}

var (
	GetWeather = getWeather
)
