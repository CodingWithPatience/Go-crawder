package seeds

import (
	"crawder/single/common"
	"crawder/single/model/douban"
	"crawder/single/parser/douban"
)

var types = model.Type{Index: 0, Movie: model.MOVIE, TV: model.TV}
var tags = model.Tag{Index: 0, Hot: model.HOT, Resent: model.RESENT, Classic: model.CLASSIC, TagCount: 3}
var condition = model.Condition{
	Type:      types.Next(),
	Tag:       tags.Next(),
	PageLimit: model.COUNT,
	PageStart: model.START,
	CurPage: 1,
}


const URL = "https://movie.douban.com/j/search_subjects"

func GetNextPageUrl() string {
	condition.NextPage()
	comUrl := URL + model.SearchCondition(condition)
	return comUrl
}

func GetNextTagUrl() string {
	condition.PageStart = 0
	condition.CurPage = 1
	condition.Tag = tags.Next()
	comUrl := URL + model.SearchCondition(condition)
	return comUrl
}

func GetNextTypeUrl() string {
	condition.PageStart = 0
	condition.CurPage = 1
	condition.Type = types.Next()
	comUrl := URL + model.SearchCondition(condition)
	return comUrl
}

func GenerateSeeds(requestChan chan common.Request) {
	request := common.Request{
		Url:        GetNextPageUrl(),
		Parser: common.NewFuncParser(parser.ParseMovieLink, "ParseMovieLink"),
	}
	go func() {
		for i := 0; i<tags.TagCount; i++ {
			for {
				result, err := common.Worker(request)
				if err != nil {
					continue
				}
				if len(result.Requests) == 0 {
					break
				}
				for _, r := range result.Requests {
					requestChan <- r
				}
				request.Url = GetNextPageUrl()
			}
			request.Url = GetNextTagUrl()
		}
	}()
}

