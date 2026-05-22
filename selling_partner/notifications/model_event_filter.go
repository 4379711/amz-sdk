package notifications

import (
	"github.com/bytedance/sonic"
)

// checks if the EventFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventFilter{}

// EventFilter A `notificationType` specific filter. This object contains all of the currently available filters and properties that you can use to define a `notificationType` specific filter.
type EventFilter struct {
	AggregationSettings *AggregationSettings `json:"aggregationSettings,omitempty"`
	// A list of marketplace identifiers to subscribe to (for example: ATVPDKIKX0DER). To receive notifications in every marketplace, do not provide this list.
	MarketplaceIds []string `json:"marketplaceIds,omitempty"`
	// A list of order change types to subscribe to (for example: `BuyerRequestedChange`). To receive notifications of all change types, do not provide this list.
	OrderChangeTypes []OrderChangeTypeEnum `json:"orderChangeTypes,omitempty"`
	// An `eventFilterType` value that is supported by the specific `notificationType`. This is used by the subscription service to determine the type of event filter. Refer to [Notification Type Values](https://developer-docs.amazon.com/sp-api/docs/notification-type-values) to determine if an `eventFilterType` is supported.
	EventFilterType string `json:"eventFilterType"`
}

type _EventFilter EventFilter

// NewEventFilter instantiates a new EventFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventFilter(eventFilterType string) *EventFilter {
	this := EventFilter{}
	this.EventFilterType = eventFilterType
	return &this
}

// NewEventFilterWithDefaults instantiates a new EventFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventFilterWithDefaults() *EventFilter {
	this := EventFilter{}
	return &this
}

// GetAggregationSettings returns the AggregationSettings field value if set, zero value otherwise.
func (o *EventFilter) GetAggregationSettings() AggregationSettings {
	if o == nil || IsNil(o.AggregationSettings) {
		var ret AggregationSettings
		return ret
	}
	return *o.AggregationSettings
}

// GetAggregationSettingsOk returns a tuple with the AggregationSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventFilter) GetAggregationSettingsOk() (*AggregationSettings, bool) {
	if o == nil || IsNil(o.AggregationSettings) {
		return nil, false
	}
	return o.AggregationSettings, true
}

// HasAggregationSettings returns a boolean if a field has been set.
func (o *EventFilter) HasAggregationSettings() bool {
	if o != nil && !IsNil(o.AggregationSettings) {
		return true
	}

	return false
}

// SetAggregationSettings gets a reference to the given AggregationSettings and assigns it to the AggregationSettings field.
func (o *EventFilter) SetAggregationSettings(v AggregationSettings) {
	o.AggregationSettings = &v
}

// GetMarketplaceIds returns the MarketplaceIds field value if set, zero value otherwise.
func (o *EventFilter) GetMarketplaceIds() []string {
	if o == nil || IsNil(o.MarketplaceIds) {
		var ret []string
		return ret
	}
	return o.MarketplaceIds
}

// GetMarketplaceIdsOk returns a tuple with the MarketplaceIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventFilter) GetMarketplaceIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.MarketplaceIds) {
		return nil, false
	}
	return o.MarketplaceIds, true
}

// HasMarketplaceIds returns a boolean if a field has been set.
func (o *EventFilter) HasMarketplaceIds() bool {
	if o != nil && !IsNil(o.MarketplaceIds) {
		return true
	}

	return false
}

// SetMarketplaceIds gets a reference to the given []string and assigns it to the MarketplaceIds field.
func (o *EventFilter) SetMarketplaceIds(v []string) {
	o.MarketplaceIds = v
}

// GetOrderChangeTypes returns the OrderChangeTypes field value if set, zero value otherwise.
func (o *EventFilter) GetOrderChangeTypes() []OrderChangeTypeEnum {
	if o == nil || IsNil(o.OrderChangeTypes) {
		var ret []OrderChangeTypeEnum
		return ret
	}
	return o.OrderChangeTypes
}

// GetOrderChangeTypesOk returns a tuple with the OrderChangeTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventFilter) GetOrderChangeTypesOk() ([]OrderChangeTypeEnum, bool) {
	if o == nil || IsNil(o.OrderChangeTypes) {
		return nil, false
	}
	return o.OrderChangeTypes, true
}

// HasOrderChangeTypes returns a boolean if a field has been set.
func (o *EventFilter) HasOrderChangeTypes() bool {
	if o != nil && !IsNil(o.OrderChangeTypes) {
		return true
	}

	return false
}

// SetOrderChangeTypes gets a reference to the given []OrderChangeTypeEnum and assigns it to the OrderChangeTypes field.
func (o *EventFilter) SetOrderChangeTypes(v []OrderChangeTypeEnum) {
	o.OrderChangeTypes = v
}

// GetEventFilterType returns the EventFilterType field value
func (o *EventFilter) GetEventFilterType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventFilterType
}

// GetEventFilterTypeOk returns a tuple with the EventFilterType field value
// and a boolean to check if the value has been set.
func (o *EventFilter) GetEventFilterTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventFilterType, true
}

// SetEventFilterType sets field value
func (o *EventFilter) SetEventFilterType(v string) {
	o.EventFilterType = v
}

func (o EventFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AggregationSettings) {
		toSerialize["aggregationSettings"] = o.AggregationSettings
	}
	if !IsNil(o.MarketplaceIds) {
		toSerialize["marketplaceIds"] = o.MarketplaceIds
	}
	if !IsNil(o.OrderChangeTypes) {
		toSerialize["orderChangeTypes"] = o.OrderChangeTypes
	}
	toSerialize["eventFilterType"] = o.EventFilterType
	return toSerialize, nil
}

type NullableEventFilter struct {
	value *EventFilter
	isSet bool
}

func (v NullableEventFilter) Get() *EventFilter {
	return v.value
}

func (v *NullableEventFilter) Set(val *EventFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableEventFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableEventFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventFilter(val *EventFilter) *NullableEventFilter {
	return &NullableEventFilter{value: val, isSet: true}
}

func (v NullableEventFilter) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableEventFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
