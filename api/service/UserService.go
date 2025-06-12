package service

import (
	"pledge-backend-test/api/common/stateCode"
	"pledge-backend-test/api/entity"
	"pledge-backend-test/api/models/response"
	"pledge-backend-test/config"
	"pledge-backend-test/db"
	"pledge-backend-test/log"
	"pledge-backend-test/utils"
)

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (u *UserService) Login(req *entity.UserEntity, data *response.UserRes) int {
	log.Logger.Sugar().Info("userService,req={}", req)
	if req.Name == "admin" && req.Password == "password" {
		//根据username生成token
		token, err := utils.CreateToken(req.Name, config.Config.Jwt.SecretKey)
		if err != nil {
			log.Logger.Error("CreateToken" + err.Error())
			return stateCode.CommonErrServerErr
		}
		data.TokenId = token
		//将token放到redis
		_ = db.RedisSet(req.Name, "login_ok", config.Config.Jwt.ExpireTime)
		return stateCode.CommonSuccess
	} else {
		return stateCode.NameOrPasswordErr
	}
}
