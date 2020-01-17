/*
@Description: just go
@Author: skipper
@Date: 2020/1/17
@Time: 3:16 PM
@ProjectName fileUpdater
*/
package core

import (
	"log"
	"path/filepath"
	"testing"
)

func TestFindOldFiles(t *testing.T) {
	paths, err := filepath.Glob("/tmp/*")
	if err != nil {
		t.Error(err)
	}
	log.Println(FindOldFiles(paths, 1))
}
