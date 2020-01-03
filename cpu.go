// Package intelcpu provides interaction with intel_pstate driver to drive CPU frequency and politics.
package intelcpu

import (
	"errors"
	"fmt"
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
//
// Nil means that driver installed, but it should be additional checked by GetStatus().
// Most of non-nil errors mean that intel_pstate driver is not installed.
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

// GetStatus - Operation mode for driver
//
// Statuses
//
// Active - all is ok, driver ready.
//
// Passive - some problems, some of features may not working or ignored.
// Usually it caused when intel_pstate can't recognize the CPU.
//
// Off - driver disabled, nothing to work.
func (cpu *CPU) GetStatus() (PStateStatus, error) {
	resp, err := statRead(cpu.path, "intel_pstate", "status")
	if err != nil {
		return "", NewCPUError(err)
	}

	return PStateStatus(resp), nil
}

// GetCore - Returns core representation
func (cpu *CPU) GetCore(num uint16) (*Core, error) {
	path := fmt.Sprintf("%s%d", path.Join(cpu.path, "cpu"), num)

	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, NewCPUError(errors.New("incorrect cpu number"))
		}
		return nil, err
	}

	core := &Core{
		Path: path,
		Num:  num,
	}

	return core, nil
}

// GetCores - Returns representation for all cores
func (cpu *CPU) GetCores() (CoreList, error) {
	cpus := make(CoreList, runtime.NumCPU())

	for i := 0; i < len(cpus); i++ {
		cpu, err := cpu.GetCore(uint16(i))
		if err != nil {
			return nil, err
		}

		cpus[i] = cpu
	}

	return cpus, nil
}
