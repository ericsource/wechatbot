package bootstrap

import (
	"fmt"
	"github.com/869413421/wechatbot/handlers"
	"github.com/eatmoreapple/openwechat"
	"log"
	"time"
)

func Run() {
	//bot := openwechat.DefaultBot()
	bot := openwechat.DefaultBot(openwechat.Desktop) // 桌面模式，上面登录不上的可以尝试切换这种模式

	// 注册消息处理函数
	bot.MessageHandler = handlers.Handler
	// 注册登陆二维码回调
	bot.UUIDCallback = openwechat.PrintlnQrcodeUrl

	// 创建热存储容器对象
	//reloadStorage := openwechat.NewJsonFileHotReloadStorage("storage.json")
	reloadStorage := openwechat.NewFileHotReloadStorage("storage.json")
	// 执行热登录
	err := bot.HotLogin(reloadStorage)
	if err != nil {
		if err = bot.Login(); err != nil {
			log.Printf("login error: %v \n", err)
			return
		}
	}

	// 获取登陆的用户
	self, err := bot.GetCurrentUser()
	if err != nil {
		fmt.Println("获取登陆用户错误:", err)
		return
	}
	// 获取所有的好友
	friends, err := self.Friends()
	//fmt.Println(friends, err)
	//指定用户发送心跳
	go func() {
		StartStatus := true
		for true {
			// 只发一次
			if StartStatus {
				for _, friend := range friends{
					//fmt.Println(friend.User.UserName, friend.User.UserName,  friend.User.ID())
					if friend.User.ID() == "787180045" {
						_, _ = friend.SendText("start success")
					}
				}
				StartStatus = false
			}

			time.Sleep(15 * time.Minute)

			for _, friend := range friends{
				//fmt.Println(friend.User.UserName, friend.User.UserName,  friend.User.ID())
				if friend.User.ID() == "787180045" {
					heartSendMsg, err := friend.SendText("1")
					if err != nil {
						fmt.Println("发送心跳错误:", err)
						return
					}
					fmt.Println("发送心跳成功:", heartSendMsg)

				}
			}
		}
	}()

	// 阻塞主goroutine, 直到发生异常或者用户主动退出
	bot.Block()
}
