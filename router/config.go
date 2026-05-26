package router

import (
	"github.com/rudderlabs/rudder-go-kit/config"
)

func getRouterConfigBool(key, destType string, defaultValue bool) bool {
	_ = "STUB: not implemented"
	return false
}

func getRouterConfigInt(key, destType string, defaultValue int) int {
	_ = "STUB: not implemented"
	return 0
}

func getHierarchicalRouterConfigInt(destType string, defaultValue int, keys ...string) int {
	_ = "STUB: not implemented"
	return 0
}

func getReloadableRouterConfigInt(key, destType string, defaultValue int) config.ValueLoader[int] {
	_ = "STUB: not implemented"
	return nil
}

func getRouterConfigKeys(key, destType string) []string { _ = "STUB: not implemented"; return nil }
