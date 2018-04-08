package client

import (
	"fmt"
	"os"
	"time"

	"golang.org/x/net/context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/metadata"

	types "github.com/gogo/protobuf/types"
	log "github.com/sirupsen/logrus"

	"github.com/pachyderm/pachyderm/src/client/admin"
	"github.com/pachyderm/pachyderm/src/client/auth"
	"github.com/pachyderm/pachyderm/src/client/deploy"
	"github.com/pachyderm/pachyderm/src/client/enterprise"
	"github.com/pachyderm/pachyderm/src/client/health"
	"github.com/pachyderm/pachyderm/src/client/pfs"
	"github.com/pachyderm/pachyderm/src/client/pkg/config"
	"github.com/pachyderm/pachyderm/src/client/pkg/grpcutil"
	"github.com/pachyderm/pachyderm/src/client/pps"
	"github.com/pachyderm/pachyderm/src/client/version/versionpb"
)

const (
	// MaxListItemsLog specifies the maximum number of items we log in response to a List* API
	MaxListItemsLog = 10
	// StorageSecretName is the name of the Kubernetes secret in which
	// storage credentials are stored.
	StorageSecretName = "pachyderm-storage-secret"
)

// PfsAPIClient is an alias for pfs.APIClient.
type PfsAPIClient pfs.APIClient

// PpsAPIClient is an alias for pps.APIClient.
type PpsAPIClient pps.APIClient

// ObjectAPIClient is an alias for pfs.ObjectAPIClient
type ObjectAPIClient pfs.ObjectAPIClient

// AuthAPIClient is an alias of auth.APIClient
type AuthAPIClient auth.APIClient

// DeployAPIClient is an alias of auth.APIClient
type DeployAPIClient deploy.APIClient

// VersionAPIClient is an alias of versionpb.APIClient
type VersionAPIClient versionpb.APIClient

// AdminAPIClient is an alias of admin.APIClient
type AdminAPIClient admin.APIClient

// An APIClient is a wrapper around pfs, pps and block APIClients.
type APIClient struct {
	PfsAPIClient
	PpsAPIClient
	ObjectAPIClient
	AuthAPIClient
	DeployAPIClient
	VersionAPIClient
	AdminAPIClient
	Enterprise enterprise.APIClient // not embedded--method name conflicts with AuthAPIClient

	// addr is a "host:port" string pointing at a pachd endpoint
	addr string

	// clientConn is a cached grpc connection to 'addr'
	clientConn *grpc.ClientConn

	// healthClient is a cached healthcheck client connected to 'addr'
	healthClient health.HealthClient

	// streamSemaphore limits the number of concurrent message streams between
	// this client and pachd
	streamSemaphore chan struct{}

	// metricsUserID is an identifier that is included in usage metrics sent to
	// Pachyderm Inc. and is used to count the number of unique Pachyderm users.
	// If unset, no usage metrics are sent back to Pachyderm Inc.
	metricsUserID string

	// metricsPrefix is used to send information from this client to Pachyderm Inc
	// for usage metrics
	metricsPrefix string

	// authenticationToken is an identifier that authenticates the caller in case
	// they want to access privileged data
	authenticationToken string

	// The context used in requests, can be set with WithCtx
	ctx context.Context
}

// GetAddress returns the pachd host:port with which 'c' is communicating. If
// 'c' was created using NewInCluster or NewOnUserMachine then this is how the
// address may be retrieved from the environment.
func (c *APIClient) GetAddress() string {
	return c.addr
}

// DefaultMaxConcurrentStreams defines the max number of Putfiles or Getfiles happening simultaneously
const DefaultMaxConcurrentStreams uint = 100

// NewFromAddressWithConcurrency constructs a new APIClient and sets the max
// concurrency of streaming requests (GetFile / PutFile)
func NewFromAddressWithConcurrency(addr string, maxConcurrentStreams uint) (*APIClient, error) {
	c := &APIClient{
		addr:            addr,
		streamSemaphore: make(chan struct{}, maxConcurrentStreams),
	}
	if err := c.connect(); err != nil {
		return nil, err
	}
	return c, nil
}

// NewFromAddress constructs a new APIClient for the server at addr.
func NewFromAddress(addr string) (*APIClient, error) {
	return NewFromAddressWithConcurrency(addr, DefaultMaxConcurrentStreams)
}

