package mygopsutil

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/mem"
)

// 采集CPU 使用率
func GetCPUUsage() (float64, error) {
	percent, err := cpu.Percent(time.Second, false)
	if err != nil {
		return 0, fmt.Errorf("cpu.Percent failed: %w", err)
	}

	if len(percent) == 0 {
		return 0, fmt.Errorf("cpu.Percent returned empty")
	}

	return percent[0], nil
}

func GetMemUsage() (total, used uint64, usage float64, err error) {
	vm, err := mem.VirtualMemory()
	if err != nil {
		return 0, 0, 0, fmt.Errorf("mem.VirtualMemory failed: %w", err)
	}

	return vm.Total, vm.Used, vm.UsedPercent, nil
}

