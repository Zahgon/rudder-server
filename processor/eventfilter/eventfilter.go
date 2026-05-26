package eventfilter

import (
	"github.com/rudderlabs/rudder-go-kit/logger"

	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
	"github.com/rudderlabs/rudder-server/processor/types"
)

const (
	hybridModeEventsFilterKey = "hybridModeCloudEventsFilter"
	hybridMode                = "hybrid"
)

var pkgLogger = logger.NewLogger().Child("eventfilter")

// GetSupportedMessageTypes returns the supported message types for the given event, based on configuration.
// If no relevant configuration is found, returns false
func GetSupportedMessageTypes(destination *backendconfig.DestinationT) ([]string, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func identifyDisabled(destination *backendconfig.DestinationT) bool {
	_ = "STUB: not implemented"
	return false
}

// GetSupportedEvents returns the supported message events for the given destination, based on configuration.
// If no relevant configuration is found, returns false
func GetSupportedMessageEvents(destination *backendconfig.DestinationT) ([]string, bool) {
	_ = "STUB: not implemented"
	// "listOfConversions": [
	//  	{
	//  		"conversions": "Credit Card Added"
	//  	},
	//  	{
	//  		"conversions": "Credit Card Removed"
	//	    }
	// ]
	return nil, false
}

type AllowTransformerEventParams struct {
	TransformerEvent      *types.TransformerEvent
	SupportedMessageTypes []string
}

type EventParams struct {
	MessageType string
}

type ConnectionModeFilterParams struct {
	Destination      *backendconfig.DestinationT
	SrcType          string
	Event            *EventParams
	DefaultBehaviour bool
}

func getMessageType(event *types.SingularEventT) string { _ = "STUB: not implemented"; return "" }

/*
AllowEventToDestTransformation lets the caller know if we need to allow the event to proceed to destination transformation.

Currently this method supports below validations(executed in the same order):

1. Validate if messageType sent in event is included in SupportedMessageTypes

2. Validate if the event is sendable to destination based on connectionMode, sourceType & messageType
*/
func AllowEventToDestTransformation(transformerEvent *types.TransformerEvent, supportedMsgTypes []string) (bool, *types.TransformerResponse) {
	_ = "STUB: not implemented"
	// MessageType filtering -- STARTS
	return false, nil
}

// We will abort the event

// We will not allow the event

// MessageType filtering -- ENDS

// hybridModeCloudEventsFilter.srcType.[eventProperty] filtering -- STARTS

// Default behavior
// When something is missing in "supportedConnectionModes" or if "supportedConnectionModes" is not defined
// We would be checking for below things
// 1. Check if the event.type value is present in destination.DestinationDefinition.Config["supportedMessageTypes"]
// 2. Check if the connectionMode of destination is cloud or hybrid(evaluated through `IsProcessorEnabled`)
// Only when 1 & 2 are true, we would allow the event to flow through to server
// As when this will be called, we would have already checked if event.type in supportedMessageTypes

// hybridModeCloudEventsFilter.srcType.[eventProperty] filtering -- ENDS

/*
FilterEventsForHybridMode lets the caller know if the event is allowed to flow through server for a `specific destination`
Introduced to support hybrid-mode event filtering on cloud-side

The template inside `destinationDefinition.Config.hybridModeCloudEventsFilter` would look like this
```

	[sourceType]: {
		[eventProperty]: [...supportedEventPropertyValues]
	}

```

Example:

		{
			...
			"hybridModeCloudEventsFilter": {
	      "web": {
	        "messageType": ["track", "page"]
	      }
	    },
			...
		}
*/
func FilterEventsForHybridMode(connectionModeFilterParams ConnectionModeFilterParams) bool {
	_ = "STUB: not implemented"
	return false
}

// Flag indicating to let the event pass through

type EventPropsTypes interface {
	~string
}

/*
* Converts interface{} to []T if the go type-assertion allows it
 */
func ConvertToArrayOfType[T EventPropsTypes](data any) []T { _ = "STUB: not implemented"; return nil }
