package cpu

import (
	"errors"
	"fmt"
	"github.com/gopowersupply/intelcpu/common"
	"github.com/gopowersupply/intelcpu/core"
	"os"
	"path"
	"runtime"
)

// CPU - Object to interact with Intel CPUs
type CPU struct {
	path string
}

// New - Returns new CPU representation
func New() *CPU {
	return &CPU{
		path: "/sys/devices/system/cpu",
	}
}

// CheckDriver - Checks for Intel Power Management driver
func (cpu *CPU) CheckDriver() error {
	_, err := os.Stat(path.Join(cpu.path, "intel_pstate"))
	if err != nil {
		if os.IsNotExist(err) {
			return NewCPUError(errors.New("intel_pstate dir isn't exist"))
		}
		return err
	}

	return nil
}

// GetStatus - Operation mode for driver. Active - all is ok. Passive - some problems. Off - driver disabled
func (cpu *CPU) GetStatus() (PStateStatus, error) {
	resp, err := common.StatRead(cpu.path, "intel_pstate", "status")
	if err != nil {
		return "", NewCPUError(err)
	}

	return PStateStatus(resp), nil
}

// GetCore - Returns core representation
func (cpu *CPU) GetCore(num uint16) (*core.Core, error) {
	path := fmt.Sprintf("%s%d", path.Join(cpu.path, "cpu"), num)

	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, NewCPUError(errors.New("incorrect cpu number"))
		}
		return nil, err
	}

	core := &core.Core{
		Path: path,
		Num:  num,
	}

	return core, nil
}

// GetCores - Returns representation for all cores
func (cpu *CPU) GetCores() (core.List, error) {
	cpus := make(core.List, runtime.NumCPU())

	for i := 0; i < len(cpus); i++ {
		cpu, err := cpu.GetCore(uint16(i))
		if err != nil {
			return nil, err
		}

		cpus[i] = cpu
	}

	return cpus, nil
}
