package handlers

import (
	"github.com/869413421/wechatbot/config"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/officialaccount"
	offConfig "github.com/silenceper/wechat/v2/officialaccount/config"
)

var officialAccount *officialaccount.OfficialAccount

func init() {
	// 使用redis保存access_token，也可选择redis或自定义cache
	wc := wechat.NewWechat()
	memory := cache.NewMemory()
	cfg := &offConfig.Config{
		AppID:          config.LoadConfig().AppID,
		AppSecret:      config.LoadConfig().AppSecret,
		Token:          config.LoadConfig().Token,
		EncodingAESKey: config.LoadConfig().EncodingAESKey,
		Cache:          memory,
	}
	officialAccount = wc.GetOfficialAccount(cfg)
}

func GetOfficialAccount() *officialaccount.OfficialAccount {
	return officialAccount
}
