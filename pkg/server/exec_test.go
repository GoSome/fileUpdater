package server

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestExec(t *testing.T) {
	type args struct {
		body string
	}
	tests := []struct {
		name string
		args args
	}{
		{"Execute pwd", args{body: `{"shell": "pwd"}`}},
		{"Execute ls", args{body: `{"shell": "ls"}`}},
	}
	for _, tt := range tests {
		var b bytes.Buffer
		b.Write([]byte(tt.args.body))

		httpRes := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(httpRes)
		c.Request, _ = http.NewRequest("POST", "/", &b)

		t.Run(tt.name, func(t *testing.T) {
			var res execRes
			Exec(c)
			_ = json.NewDecoder(httpRes.Body).Decode(&res)

			assert.Equal(t, 0, res.ExitCode)
			assert.NotEmpty(t, res.Stdout)
			assert.Empty(t, res.Stderr)
		})
	}
}
