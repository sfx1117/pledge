package controller

import (
	"github.com/gin-gonic/gin"
	"pledge-backend-test/api/common/stateCode"
	"pledge-backend-test/api/entity"
	"pledge-backend-test/api/models/response"
	"pledge-backend-test/api/service"
	"pledge-backend-test/api/validate"
	"pledge-backend-test/log"
)

type MultiSignPoolController struct {
}

func (m *MultiSignPoolController) SetMultiSign(ctx *gin.Context) {
	response := response.Gin{Res: ctx}
	request := entity.MultiSignEntity{}
	log.Logger.Sugar().Info("SetMultiSign req ", request)

	//请求参数校验
	errCode := validate.NewMultiSignVal().MultiSignEntityVal(ctx, &request)
	if errCode != stateCode.CommonSuccess {
		response.Response(ctx, errCode, nil)
		return
	}
	//业务逻辑 新增多签数据
	errCode, err := service.NewMultiSignService().SetMultiSign(&request)
	if errCode != stateCode.CommonSuccess {
		log.Logger.Error(err.Error())
		response.Response(ctx, errCode, nil)
		return
	}
	response.Response(ctx, stateCode.CommonSuccess, nil)
	return
}

func (m *MultiSignPoolController) GetMultiSign(ctx *gin.Context) {
	res := response.Gin{Res: ctx}
	req := entity.GetMultiSign{}
	data := response.MultiSignRes{}
	log.Logger.Sugar().Info("GetMultiSign req ", req)

	//请求参数校验
	errcode := validate.NewMultiSignVal().GetMultiSign(ctx, &req)
	if errcode != stateCode.CommonSuccess {
		res.Response(ctx, errcode, nil)
		return
	}
	//业务逻辑
	errcode, err := service.NewMultiSignService().GetMultiSign(&data, req.ChainId)
	if errcode != stateCode.CommonSuccess {
		log.Logger.Error(err.Error())
		res.Response(ctx, errcode, nil)
		return
	}
	res.Response(ctx, stateCode.CommonSuccess, data)
	return
}
