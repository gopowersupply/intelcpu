package intelpower

import (
	"strconv"
)

// GetTemp - Returns ACPI temperature
func (pwr *IntelPower) GetTemp() (float32, error) {
	resp, err := pwr.cmdRead("/sys/bus/acpi/devices/LNXTHERM:00/thermal_zone/temp")
	if err != nil {
		return 0, err
	}

	temp, err := strconv.ParseFloat(resp, 32)
	if err != nil {
		return 0, err
	}

	return float32(temp/1000), nil
}
