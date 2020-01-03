package intelcpu

// Way to disable half of your cores
// Remember that first core can't be disabled, so recommend to disabling the second half of cores
func ExampleCore_SetOnline() {
	cpu := New()
	cores, _ := cpu.GetCores()
	for _, core := range cores[len(cores)/2:] {
		core.SetOnline(false)
	}
}
