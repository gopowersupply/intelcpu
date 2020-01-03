package intelcpu

import (
	"io/ioutil"
	"path"
)

// statRead - Reads value from file
func statRead(filepath ...string) (string, error) {
	file := path.Join(filepath...)

	data, err := ioutil.ReadFile(file)
	if err != nil {
		return "", NewCPUError(err)
	}

	return string(data), nil
}

// statWrite - Writes value to file
func statWrite(data string, filepath ...string) error {
	file := path.Join(filepath...)

	if err := ioutil.WriteFile(file, []byte(data), 222); err != nil {
		return NewCPUError(err)
	}

	return nil
}
