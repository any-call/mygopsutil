package mygopsutil

import (
	"testing"
	"time"
)

func TestGetCPUUsage(t *testing.T) {
	ret, err := GetCPUUsage()
	if err != nil {
		t.Error(err)
		return
	}

	t.Log("cpu percent :", ret)
}

func TestGetMemUsage(t *testing.T) {
	total, used, usage, err := GetMemUsage()
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("total is :%.2fMB,used:%.2fMB,usage:%.2f",
		float64(total)/1024.0/1024.0,
		float64(used)/1024.0/1204.0,
		usage)
}

func TestGetNetSpeed(t *testing.T) {
	rx, tx, err := GetTotalNetSpeed(time.Second * 5)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log("tx is :", tx, ";rx is :", rx)
}
