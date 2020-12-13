package config

import (
	"encoding/json"
	"github.com/GoSome/fileUpdater/pkg/core"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"path"
)

var Path string
var Config core.ServerConfigs
var DaemonZ bool
var PidPath string
var LogFile string
var IncludeSelf bool
var DisableHotReload bool
var SSHUser string
var SSHPasswd string

func Parse(init bool) {
	logFunc := log.Printf
	if init {
		logFunc = log.Fatalf
	}
	if Path == "" {
		defaultPath := "config.yaml"
		log.Println("create default config, path: %s",defaultPath)
		err := CreateDefaultConfig(defaultPath)
		if err != nil {
			log.Println("create default config failed",err)
			return
		}
		Path = defaultPath
	}else if _, err := os.Stat(Path); os.IsNotExist(err) {
		logFunc("config file \"%s\" not exist", Path)
		return
	}

	configFile, err := os.Open(Path)
	if err != nil {
		logFunc("open config file \"%s\" failed", Path)
		return
	}
	defer configFile.Close()

	switch path.Ext(Path) {
	case ".json":
		err = json.NewDecoder(configFile).Decode(&Config)
		if err != nil {
			logFunc("%s", err)
		}
	case ".yml":
		fallthrough
	case ".yaml":
		err := yaml.NewDecoder(configFile).Decode(&Config)
		if err != nil {
			logFunc("%s", err)
		}
	default:
		logFunc("config file path must end with .json or .yaml")
	}

	if IncludeSelf {
		Config.FileUpdaters = append(Config.FileUpdaters, core.FileUpdater{Name: "FileUpdaterSelfConfig", FilePath: Path})
	}
}

// Inject injects config into gin's context.
func Inject(c *gin.Context) {
	c.Set("cfg", Config)
	c.Next()
}

// Watch watches config file, configs will be reloaded when config file is changed.
func Watch() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	watcher.Add(Path)

	for {
		select {
		case ev := <-watcher.Events:
			if ev.Op == fsnotify.Write {
				log.Println("config file has been changed, attempt to reload...")
				Parse(false)
			}
		}
	}
}


func CreateDefaultConfig(path string)error  {
	log.Println("fuck")
	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Println("default config not exist. auto create", path)
		sf := core.NewDefaultConfig()
		file ,err := os.OpenFile(path,os.O_CREATE|os.O_RDWR,0644)
		if err != nil {
			log.Println("create default config failed ",err)
			return err
		}
		return yaml.NewEncoder(file).Encode(&sf)
	}
	log.Println("default config file already exist.will use it.")
	return nil
}