package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
)

type execReq struct {
	Shell string `json:"shell"`
}

type execRes struct {
	ExitCode int    `json:"exit_code"`
	Stdout   string `json:"stdout"`
	Stderr   string `json:"stderr"`
}

func (a *App) Exec(c *gin.Context) {
	var req execReq
	if err := json.NewDecoder(c.Request.Body).Decode(&req); err != nil {
		c.String(http.StatusBadRequest, "Invalid body format")
	}

	var res execRes
	var stdout, stderr bytes.Buffer
	cmd := exec.Command("bash", "-c", fmt.Sprintf("\"%s\"", req.Shell))
	cmd.Stderr = &stderr
	cmd.Stdout = &stdout

	_ = cmd.Run()

	res.ExitCode = cmd.ProcessState.ExitCode()
	res.Stdout = stdout.String()
	res.Stderr = stderr.String()

	c.JSON(http.StatusOK, res)
}
