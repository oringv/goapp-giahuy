package models

import "context"

type Montitor interface {
	Name() string
	Check(ctx context.Context) string
}

type SystemStats struct {
	Label string
	Value string
}
