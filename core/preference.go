package core

import (
	"intelcpu/common"
	"strings"
)

// CPUPreference - CPU performance mode
// see: https://www.kernel.org/doc/html/v4.13/admin-guide/pm/intel_pstate.html
type CPUPreference string

const (
	// CPUPreferenceDefault - Default mode
	CPUPreferenceDefault CPUPreference = "default"
	// CPUPreferencePerformance - CPU will be acceleration as fast as possible
	CPUPreferencePerformance CPUPreference = "performance"
	// CPUPreferenceBalancePerformance - CPU will be accelerating in middle speed
	CPUPreferenceBalancePerformance CPUPreference = "balance_performance"
	// CPUPreferenceBalancePower - CPU will be accelerating lower than middle speed
	CPUPreferenceBalancePower CPUPreference = "balance_power"
	// CPUPreferencePower - CPU will be accelerating as low as possible
	CPUPreferencePower CPUPreference = "power"
)

// GetAvailablePreferences - Returns available performance preferences
func (core *Core) GetAvailablePreferences() ([]CPUPreference, error) {
	resp, err := common.StatRead(core.Path, "cpufreq", "energy_performance_available_preferences")
	if err != nil {
		return nil, err
	}

	prefs := strings.Split(resp, " ")
	preferences := make([]CPUPreference, len(prefs))
	for i, pref := range prefs {
		preferences[i] = CPUPreference(pref)
	}

	return preferences, nil
}

// GetPreference - Returns current power preference
func (core *Core) GetPreference() (CPUPreference, error) {
	resp, err := common.StatRead(core.Path, "cpufreq", "energy_performance_preference")
	if err != nil {
		return "", err
	}

	return CPUPreference(resp), nil
}

// SetPreference - Sets power preference
func (core *Core) SetPreference(preference CPUPreference) error {
	err := common.StatWrite(string(preference), core.Path, "cpufreq", "energy_performance_preference")
	if err != nil {
		return err
	}

	return nil
}
