package webutil

import (
	"github.com/blocktree/openwallet/v2/openwallet"
	"github.com/godaddy-x/jorm/log"
	"github.com/godaddy-x/jorm/util"
	"strings"
)

type WebEx struct {
	Code   uint64
	ErrMsg string
	ExtMsg string
}

func Try(err error, msg string) error {
	owerr := openwallet.ConvertError(err)
	obj := &WebEx{Code: owerr.Code(), ErrMsg: owerr.Error(), ExtMsg: msg}
	result, _ := util.ObjectToJson(obj)
	return util.Error(result)
}

func Catch(err error) *WebEx {
	obj := &WebEx{}
	msg := err.Error()
	if strings.HasPrefix(msg, "{") && strings.HasSuffix(msg, "}") {
		if err := util.JsonToObject(msg, obj); err != nil {
			log.Error("交易单异常捕获失败", 0, log.String("msg", msg), log.AddError(err))
			obj.ErrMsg = util.AddStr("交易单异常捕获失败: ", err)
		}
	}
	return obj
}
