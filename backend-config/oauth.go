package backendconfig

import (
	"github.com/rudderlabs/rudder-server/services/oauth/v2/common"
)

const (
	destinationDefinitionOAuthType = "OAuth"
	accountDefinitionOAuthType     = "oauth"
)

/*
GetAccountID Gets AccountId for OAuth destination based on if rudderFlow is `Delivery` or `Delete`

Example:
`dest.GetAccountID(common.RudderFlowDelete)` --> To be used when we make use of OAuth during regulation flow
`dest.GetAccountID(common.RudderFlowDelivery)` --> To be used when we make use of OAuth during normal event delivery
*/
func (d *DestinationT) GetAccountID(flow common.RudderFlow) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// IsOAuthDestination checks if a destination is configured for OAuth authentication.
// If the destination has an account with an account definition for the given flow,
// the account definition's AuthenticationType is used to determine if it is OAuth.
// Otherwise, it falls back to checking the destination definition config.
func (d *DestinationT) IsOAuthDestination(flow common.RudderFlow) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// valid use-case for non-OAuth destinations

// we should throw error here, as we expect authValue to be a string if present

// resolveAccount resolves the account associated with the destination based on the flow.
// It returns the account for the specified flow (delivery or delete) if it exists, otherwise it returns nil.
func (d *DestinationT) resolveAccount(flow common.RudderFlow) *Account {
	_ = "STUB: not implemented"
	return nil
}

func isOAuthSupportedForFlow(definitionConfig map[string]any, flow common.RudderFlow) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// valid use-case for non-OAuth destinations
// when the auth.type is OAuth and rudderScopes is not mentioned, we would assume oauth flow is to be used when it is in "delivery" flow
