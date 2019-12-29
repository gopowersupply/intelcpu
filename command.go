package intelpower

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"path"
	"strings"
)

func (pwr *IntelPower) command(cmd string) (string, error) {
	var bufError bytes.Buffer
	var bufOut bytes.Buffer

	c := exec.Command(pwr.shell[0], pwr.shell[1], cmd)
	c.Stderr = &bufError
	c.Stdout = &bufOut

	err := c.Run()
	if err != nil {
		return "", NewCommandError(err)
	}

	if bufError.Len() > 0 {
		return "", NewCommandError(errors.New(bufError.String()))
	}

	return strings.TrimSpace(bufOut.String()), nil
}

func (pwr *IntelPower) cmdRead(filepath ...string) (string, error) {
	file := path.Join(filepath...)

	resp, err := pwr.command("cat " + file)
	if err != nil {
		return "", err
	}

	return resp, nil
}

func (pwr *IntelPower) cmdWrite(data string, filepath ...string) error {
	file := path.Join(filepath...)

	_, err := pwr.command(fmt.Sprintf("echo %s > %s", data, file))
	if err != nil {
		return err
	}

	return nil
}