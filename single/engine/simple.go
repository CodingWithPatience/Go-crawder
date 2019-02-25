package engine

import (
	"crawder/single/common"
	"crawder/single/seeds/douban"
	"log"
)

type SimpleEngine struct {}

func (e SimpleEngine) Run()  {
	var requests []common.Request
	in := make(chan common.Request)
	seeds.GenerateSeeds(in)
	go func() {
		for {
			r := <-in
			requests = append(requests, r)
		}
	}()
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		parseResult, err := common.Worker(r)
		if err != nil {
			continue
		}
		requests = append(requests, parseResult.Requests...)
		for _, r := range parseResult.Items {
			log.Printf("Got item %+v", r)
		}
	}
}

