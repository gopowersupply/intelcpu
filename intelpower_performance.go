package intelpower

import (
	"errors"
	"strconv"
)

// GetMinPerf - Returns min performance percent
func (pwr *IntelPower) GetMinPerf() (float32, error) {
	resp, err := pwr.cmdRead(pwr.flRoot, "intel_pstate", "min_perf_pct")
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
func (pwr *IntelPower) SetMinPerf(prc float32) error {
	if prc < 0 || prc > 1 {
		return NewCommonError(errors.New("percent must be in [0..1]"))
	}

	err := pwr.cmdWrite(strconv.Itoa(int(prc*100)), pwr.flRoot, "intel_pstate", "min_perf_pct")
	if err != nil {
		return err
	}

	return nil
}

// GetMaxPerf - Returns max performance percent
func (pwr *IntelPower) GetMaxPerf() (float32, error) {
	resp, err := pwr.cmdRead(pwr.flRoot, "intel_pstate", "max_perf_pct")
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
func (pwr *IntelPower) SetMaxPerf(prc float32) error {
	if prc < 0 || prc > 1 {
		return NewCommonError(errors.New("percent must be in [0..1]"))
	}

	err := pwr.cmdWrite(strconv.Itoa(int(prc*100)), pwr.flRoot, "intel_pstate", "max_perf_pct")
	if err != nil {
		return err
	}

	return nil
}
