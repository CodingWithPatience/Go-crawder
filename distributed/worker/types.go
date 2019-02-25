package worker

import (
	"crawder/single/common"
	"crawder/single/config"
	p "crawder/single/parser/douban"
	"github.com/pkg/errors"
	"log"
)

type SerializedParser struct {
	Name string
	Args interface{}
}

type Request struct {
	Url string
	Parser SerializedParser
}

type ParseResult struct {
	Requests []Request
	Items []common.Item
}

func SerializeRequest(request common.Request) Request {
	return Request{
		Url: request.Url,
		Parser: SerializeParser(request.Parser),
	}
}

func SerializeParser(parser common.Parser) SerializedParser {
	name, args := parser.Serialize()
	return SerializedParser{
		Name: name,
		Args: args,
	}
}

func SerializeParseResult(result common.ParseResult) ParseResult {
	parseResult := ParseResult{
		Items: result.Items,
	}
	for _, r := range result.Requests {
		parseResult.Requests = append(parseResult.Requests, SerializeRequest(r))
	}
	return parseResult
}

func DeserializeRequest(request Request) (common.Request, error) {
	parser, err := DeserializeParser(request.Parser)
	if err != nil {
		return common.Request{}, err
	}
	return common.Request{
		Url: request.Url,
		Parser: parser,
	}, nil
}

func DeserializeParseResult(result ParseResult) common.ParseResult {
	parseResult := common.ParseResult{
		Items: result.Items,
	}
	for _, r := range result.Requests {
		request, err := DeserializeRequest(r)
		if err == nil {
			parseResult.Requests = append(parseResult.Requests, request)
		}else {
			log.Printf("Error deserilaizing request: %v", err)
		}
	}
	return parseResult
}

func DeserializeParser(parser SerializedParser) (common.Parser, error){
	switch parser.Name {
	case config.ParseMovieLink:
		return common.NewFuncParser(p.ParseMovieLink, config.ParseMovieLink), nil
	case config.ParseMovie:
		return common.NewFuncParser(p.ParseMovie, config.ParseMovie), nil
	default:
		return nil, errors.New("Unknown parser name")
	}
}


