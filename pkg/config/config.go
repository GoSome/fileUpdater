package config

import (
	"encoding/json"
	"github.com/GoSome/fileUpdater/pkg/types"
	"github.com/fsnotify/fsnotify"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"strings"
	"sync"
)

var Path string
var Configs types.ServerConfigs
var lock sync.Mutex

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
		lock.Lock()
		err = json.NewDecoder(configFile).Decode(&Configs)
		lock.Unlock()
		if err != nil {
			// todo
			panic(err)
		}
	} else if strings.HasSuffix(Path, ".yaml") {
		log.Println("what")
		lock.Lock()
		err := yaml.NewDecoder(configFile).Decode(&Configs)
		lock.Unlock()
		if err != nil {
			// todo
			log.Println("err: ", err.Error())
			panic(err)
		}
	} else {
		panic("config file path must end with .json or .yaml")
	}
}

func Watch() {
	go func() {
		watcher, err := fsnotify.NewWatcher()
		if err != nil {
			log.Fatal(err)
		}
		defer watcher.Close()
	}()
}