package goframework

import (
	"github.com/mobliyishengyuan/goframework/library"
)

func Run() {
	library.InitOptParser()
	
	var confPath = library.GetOpt("conf_path")
	var port = library.GetOpt("port")
	var logPath = library.GetOpt("log_path")
	
	library.InitConfiger(confPath)

	library.InitLogger(logPath)
	
	library.InitHttpServer(port)
}