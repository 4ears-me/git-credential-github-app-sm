package main

import (
	"bytes"
	"os/exec"
)

type commandRetriever struct {
	command string
}

func (c *commandRetriever) GetIdentityToken() ([]byte, error) {
	var buffer bytes.Buffer

	cmd := exec.Command(c.command)
	cmd.Stdout = &buffer
	err := cmd.Run()

	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
