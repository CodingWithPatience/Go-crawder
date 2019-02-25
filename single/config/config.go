package config

const (
	Host = "http://192.168.254.128:9200"  //elastic host
	Index = "movie"       //elastic index
	Type = "douban"       //elastic type

	//rpc server host
	SaverPort = 9000
	WorkerPort = 10000

	//rpc service name
	SaverRpcService = "MovieSaverService.MovieSave"
	WorkerRpcService = "CrawlService.Process"

	//parse function name
	ParseMovieLink = "ParseMovieLink"
	ParseMovie = "ParseMovie"
)
