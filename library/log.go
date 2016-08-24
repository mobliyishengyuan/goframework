package library

import (
	"log"
	"os"
	"strconv"
)

type LogLevel uint32


const (
	LogLevelFatal LogLevel = 1 << iota 
	LogLevelWarn
	LogLevelNotice
	LogLevelTrace
	LogLevelDebug
	DefaultLogLevel = LogLevelDebug
)

type FrLogger struct {
	// 正常日志
	logger *log.Logger
	// 异常日志
	wfLogger *log.Logger
	// 日志级别
	level LogLevel
}

var logLevelMap = map[LogLevel]string{
	LogLevelFatal : "Fatal",
	LogLevelWarn : "Warn",
	LogLevelNotice : "Notice",
	LogLevelTrace : "Trace",
	LogLevelDebug : "Debug",
}

var frLogger *FrLogger

func InitLogger(logPath string) {
	frLogger = new(FrLogger)
	
	frLogger.logger = getLogger(logPath)
	frLogger.wfLogger = getLogger(logPath + ".wf")
	frLogger.level = getLevel()
}

func getLogger(logPath string) *log.Logger {
	var logger *log.Logger
	
	file, err := os.OpenFile(logPath, os.O_WRONLY | os.O_APPEND | os.O_CREATE, 0666)
	
	if err != nil {
		log.Fatalln("log : open file error ", logPath , err)
	}
	
	logger = log.New(file, "", log.LstdFlags | log.Llongfile)
	
	return logger
}

func getLevel() LogLevel {
	var logLevel = DefaultLogLevel

	rtnLogLevel, ok := GetConfValue(SystemConfName, "conf", "level")
	
	if !ok {
		return logLevel
	}

	tmp, err := strconv.Atoi(rtnLogLevel)
	
	if err != nil {
		return logLevel
	}
	
	logLevel = (LogLevel)(tmp)
	
	return logLevel
}

func writeLog(logLevel LogLevel, content string) {
	if logLevel > frLogger.level {
		return
	}
	
	var str string
	var ok bool
	if str,ok = logLevelMap[logLevel]; !ok {
		return
	}
	
	content = str + " " + content
	
	if logLevel <= LogLevelWarn {
		frLogger.wfLogger.Println(content)
	} else {
		frLogger.logger.Println(content)
	}
}
