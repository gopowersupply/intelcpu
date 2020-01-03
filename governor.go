package intelcpu

import (
	"strings"
)

// CPUCoreGovernor - Core governor state
// see: https://www.kernel.org/doc/html/v4.12/admin-guide/pm/cpufreq.html#generic-scaling-governors
type CPUCoreGovernor string

const (
	// CPUGovernorOnDemand - CPU will be using low freq as long as possible but hi-freq also available if it really needed
	CPUGovernorOnDemand CPUCoreGovernor = "ondemand"
	// CPUGovernorPowersave - CPU will be using lowest frequency
	CPUGovernorPowersave CPUCoreGovernor = "powersave"
	// CPUGovernorConservative - Same as OnDemand but freq will be changing more smoothly
	CPUGovernorConservative CPUCoreGovernor = "conservative"
	// CPUGovernorUserspace - Allow to set your own freq by writing to /sys/devices/system/cpu/cpufreq/policy0/scaling_setspeed
	CPUGovernorUserspace CPUCoreGovernor = "userspace"
	// CPUGovernorPerformance - CPU will be using hi-freq as long as possible
	CPUGovernorPerformance CPUCoreGovernor = "performance"
	// CPUGovernorSchedutil - Something hard, see: https://www.kernel.org/doc/html/v4.12/admin-guide/pm/cpufreq.html#schedutil
	CPUGovernorSchedutil CPUCoreGovernor = "schedutil"
)

// GetAvailableGovernors - Returns available governors
func (core *Core) GetAvailableGovernors() ([]CPUCoreGovernor, error) {
	resp, err := StatRead(core.Path, "cpufreq", "scaling_available_governors")
	if err != nil {
		return nil, err
	}

	govs := strings.Split(resp, " ")
	governors := make([]CPUCoreGovernor, len(govs))
	for i, gov := range govs {
		governors[i] = CPUCoreGovernor(gov)
	}

	return governors, nil
}

// GetGovernor - Returns current governor
func (core *Core) GetGovernor() (CPUCoreGovernor, error) {
	resp, err := StatRead(core.Path, "cpufreq", "scaling_governor")
	if err != nil {
		return "", err
	}

	return CPUCoreGovernor(resp), nil
}

// SetGovernor - Sets governor
func (core *Core) SetGovernor(governor CPUCoreGovernor) error {
	err := StatWrite(string(governor), core.Path, "cpufreq", "scaling_governor")
	if err != nil {
		return err
	}

	return nil
}
