package driver

type Weixin struct{}

func NewWeixin() Weixin {
	return Weixin{}
}

func (w Weixin) CreateWeWorkTemplate(title string, messageList []string) (content string) {
	return ""
}

func (w Weixin) Notify(title, webHook string, messageList []string) bool {
	return false
}
