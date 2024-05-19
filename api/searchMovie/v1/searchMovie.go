package v1

import "github.com/gogf/gf/v2/frame/g"

type GetSearchMovieReq struct {
	g.Meta `path:"/search" tags:"search" method:"get" summary:"You first search api"`
}
type GetSearchMovieRes struct {
	g.Meta `mime:"application/json" example:"string"`
}
