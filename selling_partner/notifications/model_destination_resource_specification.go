package notifications

import (
	"github.com/bytedance/sonic"
)

// checks if the DestinationResourceSpecification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DestinationResourceSpecification{}

// DestinationResourceSpecification The information required to create a destination resource. Applications should use one resource type (sqs or eventBridge) per destination.
type DestinationResourceSpecification struct {
	Sqs         *SqsResource                      `json:"sqs,omitempty"`
	EventBridge *EventBridgeResourceSpecification `json:"eventBridge,omitempty"`
}

// NewDestinationResourceSpecification instantiates a new DestinationResourceSpecification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDestinationResourceSpecification() *DestinationResourceSpecification {
	this := DestinationResourceSpecification{}
	return &this
}

// NewDestinationResourceSpecificationWithDefaults instantiates a new DestinationResourceSpecification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDestinationResourceSpecificationWithDefaults() *DestinationResourceSpecification {
	this := DestinationResourceSpecification{}
	return &this
}

// GetSqs returns the Sqs field value if set, zero value otherwise.
func (o *DestinationResourceSpecification) GetSqs() SqsResource {
	if o == nil || IsNil(o.Sqs) {
		var ret SqsResource
		return ret
	}
	return *o.Sqs
}

// GetSqsOk returns a tuple with the Sqs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DestinationResourceSpecification) GetSqsOk() (*SqsResource, bool) {
	if o == nil || IsNil(o.Sqs) {
		return nil, false
	}
	return o.Sqs, true
}

// HasSqs returns a boolean if a field has been set.
func (o *DestinationResourceSpecification) HasSqs() bool {
	if o != nil && !IsNil(o.Sqs) {
		return true
	}

	return false
}

// SetSqs gets a reference to the given SqsResource and assigns it to the Sqs field.
func (o *DestinationResourceSpecification) SetSqs(v SqsResource) {
	o.Sqs = &v
}

// GetEventBridge returns the EventBridge field value if set, zero value otherwise.
func (o *DestinationResourceSpecification) GetEventBridge() EventBridgeResourceSpecification {
	if o == nil || IsNil(o.EventBridge) {
		var ret EventBridgeResourceSpecification
		return ret
	}
	return *o.EventBridge
}

// GetEventBridgeOk returns a tuple with the EventBridge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DestinationResourceSpecification) GetEventBridgeOk() (*EventBridgeResourceSpecification, bool) {
	if o == nil || IsNil(o.EventBridge) {
		return nil, false
	}
	return o.EventBridge, true
}

// HasEventBridge returns a boolean if a field has been set.
func (o *DestinationResourceSpecification) HasEventBridge() bool {
	if o != nil && !IsNil(o.EventBridge) {
		return true
	}

	return false
}

// SetEventBridge gets a reference to the given EventBridgeResourceSpecification and assigns it to the EventBridge field.
func (o *DestinationResourceSpecification) SetEventBridge(v EventBridgeResourceSpecification) {
	o.EventBridge = &v
}

func (o DestinationResourceSpecification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Sqs) {
		toSerialize["sqs"] = o.Sqs
	}
	if !IsNil(o.EventBridge) {
		toSerialize["eventBridge"] = o.EventBridge
	}
	return toSerialize, nil
}

type NullableDestinationResourceSpecification struct {
	value *DestinationResourceSpecification
	isSet bool
}

func (v NullableDestinationResourceSpecification) Get() *DestinationResourceSpecification {
	return v.value
}

func (v *NullableDestinationResourceSpecification) Set(val *DestinationResourceSpecification) {
	v.value = val
	v.isSet = true
}

func (v NullableDestinationResourceSpecification) IsSet() bool {
	return v.isSet
}

func (v *NullableDestinationResourceSpecification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDestinationResourceSpecification(val *DestinationResourceSpecification) *NullableDestinationResourceSpecification {
	return &NullableDestinationResourceSpecification{value: val, isSet: true}
}

func (v NullableDestinationResourceSpecification) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableDestinationResourceSpecification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
