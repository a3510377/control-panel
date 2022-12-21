package system

import (
	"time"

	"github.com/shirou/gopsutil/v3/mem"
)

const (
	CacheMax      = time.Minute * 10 // cache 10 minutes
	CacheInterval = time.Second * 10 // interval
)

var SystemTimeInfo = []SystemInfoCache{}

type Mem struct {
	Total       uint64  `json:"total"`
	Available   uint64  `json:"available"`
	UsedPercent float64 `json:"used_percent"`

	STotal     string `json:"str_total"`
	SAvailable string `json:"str_available"`
}

type CPUs struct {
	ID       int    `json:"cpu"`
	Cores    int    `json:"cores"`
	ModeName string `json:"mode"`
	MHz      int    `json:"mhz"`
}

type SystemInfoCache struct {
	Mem  Mem  `json:"mem"`
	CPUs CPUs `json:"CPUs"`

	// BootTime uint64 `json "boot_time"`
	// Platform string `json "platform"`
	// Family   string `json "family"`
}

func GetNowSystemInfo() SystemInfoCache {
	return SystemInfoCache{}
}

func GetNowMemInfo() Mem {
	mem, _ := mem.VirtualMemory()

	return Mem{
		Total:       mem.Total,
		Available:   mem.Available,
		UsedPercent: mem.UsedPercent,

		STotal:     BytesString(mem.Total),
		SAvailable: BytesString(mem.Available),
	}
}
