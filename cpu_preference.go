package intelpower

import (
	"strings"
)

type CPUPreference string

const (
	CPUPreferenceDefault            = CPUPreference("default")
	CPUPreferencePerformance        = CPUPreference("performance")
	CPUPreferenceBalancePerformance = CPUPreference("balance_performance")
	CPUPreferenceBalancePower       = CPUPreference("balance_power")
	CPUPreferencePower              = CPUPreference("power")
)

// GetAvailablePreferences - Returns available performance preferences
func (cpu *CPU) GetAvailablePreferences() ([]CPUPreference, error) {
	resp, err := cpu.pwr.cmdRead(cpu.cpuRoot, "cpufreq", "energy_performance_available_preferences")
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
func (cpu *CPU) GetPreference() (CPUPreference, error) {
	resp, err := cpu.pwr.cmdRead(cpu.cpuRoot, "cpufreq", "energy_performance_preference")
	if err != nil {
		return "", err
	}

	return CPUPreference(resp), nil
}

// SetPreference - Sets power preference
func (cpu *CPU) SetPreference(preference CPUPreference) error {
	err := cpu.pwr.cmdWrite(string(preference), cpu.cpuRoot, "cpufreq", "energy_performance_preference")
	if err != nil {
		return err
	}

	return nil
}