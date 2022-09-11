package stresser

import (
	"math/rand"
	"time"

	"github.com/teneta-io/healthchecker/pkg/model"
	"github.com/teneta-io/healthchecker/pkg/stresser/cpu"
)

// StressTester represent interface for usage in other package
type StressTester interface {
	LastResult() *model.StressTest
	Run()
}

// Stresser represent StressTester implementation
type Stresser struct {
	CPU *cpu.CPUStresser

	StressTestData *model.StressTest
}

// NewStresser create and configure new system stresser
func NewStresser() *Stresser {
	return &Stresser{
		CPU: cpu.NewCPUStreser(),
	}
}

// TODO update for real stresstest
func (s *Stresser) Run() {
	ticker := time.NewTicker(time.Second * 30)
	for {

		select {
		case <-ticker.C:
			s.StressTestData = &model.StressTest{
				VCPUCount:      s.CPU.CPUCount(),
				VCPUSpeed:      int64(rand.Intn(4000-1000) + 1000),
				VRAMCount:      int64(rand.Intn(4000-1000) + 1000),
				VRAMSpeed:      int64(rand.Intn(4000-1000) + 1000),
				VSTORAGEVolume: int64(rand.Intn(40000-1000) + 1000),
				VSTORAGESpeed:  int64(rand.Intn(4000-1000) + 1000),
				VSTORAGEIops:   int64(rand.Intn(4000-1000) + 1000),
				NetworkSpeed:   int64(rand.Intn(4000-1000) + 1000),
			}
		}
	}
}

func (s *Stresser) LastResult() *model.StressTest {
	return s.StressTestData
}
