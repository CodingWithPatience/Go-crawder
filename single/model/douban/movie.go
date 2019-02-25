package model

type Movie struct {
	Title string
	Rating float64
	RatingNumber int
	Director string
	Type []string
	Duration string
	Casting Cast
	ReleaseDate []string
}

type Cast struct {
	Actor []string
}

func (cast *Cast) AddActor(name string) {
	cast.Actor = append(cast.Actor, name)
}