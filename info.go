package pbar

import (
	"time"
)

// Info Current information about the progress bar
type Info struct {
	Refresh   int
	StartTime time.Time
}
