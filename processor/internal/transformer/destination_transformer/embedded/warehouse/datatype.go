package warehouse

func dataTypeFor(destType, key string, val any, isJSONKey bool) string {
	_ = "STUB: not implemented"
	return ""
}

func primitiveType(val any) string { _ = "STUB: not implemented"; return "" }

func getFloatType(v float64) string {
	_ = "STUB: not implemented"
	// JSON unmarshalling treats all numbers as float64 by default, even if they are whole numbers
	// So, we need to check if the float is actually an integer
	// We are using Number.isInteger(val) for detecting whether datatype is int or float in rudder-transformer
	// which has higher range then what we have in Golang (9223372036854775807), therefore using big package for determining the type
	return ""
}

func dataTypeOverride(destType, key string, val any, isJSONKey bool) string {
	_ = "STUB: not implemented"
	return ""
}

func overrideForPostgresSnowflake(key string, isJSONKey bool) string {
	_ = "STUB: not implemented"
	return ""
}

func overrideForRedshift(val any, isJSONKey bool) string { _ = "STUB: not implemented"; return "" }

func shouldUseTextForRedshift(data any) bool { _ = "STUB: not implemented"; return false }

// Javascript strings are UTF-16 encoded, use utf16 instead of utf8 package for determining the length

// Javascript strings are UTF-16 encoded, use utf16 instead of utf8 package for determining the length

func convertValIfDateTime(val any, colType string) any { _ = "STUB: not implemented"; return *new(any) }
