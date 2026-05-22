package notifications

import (
	"github.com/bytedance/sonic"
)

// checks if the EventBridgeResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventBridgeResource{}

// EventBridgeResource The Amazon EventBridge destination.
type EventBridgeResource struct {
	// The name of the partner event source associated with the destination.
	Name string `json:"name"`
	// The AWS region in which you receive the notifications. For AWS regions that are supported in Amazon EventBridge, refer to [Amazon EventBridge endpoints and quotas](https://docs.aws.amazon.com/general/latest/gr/ev.html).
	Region string `json:"region"`
	// The identifier for the AWS account that is responsible for charges related to receiving notifications.
	AccountId string `json:"accountId"`
}

type _EventBridgeResource EventBridgeResource

// NewEventBridgeResource instantiates a new EventBridgeResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventBridgeResource(name string, region string, accountId string) *EventBridgeResource {
	this := EventBridgeResource{}
	this.Name = name
	this.Region = region
	this.AccountId = accountId
	return &this
}

// NewEventBridgeResourceWithDefaults instantiates a new EventBridgeResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventBridgeResourceWithDefaults() *EventBridgeResource {
	this := EventBridgeResource{}
	return &this
}

// GetName returns the Name field value
func (o *EventBridgeResource) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EventBridgeResource) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EventBridgeResource) SetName(v string) {
	o.Name = v
}

// GetRegion returns the Region field value
func (o *EventBridgeResource) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *EventBridgeResource) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *EventBridgeResource) SetRegion(v string) {
	o.Region = v
}

// GetAccountId returns the AccountId field value
func (o *EventBridgeResource) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *EventBridgeResource) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *EventBridgeResource) SetAccountId(v string) {
	o.AccountId = v
}

func (o EventBridgeResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["region"] = o.Region
	toSerialize["accountId"] = o.AccountId
	return toSerialize, nil
}

type NullableEventBridgeResource struct {
	value *EventBridgeResource
	isSet bool
}

func (v NullableEventBridgeResource) Get() *EventBridgeResource {
	return v.value
}

func (v *NullableEventBridgeResource) Set(val *EventBridgeResource) {
	v.value = val
	v.isSet = true
}

func (v NullableEventBridgeResource) IsSet() bool {
	return v.isSet
}

func (v *NullableEventBridgeResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventBridgeResource(val *EventBridgeResource) *NullableEventBridgeResource {
	return &NullableEventBridgeResource{value: val, isSet: true}
}

func (v NullableEventBridgeResource) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableEventBridgeResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
