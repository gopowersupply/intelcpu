package intelcpu

import (
	"errors"

	"strconv"
)

// GetMinPerf - Returns min performance percent
//
// You can set any value as minPerf, but actually, CPU will not work slower than GetMinAvailableFreq().
// And value here will not be lower than 0.2 (example for i5-8300H).
func (cpu *CPU) GetMinPerf() (float32, error) {
	resp, err := statRead(cpu.path, "intel_pstate", "min_perf_pct")
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
//
// Value must be between 0 and 1.
// Actually, even if you set 0 then core freq will not be lower than GetMinAvailableFreq().
// If you want to fully disable core see: SetOnline().
func (cpu *CPU) SetMinPerf(prc float32) error {
	if prc < 0 || prc > 1 {
		return NewCPUError(errors.New("percent must be in [0..1]"))
	}

	err := statWrite(strconv.Itoa(int(prc*100)), cpu.path, "intel_pstate", "min_perf_pct")
	if err != nil {
		return err
	}

	return nil
}

// GetMaxPerf - Returns max performance percent
//
// It returns number between 0 and 1.
// Where 1 means max freq for this CPU (include TurboBoost).
// So, if you have i5-8300H where base freq is 2.3GHz and turbo is 4.0GHz then 1 mean 4.0GHz.
// If you want to get current max freq, then you should get GetMaxAvailableFreq() for any core and mult it to this percent.
func (cpu *CPU) GetMaxPerf() (float32, error) {
	resp, err := statRead(cpu.path, "intel_pstate", "max_perf_pct")
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
//
// Value must be between 0 and 1 where 1 is max CPU freq with TurboBoost.
// If you disable turbo via SetTurbo(false) then 1 will be mean max freq with turbo anyway.
//
// For example, take an i5-8300H.
// It max freq with turbo is 4.0GHz.
// Max without turbo 2.3GHz.
// And if we make turbo off and set perf to 0.8 (3.2GHz) then CPU will not be faster than 2.3GHz.
func (cpu *CPU) SetMaxPerf(prc float32) error {
	if prc < 0 || prc > 1 {
		return NewCPUError(errors.New("percent must be in [0..1]"))
	}

	err := statWrite(strconv.Itoa(int(prc*100)), cpu.path, "intel_pstate", "max_perf_pct")
	if err != nil {
		return err
	}

	return nil
}
