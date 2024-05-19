package searchMovie

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"

	"saliva-movie/api/searchMovie/v1"
)

func (c *ControllerV1) GetSearchMovie(ctx context.Context, req *v1.GetSearchMovieReq) (res *v1.GetSearchMovieRes, err error) {
	g.RequestFromCtx(ctx).Response.Writeln("Hello SEARCH!")
	return
}
