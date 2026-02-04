package mygopsutil

import (
	"fmt"
	"strings"
	"time"

	"github.com/shirou/gopsutil/v4/net"
)

func GetTotalNetBytes() (rx uint64, tx uint64, err error) {
	ios, err := net.IOCounters(true)
	if err != nil {
		return 0, 0, err
	}

	for _, io := range ios {
		name := io.Name
		// 排除不属于物理网卡的接口
		if name == "lo" ||
			strings.HasPrefix(name, "docker") ||
			strings.HasPrefix(name, "veth") ||
			strings.HasPrefix(name, "tun") ||
			strings.HasPrefix(name, "wg") {
			continue
		}

		rx += io.BytesRecv
		tx += io.BytesSent
	}

	return rx, tx, nil
}

// 读取网卡两次间隔流量差值，计算出上行，下行速率
func GetTotalNetSpeed(interval time.Duration) (rxBps, txBps uint64, err error) {
	rx1, tx1, err := GetTotalNetBytes()
	if err != nil {
		return 0, 0, err
	}

	time.Sleep(interval)

	rx2, tx2, err := GetTotalNetBytes()
	if err != nil {
		return 0, 0, err
	}

	// 防止计数器重置导致负数
	if rx2 < rx1 || tx2 < tx1 {
		return 0, 0, fmt.Errorf("net counter reset detected")
	}

	sec := interval.Seconds()
	if sec <= 0 {
		return 0, 0, fmt.Errorf("invalid interval")
	}

	rxBps = uint64(float64(rx2-rx1) / sec)
	txBps = uint64(float64(tx2-tx1) / sec)

	return rxBps, txBps, nil
}
