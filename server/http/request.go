package http

import (
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/caicaispace/gohelper/business"
	"github.com/caicaispace/gohelper/errx"
	"github.com/caicaispace/gohelper/logx"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func MarkErrors(errors []*validation.Error) {
	for _, err := range errors {
		logx.Info(err.Key, err.Message)
	}
}

// BindAndValid binds and validates data
func BindAndValid(c *gin.Context, body interface{}) (int, int) {
	err := c.Bind(body)
	if err != nil {
		return http.StatusBadRequest, errx.InvalidParams
	}
	valid := validation.Validation{}
	check, err := valid.Valid(body)
	if err != nil {
		return http.StatusInternalServerError, errx.Error
	}
	if !check {
		MarkErrors(valid.Errors)
		return http.StatusBadRequest, errx.InvalidParams
	}
	return http.StatusOK, errx.Success
}

func (c *Context) GetPager() *business.Pager {
	pager := business.GetInstance()
	pager.SetPage(com.StrTo(c.C.Query("p_page")).MustInt())
	pager.SetLimit(com.StrTo(c.C.Query("p_limit")).MustInt())
	pager.SetTotal(com.StrTo(c.C.Query("p_total")).MustInt())
	return pager
}
