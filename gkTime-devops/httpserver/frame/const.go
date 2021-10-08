package frame

import "github.com/golang/glog"

type LogLevel = glog.Level

// 自定义常量方便后续的框架扩展
const (
	LogInfo  LogLevel = 4
	LogWarn  LogLevel = 3
	LogErr   LogLevel = 2
	LogFatal LogLevel = 1
)
