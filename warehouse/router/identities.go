package router

import (
	"context"
	"sync"

	"github.com/rudderlabs/rudder-server/warehouse/internal/model"
)

var (
	populatingHistoricIdentitiesProgressMap     map[string]bool
	populatingHistoricIdentitiesProgressMapLock sync.RWMutex
	populatedHistoricIdentitiesMap              map[string]bool
	populatedHistoricIdentitiesMapLock          sync.RWMutex
)

func init() {
	populatingHistoricIdentitiesProgressMap = map[string]bool{}
	populatedHistoricIdentitiesMap = map[string]bool{}
}

func uniqueWarehouseNamespaceString(warehouse model.Warehouse) string {
	_ = "STUB: not implemented"
	return ""
}

func isDestHistoricIdentitiesPopulated(warehouse model.Warehouse) bool {
	_ = "STUB: not implemented"
	return false
}

func setDestHistoricIdentitiesPopulated(warehouse model.Warehouse) {
	_ = "STUB: not implemented"
	return
}

func setDestHistoricIdentitiesPopulateInProgress(warehouse model.Warehouse, starting bool) {
	_ = "STUB: not implemented"
	return
}

func isDestHistoricIdentitiesPopulateInProgress(warehouse model.Warehouse) bool {
	_ = "STUB: not implemented"
	return false
}

func (r *Router) getPendingPopulateIdentitiesLoad(warehouse model.Warehouse) (upload model.Upload, found bool) {
	_ = "STUB: not implemented"
	return *new(model.Upload), false
}

func (r *Router) populateHistoricIdentitiesDestType() string { _ = "STUB: not implemented"; return "" }

func (r *Router) hasLocalIdentityData(warehouse model.Warehouse) (exists bool) {
	_ = "STUB: not implemented"
	return false
}

// TODO: Handle this

func (r *Router) hasWarehouseData(ctx context.Context, warehouse model.Warehouse) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *Router) setupIdentityTables(ctx context.Context, warehouse model.Warehouse) {
	_ = "STUB: not implemented"
	return
}

// create tables

func (r *Router) initPrePopulateDestIdentitiesUpload(warehouse model.Warehouse) model.Upload {
	_ = "STUB: not implemented"
	return *new(model.Upload)
}

// TODO: DRY this code

// add rudder_identity_mappings table

func (*Router) setFailedStat(warehouse model.Warehouse, err error) {
	_ = "STUB: not implemented"
	return
}

func (r *Router) populateHistoricIdentities(ctx context.Context, warehouse model.Warehouse) {
	_ = "STUB: not implemented"
	return
}

// check for pending loads (populateHistoricIdentities)

// TODO: Handle error / Retry
