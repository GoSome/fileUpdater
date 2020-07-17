package process

import (
	"log"
	"os"
	"os/exec"
)

type Process struct {
	Command     string `json:"command" yaml:"command"`
	AutoRestart bool   `json:"auto_restart" yaml:"auto_restart"`
	LogPath     string `json:"log_path" yaml:"log_path"`
	logfile     *os.File
	Enable      bool `json:"enable" yaml:"enable"`
	Pid         int  `json:"pid" yaml:"pid"`
}

func (p Process) Go() {
	go func() {
		err := p.Run()
		if err != nil {
			log.Println("run command ",p.Command," err: ", err.Error())
		}
	}()
}
func (p *Process) Run() error {
	cmd := exec.Command("sh", "-c", p.Command)
	if p.LogPath != "" {
		logFile, err := os.OpenFile(p.LogPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		p.logfile = logFile

	}
	if p.logfile == nil {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}
	cmd.Stderr = p.logfile
	cmd.Stdout = p.logfile
	err := cmd.Start()
	if err != nil {
		return err
	}
	p.Pid = cmd.Process.Pid

	err = cmd.Wait()
	if err != nil {
		return err
	}
	return nil
}
