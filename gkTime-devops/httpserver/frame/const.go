package frame

import "github.com/golang/glog"

type LogLevel = glog.Level

const (
	LOG_INFO LogLevel = 4
	LOG_WARN LogLevel = 3
	LOG_ERR LogLevel = 2
	LOG_FATAL LogLevel = 1
)

