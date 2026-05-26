package testhelper

type OutputBuilder map[string]any

func (ob OutputBuilder) SetDataField(key string, value any) OutputBuilder {
	_ = "STUB: not implemented"
	return *new(OutputBuilder)
}

func (ob OutputBuilder) SetMetadata(key string, value any) OutputBuilder {
	_ = "STUB: not implemented"
	return *new(OutputBuilder)
}

func (ob OutputBuilder) SetColumnField(key string, value any) OutputBuilder {
	_ = "STUB: not implemented"
	return *new(OutputBuilder)
}

func (ob OutputBuilder) SetTableName(tableName string) OutputBuilder {
	_ = "STUB: not implemented"
	return *new(OutputBuilder)
}

func (ob OutputBuilder) RemoveDataFields(fields ...string) OutputBuilder {
	_ = "STUB: not implemented"
	return *new(OutputBuilder)
}

func (ob OutputBuilder) RemoveMetadata(fields ...string) OutputBuilder {
	_ = "STUB: not implemented"
	return *new(OutputBuilder)
}

func (ob OutputBuilder) RemoveColumnFields(fields ...string) OutputBuilder {
	_ = "STUB: not implemented"
	return *new(OutputBuilder)
}

func (ob OutputBuilder) AddRandomEntries(count int, predicate func(index int) (dataKey, dataValue, columnKey, columnValue string)) OutputBuilder {
	_ = "STUB: not implemented"
	return *new(OutputBuilder)
}

func (ob OutputBuilder) BuildForSnowflake() OutputBuilder {
	_ = "STUB: not implemented"
	return *new(OutputBuilder)
}
