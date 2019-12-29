package intelpower

import (
	"strconv"
)

//TODO: SET? scaling_max_freq, scaling_min_freq

// GetFreq - Returns current core frequency
func (cpu *CPU) GetFreq() (uint32, error) {
	resp, err := cpu.pwr.cmdRead(cpu.cpuRoot, "cpufreq", "scaling_cur_freq")
	if err != nil {
		return 0, err
	}

	freq, err := strconv.ParseUint(resp, 10, 32)
	if err != nil {
		return 0, err
	}

	return uint32(freq), nil
}

// GetBaseFreq - Returns base core frequency
func (cpu *CPU) GetBaseFreq() (uint32, error) {
	resp, err := cpu.pwr.cmdRead(cpu.cpuRoot, "cpufreq", "base_frequency")
	if err != nil {
		return 0, err
	}

	freq, err := strconv.ParseUint(resp, 10, 32)
	if err != nil {
		return 0, err
	}

	return uint32(freq), nil
}

// GetMaxFreq - Returns max core frequency
func (cpu *CPU) GetMaxFreq() (uint32, error) {
	resp, err := cpu.pwr.cmdRead(cpu.cpuRoot, "cpufreq", "scaling_max_freq")
	if err != nil {
		return 0, err
	}

	freq, err := strconv.ParseUint(resp, 10, 32)
	if err != nil {
		return 0, err
	}

	return uint32(freq), nil
}

// GetMinFreq - Returns min core frequency
func (cpu *CPU) GetMinFreq() (uint32, error) {
	resp, err := cpu.pwr.cmdRead(cpu.cpuRoot, "cpufreq", "scaling_min_freq")
	if err != nil {
		return 0, err
	}

	freq, err := strconv.ParseUint(resp, 10, 32)
	if err != nil {
		return 0, err
	}

	return uint32(freq), nil
}
