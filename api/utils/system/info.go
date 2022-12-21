package system

import (
	"fmt"
	"math"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
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

type CPU struct {
	ID       int32  `json:"cpu"`
	Cores    int32  `json:"cores"`
	ModeName string `json:"mode"`
	MHz      string `json:"mhz"`
}

type SystemInfoCache struct {
	Mem      Mem     `json:"mem"`
	CPUs     []CPU   `json:"CPUs"`
	CPUUsage float64 `json:"cpu_usage"`

	// BootTime uint64 `json "boot_time"`
	// Platform string `json "platform"`
	// Family   string `json "family"`
}

func GetNowSystemInfo() SystemInfoCache {
	return SystemInfoCache{}
}

// get mem info from system
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

func GetNowCPUInfo() []CPU {
	// TODO add cpu usage information
	cpuInfo := []CPU{}

	infos, _ := cpu.Info()

	for _, info := range infos {
		cpuInfo = append(cpuInfo, CPU{
			ID:       info.CPU,
			Cores:    info.Cores,
			ModeName: info.ModelName,
			MHz:      fmt.Sprintf("%f", math.Round(info.Mhz/1e3)),
		})
	}

	return cpuInfo
}

func GetNowCPUUsage() (total float64) {
	perPercents, err := cpu.Percent(time.Second, false)
	if err != nil || len(perPercents) == 0 {
		return -1
	}

	return math.Round(perPercents[0])
}
