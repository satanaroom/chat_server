package interceptor

import (
	"context"

	"github.com/fatih/color"
	"github.com/pkg/errors"
	"github.com/satanaroom/auth/pkg/logger"
	"github.com/satanaroom/chat_server/internal/sys"
	"github.com/satanaroom/chat_server/internal/sys/codes"
	"github.com/satanaroom/chat_server/internal/sys/validate"
	"google.golang.org/grpc"
	grpcCodes "google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GRPCStatus interface {
	GRPCStatus() *status.Status
}

func ErrorCodesInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	resp, err := handler(ctx, req)
	if nil == err {
		return resp, nil
	}

	logger.Errorf(color.RedString("error: %s", err.Error()))

	switch {
	case sys.IsCommonError(err):
		commonErr := sys.GetCommonError(err)
		code := toGRPCCode(commonErr.Code())
		err = status.Error(code, commonErr.Error())
	case validate.IsValidationError(err):
		err = status.Error(grpcCodes.InvalidArgument, err.Error())
	default:
		var se GRPCStatus
		if errors.As(err, &se) {
			return nil, se.GRPCStatus().Err()
		} else {
			if errors.Is(err, context.Canceled) {
				err = status.Error(grpcCodes.Canceled, err.Error())
			} else if errors.Is(err, context.DeadlineExceeded) {
				err = status.Error(grpcCodes.DeadlineExceeded, err.Error())
			} else {
				err = status.Error(grpcCodes.Internal, err.Error())
			}
		}
	}

	return nil, err
}

func toGRPCCode(code codes.Code) grpcCodes.Code {
	var res grpcCodes.Code
	switch code {
	case codes.OK:
		res = grpcCodes.OK
	case codes.Canceled:
		res = grpcCodes.Canceled
	case codes.InvalidArgument:
		res = grpcCodes.InvalidArgument
	case codes.DeadlineExceeded:
		res = grpcCodes.DeadlineExceeded
	case codes.NotFound:
		res = grpcCodes.NotFound
	case codes.AlreadyExists:
		res = grpcCodes.AlreadyExists
	case codes.PermissionDenied:
		res = grpcCodes.PermissionDenied
	case codes.ResourceExhausted:
		res = grpcCodes.ResourceExhausted
	case codes.FailedPrecondition:
		res = grpcCodes.FailedPrecondition
	case codes.Aborted:
		res = grpcCodes.Aborted
	case codes.OutOfRange:
		res = grpcCodes.OutOfRange
	case codes.Unimplemented:
		res = grpcCodes.Unimplemented
	case codes.Internal:
		res = grpcCodes.Internal
	case codes.Unavailable:
		res = grpcCodes.Unavailable
	case codes.DataLoss:
		res = grpcCodes.DataLoss
	case codes.Unauthenticated:
		res = grpcCodes.Unauthenticated
	default:
		res = grpcCodes.Unknown
	}

	return res
}
