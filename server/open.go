package server

import (
	"bytes"
	"errors"
	"io"
	"os/exec"
	"runtime"
)

var cmd map[string]string
var stderr bytes.Buffer

func init() {
	cmd = map[string]string{
		"windows": "start",
		"darwin":  "open",
		"linux":   "xdg-open",
	}
}

func Open(url string) (io.Writer, error) {
	if open, err := cmd[runtime.GOOS]; !err {
		return nil, errors.New("This operating system is not supported.")
	} else {
		cmd := exec.Command(open, url)
		cmd.Stderr = &stderr
		if err := cmd.Run(); err != nil {
			return cmd.Stderr, err
		}
	}
	return nil, nil
}
