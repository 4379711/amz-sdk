package product_fees_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the GetMyFeesEstimateResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMyFeesEstimateResult{}

// GetMyFeesEstimateResult Response schema.
type GetMyFeesEstimateResult struct {
	FeesEstimateResult *FeesEstimateResult `json:"FeesEstimateResult,omitempty"`
}

// NewGetMyFeesEstimateResult instantiates a new GetMyFeesEstimateResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMyFeesEstimateResult() *GetMyFeesEstimateResult {
	this := GetMyFeesEstimateResult{}
	return &this
}

// NewGetMyFeesEstimateResultWithDefaults instantiates a new GetMyFeesEstimateResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMyFeesEstimateResultWithDefaults() *GetMyFeesEstimateResult {
	this := GetMyFeesEstimateResult{}
	return &this
}

// GetFeesEstimateResult returns the FeesEstimateResult field value if set, zero value otherwise.
func (o *GetMyFeesEstimateResult) GetFeesEstimateResult() FeesEstimateResult {
	if o == nil || IsNil(o.FeesEstimateResult) {
		var ret FeesEstimateResult
		return ret
	}
	return *o.FeesEstimateResult
}

// GetFeesEstimateResultOk returns a tuple with the FeesEstimateResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMyFeesEstimateResult) GetFeesEstimateResultOk() (*FeesEstimateResult, bool) {
	if o == nil || IsNil(o.FeesEstimateResult) {
		return nil, false
	}
	return o.FeesEstimateResult, true
}

// HasFeesEstimateResult returns a boolean if a field has been set.
func (o *GetMyFeesEstimateResult) HasFeesEstimateResult() bool {
	if o != nil && !IsNil(o.FeesEstimateResult) {
		return true
	}

	return false
}

// SetFeesEstimateResult gets a reference to the given FeesEstimateResult and assigns it to the FeesEstimateResult field.
func (o *GetMyFeesEstimateResult) SetFeesEstimateResult(v FeesEstimateResult) {
	o.FeesEstimateResult = &v
}

func (o GetMyFeesEstimateResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FeesEstimateResult) {
		toSerialize["FeesEstimateResult"] = o.FeesEstimateResult
	}
	return toSerialize, nil
}

type NullableGetMyFeesEstimateResult struct {
	value *GetMyFeesEstimateResult
	isSet bool
}

func (v NullableGetMyFeesEstimateResult) Get() *GetMyFeesEstimateResult {
	return v.value
}

func (v *NullableGetMyFeesEstimateResult) Set(val *GetMyFeesEstimateResult) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMyFeesEstimateResult) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMyFeesEstimateResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMyFeesEstimateResult(val *GetMyFeesEstimateResult) *NullableGetMyFeesEstimateResult {
	return &NullableGetMyFeesEstimateResult{value: val, isSet: true}
}

func (v NullableGetMyFeesEstimateResult) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetMyFeesEstimateResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
