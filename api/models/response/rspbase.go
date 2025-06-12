package response

import (
	"github.com/gin-gonic/gin"
	"pledge-backend-test/api/common/stateCode"
)

type Gin struct {
	Res *gin.Context
}

type Page struct {
	Code  int         `json:"code"`
	Msg   string      `json:"message"`
	Total int         `json:"total"`
	Data  interface{} `json:"data"`
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"message"`
	Data interface{} `json:"data"`
}

func (g *Gin) ResponsePages(ctx *gin.Context, code int, total int, data interface{}) {
	lang := stateCode.LangZh         // 默认中文
	value, exists := ctx.Get("lang") //尝试从 Gin 上下文 c 中获取 lang 字段（可能由前置中间件设置）
	if exists {                      //如果存在，覆盖默认语言
		lang = value.(int)
	}
	rsp := Page{
		Code:  code,
		Msg:   stateCode.GetMsg(code, lang),
		Total: total,
		Data:  data,
	}
	g.Res.JSON(200, rsp)
	return
}

func (g *Gin) Response(ctx *gin.Context, code int, data interface{}, httpStatus ...int) {
	lang := stateCode.LangEn
	value, exists := ctx.Get("lang")
	if exists {
		lang = value.(int)
	}
	rsp := Response{
		Code: code,
		Msg:  stateCode.GetMsg(code, lang),
		Data: data,
	}
	HttpStatus := 200
	if len(httpStatus) > 0 {
		HttpStatus = httpStatus[0]
	}
	g.Res.JSON(HttpStatus, rsp)
	return
}
