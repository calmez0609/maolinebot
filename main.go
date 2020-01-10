package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
        "math/rand"
        "time"
	"github.com/line/line-bot-sdk-go/linebot"
)

var bot *linebot.Client
var  Reply string
var names = []string{
	"三山國王",
	"耶穌",
	"佛祖",
}
func main() {
	var err error
	bot, err = linebot.New(os.Getenv("ChannelSecret"), os.Getenv("ChannelAccessToken"))
	log.Println("Bot:", bot, " err:", err)
	http.HandleFunc("/callback", callbackHandler)
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)
}
func RandomMessage(MessageText string){
 if MessageText=="你好"{
   Reply:="好三小"
   } else if MessageText=="Random"{ 
	rand.Seed(time.Now().UnixNano())
	Random:=rand.Intn(3)
    Reply:=names[Random]
   }else{
	   Reply:=""
   }
}
func callbackHandler(w http.ResponseWriter, r *http.Request) {
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
				RandomMessage(message.Text)
				if err != nil {
					log.Println("Quota err:", err)
				}
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(Reply)).Do(); err != nil {
					log.Print(err)
				}
			}
		}
	}
}
