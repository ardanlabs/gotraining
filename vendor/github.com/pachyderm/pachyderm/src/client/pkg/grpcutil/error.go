package grpcutil

import (
	"errors"

	"google.golang.org/grpc/status"
)

// ScrubGRPC removes GRPC error code information from 'err' if it came from
// GRPC (and returns it unchanged otherwise)
func ScrubGRPC(err error) error {
	if err == nil {
		return nil
	}
	if s, ok := status.FromError(err); ok {
		return errors.New(s.Message())
	}
	return err
}
