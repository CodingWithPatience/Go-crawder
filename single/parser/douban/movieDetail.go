package parser

import (
	"crawder/single/common"
	"crawder/single/model/douban"
	"regexp"
	"strconv"
)
//store the movie url that parsed, in order not to parse repeatedly
var urlMap = make(map[string]string)

var titleRe = regexp.MustCompile(`<span property="v:itemreviewed">([^<]+)</span>`)
var ratingRe = regexp.MustCompile(`<strong class="ll rating_num" property="v:average">([^<]+)</strong>`)
var ratingNumRe = regexp.MustCompile(`<span property="v:votes">([^<]+)</span>`)
var directorRe = regexp.MustCompile(`<a href="/celebrity/[0-9]+/" rel="v:directedBy">([^<]+)</a>`)
var typeRe = regexp.MustCompile(`<span property="v:genre">([^<]+)</span>`)
var durationRe = regexp.MustCompile(`<span property="v:runtime" content=[^>]+>([^<]+)</span>`)
var castRe = regexp.MustCompile(`<a href="/celebrity/[0-9]+/" rel="v:starring">([^<]+)</a>`)
var releaseDateRe = regexp.MustCompile(`<span property="v:initialReleaseDate" content=[^>]+>([^<]+)</span>`)

var movieIdRe = regexp.MustCompile(`https://movie.douban.com/subject/([0-9]+)`)

var similarMovieRe = regexp.MustCompile(`<dd> <a href="(https://movie.douban.com/subject/[0-9]+)/[^>]+>[^<]+</a> </dd>`)

func ParseMovie(contents []byte, url string) common.ParseResult {
	result := common.ParseResult{}

	//similar movies url
	otherMovies := similarMovieRe.FindAllSubmatch(contents, -1)
	for _, r := range otherMovies {
		url := getInfo(r, 1)
		result.Requests = append(result.Requests, common.Request{
			Url: url,
			Parser: common.NewFuncParser(ParseMovie, "ParseMovie"),
		})
	}
	movieDetail(contents, url, &result)
	return result
}

func getInfo(content [][]byte, i int) string {
	var result = ""
	if content != nil{
		result = string(content[i])
	}
	return result
}

func movieDetail(contents []byte, url string, result *common.ParseResult)  {
	value := urlMap[url]
	if value == "" {
		//add url to urlMap
		urlMap[url] = url

		//get movie id from url
		idMatch := movieIdRe.FindStringSubmatch(url)
		//movie title
		title := titleRe.FindSubmatch(contents)
		//movie rating
		rating := ratingRe.FindSubmatch(contents)
		r := getInfo(rating, 1)
		ratingVal, _ := strconv.ParseFloat(r, 64)
		//number of voting people
		ratingNum := ratingNumRe.FindSubmatch(contents)
		n := getInfo(ratingNum, 1)
		numVal, _ := strconv.Atoi(n)
		//movie director
		director := directorRe.FindSubmatch(contents)
		//movie types
		mType := typeRe.FindAllSubmatch(contents, -1)
		types := make([]string, len(mType))
		for i, r := range mType {
			types[i] = getInfo(r, 1)
		}
		//movie duration
		duration := durationRe.FindSubmatch(contents)
		//major star of the movie
		cast := model.Cast{}
		casts := castRe.FindAllSubmatch(contents, -1)
		for _, r := range casts {
			cast.AddActor(getInfo(r, 1))
		}
		//movie release date
		dates := releaseDateRe.FindAllSubmatch(contents, -1)
		releaseDate := make([]string, len(dates))
		for i, r := range dates {
			releaseDate[i] = getInfo(r, 1)
		}

		result.Items = append(result.Items, common.Item{
			Url: url,
			Id:  idMatch[1],
			Content: model.Movie{
				Title:        getInfo(title, 1),
				Rating:       ratingVal,
				RatingNumber: numVal,
				Director:     getInfo(director, 1),
				Type:         types,
				Duration:     getInfo(duration, 1),
				Casting:      cast,
				ReleaseDate:  releaseDate,
			},
		})
	}
}
