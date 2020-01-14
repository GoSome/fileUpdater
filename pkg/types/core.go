package types

import (
	"bytes"
	"errors"
	"io"
	"log"
	"os"
	"os/exec"
	"time"
)

// server configFiles struct

type ServerConfigs struct {
	ServerHost   string        `json:"server_host" yaml:"server_host"`
	ServerPort   string        `json:"server_port" yaml:"server_port"`
	FileUpdaters []FileUpdater `json:"updaters" yaml:"updaters"`
	IncludeSelf  bool          `json:"include_self"`
}

func (s ServerConfigs) GetUpdatersByName(name string) *FileUpdater {
	for k, v := range s.FileUpdaters {
		if v.Name == name {
			return &s.FileUpdaters[k]
		}
	}
	return nil
}

type FileUpdater struct {
	Name     string      `json:"name" yaml:"name"`
	Type     string      `json:"type" yaml:"type"`
	FilePath string      `json:"path" yaml:"path"`
	PreHook  CommandHook `json:"pre_hook" yaml:"pre_hook"`
	PostHook CommandHook `json:"post_hook" yaml:"post_hook"`
}

// should close reader
func (u FileUpdater) GetFileContent() (reader *os.File, err error) {
	file, err := os.Open(u.FilePath)
	if err != nil {
		log.Println("what: ", err)
		return nil, err
	}
	return file, nil
}

// when preHook Not nil should execute pre hook
// when preHook executed and exit not 0 return
// copy origin file for backup and update file
// when file update succeeded execute the post hook
func (u FileUpdater) UpdateFile(date io.Reader) error {

	// pre hook
	err := u.execPreHook()
	if err != nil {
		return err
	}

	//Copy File for backup
	//todo
	bfp, err := BackupFile(u.FilePath, "")
	if err != nil {
		return errors.New("backup origin file failed")
	}

	//write new content to file
	file, err := os.OpenFile(u.FilePath, os.O_TRUNC|os.O_RDWR|os.O_SYNC, 0644)
	if err != nil {
		return errors.New("open file failed")
	}
	_,err = io.Copy(file, date)
	file.Close()
	// when write failed,will auto restore,when restore failed,something bad happen
	if err != nil {
		// restore
		log.Printf("restore file! origin: %s,backup: %s", u.FilePath, bfp)
		err := RestoreFile(u.FilePath, bfp)
		if err != nil {
			log.Println("restore backup file failed")
			return errors.New("restore backup file failed,please check it out manually")
		}
	}

	// pre hook
	err = u.execPostHook()
	if err != nil {
		// todo
		return err
	}
	return nil
}

func (u FileUpdater) execPreHook() error {
	// if no command in pre Hook skip
	if len(u.PreHook.Commands) != 0 {
		if u.PreHook.Mode == "strict" {
			for _, c := range u.PreHook.Commands {
				out, err := BashExec(c)
				if err != nil {
					log.Println(out)
					return errors.New("pre hook execute failed")
				}
			}
		} else {
			for _, c := range u.PreHook.Commands {
				BashExec(c)
			}
		}
	}
	return nil
}
func (u FileUpdater) execPostHook() error {
	// do the post hook
	if len(u.PostHook.Commands) != 0 {
		if u.PostHook.Mode == "strict" {
			for _, c := range u.PostHook.Commands {
				out, err := BashExec(c)
				if err != nil {
					log.Println(out)
					return errors.New("pre hook execute failed")
				}
			}
		} else {
			for _, c := range u.PostHook.Commands {
				BashExec(c)
			}
		}
	}
	return nil
}

func BashExec(cmd string) (output string, err error) {
	var stdout, stderr bytes.Buffer
	command := exec.Command("bash", "-c", cmd)
	command.Stdout = &stdout
	command.Stderr = &stderr
	log.Printf("执行命令: %s", cmd)
	err = command.Run()
	if err != nil {
		output = stderr.String()
		log.Printf("执行命令报错: %s, stderr: \n\n%s", err.Error(), output)
		return output, err
	}
	return stdout.String(), nil
}

func BackupFile(filePath, backupPath string) (newPath string, err error) {
	if backupPath == "" {
		backupPath = filePath + time.Now().Format("2006-01-02-15:04:05") + ".fub"
	}
	originFile, err := os.Open(filePath)
	defer originFile.Close()
	if err != nil {
		return "", err
	}
	backupFile, err := os.OpenFile(backupPath, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0644)
	defer backupFile.Close()
	if err != nil {
		return "", err
	}
	_, err = io.Copy(backupFile, originFile)
	if err != nil {
		return "", err
	}
	return backupPath, nil
}

func RestoreFile(filePath, BackupPath string) error {
	file, err := os.OpenFile(filePath, os.O_TRUNC|os.O_RDWR|os.O_SYNC, 0644)
	if err != nil {
		return err
	}
	backcup, err := os.Open(BackupPath)
	if err != nil {
		return errors.New("open backup file failed")
	}
	_, err = io.Copy(file, backcup)
	return err
}
