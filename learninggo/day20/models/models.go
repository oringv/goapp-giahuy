package models

import (
	"context"
	"sync"
)

type Montitor interface {
	Name() string
	Check(ctx context.Context) string
}

type SystemStats struct {
	Label string
	Value string
}

type ProStat struct {
	PID         int32
	Name        string
	CPU         float64
	Memory      uint64
	RamPercent  float64
	RunningTime string
}

var Stats = make(map[string]SystemStats)
var StatsMutex sync.Mutex
