package telegram

import (
	"crypto/tls"
	"fmt"
	"github.com/869413421/wechatbot/config"
	"github.com/869413421/wechatbot/gtp"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func Run() {
	// 设置代理http或sock5
	proxyUrl, err := url.Parse(config.LoadConfig().TelegramProxy)
	fmt.Println("proxy:", proxyUrl)
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		//使用代理
		Proxy: http.ProxyURL(proxyUrl),
	}
	myClient := &http.Client{Transport: transport}
	APIEndpoint := "https://api.telegram.org/bot%s/%s"
	bot, err := tgbotapi.NewBotAPIWithClient(config.LoadConfig().TelegramToken, APIEndpoint, myClient)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	BotUserName := "@" + bot.Self.UserName

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		// group
		if update.Message.Chat.IsGroup() {
			// 有@再回复
			if strings.Contains(update.Message.Text, BotUserName) {
				//fmt.Println("有at机器人", update.Message.Text)
				if update.Message != nil { // If we got a message
					log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

					requestText := strings.TrimSpace(strings.ReplaceAll(update.Message.Text, " " + BotUserName + " ", ""))
					if requestText == "" {
						requestText = "你好"
					}

					reply, err := gtp.Completions(requestText)
					if err != nil {
						log.Printf("gtp request error: %v \n", err)
						reply = "机器人傻了，请再试一试。"
					}
					if reply != "" {
						msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
						msg.ReplyToMessageID = update.Message.MessageID
						bot.Send(msg)
					}
				}
			}
		}

		// chat
		if update.Message.Chat.IsPrivate() {
			fmt.Println("is private")
			if update.Message != nil { // If we got a message
				log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

				// 向GPT发起请求
				requestText := strings.TrimSpace(update.Message.Text)
				requestText = strings.Trim(update.Message.Text, "\n")
				reply, err := gtp.Completions(requestText)
				if err != nil {
					log.Printf("gtp request error: %v \n", err)
					reply = "机器人傻了，请再试一试。"
				}
				if reply != "" {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
					msg.ReplyToMessageID = update.Message.MessageID
					bot.Send(msg)
				}
			}
		}
	}
}