// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package searchInto

import (
	"context"

	"saliva-movie/api/searchInto/v1"
)

type ISearchIntoV1 interface {
	GetSearchInto(ctx context.Context, req *v1.GetSearchIntoReq) (res *v1.GetSearchIntoRes, err error)
}
