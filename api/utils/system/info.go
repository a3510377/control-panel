package system

import (
	"fmt"
	"math"
	"os"
	"time"

	"github.com/a3510377/control-panel/utils/JTime"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
)

type Mem struct {
	Total       uint64  `json:"total"`
	Available   uint64  `json:"available"`
	UsedPercent float64 `json:"-"`

	STotal     string `json:"str_total"`
	SAvailable string `json:"str_available"`
	SUsed      string `json:"str_used"`
}

type CPU struct {
	ID       int32  `json:"cpu"`
	Cores    int32  `json:"cores"`
	ModeName string `json:"mode"`
	MHz      string `json:"mhz"`
}

type Host struct {
	Name     string     `json:"name"`
	Platform string     `json:"platform"`
	Version  string     `json:"version"`
	BootTime JTime.Time `json:"boot_time"`
}

type SystemInfo struct {
	Time     time.Time `json:"time"`
	Mem      Mem       `json:"mem"`
	CPUUsage float64   `json:"cpu_usage"`
	MEMUsage float64   `json:"mem_usage"`
}

type SystemBaseInfo struct {
	CPUs       []CPU      `json:"CPUs"`
	Host       Host       `json:"host"`
	SystemTime JTime.Time `json:"system_time"`
}

func GetNowSystemInfo() SystemInfo {
	MemData := GetNowMemInfo()
	return SystemInfo{
		Time:     time.Now(),
		Mem:      MemData,
		CPUUsage: GetNowCPUUsage(),
		MEMUsage: MemData.UsedPercent,
	}
}

func GetSystemBaseInfo() SystemBaseInfo {
	return SystemBaseInfo{
		CPUs:       GetNowCPUInfo(),
		Host:       GetNowHostInfo(),
		SystemTime: GetNowSystemTime(),
	}
}

// return system time
func GetNowSystemTime() JTime.Time { return JTime.Now() }

// get mem info from system
func GetNowMemInfo() Mem {
	mem, _ := mem.VirtualMemory()

	return Mem{
		Total:       mem.Total,
		Available:   mem.Available,
		UsedPercent: mem.UsedPercent,

		STotal:     BytesString(mem.Total),
		SAvailable: BytesString(mem.Available),
		SUsed:      BytesString(mem.Used),
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
			MHz:      fmt.Sprintf("%.1f", math.Round(info.Mhz/1e3)),
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

func GetNowHostInfo() Host {
	timestamp, _ := host.BootTime()
	platform, _, version, _ := host.PlatformInformation()
	name, _ := os.Hostname()

	return Host{
		Name:     name,
		Platform: platform,
		Version:  version,
		BootTime: JTime.Time(time.Unix(int64(timestamp), 0).Local()),
	}
}
