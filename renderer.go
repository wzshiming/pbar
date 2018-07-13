package pbar

import (
	"io"
	"os"
	"strings"
	"time"
	"unsafe"

	"github.com/wzshiming/task"
)

// Renderer is render progress bar
type Renderer struct {
	pbs     rendererMark
	task    *task.Task
	hz      uint8
	isStart int
	output  io.Writer
}

// NewRenderer creating a new progress bar renderer must have only one working renderer
func NewRenderer() *Renderer {
	return &Renderer{
		task:   task.NewTask(1),
		hz:     6,
		output: os.Stdout,
	}
}

// print output
func (r *Renderer) print(s string) {
	r.output.Write(*(*[]byte)(unsafe.Pointer(&s)))
}

// Clear clear all progress bars
func (r *Renderer) Clear() {
	r.print(r.pbs.String())
	r.print("\n")
	r.task.CancelAll()
	r.pbs = r.pbs[:0]
	r.isStart = 0
}

// New adds a new default progress bar
func (r *Renderer) New(name string) *Info {
	pb := NewNormalStyle(name)
	return r.Add(pb)
}

// Add adds a new progress bar
func (r *Renderer) Add(pb Mark) *Info {
	n := &infoMark{
		Mark: pb,
	}
	r.pbs = append(r.pbs, n)
	return &n.Info
}

// SetHZ sets refresh frequency HZ
func (r *Renderer) SetHZ(hz uint8) {
	r.hz = hz
}

// Start rendering until complete
func (r *Renderer) Start() {
	if r.isStart != 0 {
		return
	}
	r.isStart = 1
	r.print(strings.Repeat("\n", r.pbs.Count()-1))
	r.print(r.pbs.String())
	var node *task.Node
	node = task.AddPeriodic(task.PeriodicInterval(0, time.Second/time.Duration(r.hz)), func() {
		r.print(r.pbs.String())
		if r.pbs.IsComplete() {
			r.print("\n")
			task.Cancel(node)
			r.isStart = 0
		}
	})
}

// Wait until complete
func (r *Renderer) Wait() {
	for !r.pbs.IsComplete() {
		r.task.Join()
	}
	r.Clear()
}
