package intelcpu

// How to know what is current max freq?
func ExampleCPU_GetMaxPerf_GetCurrentMaxFreq() {
	cpu := New()
	cores, _ := cpu.GetCores()
	maxPerf, _ := cpu.GetMaxPerf()
	maxAvailFreq, _ := cores[0].GetMaxAvailableFreq()

	// For i5-8300H maxAvailFreq - 4.0GHz.
	// Then, if maxPerf is 0.8 then maxFreq is 3.2GHz
	maxFreq := float32(maxAvailFreq) * maxPerf
	_ = maxFreq
}

// How to know what is a real lowest perf?
func ExampleCPU_SetMinPerf_RealLowestPerf() {
	cpu := New()
	cores, _ := cpu.GetCores()

	// For i5-8300H it is 800MHz
	minAvailFreq, _ := cores[0].GetMinAvailableFreq()
	// And 4GHz for max with turbo
	maxAvailFreq, _ := cores[0].GetMaxAvailableFreq()

	// It is real min perf. Value is 0.2
	minPerf := float32(minAvailFreq) / float32(maxAvailFreq)
	_ = minPerf
}