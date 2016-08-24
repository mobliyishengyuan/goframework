package library

import (
	"github.com/mobliyishengyuan/goconf"
	"io/ioutil"
	"log"
	"path/filepath"
)

const (
	SystemConfName = "system.ini"
)

var confs map[string]*goconf.Conf = make(map[string]*goconf.Conf)

func InitConfiger(confDirPath string) {
	list, err := ioutil.ReadDir(confDirPath)
	
	if err != nil {
		log.Fatal("library InitConfiger err", err)
	}
	
	var config *goconf.Conf
	
	for _, fileInfo := range list {
		fileName := fileInfo.Name()
		
		if filepath.Ext(fileName) == ".ini" {
			confAbsPath := filepath.Join(confDirPath, fileName)
			
			config = goconf.GetNewConfig()
			config.Read(confAbsPath)
			
			confs[fileName] = config
		}
	}
}

func GetConfValue(name, section , key string) (string, bool) {
	var config,status = confs[name]
	
	if !status {
		return "", false
	}
	
	return config.Get(section, key)
}