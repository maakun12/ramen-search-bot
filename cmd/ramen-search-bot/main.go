package main

import (
	"net/http"
	"os"

	"github.com/maakun12/ramen-search-bot/internal/handler"
	_ "github.com/maakun12/ramen-search-bot/internal/handler"
)

func main() {

	http.HandleFunc("/callback", handler.LineHandler)

	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
