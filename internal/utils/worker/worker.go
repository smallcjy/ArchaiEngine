package worker

import (
	"github.com/panjf2000/ants/v2"
	"github.com/panjf2000/gnet/pkg/pool/goroutine"
)

// Dispatcher 调度器
type Dispatcher struct {
	//workerList []*Worker
	// 调度的工作队列池
	workerPool *ants.Pool
}

func NewDispatcher() (*Dispatcher, error) {
	pool, err := NewWorkPool()
	return &Dispatcher{
		workerPool: pool,
	}, err
}

// AddTask 添加任务，把任务压入任务缓冲池
func (d *Dispatcher) AddTask(data func()) {
	d.workerPool.Submit(data)
	//d.jobQueue <- data
}

// NewWorkPool
func NewWorkPool() (*ants.Pool, error) {
	workerPool, err := ants.NewPool(goroutine.DefaultAntsPoolSize,
		ants.WithOptions(ants.Options{ExpiryDuration: goroutine.ExpiryDuration,
			Nonblocking:  goroutine.Nonblocking,
			PanicHandler: nil,
		}))
	return workerPool, err
}
