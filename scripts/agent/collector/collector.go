package collector

import (
	"os"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"

	"monitoring-system/scripts/agent/model"
)

func GetHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		return "unknown"
	}
	return hostname
}

func GetCPU() float64 {
	percent, err := cpu.Percent(0, false)
	if err != nil || len(percent) == 0 {
		return 0
	}
	return percent[0]
}

func GetMemory() float64 {
	vmStat, err := mem.VirtualMemory()
	if err != nil {
		return 0
	}
	return vmStat.UsedPercent
}

func GetDisk() float64 {
	diskStat, err := disk.Usage("C:\\") // Windows
	if err != nil {
		return 0
	}
	return diskStat.UsedPercent
}

func Collect() model.Metric {
	return model.Metric{
		Hostname: GetHostname(),
		CPU:      GetCPU(),
		Memory:   GetMemory(),
		Disk:     GetDisk(),
	}
}
