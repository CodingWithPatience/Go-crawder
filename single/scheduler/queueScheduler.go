package scheduler

import (
	"crawder/single/common"
)

type QueueScheduler struct {
	requestChan chan common.Request
	workerChan chan chan common.Request
}

func (s *QueueScheduler) WorkerChan() chan common.Request {
    return make(chan common.Request)
}

func (s *QueueScheduler) Submit(r common.Request) {
	s.requestChan <- r
}

func (s *QueueScheduler) WorkReady(w chan common.Request)  {
	s.workerChan <- w
}

func (s *QueueScheduler) Run()  {
	s.requestChan = make(chan common.Request)
	s.workerChan = make(chan chan common.Request)
	go func() {
		var requestQ []common.Request
		var workerQ []chan common.Request
		for {
			var activeRequest common.Request
			var activeWorker chan common.Request
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeRequest = requestQ[0]
				activeWorker = workerQ[0]
			}
			select {
			case r := <- s.requestChan:
				requestQ = append(requestQ, r)
			case w := <- s.workerChan:
				workerQ = append(workerQ, w)
			case activeWorker <- activeRequest:
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]
			}
		}
	}()
}

