package notify

import (
	"errors"

	"github.com/caicaispace/gohelper/notify/driver"
)

const (
	WEIXIN   = "weixin"
	FEISHU   = "feishu"
	DINGDING = "dingding"
)

var WebHookList = map[string]string{
	WEIXIN:   "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=xxxxx",
	FEISHU:   "https://open.feishu.cn/open-apis/bot/v2/hook/xxx",
	DINGDING: "https://oapi.dingtalk.com/robot/send?access_token=xxxxx",
}

type Notify struct {
	dirver driver.Driver
}

func New(env string) *Notify {
	return &Notify{
		dirver: driver.NewFeishu(env),
	}
}

func (n *Notify) Notify(platform string, msg string) error {
	webHook, ok := WebHookList[platform]
	if !ok {
		return errors.New("platform not found")
	}
	title := "错误通知"
	n.dirver.Notify(title, webHook, []string{msg})
	return nil
}
