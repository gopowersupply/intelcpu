package intelpower

import (
	"os"
	"path"
)

// IsOfflineAvailable - Returns cpu offline availability
func (cpu *CPU) IsOfflineAvailable() (bool, error) {
	_, err := os.Stat(path.Join(cpu.cpuRoot, "online"))

	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

// IsOnline - Returns current cpu online status
func (cpu *CPU) IsOnline() (bool, error) {
	canBeOffline, err := cpu.IsOfflineAvailable()
	if err != nil {
		return false, err
	}

	if !canBeOffline {
		return true, nil
	}

	resp, err := cpu.pwr.cmdRead(cpu.cpuRoot, "online")
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
func (cpu *CPU) SetOnline(online bool) error {
	stat := "1"
	if !online {
		stat = "0"
	}

	err := cpu.pwr.cmdWrite(stat, cpu.cpuRoot, "online")
	if err != nil {
		return err
	}

	return nil
}