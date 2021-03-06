package intelcpu

import (
	"strings"
)

// CPUPreference - CPU performance mode
//
// In simple, it means how fast the CPU will be accelerating.
// see: https://www.kernel.org/doc/html/v4.13/admin-guide/pm/intel_pstate.html
type CPUPreference string

const (
	// CPUPreferenceDefault - Default mode. OS will be select modes based on power mode.
	CPUPreferenceDefault CPUPreference = "default"
	// CPUPreferencePerformance - Max performance. CPU will be acceleration as fast as possible
	CPUPreferencePerformance CPUPreference = "performance"
	// CPUPreferenceBalancePerformance - Performance preferred. CPU will be accelerating in middle speed
	CPUPreferenceBalancePerformance CPUPreference = "balance_performance"
	// CPUPreferenceBalancePower - Power saving preferred. CPU will be accelerating lower than middle speed
	CPUPreferenceBalancePower CPUPreference = "balance_power"
	// CPUPreferencePower - Power saving. CPU will be accelerating as low as possible
	CPUPreferencePower CPUPreference = "power"
)

// GetAvailablePreferences - Returns available performance preferences
func (core *Core) GetAvailablePreferences() ([]CPUPreference, error) {
	resp, err := statRead(core.Path, "cpufreq", "energy_performance_available_preferences")
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
	resp, err := statRead(core.Path, "cpufreq", "energy_performance_preference")
	if err != nil {
		return "", err
	}

	return CPUPreference(resp), nil
}

// SetPreference - Sets power preference
func (core *Core) SetPreference(preference CPUPreference) error {
	err := statWrite(string(preference), core.Path, "cpufreq", "energy_performance_preference")
	if err != nil {
		return err
	}

	return nil
}
