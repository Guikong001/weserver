package common

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type WeChatMessageRequest struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `xml:"ToUserName"`
	FromUserName string   `xml:"FromUserName"`
	CreateTime   int64    `xml:"CreateTime"`
	MsgType      string   `xml:"MsgType"`
	Content      string   `xml:"Content"`
	MsgId        int64    `xml:"MsgId"`
	MsgDataId    int64    `xml:"MsgDataId"`
	Idx          int64    `xml:"Idx"`
}

type WeChatMessageResponse struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `xml:"ToUserName"`
	FromUserName string   `xml:"FromUserName"`
	CreateTime   int64    `xml:"CreateTime"`
	MsgType      string   `xml:"MsgType"`
	Content      string   `xml:"Content"`
}

// AI API 请求的数据结构
type AIRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// AI API 返回的数据结构
type AIResponse struct {
	Choices []Choice `json:"choices"`
}

type Choice struct {
	Message Message `json:"message"`
}

// 调用AI API函数
func CallAIAPI(userInput string) (string, error) {
	url := "https://api.nicrik.com/v1/chat/completions"
	apiKey := "API-key" // 请替换为你的 API 密钥

	// 构造请求体
	aiRequest := AIRequest{
		Model: "gpt-4o-mini",
		Messages: []Message{
			{
				Role:    "system",
				Content: "As an AI assistant.",
			},
			{
				Role:    "user",
				Content: userInput,
			},
		},
	}

	// 将请求体转换为 JSON
	jsonData, err := json.Marshal(aiRequest)
	if err != nil {
		return "", err
	}

	// 创建 HTTP 请求
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	// 发送请求
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// 解析响应体
	var aiResponse AIResponse
	if err := json.Unmarshal(body, &aiResponse); err != nil {
		return "", err
	}

	// 获取 AI 回复的内容
	if len(aiResponse.Choices) > 0 {
		return aiResponse.Choices[0].Message.Content, nil
	}

	return "抱歉，我无法处理您的请求。", nil
}

// ProcessWeChatMessage 处理微信消息并根据关键词回复
func ProcessWeChatMessage(req *WeChatMessageRequest, res *WeChatMessageResponse) {
	// 处理验证码逻辑，优先处理验证码
	if strings.Contains(req.Content, "验证码") {
		code := GenerateAllNumberVerificationCode(6)
		RegisterWeChatCodeAndID(code, req.FromUserName)
		res.Content = code
		return
	}

	// 调用 AI API 处理用户的所有其他消息
	aiReply, err := CallAIAPI(req.Content)
	if err != nil {
		res.Content = "抱歉，AI 服务暂时不可用，请稍后再试。"
	} else {
		res.Content = aiReply
	}
}
