package driver

type Dingding struct{}

func NewDingding() Dingding {
	return Dingding{}
}

func (d Dingding) CreateWeWorkTemplate(title string, messageList []string) (content string) {
	return ""
}

func (d Dingding) Notify(title, webHook string, messageList []string) bool {
	return false
}
