package library

import (
	"flag"
	"os"
	"log"
	"path/filepath"
)

var optMap map[string]string = make(map[string]string)

func InitOptParser() {
	curDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal("library InitoptParser", err)
	}
	
	var confPath, logPath string
	var port string
	
	var defaultConfPath = filepath.Join(curDir, "conf")
	var defaultLogPath = filepath.Join(curDir, "server.log")
	
	flag.StringVar(&confPath, "conf_path", defaultConfPath, "please input conf path")
	flag.StringVar(&logPath, "log_path", defaultLogPath, "please input log path")
	flag.StringVar(&port, "port", "8000", "please input port")
	
	flag.Parse()
	
	optMap["conf_path"] = confPath
	optMap["log_path"] = logPath
	optMap["port"] = port
}

func GetOpt(name string) (string)  {
	return optMap[name]
}