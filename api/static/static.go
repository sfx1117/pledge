package static

import (
	"path"
	"runtime"
)

func GetCurrentAbPathByCaller() string {
	var abPath string
	_, file, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(file)
	}
	return abPath
}
