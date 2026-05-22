package notifications

import (
	"github.com/bytedance/sonic"
)

// checks if the SqsResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SqsResource{}

// SqsResource The information required to create an Amazon Simple Queue Service (Amazon SQS) queue destination.
type SqsResource struct {
	// The Amazon Resource Name (ARN) associated with the SQS queue.
	Arn string `json:"arn" validate:"regexp=^arn:aws:sqs:\\\\S+:\\\\S+:\\\\S+"`
}

type _SqsResource SqsResource

// NewSqsResource instantiates a new SqsResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSqsResource(arn string) *SqsResource {
	this := SqsResource{}
	this.Arn = arn
	return &this
}

// NewSqsResourceWithDefaults instantiates a new SqsResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSqsResourceWithDefaults() *SqsResource {
	this := SqsResource{}
	return &this
}

// GetArn returns the Arn field value
func (o *SqsResource) GetArn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Arn
}

// GetArnOk returns a tuple with the Arn field value
// and a boolean to check if the value has been set.
func (o *SqsResource) GetArnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Arn, true
}

// SetArn sets field value
func (o *SqsResource) SetArn(v string) {
	o.Arn = v
}

func (o SqsResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["arn"] = o.Arn
	return toSerialize, nil
}

type NullableSqsResource struct {
	value *SqsResource
	isSet bool
}

func (v NullableSqsResource) Get() *SqsResource {
	return v.value
}

func (v *NullableSqsResource) Set(val *SqsResource) {
	v.value = val
	v.isSet = true
}

func (v NullableSqsResource) IsSet() bool {
	return v.isSet
}

func (v *NullableSqsResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSqsResource(val *SqsResource) *NullableSqsResource {
	return &NullableSqsResource{value: val, isSet: true}
}

func (v NullableSqsResource) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSqsResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
