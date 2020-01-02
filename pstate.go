package intelcpu

import (
	"intelcpu/common"
	"strconv"
)

// PStateStatus - Status of intel_pstate driver
type PStateStatus string

const (
	// PStateStatusActive - Active status means that all is ok and driver is ready
	PStateStatusActive PStateStatus = "active"
	// PStateStatusPassive - Passive status means that something wrong and driver working partially
	PStateStatusPassive PStateStatus = "passive"
	// PStateStatusOff - Off status means that driver disabled and the package won't work
	PStateStatusOff PStateStatus = "off"
)

// GetPStatesNum - Returns number of P-States
func (cpu *CPU) GetPStatesNum() (uint8, error) {
	resp, err := common.StatRead(cpu.path, "intel_pstate", "num_pstates")
	if err != nil {
		return 0, err
	}

	states, err := strconv.ParseUint(resp, 10, 8)
	if err != nil {
		return 0, err
	}

	return uint8(states), nil
}
