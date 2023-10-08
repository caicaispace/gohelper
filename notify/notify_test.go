package notify

import (
	"testing"

	"github.com/caicaispace/gohelper/notify/driver"
)

func TestFeishuNotify(t *testing.T) {
	n := New(driver.ENV_DEV)
	n.Notify(FEISHU, "test")
}
