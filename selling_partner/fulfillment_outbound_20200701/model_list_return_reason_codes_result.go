package fulfillment_outbound_20200701

import (
	"github.com/bytedance/sonic"
)

// checks if the ListReturnReasonCodesResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListReturnReasonCodesResult{}

// ListReturnReasonCodesResult The request for the listReturnReasonCodes operation.
type ListReturnReasonCodesResult struct {
	// An array of return reason code details.
	ReasonCodeDetails []ReasonCodeDetails `json:"reasonCodeDetails,omitempty"`
}

// NewListReturnReasonCodesResult instantiates a new ListReturnReasonCodesResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListReturnReasonCodesResult() *ListReturnReasonCodesResult {
	this := ListReturnReasonCodesResult{}
	return &this
}

// NewListReturnReasonCodesResultWithDefaults instantiates a new ListReturnReasonCodesResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListReturnReasonCodesResultWithDefaults() *ListReturnReasonCodesResult {
	this := ListReturnReasonCodesResult{}
	return &this
}

// GetReasonCodeDetails returns the ReasonCodeDetails field value if set, zero value otherwise.
func (o *ListReturnReasonCodesResult) GetReasonCodeDetails() []ReasonCodeDetails {
	if o == nil || IsNil(o.ReasonCodeDetails) {
		var ret []ReasonCodeDetails
		return ret
	}
	return o.ReasonCodeDetails
}

// GetReasonCodeDetailsOk returns a tuple with the ReasonCodeDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListReturnReasonCodesResult) GetReasonCodeDetailsOk() ([]ReasonCodeDetails, bool) {
	if o == nil || IsNil(o.ReasonCodeDetails) {
		return nil, false
	}
	return o.ReasonCodeDetails, true
}

// HasReasonCodeDetails returns a boolean if a field has been set.
func (o *ListReturnReasonCodesResult) HasReasonCodeDetails() bool {
	if o != nil && !IsNil(o.ReasonCodeDetails) {
		return true
	}

	return false
}

// SetReasonCodeDetails gets a reference to the given []ReasonCodeDetails and assigns it to the ReasonCodeDetails field.
func (o *ListReturnReasonCodesResult) SetReasonCodeDetails(v []ReasonCodeDetails) {
	o.ReasonCodeDetails = v
}

func (o ListReturnReasonCodesResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ReasonCodeDetails) {
		toSerialize["reasonCodeDetails"] = o.ReasonCodeDetails
	}
	return toSerialize, nil
}

type NullableListReturnReasonCodesResult struct {
	value *ListReturnReasonCodesResult
	isSet bool
}

func (v NullableListReturnReasonCodesResult) Get() *ListReturnReasonCodesResult {
	return v.value
}

func (v *NullableListReturnReasonCodesResult) Set(val *ListReturnReasonCodesResult) {
	v.value = val
	v.isSet = true
}

func (v NullableListReturnReasonCodesResult) IsSet() bool {
	return v.isSet
}

func (v *NullableListReturnReasonCodesResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListReturnReasonCodesResult(val *ListReturnReasonCodesResult) *NullableListReturnReasonCodesResult {
	return &NullableListReturnReasonCodesResult{value: val, isSet: true}
}

func (v NullableListReturnReasonCodesResult) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableListReturnReasonCodesResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
