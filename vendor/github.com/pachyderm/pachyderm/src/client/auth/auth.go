package auth

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"google.golang.org/grpc/metadata"
)

const (
	// ContextTokenKey is the key of the auth token in an
	// authenticated context
	ContextTokenKey = "authn-token"

	// The following constants are Subject prefixes. These are prepended to
	// Subjects in the 'tokens' collection, and Principals in 'admins' and on ACLs
	// to indicate what type of Subject or Principal they are (every Pachyderm
	// Subject has a logical Principal with the same name).

	// GitHubPrefix indicates that this Subject is a GitHub user (because users
	// can authenticate via GitHub, and Pachyderm doesn't have a users table,
	// every GitHub user is also a logical Pachyderm user (but most won't be on
	// any ACLs)
	GitHubPrefix = "github:"

	// RobotPrefix indicates that this Subject is a Pachyderm robot user. Any
	// string (with this prefix) is a logical Pachyderm robot user.
	RobotPrefix = "robot:"

	// PipelinePrefix indicates that this Subject is a PPS pipeline. Any string
	// (with this prefix) is a logical PPS pipeline (even though the pipeline may
	// not exist).
	PipelinePrefix = "pipeline:"
)

// ParseScope parses the string 's' to a scope (for example, parsing a command-
// line argument.
func ParseScope(s string) (Scope, error) {
	for name, value := range Scope_value {
		if strings.EqualFold(s, name) {
			return Scope(value), nil
		}
	}
	return Scope_NONE, fmt.Errorf("unrecognized scope: %s", s)
}

// In2Out converts an incoming context containing auth information into an
// outgoing context containing auth information, stripping other keys (e.g.
// for metrics) in the process. If the incoming context doesn't have any auth
// information, then the returned context won't either.
func In2Out(ctx context.Context) context.Context {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ctx
	}
	mdOut := make(metadata.MD)
	if value, ok := md[ContextTokenKey]; ok {
		mdOut[ContextTokenKey] = value
	}
	return metadata.NewOutgoingContext(ctx, mdOut)
}

var (
	// ErrNotActivated is returned by an Auth API if the Auth service
	// has not been activated.
	//
	// Note: This error message string is matched in the UI. If edited,
	// it also needs to be updated in the UI code
	ErrNotActivated = errors.New("the auth service is not activated")

	// ErrNotSignedIn indicates that the caller isn't signed in
	//
	// Note: This error message string is matched in the UI. If edited,
	// it also needs to be updated in the UI code
	ErrNotSignedIn = errors.New("auth token not found in context (user may not be signed in)")

	// ErrNoToken is returned by the Auth API if the caller sent a request
	// containing no auth token.
	ErrNoToken = errors.New("no authentication metadata found in context")

	// ErrBadToken is returned by the Auth API if the caller's token is corruped
	// or has expired.
	ErrBadToken = errors.New("provided auth token is corrupted or has expired (try logging in again)")
)

// IsErrNotActivated checks if an error is a ErrNotActivated
func IsErrNotActivated(err error) bool {
	if err == nil {
		return false
	}
	// TODO(msteffen) This is unstructured because we have no way to propagate
	// structured errors across GRPC boundaries. Fix
	return strings.Contains(err.Error(), ErrNotActivated.Error())
}

// ErrNotAuthorized is returned if the user is not authorized to perform
// a certain operation. Either
// 1) the operation is a user operation, in which case 'Repo' and/or 'Required'
// 		should be set (indicating that the user needs 'Required'-level access to
// 		'Repo').
// 2) the operation is an admin-only operation (e.g. DeleteAll), in which case
//    AdminOp should be set
type ErrNotAuthorized struct {
	Subject string // subject trying to perform blocked operation -- always set

	Repo     string // Repo that the user is attempting to access
	Required Scope  // Caller needs 'Required'-level access to 'Repo'

	// Group 2:
	// AdminOp indicates an operation that the caller couldn't perform because
	// they're not an admin
	AdminOp string
}

// This error message string is matched in the UI. If edited,
// it also needs to be updated in the UI code
const errNotAuthorizedMsg = "not authorized to perform this operation"

func (e *ErrNotAuthorized) Error() string {
	var msg string
	if e.Subject != "" {
		msg += e.Subject + " is "
	}
	msg += errNotAuthorizedMsg
	if e.Repo != "" {
		msg += " on the repo " + e.Repo
	}
	if e.Required != Scope_NONE {
		msg += ", must have at least " + e.Required.String() + " access"
	}
	if e.AdminOp != "" {
		msg += "; must be an admin to call " + e.AdminOp
	}
	return msg
}

// IsErrNotAuthorized checks if an error is a ErrNotAuthorized
func IsErrNotAuthorized(err error) bool {
	if err == nil {
		return false
	}
	// TODO(msteffen) This is unstructured because we have no way to propagate
	// structured errors across GRPC boundaries. Fix
	return strings.Contains(err.Error(), errNotAuthorizedMsg)
}

// IsErrNotSignedIn returns true if 'err' is a ErrNotSignedIn
func IsErrNotSignedIn(err error) bool {
	// TODO(msteffen) This is unstructured because we have no way to propagate
	// structured errors across GRPC boundaries. Fix
	return strings.Contains(err.Error(), ErrNotSignedIn.Error())
}

// ErrInvalidPrincipal indicates that a an argument to e.g. GetScope,
// SetScope, or SetACL is invalid
type ErrInvalidPrincipal struct {
	Principal string
}

func (e *ErrInvalidPrincipal) Error() string {
	return fmt.Sprintf("invalid principal \"%s\"; must start with one of \"pipeline:\", \"github:\", or \"robot:\", or have no \":\"", e.Principal)
}

// IsErrInvalidPrincipal returns true if 'err' is an ErrInvalidPrincipal
func IsErrInvalidPrincipal(err error) bool {
	if err == nil {
		return false
	}
	return strings.Contains(err.Error(), "invalid principal \"") &&
		strings.Contains(err.Error(), "\"; must start with one of \"pipeline:\", \"github:\", or \"robot:\", or have no \":\"")
}

// IsErrNoToken returns true if 'err' is a ErrNoToken (uses string
// comparison to work across RPC boundaries)
func IsErrNoToken(err error) bool {
	if err == nil {
		return false
	}
	return strings.Contains(err.Error(), ErrNoToken.Error())
}

// IsErrBadToken returns true if 'err' is a ErrBadToken
func IsErrBadToken(err error) bool {
	if err == nil {
		return false
	}
	return strings.Contains(err.Error(), ErrBadToken.Error())
}

// ErrTooShortTTL is returned by the ExtendAuthToken if request.Token already
// has a TTL longer than request.TTL.
type ErrTooShortTTL struct {
	RequestTTL, ExistingTTL int64
}

const errTooShortTTLMsg = "provided TTL (%d) is shorter than token's existing TTL (%d)"

func (e ErrTooShortTTL) Error() string {
	return fmt.Sprintf(errTooShortTTLMsg, e.RequestTTL, e.ExistingTTL)
}

// IsErrTooShortTTL returns true if 'err' is a ErrTooShortTTL
func IsErrTooShortTTL(err error) bool {
	if err == nil {
		return false
	}
	errMsg := err.Error()
	return strings.Contains(errMsg, "provided TTL (") &&
		strings.Contains(errMsg, ") is shorter than token's existing TTL (") &&
		strings.Contains(errMsg, ")")
}
