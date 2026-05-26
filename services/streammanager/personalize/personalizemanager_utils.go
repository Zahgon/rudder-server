package personalize

import (
	"encoding/json"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/personalizeevents"
	"github.com/aws/aws-sdk-go-v2/service/personalizeevents/types"

	"github.com/rudderlabs/rudder-go-kit/logger"
)

var pkgLogger logger.Logger

func init() {
	pkgLogger = logger.NewLogger().Child("streammanager").Child("personalize")
}

type PersonalizeEvent struct {
	EventList  []Event
	SessionId  *string
	TrackingId *string
	UserId     *string
}

type Event struct {
	EventType         *string
	SentAt            *time.Time
	EventId           *string
	EventValue        *float32
	Impression        []string
	ItemId            *string
	MetricAttribution *MetricAttribution
	Properties        *json.RawMessage
	RecommendationId  *string
}

type MetricAttribution struct {
	EventAttributionSource *string
}

func (ma *MetricAttribution) ToAWSEventAttribution() *types.MetricAttribution {
	_ = "STUB: not implemented"
	return nil
}

func (e *Event) ToAWSEvent() types.Event { _ = "STUB: not implemented"; return *new(types.Event) }

func (p *PersonalizeEvent) ToPutEventsInput() *personalizeevents.PutEventsInput {
	_ = "STUB: not implemented"
	return nil
}

type Users struct {
	Users      []User
	DatasetArn *string
}

type User struct {
	UserId     *string
	Properties *json.RawMessage
}

func (u *User) ToAWSUser() types.User { _ = "STUB: not implemented"; return *new(types.User) }

func (u *Users) ToPutUsersInput() *personalizeevents.PutUsersInput {
	_ = "STUB: not implemented"
	return nil
}

type Items struct {
	DatasetArn *string
	Items      []Item
}

type Item struct {
	ItemId     *string
	Properties *json.RawMessage
}

func (i *Item) ToAWSItem() types.Item { _ = "STUB: not implemented"; return *new(types.Item) }

func (i *Items) ToPutItemsInput() *personalizeevents.PutItemsInput {
	_ = "STUB: not implemented"
	return nil
}

func stringifyJsonRaw(raw *json.RawMessage) *string { _ = "STUB: not implemented"; return nil }
