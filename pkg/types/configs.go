package types

import (
	"log"
	"os"
)

// server configFiles struct

type ServerConfigs struct {
	ServerHost   string        `json:"server_host"`
	ServerPort   string        `json:"server_port"`
	FileUpdaters []FileUpdater `json:"updaters"`
	IncludeSelf  bool          `json:"include_self"`
}

type FileUpdater struct {
	Name     string        `json:"name"`
	Type     string        `json:"type"`
	FilePath string         `json:"file_path"`
	PreHook  CommandHook `json:"pre_hook"`
	PostHook CommandHook `json:"post_hook"`
}

// should close reader
func (u FileUpdater) GetFileContent() (reader *os.File, err error) {
	file, err := os.Open(u.FilePath)
	if err != nil {
		log.Println("what: ",err)
		return nil, err
	}
	return file, nil
}
func (u FileUpdater) UpdateFile(date []byte) error {
	return nil
}

func (u FileUpdater) execPreHook() {

}
func (u FileUpdater) execPostHook() {

}