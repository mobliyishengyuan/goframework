package goframework

import (
	"github.com/mobliyishengyuan/goframework/library"
)

func Run() {
	library.QueryPlugin(library.PluginPreOptParse)
	
	library.InitOptParser()
	
	library.QueryPlugin(library.PluginPostOptParse)
	
	
	var confPath = library.GetOpt("conf_path")
	var port = library.GetOpt("port")
	var logPath = library.GetOpt("log_path")
	
	
	library.InitConfiger(confPath)
	

	library.InitLogger(logPath)
	
	
	library.QueryPlugin(library.PluginPreServerStart)
	
	library.InitHttpServer(port)
}