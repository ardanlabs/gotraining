package grpcutil

import (
	"errors"
	"fmt"
	"math"
	"net"
	"time"

	"github.com/pachyderm/pachyderm/src/client/version"
	"github.com/pachyderm/pachyderm/src/client/version/versionpb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

var (
	// ErrMustSpecifyRegisterFunc is used when a register func is nil.
	ErrMustSpecifyRegisterFunc = errors.New("must specify registerFunc")
)

// ServeOptions represent optional fields for serving.
type ServeOptions struct {
	Version    *versionpb.Version
	MaxMsgSize int
	Cancel     chan struct{}
}

// ServeEnv are environment variables for serving.
type ServeEnv struct {
	// Default is 7070.
	GRPCPort uint16 `env:"GRPC_PORT,default=7070"`
}

// Serve serves stuff.
func Serve(
	registerFunc func(*grpc.Server),
	options ServeOptions,
	serveEnv ServeEnv,
) (retErr error) {
	if registerFunc == nil {
		return ErrMustSpecifyRegisterFunc
	}
	if serveEnv.GRPCPort == 0 {
		serveEnv.GRPCPort = 7070
	}
	grpcServer := grpc.NewServer(
		grpc.MaxConcurrentStreams(math.MaxUint32),
		grpc.MaxRecvMsgSize(options.MaxMsgSize),
		grpc.MaxSendMsgSize(options.MaxMsgSize),
		grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
			MinTime:             5 * time.Second,
			PermitWithoutStream: true,
		}),
	)
	registerFunc(grpcServer)
	if options.Version != nil {
		versionpb.RegisterAPIServer(grpcServer, version.NewAPIServer(options.Version, version.APIServerOptions{}))
	}
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", serveEnv.GRPCPort))
	if err != nil {
		return err
	}
	if options.Cancel != nil {
		go func() {
			<-options.Cancel
			if err := listener.Close(); err != nil {
				fmt.Printf("listener.Close(): %v\n", err)
			}
		}()
	}
	return grpcServer.Serve(listener)
}
