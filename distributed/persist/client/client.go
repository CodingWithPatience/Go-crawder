package main

import (
	"crawder/distributed/rpcsupport"
	"crawder/single/common"
	"crawder/single/config"
	"crawder/single/model/douban"
	"fmt"
	"log"
)

func main() {
	client, err := rpcsupport.NewClient(config.SaverPort)
	if err != nil {
		log.Printf("Error connecting rpc server %v", err)
	}
	var result string
	item := common.Item {
		Url:     "http://xxxx",
		Id:      "123",
		Content: model.Movie{
			Title: "test",
		},
	}
	err = client.Call("MovieSaverService.MovieSave", item, &result)
	if err != nil {
		log.Printf("Error accur during rpc call %v", err)
	}
	fmt.Println(result)
}
