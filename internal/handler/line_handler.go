package handler

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/maakun12/ramen-search-bot/internal/api"
	_ "github.com/maakun12/ramen-search-bot/internal/api"
)

func LineHandler(w http.ResponseWriter, r *http.Request) {
	secret := os.Getenv("LINE_SECRET")
	accessToken := os.Getenv("LINE_ACCESS_TOKEN")

	bot, err := linebot.New(
		secret,
		accessToken,
	)

	if err != nil {
		log.Fatal(err)
	}

	events, err := bot.ParseRequest(r)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}
	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				// parrot bot
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
					log.Print(err)
				}
			case *linebot.LocationMessage:
				msg := event.Message.(*linebot.LocationMessage)
				lat := strconv.FormatFloat(msg.Latitude, 'f', 2, 64)
				lon := strconv.FormatFloat(msg.Longitude, 'f', 2, 64)

				carouselColumns := api.GetRamenInfo(lat, lon)
				res := linebot.NewTemplateMessage(
					"ラーメン屋一覧",
					linebot.NewCarouselTemplate(carouselColumns...).WithImageOptions("rectangle", "cover"),
				)
				if _, err = bot.ReplyMessage(event.ReplyToken, res).Do(); err != nil {
					log.Print(err)
				}
			}
		}
	}
}
