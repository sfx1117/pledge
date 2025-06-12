package common

import (
	"os"
	"pledge-backend-test/log"
)

var PledgeAdminPrivateKey string

func GetEnv() {
	var ok bool
	PledgeAdminPrivateKey, ok = os.LookupEnv("pledge_admin_private_key")
	if !ok {
		log.Logger.Sugar().Error("environment variable is not set")
		panic("environment variable is not set")
	}
}
