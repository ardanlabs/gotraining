package client

import (
	"github.com/pachyderm/pachyderm/src/client/pkg/grpcutil"

	"github.com/gogo/protobuf/types"
)

// Health health checks pachd, it returns an error if pachd isn't healthy.
func (c APIClient) Health() error {
	_, err := c.healthClient.Health(c.Ctx(), &types.Empty{})
	return grpcutil.ScrubGRPC(err)
}
