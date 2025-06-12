package controller

import (
	"github.com/gin-gonic/gin"
	"pledge-backend-test/api/common/stateCode"
	"pledge-backend-test/api/dao"
	"pledge-backend-test/api/entity"
	"pledge-backend-test/api/models/response"
	"pledge-backend-test/api/service"
	"pledge-backend-test/api/validate"
	"pledge-backend-test/config"
	"regexp"
	"strings"
	"time"
)

type PoolController struct {
}

/*
*
查询poolBaseInfo信息
*/
func (p *PoolController) SelectPoolBaseInfo(ctx *gin.Context) {
	res := response.Gin{Res: ctx}
	req := entity.PoolBaseInfoEntity{}
	data := []response.PoolBaseInfoRes{}
	//参数绑定校验
	errcode := validate.NewPoolBaseInfoValidate().PoolBaseInfo(ctx, &req)
	if errcode != stateCode.CommonSuccess {
		res.Response(ctx, errcode, nil)
		return
	}
	//业务逻辑，查询poolBaseInfo
	errcode = service.NewPoolBaseInfoService().SelectPoolBaseInfo(req.ChainId, &data)
	if errcode != stateCode.CommonSuccess {
		res.Response(ctx, errcode, nil)
		return
	}
	res.Response(ctx, stateCode.CommonSuccess, data)
	return
}

/*
*
查询poolData controller
*/
func (p *PoolController) SelectPoolDataInfo(ctx *gin.Context) {
	res := response.Gin{Res: ctx}
	req := entity.PoolDataInfoEntity{}
	data := []response.PoolDataInfoRes{}

	//参数绑定校验
	errcode := validate.NewPoolDataInfoVal().PoolDataInfo(ctx, &req)
	if errcode != stateCode.CommonSuccess {
		res.Response(ctx, errcode, nil)
		return
	}
	//业务逻辑
	errcode = service.NewPoolDataInfoService().SelectPoolDataInfo(req.ChainId, &data)
	if errcode != stateCode.CommonSuccess {
		res.Response(ctx, errcode, nil)
		return
	}
	res.Response(ctx, stateCode.CommonSuccess, data)
}

/*
*
查询tokenList controller
*/
func (p *PoolController) SelectTokenList(ctx *gin.Context) {
	res := response.Gin{Res: ctx}
	req := entity.TokenListEntity{}
	data := response.TokenListRes{}
	//参数绑定校验
	errcode := validate.NewTokenListVal().TokenListVal(ctx, &req)
	if errcode != stateCode.CommonSuccess {
		res.Response(ctx, errcode, nil)
		return
	}
	//业务逻辑
	errcode, tokenList := service.NewTokenListService().GetTokenList(req.ChainId)
	if errcode != stateCode.CommonSuccess {
		res.Response(ctx, errcode, nil)
		return
	}
	baseUrl := p.GetBaseUrl()
	data.Name = "Pledge Token List"
	data.LogoURI = baseUrl + "static/img/Pledge-project-logo.png"
	data.Timestamp = time.Now()
	data.Version = response.Version{
		Major: 2,
		Minor: 16,
		Patch: 12,
	}
	for _, v := range tokenList {
		data.Tokens = append(data.Tokens, response.Token{
			Name:     v.Symbol,
			Symbol:   v.Symbol,
			Decimals: v.Decimals,
			Address:  v.Token,
			ChainId:  v.ChainId,
			LogoURI:  v.Logo,
		})
	}
	res.Response(ctx, stateCode.CommonSuccess, data)
	return
}
func (p *PoolController) GetBaseUrl() string {
	domainName := config.Config.Env.DomainName
	domainNameSlice := strings.Split(domainName, "")
	pattern := "\\d+"
	isNumber, _ := regexp.MatchString(pattern, domainNameSlice[0])
	if isNumber {
		return config.Config.Env.Protocol + "://" + config.Config.Env.DomainName + ":" + config.Config.Env.Port + "/"
	} else {
		return config.Config.Env.Protocol + "://" + config.Config.Env.DomainName + "/"
	}
}

/*
*
分页查询poolBase  controller
*/
func (p *PoolController) Search(ctx *gin.Context) {
	res := response.Gin{Res: ctx}
	req := entity.SearchEntity{}
	data := dao.SearchRes{}
	//参数绑定校验
	errCode := validate.NewSearchVal().Search(ctx, &req)
	if errCode != stateCode.CommonSuccess {
		res.Response(ctx, errCode, nil)
		return
	}
	//业务逻辑
	errCode, total, poolList := service.NewSearchService().Search(&req)
	if errCode != stateCode.CommonSuccess {
		res.Response(ctx, errCode, nil)
		return
	}
	data.Count = total
	data.Rows = poolList
	res.Response(ctx, stateCode.CommonSuccess, data)
	return
}

func (p *PoolController) DebtTokenList(ctx *gin.Context) {
	res := response.Gin{Res: ctx}
	req := entity.TokenListEntity{}
	data := []response.TokenInfo{}

	errcode := validate.NewTokenListVal().TokenListVal(ctx, &req)
	if errcode != stateCode.CommonSuccess {
		res.Response(ctx, errcode, nil)
		return
	}

	errcode, tokenInfo := service.NewTokenListService().DebtTokenList(req.ChainId)
	if errcode != stateCode.CommonSuccess {
		res.Response(ctx, errcode, nil)
		return
	}
	for _, v := range tokenInfo {
		data = append(data, response.TokenInfo{
			Id:      v.Id,
			Symbol:  v.Symbol,
			Token:   v.Token,
			ChainId: v.ChainId,
		})
	}
	res.Response(ctx, stateCode.CommonSuccess, data)
	return
}
