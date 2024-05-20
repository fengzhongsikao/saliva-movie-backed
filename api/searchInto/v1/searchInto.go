package v1

import "github.com/gogf/gf/v2/frame/g"

type GetSearchIntoReq struct {
	g.Meta `path:"/searchInto" tags:"searchInto" method:"get" summary:"You first searchInto api"`
}
type GetSearchIntoRes struct {
	g.Meta `mime:"application/json" example:"string"`
}
