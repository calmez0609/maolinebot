package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
	"database/sql"
	 "github.com/lib/pq" //1.go get github.com/lib/pq //2.export GOPATH=$HOME
	"github.com/line/line-bot-sdk-go/linebot"
)

const (
	// Initialize connection constants.
	HOST     = "172.18.0.10" //固定
	DATABASE = "class_db"    //固定
	USER     = "calmez"      //ch2
	PASSWORD = "dbuser123"   //
)

var bot *linebot.Client
var Reply string
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
func RandomMessage(MessageText string) {
	if MessageText == "你好" {
		Reply = "你好我可以為您提供服務"
	} else if MessageText == "Random" {
		rand.Seed(time.Now().UnixNano())
		Random := rand.Intn(3)
		Reply = names[Random]
	} else {
		var connectionString string = fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=require", HOST, USER, PASSWORD, DATABASE)
		db, err := sql.Open("postgres", connectionString)
		err = db.Ping()
		sql_statement := "INSERT INTO Account (account, password) VALUES ($1,$2);"
		_, err = db.Exec(sql_statement,Message.Text, ,Message.Text)
		Reply = "不懂" + MessageText + "的意思,+正在新增"
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
