package scheduler

import (
	"crawder/single/common"
)

type SimpleScheduler struct {
	workChan chan common.Request
}

func (s *SimpleScheduler) WorkReady(chan common.Request) {
}

func (s *SimpleScheduler) WorkerChan() chan common.Request {
    return s.workChan
}

func (s *SimpleScheduler) Run() {
    s.workChan = make(chan common.Request)
}

func (s *SimpleScheduler) Submit(r common.Request) {
    go func() {
    	s.workChan <- r
	}()
}

