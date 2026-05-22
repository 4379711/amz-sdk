package notifications

import (
	"github.com/bytedance/sonic"
)

// checks if the DestinationResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DestinationResource{}

// DestinationResource The destination resource types.
type DestinationResource struct {
	Sqs         *SqsResource         `json:"sqs,omitempty"`
	EventBridge *EventBridgeResource `json:"eventBridge,omitempty"`
}

// NewDestinationResource instantiates a new DestinationResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDestinationResource() *DestinationResource {
	this := DestinationResource{}
	return &this
}

// NewDestinationResourceWithDefaults instantiates a new DestinationResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDestinationResourceWithDefaults() *DestinationResource {
	this := DestinationResource{}
	return &this
}

// GetSqs returns the Sqs field value if set, zero value otherwise.
func (o *DestinationResource) GetSqs() SqsResource {
	if o == nil || IsNil(o.Sqs) {
		var ret SqsResource
		return ret
	}
	return *o.Sqs
}

// GetSqsOk returns a tuple with the Sqs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DestinationResource) GetSqsOk() (*SqsResource, bool) {
	if o == nil || IsNil(o.Sqs) {
		return nil, false
	}
	return o.Sqs, true
}

// HasSqs returns a boolean if a field has been set.
func (o *DestinationResource) HasSqs() bool {
	if o != nil && !IsNil(o.Sqs) {
		return true
	}

	return false
}

// SetSqs gets a reference to the given SqsResource and assigns it to the Sqs field.
func (o *DestinationResource) SetSqs(v SqsResource) {
	o.Sqs = &v
}

// GetEventBridge returns the EventBridge field value if set, zero value otherwise.
func (o *DestinationResource) GetEventBridge() EventBridgeResource {
	if o == nil || IsNil(o.EventBridge) {
		var ret EventBridgeResource
		return ret
	}
	return *o.EventBridge
}

// GetEventBridgeOk returns a tuple with the EventBridge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DestinationResource) GetEventBridgeOk() (*EventBridgeResource, bool) {
	if o == nil || IsNil(o.EventBridge) {
		return nil, false
	}
	return o.EventBridge, true
}

// HasEventBridge returns a boolean if a field has been set.
func (o *DestinationResource) HasEventBridge() bool {
	if o != nil && !IsNil(o.EventBridge) {
		return true
	}

	return false
}

// SetEventBridge gets a reference to the given EventBridgeResource and assigns it to the EventBridge field.
func (o *DestinationResource) SetEventBridge(v EventBridgeResource) {
	o.EventBridge = &v
}

func (o DestinationResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Sqs) {
		toSerialize["sqs"] = o.Sqs
	}
	if !IsNil(o.EventBridge) {
		toSerialize["eventBridge"] = o.EventBridge
	}
	return toSerialize, nil
}

type NullableDestinationResource struct {
	value *DestinationResource
	isSet bool
}

func (v NullableDestinationResource) Get() *DestinationResource {
	return v.value
}

func (v *NullableDestinationResource) Set(val *DestinationResource) {
	v.value = val
	v.isSet = true
}

func (v NullableDestinationResource) IsSet() bool {
	return v.isSet
}

func (v *NullableDestinationResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDestinationResource(val *DestinationResource) *NullableDestinationResource {
	return &NullableDestinationResource{value: val, isSet: true}
}

func (v NullableDestinationResource) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableDestinationResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
