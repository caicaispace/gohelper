package driver

import (
	"encoding/json"
	"time"
)

// 卡片消息
type messaegContentCard struct {
	MsgType string `json:"msg_type"`
	Card    Card   `json:"card"`
}

type Text struct {
	Content string `json:"content"`
	Tag     string `json:"tag"`
}

type Value struct{}

type Actions struct {
	Tag   string `json:"tag"`
	Text  Text   `json:"text"`
	URL   string `json:"url"`
	Type  string `json:"type"`
	Value Value  `json:"value"`
}

type Elements struct {
	Tag     string    `json:"tag"`
	Text    Text      `json:"text,omitempty"`
	Actions []Actions `json:"actions,omitempty"`
}

type Title struct {
	Content string `json:"content"`
	Tag     string `json:"tag"`
}

type Header struct {
	Title Title `json:"title"`
}

type Card struct {
	Elements []Elements `json:"elements"`
	Header   Header     `json:"header"`
}

// 简单文本消息
type messaegContent struct {
	MsgType string            `json:"msg_type"`
	Content map[string]string `json:"content"`
}

type Feishu struct {
	env string
}

func NewFeishu(env string) Feishu {
	return Feishu{
		env: env,
	}
}

func (f Feishu) CreateWeWorkTemplate(title string, messageList []string) (content string) {
	str := "> 操作环境：" + getEnv(f.env) + "\n"
	str += "> 操作时间：" + time.Now().Format("2006-01-02 15:04:05") + "\n"
	for _, message := range messageList {
		str += "> " + message + "\n"
	}
	contentStr := "**" + title + "**\n" + str
	// messaegContent := messaegContent{
	// 	MsgType: "text",
	// 	Content: map[string]string{
	// 		"text": contentStr,
	// 	},
	// }
	// 发送卡片消息
	messaegContentCard := messaegContentCard{
		MsgType: "interactive",
		Card: Card{
			Header: Header{
				Title: Title{
					Content: title,
					Tag:     "plain_text",
				},
			},
			Elements: []Elements{
				{
					Tag: "div",
					Text: Text{
						Content: contentStr,
						Tag:     "lark_md",
					},
				},
			},
		},
	}
	jsonStr, _ := json.Marshal(messaegContentCard)
	return string(jsonStr)
}

func (f Feishu) Notify(title, webHook string, messageList []string) bool {
	return notify(title, webHook, messageList, f)
}
