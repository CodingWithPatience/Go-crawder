package common

import (
	"crawder/single/fetcher"
	"log"
)

type Item struct {
	Url string
	Id string
	Content interface{}
}

type ParseResult struct {
	Requests []Request
	Items []Item
}

type Request struct {
	Url string
	Parser Parser
}

type ParserFunc func([]byte, string) ParseResult

type Parser interface {
	Parse(contents []byte, url string) ParseResult
	Serialize() (name string, args interface{})
}

type FuncParser struct {
	parser ParserFunc
	name string
	args interface{}
}

func (f *FuncParser) Parse(contents []byte, url string) ParseResult {
    return f.parser(contents, url)
}

func (f *FuncParser) Serialize() (name string, args interface{}) {
    return f.name, nil
}

func NewFuncParser(p ParserFunc, name string) *FuncParser {
	return &FuncParser{
		parser: p,
		name: name,
	}
}

func Worker(request Request) (ParseResult, error) {
	resp, err := fetcher.Fetch(request.Url)
	log.Printf("Fetching %s\n", request.Url)
	if err != nil {
		log.Printf("Error fetching url %s! error %v\n", request.Url, err)
		return ParseResult{}, err
	}
	parseResult := request.Parser.Parse(resp, request.Url)
	return parseResult, nil
}

