package intelpower

import (
	"fmt"
)

// IsTurbo - TurboBoost status
func (pwr *IntelPower) IsTurbo() (bool, error) {
	resp, err := pwr.cmdRead(pwr.flRoot, "intel_pstate", "no_turbo")
	if err != nil {
		return false, err
	}

	switch resp {
	case "1":
		return false, nil
	case "0":
		return true, nil
	}

	return false, NewCommonError(fmt.Errorf("unknown no_turbo status: %s", resp))
}

// SetTurbo - Changes TurboBoost status
func (pwr *IntelPower) SetTurbo(status bool) error {
	stat := "1"
	if status {
		stat = "0"
	}

	err := pwr.cmdWrite(stat, pwr.flRoot, "intel_pstate", "no_turbo")
	if err != nil {
		return err
	}

	return nil
}