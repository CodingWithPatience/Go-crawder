package parser

import (
	"crawder/single/common"
	"regexp"
)

const linkRe  = `https://movie.douban.com/subject/[0-9]+`
//const titleRe_  = `"title":"([^"]+)"`

func ParseMovieLink(contents []byte, _ string) common.ParseResult {
	compile := regexp.MustCompile(linkRe)
	urls := compile.FindAllSubmatch(contents, -1)

	result := common.ParseResult{}
	for _, r := range urls {
		for _, e := range r {
			result.Requests = append(result.Requests,
				common.Request{
					Url: string(e),
					Parser: common.NewFuncParser(ParseMovie, "ParseMovie"),
				})
		}
	}
	return result
}
