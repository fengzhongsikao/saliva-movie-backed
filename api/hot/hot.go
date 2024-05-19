// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package hot

import (
	"context"

	"saliva-movie/api/hot/v1"
)

type IHotV1 interface {
	GetHotList(ctx context.Context, req *v1.GetHotListReq) (res *v1.GetHotListRes, err error)
}
