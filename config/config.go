package config

import (
	"encoding/json"
	"log"
	"os"
	"sync"
)

// Configuration 项目配置
type Configuration struct {
	// gtp apikey
	ApiKey string `json:"api_key"`
	// 自动通过好友
	AutoPass bool `json:"auto_pass"`

	// telegram 代理
	TelegramProxy string `json:"telegram_proxy"`
	// telegram token
	TelegramToken string `json:"telegram_token"`

	// template
	HealTemplate   string `json:"heal_template"`
	PromptTemplate string `json:"prompt_template"`

	// official account
	AppID          string `json:"app_id"`
	AppSecret      string `json:"app_secret"`
	Token          string `json:"token"`
	EncodingAESKey string `json:"encoding_aes_key"`
}

var config *Configuration
var once sync.Once

// LoadConfig 加载配置
func LoadConfig() *Configuration {
	once.Do(func() {
		// 从文件中读取
		config = &Configuration{}
		f, err := os.Open("config.json")
		if err != nil {
			log.Fatalf("open config err: %v", err)
			return
		}
		defer f.Close()
		encoder := json.NewDecoder(f)
		err = encoder.Decode(config)
		if err != nil {
			log.Fatalf("decode config err: %v", err)
			return
		}

		// 如果环境变量有配置，读取环境变量
		ApiKey := os.Getenv("ApiKey")
		AutoPass := os.Getenv("AutoPass")
		if ApiKey != "" {
			config.ApiKey = ApiKey
		}
		if AutoPass == "true" {
			config.AutoPass = true
		}

		//telegram
		TelegramProxy := os.Getenv("TelegramProxy")
		TelegramToken := os.Getenv("TelegramToken")
		if TelegramProxy != "" {
			config.TelegramProxy = TelegramProxy
		}
		if TelegramToken != "" {
			config.TelegramToken = TelegramToken
		}

		// heal
		HealTemplate := os.Getenv("HealTemplate")
		PromptTemplate := os.Getenv("PromptTemplate")
		if HealTemplate != "" {
			config.HealTemplate = HealTemplate
		}
		if PromptTemplate != "" {
			config.PromptTemplate = PromptTemplate
		}

		// official account
		AppID := os.Getenv("AppID")
		AppSecret := os.Getenv("AppSecret")
		Token := os.Getenv("Token")
		EncodingAESKey := os.Getenv("EncodingAESKey")
		config.AppID = AppID
		config.AppSecret = AppSecret
		config.Token = Token
		config.EncodingAESKey = EncodingAESKey

	})
	return config
}
