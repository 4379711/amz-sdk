package notifications

import (
	"github.com/bytedance/sonic"
)

// checks if the EventBridgeResourceSpecification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventBridgeResourceSpecification{}

// EventBridgeResourceSpecification The information required to create an Amazon EventBridge destination.
type EventBridgeResourceSpecification struct {
	// The AWS region in which you will be receiving the notifications.
	Region string `json:"region"`
	// The identifier for the AWS account that is responsible for charges related to receiving notifications.
	AccountId string `json:"accountId"`
}

type _EventBridgeResourceSpecification EventBridgeResourceSpecification

// NewEventBridgeResourceSpecification instantiates a new EventBridgeResourceSpecification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventBridgeResourceSpecification(region string, accountId string) *EventBridgeResourceSpecification {
	this := EventBridgeResourceSpecification{}
	this.Region = region
	this.AccountId = accountId
	return &this
}

// NewEventBridgeResourceSpecificationWithDefaults instantiates a new EventBridgeResourceSpecification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventBridgeResourceSpecificationWithDefaults() *EventBridgeResourceSpecification {
	this := EventBridgeResourceSpecification{}
	return &this
}

// GetRegion returns the Region field value
func (o *EventBridgeResourceSpecification) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *EventBridgeResourceSpecification) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *EventBridgeResourceSpecification) SetRegion(v string) {
	o.Region = v
}

// GetAccountId returns the AccountId field value
func (o *EventBridgeResourceSpecification) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *EventBridgeResourceSpecification) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *EventBridgeResourceSpecification) SetAccountId(v string) {
	o.AccountId = v
}

func (o EventBridgeResourceSpecification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["region"] = o.Region
	toSerialize["accountId"] = o.AccountId
	return toSerialize, nil
}

type NullableEventBridgeResourceSpecification struct {
	value *EventBridgeResourceSpecification
	isSet bool
}

func (v NullableEventBridgeResourceSpecification) Get() *EventBridgeResourceSpecification {
	return v.value
}

func (v *NullableEventBridgeResourceSpecification) Set(val *EventBridgeResourceSpecification) {
	v.value = val
	v.isSet = true
}

func (v NullableEventBridgeResourceSpecification) IsSet() bool {
	return v.isSet
}

func (v *NullableEventBridgeResourceSpecification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventBridgeResourceSpecification(val *EventBridgeResourceSpecification) *NullableEventBridgeResourceSpecification {
	return &NullableEventBridgeResourceSpecification{value: val, isSet: true}
}

func (v NullableEventBridgeResourceSpecification) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableEventBridgeResourceSpecification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
