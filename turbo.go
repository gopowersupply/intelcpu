package intelcpu

import (
	"fmt"
)

// IsTurbo - TurboBoost status
func (cpu *CPU) IsTurbo() (bool, error) {
	resp, err := statRead(cpu.path, "intel_pstate", "no_turbo")
	if err != nil {
		return false, err
	}

	switch resp {
	case "1":
		return false, nil
	case "0":
		return true, nil
	}

	return false, NewCPUError(fmt.Errorf("unknown no_turbo status: %s", resp))
}

// SetTurbo - Changes TurboBoost status
func (cpu *CPU) SetTurbo(status bool) error {
	stat := "1"
	if status {
		stat = "0"
	}

	err := statWrite(stat, cpu.path, "intel_pstate", "no_turbo")
	if err != nil {
		return err
	}

	return nil
}
