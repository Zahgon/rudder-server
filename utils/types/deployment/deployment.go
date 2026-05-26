package deployment

import (
	"github.com/rudderlabs/rudder-go-kit/logger"
)

type Type string // skipcq: RVV-B0009

const (
	DedicatedType   Type = "DEDICATED"
	MultiTenantType Type = "MULTITENANT"
)

// Types of tokens that can be used to authenticate with CP router
const (
	workspaceToken = "WORKSPACE_TOKEN"
	namespace      = "NAMESPACE"
)

const defaultClusterType = DedicatedType

var pkgLogger = logger.NewLogger().Child("deployment")

func GetFromEnv() (Type, error) { _ = "STUB: not implemented"; return *new(Type), nil }

func (t Type) Valid() bool { _ = "STUB: not implemented"; return false }

func GetConnectionToken() (string, string, bool, error) {
	_ = "STUB: not implemented"
	return "", "", false, nil
}
