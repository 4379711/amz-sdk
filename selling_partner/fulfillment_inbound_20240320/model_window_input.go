package fulfillment_inbound_20240320

import (
	"time"

	"github.com/bytedance/sonic"
)

// checks if the WindowInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WindowInput{}

// WindowInput Contains only a starting DateTime.
type WindowInput struct {
	// The start date of the window. In [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) datetime format with minute precision. Supports patterns `yyyy-MM-ddTHH:mmZ`, `yyyy-MM-ddTHH:mm:ssZ`, or `yyyy-MM-ddTHH:mm:ss.sssZ`. Note that non-zero second and millisecond components are removed.
	Start time.Time `json:"start"`
}

type _WindowInput WindowInput

// NewWindowInput instantiates a new WindowInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWindowInput(start time.Time) *WindowInput {
	this := WindowInput{}
	this.Start = start
	return &this
}

// NewWindowInputWithDefaults instantiates a new WindowInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWindowInputWithDefaults() *WindowInput {
	this := WindowInput{}
	return &this
}

// GetStart returns the Start field value
func (o *WindowInput) GetStart() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Start
}

// GetStartOk returns a tuple with the Start field value
// and a boolean to check if the value has been set.
func (o *WindowInput) GetStartOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Start, true
}

// SetStart sets field value
func (o *WindowInput) SetStart(v time.Time) {
	o.Start = v
}

func (o WindowInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["start"] = o.Start
	return toSerialize, nil
}

type NullableWindowInput struct {
	value *WindowInput
	isSet bool
}

func (v NullableWindowInput) Get() *WindowInput {
	return v.value
}

func (v *NullableWindowInput) Set(val *WindowInput) {
	v.value = val
	v.isSet = true
}

func (v NullableWindowInput) IsSet() bool {
	return v.isSet
}

func (v *NullableWindowInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWindowInput(val *WindowInput) *NullableWindowInput {
	return &NullableWindowInput{value: val, isSet: true}
}

func (v NullableWindowInput) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableWindowInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
