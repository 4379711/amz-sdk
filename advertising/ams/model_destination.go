package ams

import (
	"github.com/bytedance/sonic"
)

// checks if the Destination type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Destination{}

// Destination struct for Destination
type Destination struct {
	FirehoseDestination *FirehoseDestination `json:"firehoseDestination,omitempty"`
	SqsDestination      *SqsDestination      `json:"sqsDestination,omitempty"`
}

// NewDestination instantiates a new Destination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDestination() *Destination {
	this := Destination{}
	return &this
}

// NewDestinationWithDefaults instantiates a new Destination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDestinationWithDefaults() *Destination {
	this := Destination{}
	return &this
}

// GetFirehoseDestination returns the FirehoseDestination field value if set, zero value otherwise.
func (o *Destination) GetFirehoseDestination() FirehoseDestination {
	if o == nil || IsNil(o.FirehoseDestination) {
		var ret FirehoseDestination
		return ret
	}
	return *o.FirehoseDestination
}

// GetFirehoseDestinationOk returns a tuple with the FirehoseDestination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Destination) GetFirehoseDestinationOk() (*FirehoseDestination, bool) {
	if o == nil || IsNil(o.FirehoseDestination) {
		return nil, false
	}
	return o.FirehoseDestination, true
}

// HasFirehoseDestination returns a boolean if a field has been set.
func (o *Destination) HasFirehoseDestination() bool {
	if o != nil && !IsNil(o.FirehoseDestination) {
		return true
	}

	return false
}

// SetFirehoseDestination gets a reference to the given FirehoseDestination and assigns it to the FirehoseDestination field.
func (o *Destination) SetFirehoseDestination(v FirehoseDestination) {
	o.FirehoseDestination = &v
}

// GetSqsDestination returns the SqsDestination field value if set, zero value otherwise.
func (o *Destination) GetSqsDestination() SqsDestination {
	if o == nil || IsNil(o.SqsDestination) {
		var ret SqsDestination
		return ret
	}
	return *o.SqsDestination
}

// GetSqsDestinationOk returns a tuple with the SqsDestination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Destination) GetSqsDestinationOk() (*SqsDestination, bool) {
	if o == nil || IsNil(o.SqsDestination) {
		return nil, false
	}
	return o.SqsDestination, true
}

// HasSqsDestination returns a boolean if a field has been set.
func (o *Destination) HasSqsDestination() bool {
	if o != nil && !IsNil(o.SqsDestination) {
		return true
	}

	return false
}

// SetSqsDestination gets a reference to the given SqsDestination and assigns it to the SqsDestination field.
func (o *Destination) SetSqsDestination(v SqsDestination) {
	o.SqsDestination = &v
}

func (o Destination) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FirehoseDestination) {
		toSerialize["firehoseDestination"] = o.FirehoseDestination
	}
	if !IsNil(o.SqsDestination) {
		toSerialize["sqsDestination"] = o.SqsDestination
	}
	return toSerialize, nil
}

type NullableDestination struct {
	value *Destination
	isSet bool
}

func (v NullableDestination) Get() *Destination {
	return v.value
}

func (v *NullableDestination) Set(val *Destination) {
	v.value = val
	v.isSet = true
}

func (v NullableDestination) IsSet() bool {
	return v.isSet
}

func (v *NullableDestination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDestination(val *Destination) *NullableDestination {
	return &NullableDestination{value: val, isSet: true}
}

func (v NullableDestination) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableDestination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
