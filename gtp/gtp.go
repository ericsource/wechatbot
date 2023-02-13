package gtp

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/869413421/wechatbot/config"
	"io/ioutil"
	"log"
	"io"
	"net/http"
)

const BASEURL = "https://api.openai.com/v1/"

// ChatGPTResponseBody 请求体
type ChatGPTResponseBody struct {
	ID      string                   `json:"id"`
	Object  string                   `json:"object"`
	Created int                      `json:"created"`
	Model   string                   `json:"model"`
	Choices []map[string]interface{} `json:"choices"`
	Usage   map[string]interface{}   `json:"usage"`
	// 错误处理
	Error   map[string]interface{}   `json:"error"`
}

type ChoiceItem struct {
}

// ChatGPTRequestBody 响应体
type ChatGPTRequestBody struct {
	Model            string  `json:"model"`
	Prompt           string  `json:"prompt"`
	MaxTokens        int     `json:"max_tokens"`
	Temperature      float32 `json:"temperature"`
	TopP             int     `json:"top_p"`
	FrequencyPenalty int     `json:"frequency_penalty"`
	PresencePenalty  int     `json:"presence_penalty"`
}

// Completions gtp文本模型回复
//curl https://api.openai.com/v1/completions
//-H "Content-Type: application/json"
//-H "Authorization: Bearer your chatGPT key"
//-d '{"model": "text-davinci-003", "prompt": "give me good song", "temperature": 0, "max_tokens": 7}'
func Completions(msg string) (string, error) {
	requestBody := ChatGPTRequestBody{
		Model:            "text-davinci-003",
		Prompt:           msg,
		MaxTokens:        2048,
		Temperature:      0.7,
		TopP:             1,
		FrequencyPenalty: 0,
		PresencePenalty:  0,
	}
	requestData, err := json.Marshal(requestBody)

	if err != nil {
		return "", err
	}
	log.Printf("request gtp json string : %v", string(requestData))
	req, err := http.NewRequest("POST", BASEURL+"completions", bytes.NewBuffer(requestData))
	if err != nil {
		return "", err
	}

	apiKey := config.LoadConfig().ApiKey
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	gptResponseBody := &ChatGPTResponseBody{}
	log.Println(string(body))
	err = json.Unmarshal(body, gptResponseBody)
	if err != nil {
		return "", err
	}

	// 错误返回
	if len(gptResponseBody.Error) > 0 {
		if message, ok := gptResponseBody.Error["message"]; ok {
			return "", errors.New(message.(string))
		} else {
			return "", errors.New("ChatGPT server error")
		}
	}

	var reply string
	if len(gptResponseBody.Choices) > 0 {
		for _, v := range gptResponseBody.Choices {
			reply = v["text"].(string)
			break
		}
	}
	log.Printf("gpt response text: %s \n", reply)
	return reply, nil
}


// ImageGenerationRequestBody 图片响应体
type ImageGenerationRequestBody struct {
	Prompt           string  `json:"prompt"`
	N        		 int     `json:"n"`
	Size      		 string `json:"size"`
}
// ImageGenerationResponseBody 请求体
type ImageGenerationResponseBody struct {
	Created int                      `json:"created"`
	Data []map[string]interface{} 	 `json:"data"`
	// 错误处理
	Error   map[string]interface{}   `json:"error"`
}
func ImagesGenerations(prompt string) (file io.Reader, err error) {
	requestBody := ImageGenerationRequestBody{
		Prompt:           prompt,
		N: 				  1,
		Size: 			  "512x512",
	}
	requestData, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}
	log.Printf("request gtp json string : %v", string(requestData))
	req, err := http.NewRequest("POST", BASEURL+"images/generations", bytes.NewBuffer(requestData))
	if err != nil {
		return nil, err
	}

	apiKey := config.LoadConfig().ApiKey
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	gptResponseBody := &ChatGPTResponseBody{}
	log.Println(string(body))
	err = json.Unmarshal(body, gptResponseBody)
	if err != nil {
		return nil, err
	}

	// 错误返回
	if len(gptResponseBody.Error) > 0 {
		if message, ok := gptResponseBody.Error["message"]; ok {
			return nil, errors.New(message.(string))
		} else {
			return nil, errors.New("ChatGPT server error")
		}
	}

	var url string
	if len(gptResponseBody.Choices) > 0 {
		for _, v := range gptResponseBody.Choices {
			url = v["url"].(string)
			break
		}
	}
	log.Printf("gpt response url: %s \n", url)
	urlRes, err := http.Get(url)
	if err != nil{
		return nil, err
	}

	return urlRes.Body, nil
}