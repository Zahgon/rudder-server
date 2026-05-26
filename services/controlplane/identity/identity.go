package identity

import "github.com/rudderlabs/rudder-server/utils/types/deployment"

// Authorizer abstracts how data-plane can be authorized by the control-plane.
type Authorizer interface {
	BasicAuth() (string, string)
}

// Identifier abstracts how data-plane can be identified to the control-plane.
//
//	Including both a unique identifier and the authentication method.
type Identifier interface {
	Authorizer
	ID() string
	Type() deployment.Type
}

var (
	_ Identifier = (*Namespace)(nil)
	_ Identifier = (*Workspace)(nil)
	_ Identifier = (*NOOP)(nil)
)

// Workspace identifier represents a single customer's workspace.
// Less flexible than a namespace, it does not allow for multitenant.
type Workspace struct {
	WorkspaceID    string
	WorkspaceToken string
}

func (w *Workspace) ID() string { _ = "STUB: not implemented"; return "" }

func (w *Workspace) BasicAuth() (string, string) { _ = "STUB: not implemented"; return "", "" }

func (*Workspace) Type() deployment.Type { _ = "STUB: not implemented"; return *new(deployment.Type) }

// Namespace identifier represents a group of workspaces that share a common resource.
//
//	Namespace is used but is not limited to implemented multi-tenancy.
//	It also allows for more complex entity relations.
type Namespace struct {
	Namespace    string
	HostedSecret string
}

func (n *Namespace) ID() string { _ = "STUB: not implemented"; return "" }

func (n *Namespace) BasicAuth() (string, string) { _ = "STUB: not implemented"; return "", "" }

func (*Namespace) Type() deployment.Type { _ = "STUB: not implemented"; return *new(deployment.Type) }

// NOOP is a no-op implementation of the Identifier interface.
// Used only for testing purposes.
type NOOP struct{}

func (*NOOP) ID() string { _ = "STUB: not implemented"; return "" }

func (*NOOP) BasicAuth() (string, string) { _ = "STUB: not implemented"; return "", "" }

func (*NOOP) Type() deployment.Type {
	_ = "STUB: not implemented"

	// Admin is an implementation of the Authorizer interface for data-plane admin endpoints.
	return *new(deployment.Type)
}

type Admin struct {
	Username, Password string
}

func (a *Admin) BasicAuth() (string, string) { _ = "STUB: not implemented"; return "", "" }

type IdentifierDecorator struct {
	Identifier
	Id string
}

func (d *IdentifierDecorator) ID() string { _ = "STUB: not implemented"; return "" }
