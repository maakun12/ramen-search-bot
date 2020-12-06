package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"unicode/utf8"

	"github.com/line/line-bot-sdk-go/linebot"
)

type response struct {
	Results results `json:"results"`
}

type results struct {
	Shop []shop `json:"shop"`
}

type shop struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Photo   photo  `json:"photo"`
	URLs    urls   `json:"urls"`
}

type photo struct {
	Mobile mobile `json:"mobile"`
}

type mobile struct {
	L string `json:"l"`
}

type urls struct {
	PC string `json:"pc"`
}

func GetRamenInfo(lat, lon string) []*linebot.CarouselColumn {
	apiKey := os.Getenv("HOTPEPPER_API_KEY")
	url := fmt.Sprintf(
		"https://webservice.recruit.co.jp/hotpepper/gourmet/v1/?format=json&key=%s&lat=%s&lng=%s&keyword=ラーメン",
		apiKey, lat, lon)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var data response
	if err := json.Unmarshal(body, &data); err != nil {
		log.Fatal(err)
	}

	var carouselColumns []*linebot.CarouselColumn
	for _, shop := range data.Results.Shop {
		addr := shop.Address
		if utf8.RuneCountInString(addr) > 60 {
			addr = string([]rune(addr)[:60])
		}
		column := linebot.NewCarouselColumn(
			shop.Photo.Mobile.L,
			shop.Name,
			addr,
			linebot.NewURIAction("HotPepperで開く", shop.URLs.PC),
		).WithImageOptions("#FFFFFF")
		carouselColumns = append(carouselColumns, column)
	}
	return carouselColumns
}
