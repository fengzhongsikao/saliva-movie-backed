// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package searchMovie

import (
	"context"

	"saliva-movie/api/searchMovie/v1"
)

type ISearchMovieV1 interface {
	GetSearchMovie(ctx context.Context, req *v1.GetSearchMovieReq) (res *v1.GetSearchMovieRes, err error)
}
