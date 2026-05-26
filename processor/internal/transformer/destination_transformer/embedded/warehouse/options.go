package warehouse

func extractIntrOpts(destType string, message map[string]any) intrOptions {
	_ = "STUB: not implemented"
	return *new(intrOptions)
}

func mergeDataWarehouseIntrOpts(destType string, message map[string]any, opts intrOptions) intrOptions {
	_ = "STUB: not implemented"
	return *new(intrOptions)
}

func extractDestOpts(destType string, destConfig map[string]any) destOptions {
	_ = "STUB: not implemented"
	return *new(destOptions)
}

func setOption[T any](src map[string]any, key string, dest *T) { _ = "STUB: not implemented"; return }
