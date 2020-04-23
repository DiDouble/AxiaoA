package myloger

import (
	"fmt"
	"time"
)

// Logger 定义了一个Logger的结构体
type Logger struct {
	Level LogLevel
}

func (l Logger) enable(LogLevel LogLevel) bool {
	return LogLevel >= l.Level
}

//NewLog 调用了一个方法
func NewLog(level string) Logger {
	levelStr, err := parseLogLevel(level)
	if err != nil {
		panic(err)
	}
	return Logger{
		Level: levelStr,
	}
}

// 解析传过来的值是什么类型的.

//定义一个日志输出的方法
func log(lv LogLevel, msg string) {
	now := time.Now()
	funcName, fileName, lineNo := getInfo(3)
	fmt.Printf("%s [%s][%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), funcName, fileName, lineNo, msg)
}

//Debug 是Logger 结构体方法
func (L Logger) Debug(msg string) {
	if L.enable(DEBUG) {
		log(DEBUG, msg)
	}
}
func (L Logger) Info(msg string) {
	if L.enable(INFO) {
		now := time.Now()
		fmt.Printf("%s [Info] %s\n", now.Format("2006-01-02 15:04:05"), msg)
	}
}
func (L Logger) Warning(msg string) {
	if L.enable(WARNING) {
		now := time.Now()
		fmt.Printf("%s [Warning] %s\n", now.Format("2006-01-02 15:04:05"), msg)
	}
}
func (L Logger) Error(msg string) {
	if L.enable(ERROR) {
		now := time.Now()
		fmt.Printf("%s [Error] %s\n", now.Format("2006-01-02 15:04:05"), msg)
	}
}
