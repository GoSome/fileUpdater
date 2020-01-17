package core

import (
	"bytes"
	"errors"
	"gopkg.in/djherbis/times.v1"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

// server configFiles struct
type ServerConfigs struct {
	ServerHost   string        `json:"server_host" yaml:"server_host"`
	ServerPort   string        `json:"server_port" yaml:"server_port"`
	FileUpdaters []FileUpdater `json:"updaters" yaml:"updaters"`
	IncludeSelf  bool          `json:"include_self"`
}

func (s ServerConfigs) GetUpdaterByName(name string) *FileUpdater {
	for k, v := range s.FileUpdaters {
		if v.Name == name {
			return &s.FileUpdaters[k]
		}
	}
	return nil
}

type FileUpdater struct {
	Name     string      `json:"name" yaml:"name"`
	FilePath string      `json:"path" yaml:"path"`
	Backup   bool        `json:"backup" yaml:"backup"`
	PreHook  CommandHook `json:"pre_hook" yaml:"pre_hook"`
	PostHook CommandHook `json:"post_hook" yaml:"post_hook"`
}

// should close reader
func (u FileUpdater) GetFile() (reader *os.File, err error) {
	return os.Open(u.FilePath)
}

func (u FileUpdater) GetFileContent() ([]byte, error) {
	file, err := u.GetFile()
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(file)
}

func (u FileUpdater) GetFileContentAsString() (string, error) {
	content, err := u.GetFileContent()
	if err != nil {
		return "", err
	}
	return string(content), nil
}

// when preHook Not nil should execute pre hook
// when preHook executed and exit not 0 return
// copy origin file for backup and update file
// when file update succeeded execute the post hook
// pre->write->post
func (u FileUpdater) UpdateFile(date io.Reader) error {
	var err error

	// pre hook
	err = u.execPreHook()
	if err != nil {
		return err
	}

	// Copy File for backup
	// todo
	var bfp string
	if u.Backup {
		bfp, err = BackupFile(u.FilePath)
		if err != nil {
			log.Println("backup err: ", err)
			return errors.New("backup origin file failed")
		}
	}
	// write new content to file
	file, err := os.OpenFile(u.FilePath, os.O_TRUNC|os.O_RDWR|os.O_SYNC, 0644)
	if err != nil {
		return errors.New("open file failed")
	}
	_, err = io.Copy(file, date)
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


	// post hook
	err = u.execPostHook()
	if err != nil {
		// todo
		log.Println("run post hook failed ", err.Error())
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

func BackupFile(filePath string) (newPath string, err error) {
	KeepBackup(filePath, 2)
	backupPath := genBackupFilePath(filePath)
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

func KeepBackup(path string, num int) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		log.Println("get  file abs path err: ", err.Error())
		return
	}
	dir := filepath.Dir(absPath)
	allFiles, err := filepath.Glob(dir + "/*")
	if err != nil {
		log.Println("get  files err: ", err.Error())
		return
	}
	var backFilesPath []string
	for _, f := range allFiles {
		if isBackupFile(f, path) {
			backFilesPath = append(backFilesPath, f)
		}
	}
	needRemoveBack := FindOldFiles(backFilesPath, len(backFilesPath)-num)
	log.Println("remove the old backup files: ", needRemoveBack)
	for _, f := range needRemoveBack {
		os.Remove(f)
	}
}

// return the most older files in given files
func FindOldFiles(Paths []string, n int) (oldFiles []string) {
	fsa := make(FilesAtime, len(Paths))
	for k, f := range Paths {
		t, err := times.Stat(f)
		// if get files atime failed use now instead
		if err != nil {
			fsa[k] = fileAtime{
				FilePath: f,
				Atime:    time.Now(),
			}
			continue
		}
		fsa[k] = fileAtime{
			FilePath: f,
			Atime:    t.ChangeTime(),
		}
	}
	// get the older
	sort.Reverse(fsa)
	if len(Paths) > n {
		for i, f := range fsa {
			if i < n {
				oldFiles = append(oldFiles, f.FilePath)
				continue
			}
			return
		}
	}
	for _, f := range fsa {
		oldFiles = append(oldFiles, f.FilePath)
	}
	return
}

// for sort
type fileAtime struct {
	FilePath string
	Atime    time.Time
}
type FilesAtime []fileAtime

func (f FilesAtime) Len() int {
	return len(f)
}
func (f FilesAtime) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}
func (f FilesAtime) Less(i, j int) bool {
	ti := f[i].Atime
	tj := f[j].Atime

	return ti.Before(tj)
}

//
func genBackupFilePath(path string) string {
	absPath, err := filepath.Abs(path)
	if err != nil {
		log.Println("err: ", err.Error())
		return path
	}
	dir := filepath.Dir(absPath)

	return dir + backupFlag(path) + time.Now().Format("2006-01-02-15:04:05")
}
func backupFlag(path string) string {
	return "/." + filepath.Base(path) + "-fub" + "."
}

// is f1 an backup of f2
func isBackupFile(f1, f2 string) bool {
	flag := backupFlag(f2)
	res := strings.Contains(f1, flag)
	return res
}
