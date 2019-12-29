package intelpower

type CPUList []*CPU

// CPU - CPU core presentation
type CPU struct {
	pwr     *IntelPower
	cpuRoot string
	num     uint16
}