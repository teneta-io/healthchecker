package model

type StressTest struct {
	VCPUCount      int   `json:"cpu_count"`
	VCPUSpeed      int64 `json:"cpu_speed"`
	VRAMCount      int64 `json:"ram_count"`
	VRAMSpeed      int64 `json:"ram_speed"`
	VSTORAGEVolume int64 `json:"storage_volume"`
	VSTORAGESpeed  int64 `json:"storage_speed"`
	VSTORAGEIops   int64 `json:"storage_iops"`
	NetworkSpeed   int64 `json:"network_speed"`
}
