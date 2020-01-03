package intelcpu

// CoreList - Array of cores, some actions can be applied to all cores at the same time
type CoreList []*Core

// Core - Core presentation
type Core struct {
	Path string
	Num  uint16
}

// SetGovernor - Sets governor to all cores
func (list *CoreList) SetGovernor(governor CPUCoreGovernor) error {
	for _, core := range *list {
		if err := core.SetGovernor(governor); err != nil {
			return err
		}
	}

	return nil
}

// SetPreference - Sets preference to all cores
func (list *CoreList) SetPreference(preference CPUPreference) error {
	for _, core := range *list {
		if err := core.SetPreference(preference); err != nil {
			return err
		}
	}

	return nil
}
