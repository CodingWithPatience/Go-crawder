package engine

import (
	"crawder/single/common"
	"crawder/single/seeds/douban"
)

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
	ItemChan chan common.Item
	Processor Processor
}

type Processor func(request common.Request) (common.ParseResult, error)

type Scheduler interface {
	NotifyReady
	Submit(common.Request)
	WorkerChan() chan common.Request
	Run()
}

type NotifyReady interface {
	WorkReady(chan common.Request)
}

func (e ConcurrentEngine) Run()  {
	out := make(chan common.ParseResult)
	e.Scheduler.Run()
	for i := 0; i<e.WorkerCount; i++ {
		e.createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	in := make(chan common.Request)
	seeds.GenerateSeeds(in)
	go func() {
		for {
			r := <- in
			e.Scheduler.Submit(r)
		}
	}()

	for {
		result := <- out
		for _, item := range result.Items {
			go func() {e.ItemChan <- item}()
		}
		for _, r := range result.Requests {
			e.Scheduler.Submit(r)
		}
	}
}

func (e *ConcurrentEngine) createWorker(in chan common.Request, result chan common.ParseResult, s Scheduler) {
	go func() {
		for {
			s.WorkReady(in)
			r := <- in
			parseResult, err := e.Processor(r)
			if err != nil {
				continue
			}
			result <- parseResult
		}
	}()
}
