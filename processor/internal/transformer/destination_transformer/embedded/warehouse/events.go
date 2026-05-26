package warehouse

func (t *Transformer) trackEvents(tec *transformEventContext) ([]map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *Transformer) trackCommonProps(tec *transformEventContext) (map[string]any, map[string]string, string, error) {
	_ = "STUB: not implemented"
	return nil, nil, "", nil
}

func (t *Transformer) tracksResponse(tec *transformEventContext, commonData map[string]any, commonMetadata map[string]string) ([]map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *Transformer) trackEventsResponse(tec *transformEventContext, transformerEventName string, commonData map[string]any, commonMetadata map[string]string) ([]map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *Transformer) extractEvents(tec *transformEventContext) ([]map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *Transformer) extractCommonProps(tec *transformEventContext) (map[string]any, map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func excludeRudderCreatedTableNames(name string, skipReservedKeywordsEscaping bool) string {
	_ = "STUB: not implemented"
	return ""
}

func (t *Transformer) identifyEvents(tec *transformEventContext) ([]map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *Transformer) identifyCommonProps(tec *transformEventContext) (map[string]any, map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (t *Transformer) identifiesResponse(tec *transformEventContext, commonData map[string]any, commonMetadata map[string]string) ([]map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *Transformer) usersResponse(tec *transformEventContext, commonData map[string]any, commonMetadata map[string]string) ([]map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func shouldSkipUsersTable(tec *transformEventContext) bool { _ = "STUB: not implemented"; return false }

func (t *Transformer) pageEvents(tec *transformEventContext) ([]map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *Transformer) screenEvents(tec *transformEventContext) ([]map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *Transformer) groupEvents(tec *transformEventContext) ([]map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *Transformer) aliasEvents(tec *transformEventContext) ([]map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
