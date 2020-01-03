package intelcpu

import (
	"os"
	"path"
)

// IsOfflineAvailable - Returns cpu offline availability
//
// Some of cores, e.g. Core_0, can't be offline.
func (core *Core) IsOfflineAvailable() (bool, error) {
	_, err := os.Stat(path.Join(core.Path, "online"))
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

// IsOnline - Returns current cpu online status
//
// Some cores, like first, online always.
func (core *Core) IsOnline() (bool, error) {
	canBeOffline, err := core.IsOfflineAvailable()
	if err != nil {
		return false, err
	}

	if !canBeOffline {
		return true, nil
	}

	resp, err := statRead(core.Path, "online")
	if err != nil {
		return false, err
	}

	online := false
	if resp == "1" {
		online = true
	}

	return online, nil
}

// SetOnline - Sets cpu online status
//
// It possible to disable (make core offline) any number of cores and stay you system running on 4 cores of 8 for example.
func (core *Core) SetOnline(online bool) error {
	stat := "1"
	if !online {
		stat = "0"
	}

	err := statWrite(stat, core.Path, "online")
	if err != nil {
		return err
	}

	return nil
}