// GetAddressFromUserMachine interprets the Pachyderm config in 'cfg' in the
// context of local environment variables and returns a "host:port" string
// pointing at a Pachd target.
func GetAddressFromUserMachine(cfg *config.Config) string {
	address := "0.0.0.0:30650"
	if cfg != nil && cfg.V1 != nil && cfg.V1.PachdAddress != "" {
		address = cfg.V1.PachdAddress
	}
	// ADDRESS environment variable (shell-local) overrides global config
	if envAddr := os.Getenv("ADDRESS"); envAddr != "" {
		address = envAddr
	}
	return address
}

// NewOnUserMachine constructs a new APIClient using env vars that may be set
// on a user's machine (i.e. ADDRESS), as well as $HOME/.pachyderm/config if it
// exists. This is primarily intended to be used with the pachctl binary, but
// may also be useful in tests.
//
// TODO(msteffen) this logic is fairly linux/unix specific, and makes the
// pachyderm client library incompatible with Windows. We may want to move this
// (and similar) logic into src/server and have it call a NewFromOptions()
// constructor.
func NewOnUserMachine(reportMetrics bool, prefix string) (*APIClient, error) {
	return NewOnUserMachineWithConcurrency(reportMetrics, prefix, DefaultMaxConcurrentStreams)
}

// NewOnUserMachineWithConcurrency is identical to NewOnUserMachine, but
// explicitly sets a limit on the number of RPC streams that may be open
// simultaneously
func NewOnUserMachineWithConcurrency(reportMetrics bool, prefix string, maxConcurrentStreams uint) (*APIClient, error) {
	cfg, err := config.Read()
	if err != nil {
		// metrics errors are non fatal
		log.Warningf("error loading user config from ~/.pachderm/config: %v", err)
	}

	// create new pachctl client
	client, err := NewFromAddress(GetAddressFromUserMachine(cfg))
	if err != nil {
		return nil, err
	}

	// Add metrics info & authentication token
	client.metricsPrefix = prefix
	if cfg.UserID != "" && reportMetrics {
		client.metricsUserID = cfg.UserID
	}
	if cfg.V1 != nil && cfg.V1.SessionToken != "" {
		client.authenticationToken = cfg.V1.SessionToken
	}
	return client, nil
}

// NewInCluster constructs a new APIClient using env vars that Kubernetes creates.
// This should be used to access Pachyderm from within a Kubernetes cluster
// with Pachyderm running on it.
func NewInCluster() (*APIClient, error) {
	if addr := os.Getenv("PACHD_PORT_650_TCP_ADDR"); addr != "" {
		return NewFromAddress(fmt.Sprintf("%v:650", addr))
	}
	return nil, fmt.Errorf("PACHD_PORT_650_TCP_ADDR not set")
}

// Close the connection to gRPC
func (c *APIClient) Close() error {
	return c.clientConn.Close()
}

// DeleteAll deletes everything in the cluster.
// Use with caution, there is no undo.
func (c APIClient) DeleteAll() error {
	if _, err := c.AuthAPIClient.Deactivate(
		c.Ctx(),
		&auth.DeactivateRequest{},
	); err != nil && !auth.IsErrNotActivated(err) {
		return grpcutil.ScrubGRPC(err)
	}
	if _, err := c.PpsAPIClient.DeleteAll(
		c.Ctx(),
		&types.Empty{},
	); err != nil {
		return grpcutil.ScrubGRPC(err)
	}
	if _, err := c.PfsAPIClient.DeleteAll(
		c.Ctx(),
		&types.Empty{},
	); err != nil {
		return grpcutil.ScrubGRPC(err)
	}
	return nil
}

// SetMaxConcurrentStreams Sets the maximum number of concurrent streams the
// client can have. It is not safe to call this operations while operations are
// outstanding.
func (c APIClient) SetMaxConcurrentStreams(n int) {
	c.streamSemaphore = make(chan struct{}, n)
}

// EtcdDialOptions is a helper returning a slice of grpc.Dial options
// such that grpc.Dial() is synchronous: the call doesn't return until
// the connection has been established and it's safe to send RPCs
func EtcdDialOptions() []grpc.DialOption {
	return []grpc.DialOption{
		// Don't return from Dial() until the connection has been established
		grpc.WithBlock(),

		// If no connection is established in 30s, fail the call
		grpc.WithTimeout(30 * time.Second),

		grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(grpcutil.MaxMsgSize),
			grpc.MaxCallSendMsgSize(grpcutil.MaxMsgSize),
		),
	}
}

