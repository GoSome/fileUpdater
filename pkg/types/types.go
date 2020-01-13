/*
@Description: just go
@Author: skipper
@Date: 2020/1/13
@Time: 3:50 PM
@ProjectName fileUpdater
*/
package types

import (
	"io"
)

type Updater interface {
	GetFileContent() (io.Reader, error)
}

type Hook interface {
	Do() error
}

type CommandHook struct {
	Commands []string `json:"commands"`
	Mode     string   `json:"mode"`
}

func (c CommandHook) Do() error {
	return nil
}
