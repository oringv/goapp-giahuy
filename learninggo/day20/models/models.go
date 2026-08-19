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

var Stats = make(map[string]SystemStats)
var StatsMutex sync.Mutex
