package abifile

import (
	"os"
	"path"
	"runtime"
)

func GetAbiByToken(token string) (string, error) {
	currentAbPath := GetCurrentAbPathByCaller()
	byte, err := os.ReadFile(currentAbPath + "/" + token + ".abi")
	if err != nil {
		return "", err
	}
	return string(byte), nil
}

func GetCurrentAbPathByCaller() string {
	var abPath string
	_, file, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(file)
	}
	return abPath
}
