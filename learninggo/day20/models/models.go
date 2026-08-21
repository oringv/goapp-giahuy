package models

import "sync"

type SystemStats struct {
	Label   string
	Value   string
	IsAlert bool
}

type ProStat struct {
	PID         int32
	Name        string
	CPU         float64
	Memory      uint64
	RamPercent  float64
	RunningTime string
}

var (
	Stats      = make(map[string]SystemStats)
	StatsMutex sync.Mutex
)
