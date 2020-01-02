package intelcpu

import (
	"errors"
	"github.com/gopowersupply/intelcpu/common"
	"github.com/gopowersupply/intelcpu/errs"
	"strconv"
)

// GetMinPerf - Returns min performance percent
func (cpu *CPU) GetMinPerf() (float32, error) {
	resp, err := common.StatRead(cpu.path, "intel_pstate", "min_perf_pct")
	if err != nil {
		return 0, err
	}

	pct, err := strconv.ParseFloat(resp, 32)
	if err != nil {
		return 0, err
	}

	return float32(pct / 100), nil
}

// SetMinPerf - Sets min performance percent
func (cpu *CPU) SetMinPerf(prc float32) error {
	if prc < 0 || prc > 1 {
		return errs.NewCPUError(errors.New("percent must be in [0..1]"))
	}

	err := common.StatWrite(strconv.Itoa(int(prc*100)), cpu.path, "intel_pstate", "min_perf_pct")
	if err != nil {
		return err
	}

	return nil
}

// GetMaxPerf - Returns max performance percent
func (cpu *CPU) GetMaxPerf() (float32, error) {
	resp, err := common.StatRead(cpu.path, "intel_pstate", "max_perf_pct")
	if err != nil {
		return 0, err
	}

	pct, err := strconv.ParseFloat(resp, 32)
	if err != nil {
		return 0, err
	}

	return float32(pct / 100), nil
}

// SetMaxPerf - Sets max performance percent
func (cpu *CPU) SetMaxPerf(prc float32) error {
	if prc < 0 || prc > 1 {
		return errs.NewCPUError(errors.New("percent must be in [0..1]"))
	}

	err := common.StatWrite(strconv.Itoa(int(prc*100)), cpu.path, "intel_pstate", "max_perf_pct")
	if err != nil {
		return err
	}

	return nil
}
