package xerr

import (
	"context"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/status"
)

func Response(ctx context.Context, w http.ResponseWriter, resp interface{}, err error) {
	body := NewErrorCode(OK)
	if err != nil {
		if x, ok := err.(*Error); ok {
			body = x
		} else if x, ok := status.FromError(err); ok {
			body = NewError(uint32(x.Code()), x.Message())
		} else {
			body = NewError(UNKNOW_ERROR, err.Error())
		}
	} else {
		body.Data = resp
	}

	httpx.OkJsonCtx(ctx, w, body)
}
