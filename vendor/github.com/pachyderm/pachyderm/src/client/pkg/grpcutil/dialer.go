package grpcutil

import (
	"sync"

	"google.golang.org/grpc"
)

// Dialer defines a grpc.ClientConn connection dialer.
type Dialer interface {
	Dial(address string) (*grpc.ClientConn, error)
	CloseConns() error
}

// NewDialer creates a Dialer.
func NewDialer(opts ...grpc.DialOption) Dialer {
	return newDialer(opts...)
}

type dialer struct {
	opts []grpc.DialOption
	// A map from addresses to connections
	connMap map[string]*grpc.ClientConn
	lock    sync.Mutex
}

func newDialer(opts ...grpc.DialOption) *dialer {
	return &dialer{
		opts:    opts,
		connMap: make(map[string]*grpc.ClientConn),
	}
}

func (d *dialer) Dial(addr string) (*grpc.ClientConn, error) {
	d.lock.Lock()
	defer d.lock.Unlock()
	if conn, ok := d.connMap[addr]; ok {
		return conn, nil
	}
	conn, err := grpc.Dial(addr, d.opts...)
	if err != nil {
		return nil, err
	}
	d.connMap[addr] = conn
	return conn, nil
}

func (d *dialer) CloseConns() error {
	d.lock.Lock()
	defer d.lock.Unlock()
	for addr, conn := range d.connMap {
		if err := conn.Close(); err != nil {
			return err
		}
		delete(d.connMap, addr)
	}
	return nil
}
