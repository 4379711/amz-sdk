package orders_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the RejectionReason type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RejectionReason{}

// RejectionReason The reason for rejecting the order's regulated information. This is only present if the order is rejected.
type RejectionReason struct {
	// The unique identifier for the rejection reason.
	RejectionReasonId string `json:"RejectionReasonId"`
	// The description of this rejection reason.
	RejectionReasonDescription string `json:"RejectionReasonDescription"`
}

type _RejectionReason RejectionReason

// NewRejectionReason instantiates a new RejectionReason object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRejectionReason(rejectionReasonId string, rejectionReasonDescription string) *RejectionReason {
	this := RejectionReason{}
	this.RejectionReasonId = rejectionReasonId
	this.RejectionReasonDescription = rejectionReasonDescription
	return &this
}

// NewRejectionReasonWithDefaults instantiates a new RejectionReason object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRejectionReasonWithDefaults() *RejectionReason {
	this := RejectionReason{}
	return &this
}

// GetRejectionReasonId returns the RejectionReasonId field value
func (o *RejectionReason) GetRejectionReasonId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RejectionReasonId
}

// GetRejectionReasonIdOk returns a tuple with the RejectionReasonId field value
// and a boolean to check if the value has been set.
func (o *RejectionReason) GetRejectionReasonIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RejectionReasonId, true
}

// SetRejectionReasonId sets field value
func (o *RejectionReason) SetRejectionReasonId(v string) {
	o.RejectionReasonId = v
}

// GetRejectionReasonDescription returns the RejectionReasonDescription field value
func (o *RejectionReason) GetRejectionReasonDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RejectionReasonDescription
}

// GetRejectionReasonDescriptionOk returns a tuple with the RejectionReasonDescription field value
// and a boolean to check if the value has been set.
func (o *RejectionReason) GetRejectionReasonDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RejectionReasonDescription, true
}

// SetRejectionReasonDescription sets field value
func (o *RejectionReason) SetRejectionReasonDescription(v string) {
	o.RejectionReasonDescription = v
}

func (o RejectionReason) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["RejectionReasonId"] = o.RejectionReasonId
	toSerialize["RejectionReasonDescription"] = o.RejectionReasonDescription
	return toSerialize, nil
}

type NullableRejectionReason struct {
	value *RejectionReason
	isSet bool
}

func (v NullableRejectionReason) Get() *RejectionReason {
	return v.value
}

func (v *NullableRejectionReason) Set(val *RejectionReason) {
	v.value = val
	v.isSet = true
}

func (v NullableRejectionReason) IsSet() bool {
	return v.isSet
}

func (v *NullableRejectionReason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRejectionReason(val *RejectionReason) *NullableRejectionReason {
	return &NullableRejectionReason{value: val, isSet: true}
}

func (v NullableRejectionReason) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableRejectionReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
