package ams

import (
	"github.com/bytedance/sonic"
)

// checks if the SqsDestination type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SqsDestination{}

// SqsDestination struct for SqsDestination
type SqsDestination struct {
	QueueArn string `json:"queueArn"`
}

type _SqsDestination SqsDestination

// NewSqsDestination instantiates a new SqsDestination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSqsDestination(queueArn string) *SqsDestination {
	this := SqsDestination{}
	this.QueueArn = queueArn
	return &this
}

// NewSqsDestinationWithDefaults instantiates a new SqsDestination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSqsDestinationWithDefaults() *SqsDestination {
	this := SqsDestination{}
	return &this
}

// GetQueueArn returns the QueueArn field value
func (o *SqsDestination) GetQueueArn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.QueueArn
}

// GetQueueArnOk returns a tuple with the QueueArn field value
// and a boolean to check if the value has been set.
func (o *SqsDestination) GetQueueArnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueueArn, true
}

// SetQueueArn sets field value
func (o *SqsDestination) SetQueueArn(v string) {
	o.QueueArn = v
}

func (o SqsDestination) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["queueArn"] = o.QueueArn
	return toSerialize, nil
}

type NullableSqsDestination struct {
	value *SqsDestination
	isSet bool
}

func (v NullableSqsDestination) Get() *SqsDestination {
	return v.value
}

func (v *NullableSqsDestination) Set(val *SqsDestination) {
	v.value = val
	v.isSet = true
}

func (v NullableSqsDestination) IsSet() bool {
	return v.isSet
}

func (v *NullableSqsDestination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSqsDestination(val *SqsDestination) *NullableSqsDestination {
	return &NullableSqsDestination{value: val, isSet: true}
}

func (v NullableSqsDestination) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSqsDestination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
