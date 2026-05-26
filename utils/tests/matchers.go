package testutils

import (
	"github.com/onsi/gomega/types"
)

type beValidUUIDMatcher struct{}

func (*beValidUUIDMatcher) Match(actual any) (success bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (*beValidUUIDMatcher) FailureMessage(actual any) (message string) {
	_ = "STUB: not implemented"
	return ""
}

func (*beValidUUIDMatcher) NegatedFailureMessage(actual any) (message string) {
	_ = "STUB: not implemented"
	return ""
}

/*
BeValidUUID returns a gomega matcher that checks validity of a UUID string
*/
func BeValidUUID() types.GomegaMatcher { _ = "STUB: not implemented"; return *new(types.GomegaMatcher) }
