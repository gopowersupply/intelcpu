package intelcpu

import "testing"

func TestCPU_CheckDriver(t *testing.T) {
	cpu := New()
	if err := cpu.CheckDriver(); err != nil {
		t.Fatal(err)
	}
}

func TestCPU_GetStatus(t *testing.T) {
	cpu := New()

	status, err := cpu.GetStatus()
	if err != nil {
		t.Fatal(err)
	}

	switch status {
	case PStateStatusPassive:
		t.Log("driver status is passive")
	case PStateStatusOff:
		t.Fatal("driver status is off")
	}
}

func TestCPU_GetCores(t *testing.T) {
	cpu := New()
	if _, err := cpu.GetCores(); err != nil {
		t.Error(err)
	}
}

func TestCPU_GetCore(t *testing.T) {
	cpu := New()
	if _, err := cpu.GetCore(0); err != nil {
		t.Error(err)
	}

	if _, err := cpu.GetCore(9999); err == nil {
		t.Error(err)
	}
}
