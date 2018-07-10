package pbar

import (
	"fmt"
	"strings"
	"time"

	"github.com/wzshiming/task"
)

type Renderer struct {
	pbs     ProgressBarGroup
	task    *task.Task
	hz      uint8
	isStart int
}

func NewRenderer() *Renderer {
	return &Renderer{
		task: task.NewTask(1),
		hz:   6,
	}
}

func (r *Renderer) Clear() {
	fmt.Println(r.pbs.Format())
	r.pbs = r.pbs[:0]
}

func (r *Renderer) New(name string) *BaseInfo {
	pb := NewNormalStyle(name)
	return r.Add(pb)
}

func (r *Renderer) Add(pb ProgressBar) *BaseInfo {
	r.pbs = append(r.pbs, pb)
	task.Add(time.Now(), r.start)
	return pb.Info()
}

func (r *Renderer) start() {
	if r.isStart != 0 {
		return
	}
	r.isStart = 1
	fmt.Print(strings.Repeat("\n", r.pbs.Count()-1), r.pbs.Format())
	var node *task.Node
	node = task.AddPeriodic(task.PeriodicInterval(0, time.Second/time.Duration(r.hz)), func() {
		fmt.Print(r.pbs.Format())
		if r.pbs.IsComplete() {
			fmt.Print("\n")
			task.Cancel(node)
			r.isStart = 0
		}
	})
}

func (r *Renderer) Wait() {
	for !r.pbs.IsComplete() {
		r.task.Join()
	}
	r.Clear()
}
