package cpu

import "runtime"

// CPUStresser represent benchmark stress tester for CPU
type CPUStresser struct {
	vCPUCount int
}

// NewCPUStreser represent configuration and creation for CPUStresser
func NewCPUStreser() *CPUStresser {
	vCPU := runtime.NumCPU()

	return &CPUStresser{
		vCPUCount: vCPU,
	}
}

func (s *CPUStresser) CPUCount() int {
	return s.vCPUCount
}
