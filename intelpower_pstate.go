package intelpower

import (
	"path"
	"strconv"
)

type PStateStatus string

const (
	PStateStatusActive  = PStateStatus("active")
	PStateStatusPassive = PStateStatus("passive")
	PStateStatusOff     = PStateStatus("off")
)

// GetPStatesNum - Returns number of P-States
func (pwr *IntelPower) GetPStatesNum() (uint8, error) {
	path := path.Join(pwr.flRoot, "intel_pstate", "num_pstates")

	resp, err := pwr.command("cat " + path)
	if err != nil {
		return 0, err
	}

	states, err := strconv.ParseUint(resp, 10, 8)
	if err != nil {
		return 0, err
	}

	return uint8(states), nil
}
