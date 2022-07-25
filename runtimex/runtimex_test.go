package runtimex_test

import (
	"testing"

	"github.com/caicaispace/gohelper/runtimex"
)

func Test_GetCurrentAbPath(t *testing.T) {
	path := runtimex.GetCurrentAbPath()
	t.Log(path)
}

func Test_GetRootPath(t *testing.T) {
	path := runtimex.GetRootPath()
	t.Log(path)
}

func Test_GetRootParentPath(t *testing.T) {
	path := runtimex.GetRootParentPath()
	t.Log(path)
}

func Test_GetAppRootPath(t *testing.T) {
	path := runtimex.GetAppRootPath()
	t.Log(path)
}

func Test_GetAppRootPath2(t *testing.T) {
	path := runtimex.GetAppRootPath2()
	t.Log(path)
}
