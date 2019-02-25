package model

import "strconv"

const (
	//type
	MOVIE = "movie"
	TV = "tv"

	//movie tag
	HOT = "热门"
	RESENT = "最新"
	CLASSIC = "经典"

	//page limit
	COUNT = 50
	//page start
	START = 0
)

type Condition struct {
	Type string
	Tag string
	PageLimit int
	PageStart int
	CurPage int
}

func SearchCondition(c Condition) string {
	query := string("?type=" + c.Type + "&tag=" + c.Tag +
		"&page_limit=" + strconv.Itoa(c.PageLimit) + "&page_start=" + strconv.Itoa(c.PageStart))
	return query
}

type Type struct {
	Index int
	Movie string
	TV string
}

type Tag struct {
	Index int
	Hot string
	Resent string
	Classic string
	TagCount int
}
func (t *Type) Next() string {
	if t.Index == 0 {
		t.Index++
		return t.Movie
	}
	t.Index = 0
	return t.TV
}

func (t *Tag) Next() string {
	if t.Index == 0 {
		t.Index++
		return t.Hot
	}
	if t.Index == 1 {
		t.Index++
		return t.Resent
	}
	t.Index = 0
	return t.Classic
}

func (c *Condition) NextPage()  {
	c.PageStart = c.PageLimit * (c.CurPage - 1)
	c.CurPage++
}

func (c *Condition) NextTag(tag Tag)  {
	c.Tag = tag.Next()
}

func (c *Condition) NextType(t Type)  {
	c.Type = t.Next()
}

