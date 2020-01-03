package intelcpu

import (
	"strconv"
)

// GetFreq - Returns current core frequency in KHz
//
// For i5-8300H value will be between 800000 and 4000000 with turbo and max at 2300000 without.
func (core *Core) GetFreq() (uint32, error) {
	resp, err := statRead(core.Path, "cpufreq", "scaling_cur_freq")
	if err != nil {
		return 0, err
	}

	freq, err := strconv.ParseUint(resp, 10, 32)
	if err != nil {
		return 0, err
	}

	return uint32(freq), nil
}

// GetBaseFreq - Returns base core frequency in KHz
//
// For i5-8300H it will be 2300000
func (core *Core) GetBaseFreq() (uint32, error) {
	resp, err := statRead(core.Path, "cpufreq", "base_frequency")
	if err != nil {
		return 0, err
	}

	freq, err := strconv.ParseUint(resp, 10, 32)
	if err != nil {
		return 0, err
	}

	return uint32(freq), nil
}

// GetMaxAvailableFreq - Returns max available core frequency in KHz
//
// For i5-8300H it will be 4000000, even if turbo is off.
func (core *Core) GetMaxAvailableFreq() (uint32, error) {
	resp, err := statRead(core.Path, "cpufreq", "cpuinfo_max_freq")
	if err != nil {
		return 0, err
	}

	freq, err := strconv.ParseUint(resp, 10, 32)
	if err != nil {
		return 0, err
	}

	return uint32(freq), nil
}

// GetMinAvailableFreq - Returns min available core frequency in KHz
//
// For i5-8300H it will be 800000.
func (core *Core) GetMinAvailableFreq() (uint32, error) {
	resp, err := statRead(core.Path, "cpufreq", "cpuinfo_min_freq")
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
//
// Deprecated: In most cases same as GetMaxAvailableFreq()
// Will be works only with userspace governor
func (core *Core) GetMaxFreq() (uint32, error) {
	resp, err := statRead(core.Path, "cpufreq", "scaling_max_freq")
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
//
// Deprecated: In most cases same as GetMinAvailableFreq()
// Will be works only with userspace governor
func (core *Core) GetMinFreq() (uint32, error) {
	resp, err := statRead(core.Path, "cpufreq", "scaling_min_freq")
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
//
// Deprecated: Will be works only with userspace governor
func (core *Core) SetMaxFreq(freq uint32) error {
	err := statWrite(strconv.FormatUint(uint64(freq), 10), core.Path, "cpufreq", "scaling_max_freq")
	if err != nil {
		return err
	}

	return nil
}

// SetMinFreq - Sets min core frequency
//
// Deprecated: Will be works only with userspace governor
func (core *Core) SetMinFreq(freq uint32) error {
	err := statWrite(strconv.FormatUint(uint64(freq), 10), core.Path, "cpufreq", "scaling_min_freq")
	if err != nil {
		return err
	}

	return nil
}

// GetSpeed - Returns core speed
//
// Deprecated: Will be works only with userspace governor
func (core *Core) GetSpeed() (uint32, error) {
	resp, err := statRead(core.Path, "cpufreq", "scaling_setspeed")
	if err != nil {
		return 0, err
	}

	freq, err := strconv.ParseUint(resp, 10, 32)
	if err != nil {
		return 0, err
	}

	return uint32(freq), nil
}

// SetSpeed - Sets core speed
//
// Deprecated: Will be works only with userspace governor
func (core *Core) SetSpeed(freq uint32) error {
	err := statWrite(strconv.FormatUint(uint64(freq), 10), core.Path, "cpufreq", "scaling_setspeed")
	if err != nil {
		return err
	}

	return nil
}
