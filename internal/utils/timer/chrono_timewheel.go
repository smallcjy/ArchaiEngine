package timer

import (
	"errors"
	"fmt"
	"game/internal/utils/timer/chrono"
	"sync/atomic"
	"time"
)

var tw2 *ChronoTimeWheel

// NewTimeWheel 用来实现TimeWheel的单例模式
func NewTimeWheel(interval time.Duration, slotNums, workerNum int) (*ChronoTimeWheel, error) {
	if interval <= 0 || slotNums <= 0 || workerNum <= 0 {
		return nil, errors.New("param err")
	}
	once.Do(func() {
		tw2 = &ChronoTimeWheel{Scheduler: chrono.NewDefaultScheduler()}
	})
	return tw2, nil
}

type ChronoTimeWheel struct {
	*chrono.Scheduler
	guid atomic.Uint64
}

func (tw *ChronoTimeWheel) AddTimer(interval time.Duration, times int, job Job) uint64 {
	guid := tw.guid.Add(1)
	tw.RegisterRepeatedTask(fmt.Sprint(guid), interval, interval, times, job)
	return guid
}

func (tw *ChronoTimeWheel) RemoveTimer(key uint64) {
	tw.UnregisterTask(fmt.Sprint(key))
}

func (tw *ChronoTimeWheel) Start() {

}
