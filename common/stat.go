package common

import (
	"io/ioutil"
	"path"
)

// StatRead - Reads value from file
func StatRead(filepath ...string) (string, error) {
	file := path.Join(filepath...)

	data, err := ioutil.ReadFile(file)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

// StatWrite - Writes value to file
func StatWrite(data string, filepath ...string) error {
	file := path.Join(filepath...)

	if err := ioutil.WriteFile(file, []byte(data), 222); err != nil {
		return err
	}

	return nil
}
