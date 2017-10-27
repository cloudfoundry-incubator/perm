package rpc

import (
	"code.cloudfoundry.org/perm/errdefs"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func togRPCError(err error) error {
	switch err.(type) {
	case nil:
		return nil
	case errdefs.ErrNotFound:
		return status.Errorf(codes.NotFound, err.Error())
	case errdefs.ErrAlreadyExists:
		return status.Errorf(codes.AlreadyExists, err.Error())
	default:
		return status.Errorf(codes.Unknown, err.Error())
	}
}
