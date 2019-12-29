package intelpower

import (
	"strings"
)

type CPUGovernor string

const (
	CPUGovernorOnDemand     = CPUGovernor("ondemand")
	CPUGovernorPowersave    = CPUGovernor("powersave")
	CPUGovernorConservative = CPUGovernor("conservative")
	CPUGovernorUserspace    = CPUGovernor("userspace")
	CPUGovernorPerformance  = CPUGovernor("performance")
)

// GetAvailableGovernors - Returns available governors
func (cpu *CPU) GetAvailableGovernors() ([]CPUGovernor, error) {
	resp, err := cpu.pwr.cmdRead(cpu.cpuRoot, "cpufreq", "scaling_available_governors")
	if err != nil {
		return nil, err
	}

	govs := strings.Split(resp, " ")
	governors := make([]CPUGovernor, len(govs))
	for i, gov := range govs {
		governors[i] = CPUGovernor(gov)
	}

	return governors, nil
}

// GetGovernor - Returns current governor
func (cpu *CPU) GetGovernor() (CPUGovernor, error) {
	resp, err := cpu.pwr.cmdRead(cpu.cpuRoot, "cpufreq", "scaling_governor")
	if err != nil {
		return "", err
	}

	return CPUGovernor(resp), nil
}

// SetGovernor - Sets governor
func (cpu *CPU) SetGovernor(governor CPUGovernor) error {
	err := cpu.pwr.cmdWrite(string(governor), cpu.cpuRoot, "cpufreq", "scaling_governor")
	if err != nil {
		return err
	}

	return nil
}
