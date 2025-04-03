package xerr

import (
	"context"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func RpcErrorInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	resp, err = handler(ctx, req)
	if err != nil {
		causeErr := errors.Cause(err)       // err类型
		if e, ok := causeErr.(*Error); ok { //自定义错误类型
			//转成grpc err
			err = status.Error(codes.Code(e.GetCode()), e.GetMsg())
		}
		logx.WithContext(ctx).Errorf("【RPC-SRV-ERR】 %+v", err)
	}

	return resp, err
}