// PachDialOptions is a helper returning a slice of grpc.Dial options
// such that
// - TLS is disabled
// - Dial is synchronous: the call doesn't return until the connection has been
//                        established and it's safe to send RPCs
//
// This is primarily useful for Pachd and Worker clients
func PachDialOptions() []grpc.DialOption {
	return append(EtcdDialOptions(), grpc.WithInsecure())
}

func (c *APIClient) connect() error {
	keepaliveOpt := grpc.WithKeepaliveParams(keepalive.ClientParameters{
		Time:                20 * time.Second, // if 20s since last msg (any kind), ping
		Timeout:             20 * time.Second, // if no response to ping for 20s, reset
		PermitWithoutStream: true,             // send ping even if no active RPCs
	})
	dialOptions := append(PachDialOptions(), keepaliveOpt)
	clientConn, err := grpc.Dial(c.addr, dialOptions...)
	if err != nil {
		return grpcutil.ScrubGRPC(err)
	}
	c.PfsAPIClient = pfs.NewAPIClient(clientConn)
	c.PpsAPIClient = pps.NewAPIClient(clientConn)
	c.ObjectAPIClient = pfs.NewObjectAPIClient(clientConn)
	c.AuthAPIClient = auth.NewAPIClient(clientConn)
	c.Enterprise = enterprise.NewAPIClient(clientConn)
	c.DeployAPIClient = deploy.NewAPIClient(clientConn)
	c.VersionAPIClient = versionpb.NewAPIClient(clientConn)
	c.AdminAPIClient = admin.NewAPIClient(clientConn)
	c.clientConn = clientConn
	c.healthClient = health.NewHealthClient(clientConn)
	return nil
}

// AddMetadata adds necessary metadata (including authentication credentials)
// to the context 'ctx', preserving any metadata that is present in either the
// incoming or outgoing metadata of 'ctx'.
func (c *APIClient) AddMetadata(ctx context.Context) context.Context {
	// TODO(msteffen): There are several places in this client where it's possible
	// to set per-request metadata (specifically auth tokens): client.WithCtx(),
	// client.SetAuthToken(), etc. These should be consolidated, as this API
	// doesn't make it obvious how these settings are resolved when they conflict.
	clientData := make(map[string]string)
	if c.authenticationToken != "" {
		clientData[auth.ContextTokenKey] = c.authenticationToken
	}
	// metadata API downcases all the key names
	if c.metricsUserID != "" {
		clientData["userid"] = c.metricsUserID
		clientData["prefix"] = c.metricsPrefix
	}

	// Rescue any metadata pairs already in 'ctx' (otherwise
	// metadata.NewOutgoingContext() would drop them). Note that this is similar
	// to metadata.Join(), but distinct because it discards conflicting k/v pairs
	// instead of merging them)
	incomingMD, _ := metadata.FromIncomingContext(ctx)
	outgoingMD, _ := metadata.FromOutgoingContext(ctx)
	clientMD := metadata.New(clientData)
	finalMD := make(metadata.MD) // Collect k/v pairs
	for _, md := range []metadata.MD{incomingMD, outgoingMD, clientMD} {
		for k, v := range md {
			finalMD[k] = v
		}
	}
	return metadata.NewOutgoingContext(ctx, finalMD)
}

// Ctx is a convenience function that returns adds Pachyderm authn metadata
// to context.Background().
func (c *APIClient) Ctx() context.Context {
	if c.ctx == nil {
		return c.AddMetadata(context.Background())
	}
	return c.AddMetadata(c.ctx)
}

// WithCtx returns a new APIClient that uses ctx for requests it sends. Note
// that the new APIClient will still use the authentication token and metrics
// metadata of this client, so this is only useful for propagating other
// context-associated metadata.
func (c *APIClient) WithCtx(ctx context.Context) *APIClient {
	result := *c // copy c
	result.ctx = ctx
	return &result
}

// SetAuthToken sets the authentication token that will be used for all
// API calls for this client.
func (c *APIClient) SetAuthToken(token string) {
	c.authenticationToken = token
}
