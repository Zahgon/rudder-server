package warehouse

import (
	wtypes "github.com/rudderlabs/rudder-server/processor/internal/transformer/destination_transformer/embedded/warehouse/internal/types"
	"github.com/rudderlabs/rudder-server/processor/types"
)

func (t *Transformer) mergeEvents(tec *transformEventContext) ([]map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func mergeProps(message types.SingularEventT, metadata wtypes.Metadata) (*mergeRule, *mergeRule, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func mergePropsForMergeEventType(message types.SingularEventT) (*mergeRule, *mergeRule, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func mergePropsForAliasEventType(message types.SingularEventT) (*mergeRule, *mergeRule, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func mergePropsForDefaultEventType(message types.SingularEventT) (*mergeRule, *mergeRule, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func mergeRuleTable(tec *transformEventContext) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func mergeRuleColumns(tec *transformEventContext) (*mergeRulesColumns, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func isMergePropEmpty(mergeProp *mergeRule) bool { _ = "STUB: not implemented"; return false }
