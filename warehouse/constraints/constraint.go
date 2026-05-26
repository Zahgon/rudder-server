package constraints

import (
	"github.com/rudderlabs/rudder-go-kit/config"

	"github.com/rudderlabs/rudder-server/warehouse/utils/types"
)

type constraints interface {
	violates(brEvent *types.BatchRouterEvent, columnName string) (cv *Violation)
}

type Violation struct {
	IsViolated         bool
	ViolatedIdentifier string
	Reason             string
}

type indexConstraint struct {
	tableName    string
	columnName   string
	indexColumns []string
	limit        int
	reason       string
}

type Manager struct {
	constraintsMap              map[string][]constraints
	enableConstraintsViolations config.ValueLoader[bool]
}

func New(conf *config.Config) *Manager { _ = "STUB: not implemented"; return nil }

func (cm *Manager) ViolatedConstraints(destinationType string, brEvent *types.BatchRouterEvent, columnName string) (cv *Violation) {
	_ = "STUB: not implemented"
	return nil
}

func (ic *indexConstraint) violates(brEvent *types.BatchRouterEvent, columnName string) *Violation {
	_ = "STUB: not implemented"
	return nil
}
