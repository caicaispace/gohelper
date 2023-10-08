package driver

import (
	"net/http"
	"strings"
	"time"
)

const (
	ENV_DEV  = "dev"
	ENV_TEST = "test"
	ENV_PRE  = "pre"
	ENV_PROD = "prod"
)

var ENV_MAP = map[string]string{
	ENV_DEV:  "开发环境",
	ENV_TEST: "测试环境",
	ENV_PRE:  "预发布环境",
	ENV_PROD: "生产环境",
}

type Driver interface {
	CreateWeWorkTemplate(title string, messageList []string) (content string)
	Notify(title, webHook string, messageList []string) bool
}

func notify(title, webHook string, messageList []string, driver Driver) bool {
	if webHook == "" || title == "" {
		return false
	}
	template := driver.CreateWeWorkTemplate(title, messageList)
	return send(webHook, template)
}

func send(webHook string, template string) bool {
	client := &http.Client{
		Timeout: 3 * time.Second,
	}
	req, err := http.NewRequest("POST", webHook, strings.NewReader(template))
	if err != nil {
		return false
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	return true
}

func getEnv(env string) string {
	if env == "" {
		return ""
	}
	return ENV_MAP[env]
}
