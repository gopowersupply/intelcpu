package intelpower

import (
	"errors"
	"fmt"
	"os"
	"path"
	"regexp"
	"runtime"
)

// IntelPower - Object to interact with Intel CPUs
type IntelPower struct {
	shell  []string
	flRoot string
}

// IntelPower - New IntelPower constructor
func New() *IntelPower {
	return &IntelPower{
		shell:  []string{"/bin/sh", "-c"},
		flRoot: "/sys/devices/system/cpu",
	}
}

// CheckDriver - Checks for Intel Power Management driver
func (pwr *IntelPower) CheckDriver() error {
	_, err := os.Stat(path.Join(pwr.flRoot, "intel_pstate"))
	if err != nil {
		if os.IsNotExist(err) {
			return NewCommonError(errors.New("intel_pstate dir isn't exist"))
		}
		return err
	}

	return nil
}

// GetStatus - Operation mode for driver. Active - all is ok. Passive - some problems. Off - driver disabled
func (pwr *IntelPower) GetStatus() (PStateStatus, error) {
	resp, err := pwr.cmdRead(pwr.flRoot, "intel_pstate", "status")
	if err != nil {
		return "", err
	}

	return PStateStatus(resp), nil
}

func (pwr *IntelPower) GetCPUName() (string, error) {
	resp, err := pwr.command("cat /proc/cpuinfo")
	if err != nil {
		return "", err
	}

	modelNameRegExp := regexp.MustCompilePOSIX(`^model name	: (.+)$`)
	modelNameParsed := modelNameRegExp.FindStringSubmatch(resp)

	if len(modelNameParsed) > 1 {
		return modelNameParsed[1], nil
	}

	return "", NewCommonError(errors.New("can't find model name in cpuinfo"))
}

// GetCPU - Returns CPU representation
func (pwr *IntelPower) GetCPU(num uint16) (*CPU, error) {
	path := fmt.Sprintf("%s%d", path.Join(pwr.flRoot, "cpu"), num)

	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, NewCommonError(errors.New("incorrect cpu number"))
		}
		return nil, err
	}

	cpu := &CPU{
		pwr:     pwr,
		cpuRoot: path,
		num:     num,
	}

	return cpu, nil
}

// GetCPUs - Returns representation for all CPUs
func (pwr *IntelPower) GetCPUs() (CPUList, error) {
	cpus := make(CPUList, runtime.NumCPU())

	for i := 0; i < len(cpus); i++ {
		cpu, err := pwr.GetCPU(uint16(i))
		if err != nil {
			return nil, err
		}

		cpus[i] = cpu
	}

	return cpus, nil
}
