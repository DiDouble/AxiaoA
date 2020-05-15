package myloger

import (
	"errors"
	"fmt"
	"path"
	"runtime"
	"strings"
)


// 重命名uint16型
type LogLevel uint16

const (  //定义一个不变函数
	//定义日志级别
	UNKOWN LogLevel = iota
	DEBUG
	TRACE
	INFO
	WARNING
	ERROR
	FATAL
)
// 通过parseLogLevel方法将传进来的字符串转换成LogLevel返回去
func parseLogLevel(s string) (LogLevel, error) {
	s = strings.ToLower(s) //将传过来的string类型的字符串转换为小写
	switch s {
	case "debug":
		return DEBUG, nil //如果没有错,错误返回nil
	case "trace":
		return TRACE, nil
	case "info":
		return INFO, nil
	case "warning":
		return WARNING, nil
	case "error":
		return ERROR, nil
	case "fatal":
		return FATAL, nil 
	default:
		// fmt.Println("s")
		err := errors.New("新的日志级别")
		return UNKOWN, err

	}
}
//将传入的LogLevel型的参数以字符串的形式输出。
func getLogString(lv LogLevel) string {
	switch lv {
	case DEBUG:
		return "DEBUG"
	case TRACE:
		return "TRACE"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	default:
		return "DEBUG"
	}
}

func getInfo(skip int) (funcname, filename string, fileline int) {
	pc, file, lineNo, ok := runtime.Caller(skip)
	if !ok {
		fmt.Println("failed")
		return
	}
	funcname = runtime.FuncForPC(pc).Name()
	filename = path.Base(file)
	fileline = lineNo
	funcname = strings.Split(filename, ".")[1]
	return
}
