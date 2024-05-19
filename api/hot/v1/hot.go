package v1

import "github.com/gogf/gf/v2/frame/g"

type GetHotListReq struct {
	g.Meta `path:"/hot" tags:"hot" method:"get" summary:"You first hot api"`
}
type GetHotListRes struct {
	g.Meta `mime:"application/json" example:"string"`
}
