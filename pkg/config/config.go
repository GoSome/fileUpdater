package config

import (
	"encoding/json"
	"github.com/GoSome/fileUpdater/pkg/types"
	"github.com/fsnotify/fsnotify"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"path"
)

var Path string
var Configs types.ServerConfigs
var lastConfigs types.ServerConfigs

// Load loads configs from config file.
func Load(init bool) {
	logFunc := log.Printf
	if init {
		logFunc = log.Fatalf
	}
	if _, err := os.Stat(Path); os.IsNotExist(err) {
		log.Fatalf("config file \"%s\" not exist", Path)
	}

	configFile, err := os.Open(Path)
	if err != nil {
		logFunc("open config file \"%s\" failed", Path)
	}

	deepCopy(&lastConfigs, &Configs)

	switch path.Ext(Path) {
	case ".json":
		err = json.NewDecoder(configFile).Decode(&Configs)
		if err != nil {
			logFunc("%s", err)
		}
	case ".yml":
		fallthrough
	case ".yaml":
		err := yaml.NewDecoder(configFile).Decode(&Configs)
		if err != nil {
			logFunc("%s", err)
		}
	default:
		logFunc("config file path must end with .json or .yaml")
	}
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
				Load(false)
			}
		}
	}
}
