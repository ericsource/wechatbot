package main

import (
	"fmt"
	"github.com/869413421/wechatbot/gtp"
	_ "github.com/869413421/wechatbot/gtp"
	"github.com/869413421/wechatbot/handlers"
	_ "github.com/869413421/wechatbot/handlers"
	"github.com/gin-gonic/gin"
	"github.com/silenceper/wechat/v2/officialaccount"
	"github.com/silenceper/wechat/v2/officialaccount/message"
	"strings"
)

var officialAccount *officialaccount.OfficialAccount

func init() {
	officialAccount = handlers.GetOfficialAccount()
}

func main() {
	r := gin.Default()

	// /official-account/message GET 验证
	r.GET("/official-account/message", func(c *gin.Context) {
		verifyEchoStr, _ := c.GetQuery("echostr")

		verifyMsgSign, _ := c.GetQuery("msg_signature")

		verifyTimestamp, _ := c.GetQuery("timestamp")

		verifyNonce, _ := c.GetQuery("nonce")

		fmt.Println("verifyEchoStr", verifyEchoStr, " verifyMsgSign:", verifyMsgSign, " verifyTimestamp:", verifyTimestamp, " verifyNonce", verifyNonce)

		c.Writer.WriteString(verifyEchoStr)
	})

	// /official-account/message POST 消息接收
	r.POST("/official-account/message", func(c *gin.Context) {

		// 传入request和responseWriter
		server := officialAccount.GetServer(c.Request, c.Writer)

		server.SetMessageHandler(func(msg *message.MixMessage) *message.Reply {
			fmt.Println("openid:", msg.FromUserName, "content:", msg.Content)
			if msg.MsgType == "text" {
				go sendMsg(string(msg.FromUserName), msg.Content)
				return nil
			} else {
				text := message.NewText("目前只支持文本消息")
				return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
			}
		})

		//处理消息接收
		err := server.Serve()
		if err != nil {
			fmt.Println("server serve err:", err)
			return
		}

		// 直接回复success（推荐方式）, 结果异步回复
		c.Writer.WriteString("success")
	})

	r.Run(":12004")
}

func sendMsg(openid string, content string) {
	var replyMsg string = ""
	var err error

	requestText := strings.TrimSpace(content)
	//requestText = strings.Trim(content, "\n")
	replyMsg, err = gtp.ChatCompletions(requestText)
	if err != nil {
		replyMsg = "出错了，请再试一试！"
	}
	replyMsg = strings.Trim(replyMsg, "\n")
	replyMsg = strings.Trim(replyMsg, "\n\n")
	fmt.Println("msg:", replyMsg)
	manager := officialAccount.GetCustomerMessageManager()
	msg := &message.CustomerMessage{
		ToUser:  openid,
		Msgtype: "text",
		Text: &message.MediaText{
			Content: replyMsg,
		},
	}
	err = manager.Send(msg)
	if err != nil {
		println("manager send err:", err.Error())
		return
	}
	println("manager send success")
}
