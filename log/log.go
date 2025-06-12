package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path"
	"runtime"
)

var Logger *zap.Logger

func init() {
	//由于zap 不支持文件归档，如果需要将日志文件按大小和时间归档，则需要使用liumberjack
	hook := lumberjack.Logger{
		Filename:   getCurrentAbPathCaller() + "/logs/log.log", //日志路径
		MaxSize:    50,                                         //每个日志文件大最大尺寸，单位M
		MaxBackups: 20,                                         //日志文件最多保存多少个备份
		MaxAge:     7,                                          //日志文件保存多少天
		Compress:   true,                                       //是否压缩
	}
	//Zap 日志库的编码配置
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",                         //指定日志时间戳的键名	"time" → {"time": "2023-01-01T12:00:00Z"}
		LevelKey:       "level",                        //指定日志级别的键名	"level" → {"level": "info"}
		NameKey:        "logger",                       //指定日志记录器名称的键名	"logger" → {"logger": "main"}
		CallerKey:      "line",                         //指定调用者信息的键名	"line" → {"line": "main.go:42"}
		MessageKey:     "msg",                          //指定日志消息的键名	"msg" → {"msg": "user logged in"}
		StacktraceKey:  "stacktrace",                   //指定堆栈跟踪的键名	"stacktrace" → 仅在错误时显示
		LineEnding:     zapcore.DefaultLineEnding,      //行尾符	zapcore.DefaultLineEnding（\n）
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  //日志级别的编码格式 小写（info）
		EncodeTime:     zapcore.ISO8601TimeEncoder,     //时间戳的编码格式 （"2023-01-01T12:00:00Z"）
		EncodeDuration: zapcore.SecondsDurationEncoder, //耗时的编码格式 （秒，如 1.5）
		EncodeCaller:   zapcore.FullCallerEncoder,      //调用者信息的编码格式 完整路径，如 github.com/foo/bar/main.go:42）
		EncodeName:     zapcore.FullNameEncoder,        //日志记录器名称的编码格式 （完整名称）
	}

	//日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zap.InfoLevel)
	//Zap 日志库 创建了一个日志核心
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig), //编码器，将日志转换为 JSON 格式。
		//zapcore.AddSync(os.Stdout)：输出到控制台。
		//zapcore.AddSync(&hook)：输出到文件（hook 是之前配置的 lumberjack.Logger，支持日志轮转）
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)),
		atomicLevel, //日志级别
	)
	//在日志中记录调用者的文件名和行号，便于快速定位日志产生的代码位置
	caller := zap.AddCaller()
	//启用开发模式
	development := zap.Development()
	//为所有日志添加全局字段，相当于日志的“默认标签”
	fields := zap.Fields(zap.String("serviceName", "pledge"))

	Logger = zap.New(core, caller, development, fields)
}

// 获取当前文件的绝对路径
func getCurrentAbPathCaller() string {
	var abPath string
	_, file, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(file)
	}
	return abPath
}
