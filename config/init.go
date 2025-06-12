package config

import (
	"github.com/BurntSushi/toml"
	"path"
	"path/filepath"
	"pledge-backend-test/log"
	"runtime"
)

func init() {
	abPath := getCurrentAbPathByCaller()
	tomFile, err := filepath.Abs(abPath + "/config.toml")
	if err != nil {
		log.Logger.Error("read toml file err: " + err.Error())
		return
	}
	_, err = toml.DecodeFile(tomFile, &Config)
	if err != nil {
		log.Logger.Error("read toml file err: " + err.Error())
		return
	}
}
func getCurrentAbPathByCaller() string {
	var abPath string
	_, file, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(file)
	}
	return abPath
}
