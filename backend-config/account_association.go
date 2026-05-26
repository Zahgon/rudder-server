package backendconfig

import (
	"github.com/rudderlabs/rudder-go-kit/logger"
)

// processAccountAssociations processes account configurations and merges them with their corresponding
// account definitions. It then associates these merged accounts with destinations based on
// their configuration settings.
//
// The process involves:
// 1. Iterating through all sources and their destinations
// 2. For each destination, setting up account associations using setDestinationAccounts
func (c *ConfigT) processAccountAssociations() {
	_ = "STUB: not implemented"
	// Iterate through all sources and their destinations to set up account associations
	return
}

// getAccountDefinition returns a pointer to the account definition for the given account definition name.
// If the account definition name is empty or the account definition doesn't exist, it returns nil.
//
// Parameters:
//   - accountDefinitionName: The name of the account definition to look up
//   - accountAssociationLogger: The logger instance to use for logging
//
// Returns:
//   - *AccountDefinition: A pointer to the account definition or nil if not found
func (c *ConfigT) getAccountDefinition(accountDefinitionName string, accountAssociationLogger logger.Logger) *AccountDefinition {
	_ = "STUB: not implemented"
	return nil
}

// Log error if account definition not found

// enrichAccountWithDefinition populates an account to a destination field based on the account ID.
// It creates an Account object by combining account details with its definition.
//
// Parameters:
//   - accountID: The ID of the account to populate
//   - accountAssociationLogger: The logger instance to use for logging
//
// Returns:
//   - *Account: The populated account object or nil if the account doesn't exist
func (c *ConfigT) enrichAccountWithDefinition(accountID string, accountAssociationLogger logger.Logger) *Account {
	_ = "STUB: not implemented"
	return nil
}

// enrichDestinationWithAccounts assigns accounts to a destination based on its configuration.
// It handles two types of account associations:
// 1. Regular account (rudderAccountId) - Used for normal event delivery flow
// 2. Delete account (rudderDeleteAccountId) - Used for data deletion/regulation flow
//
// For each account type, it:
// 1. Checks if the corresponding account ID exists in the destination config
// 2. Verifies the account exists in the accounts map
// 3. Creates an account by combining account details with its definition
// 4. Assigns it to the appropriate field in the destination
//
// Parameters:
//   - dest: Pointer to the destination being configured
func (c *ConfigT) enrichDestinationWithAccounts(dest *DestinationT) {
	_ = "STUB: not implemented"
	return
}

// Check and set the delivery account if specified in the destination config

// Check and set the delete account if specified in the destination config
