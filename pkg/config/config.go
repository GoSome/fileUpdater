package config

import (
	"encoding/json"
	"github.com/GoSome/fileUpdater/pkg/types"
	"github.com/fsnotify/fsnotify"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"strings"
)

var Path string
var Configs types.ServerConfigs

// Load loads configs from config file.
func Load() {
	if _, err := os.Stat(Path); os.IsNotExist(err) {
		log.Fatalf("config file \"%s\" not exist", Path)
	}

	configFile, err := os.Open(Path)
	if err != nil {
		log.Fatalf("open config file \"%s\" failed", Path)
	}

	if strings.HasSuffix(Path, ".json") {
		//
		err = json.NewDecoder(configFile).Decode(&Configs)
		if err != nil {
			// todo
			panic(err)
		}
	} else if strings.HasSuffix(Path, ".yaml") {
		log.Println("what")
		err := yaml.NewDecoder(configFile).Decode(&Configs)
		if err != nil {
			// todo
			log.Println("err: ", err.Error())
			panic(err)
		}
	} else {
		panic("config file path must end with .json or .yaml")
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
				Load()
				log.Println("config reloaded")
			}
		}
	}
}
