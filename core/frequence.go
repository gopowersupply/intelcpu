package core

import (
	"github.com/gopowersupply/intelcpu/common"
	"strconv"
)

// GetFreq - Returns current core frequency
func (core *Core) GetFreq() (uint32, error) {
	resp, err := common.StatRead(core.Path, "cpufreq", "scaling_cur_freq")
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
func (core *Core) GetBaseFreq() (uint32, error) {
	resp, err := common.StatRead(core.Path, "cpufreq", "base_frequency")
	if err != nil {
		return 0, err
	}

	freq, err := strconv.ParseUint(resp, 10, 32)
	if err != nil {
		return 0, err
	}

	return uint32(freq), nil
}

// GetMaxAvailableFreq - Returns max available core frequency
func (core *Core) GetMaxAvailableFreq() (uint32, error) {
	resp, err := common.StatRead(core.Path, "cpufreq", "cpuinfo_max_freq")
	if err != nil {
		return 0, err
	}

	freq, err := strconv.ParseUint(resp, 10, 32)
	if err != nil {
		return 0, err
	}

	return uint32(freq), nil
}

// GetMinAvailableFreq - Returns min available core frequency
func (core *Core) GetMinAvailableFreq() (uint32, error) {
	resp, err := common.StatRead(core.Path, "cpufreq", "cpuinfo_min_freq")
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
func (core *Core) GetMaxFreq() (uint32, error) {
	resp, err := common.StatRead(core.Path, "cpufreq", "scaling_max_freq")
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
func (core *Core) GetMinFreq() (uint32, error) {
	resp, err := common.StatRead(core.Path, "cpufreq", "scaling_min_freq")
	if err != nil {
		return 0, err
	}

	freq, err := strconv.ParseUint(resp, 10, 32)
	if err != nil {
		return 0, err
	}

	return uint32(freq), nil
}

// SetMaxFreq - Sets max core frequency
func (core *Core) SetMaxFreq(freq uint32) error {
	err := common.StatWrite(strconv.FormatUint(uint64(freq), 10), core.Path, "cpufreq", "scaling_max_freq")
	if err != nil {
		return err
	}

	return nil
}

// SetMinFreq - Sets min core frequency
func (core *Core) SetMinFreq(freq uint32) error {
	err := common.StatWrite(strconv.FormatUint(uint64(freq), 10), core.Path, "cpufreq", "scaling_min_freq")
	if err != nil {
		return err
	}

	return nil
}