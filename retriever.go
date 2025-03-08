package main

import (
	"bytes"
	"os/exec"
	"strings"
)

type commandRetriever struct {
	command string
}

func (c *commandRetriever) GetIdentityToken() ([]byte, error) {
	var buffer bytes.Buffer

	cmdParts := strings.Split(c.command, " ") // This is not idea, but it should do for now.

	cmd := exec.Command(cmdParts[0], cmdParts[1:]...)
	cmd.Stdout = &buffer
	err := cmd.Run()

	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
