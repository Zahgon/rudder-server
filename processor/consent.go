package processor

import (
	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
	"github.com/rudderlabs/rudder-server/processor/types"
)

type ConsentManagementInfo struct {
	DeniedConsentIDs   []string `json:"deniedConsentIds"`
	AllowedConsentIDs  any      `json:"allowedConsentIds"` // Not used currently but added for future use
	Provider           string   `json:"provider"`
	ResolutionStrategy string   `json:"resolutionStrategy"`
}

type GenericConsentManagementProviderData struct {
	ResolutionStrategy string
	Consents           []string
}

type GenericConsentsConfig struct {
	Consent string `json:"consent"`
}

type GenericConsentManagementProviderConfig struct {
	Provider           string                  `json:"provider"`
	ResolutionStrategy string                  `json:"resolutionStrategy"`
	Consents           []GenericConsentsConfig `json:"consents"`
}

/*
Filters and returns destinations based on the consents configured for the destination and the user consents present in the event.

Supports legacy and generic consent management.
For GCM based filtering, uses source and destination IDs to fetch the appropriate GCM data from the config.
*/
func (proc *Handle) getConsentFilteredDestinations(event types.SingularEventT, sourceID string, destinations []backendconfig.DestinationT) []backendconfig.DestinationT {
	_ = "STUB: not implemented"
	// If the event does not have denied consent IDs, do not filter any destinations
	return nil
}

// Log the error for debugging purposes

// Generic consent management

// For custom provider, the resolution strategy is to be picked from the destination config

// The user must consent to at least one of the configured consents in the destination

// The user must consent to all of the configured consents in the destination
// "and"

// Legacy consent management

// If the destination has oneTrustCookieCategories, returns false if any of the oneTrustCategories are present in deniedCategories

// If the destination has ketchConsentPurposes, returns false if all ketchPurposes are present in deniedCategories

func (proc *Handle) getOneTrustConsentData(destinationID string) []string {
	_ = "STUB: not implemented"
	return nil
}

func (proc *Handle) getKetchConsentData(destinationID string) []string {
	_ = "STUB: not implemented"
	return nil
}

func (proc *Handle) getGCMData(sourceID, destinationID, provider string) GenericConsentManagementProviderData {
	_ = "STUB: not implemented"
	return *new(GenericConsentManagementProviderData)
}

func getOneTrustConsentCategories(dest *backendconfig.DestinationT) []string {
	_ = "STUB: not implemented"
	return nil
}

// Handle the case where oneTrustCookieCategories is not a slice

func getKetchConsentCategories(dest *backendconfig.DestinationT) []string {
	_ = "STUB: not implemented"
	return nil
}

// Handle the case where ketchConsentPurposes is not a slice

func getGenericConsentManagementData(dest *backendconfig.DestinationT) (ConsentProviderMap, error) {
	_ = "STUB: not implemented"
	return *new(ConsentProviderMap), nil
}

func getConsentManagementInfo(event types.SingularEventT) (ConsentManagementInfo, error) {
	_ = "STUB: not implemented"
	return *new(ConsentManagementInfo), nil
}
