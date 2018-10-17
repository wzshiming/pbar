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

// Print renderer
func (r *Renderer) Print() {
	r.print(r.pbs.String())
}

// Clear clear all progress bars
func (r *Renderer) Clear() {
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

	if r.isStart == 0 && r.task.Len() == 0 {
		r.isStart = 1
		r.task.Add(time.Now(), r.start)
	}
	return &n.Info
}

// SetHZ sets refresh frequency HZ
func (r *Renderer) SetHZ(hz uint8) {
	r.hz = hz
}

// start rendering until complete
func (r *Renderer) start() {
	count := r.pbs.Count()
	if count > 1 {
		r.print(strings.Repeat("\n", count-1))
	}
	r.Print()
	task.AddPeriodic(task.PeriodicInterval(0, time.Second/time.Duration(r.hz)), r.Print)
}

// Wait until complete
func (r *Renderer) Wait() {
	for !r.pbs.IsComplete() {
		r.task.Join()
	}
	r.task.CancelAll()
	r.Print()
	r.print("\n")
	r.Clear()
}
