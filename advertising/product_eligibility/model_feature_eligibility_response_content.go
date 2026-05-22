package product_eligibility

import (
	"github.com/bytedance/sonic"
)

// checks if the FeatureEligibilityResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FeatureEligibilityResponseContent{}

// FeatureEligibilityResponseContent A list of feature eligibility responses for an advertiser, split by success and error
type FeatureEligibilityResponseContent struct {
	Success []FeatureEligibilityItem  `json:"success,omitempty"`
	Error   []FeatureEligibilityError `json:"error,omitempty"`
}

// NewFeatureEligibilityResponseContent instantiates a new FeatureEligibilityResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeatureEligibilityResponseContent() *FeatureEligibilityResponseContent {
	this := FeatureEligibilityResponseContent{}
	return &this
}

// NewFeatureEligibilityResponseContentWithDefaults instantiates a new FeatureEligibilityResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeatureEligibilityResponseContentWithDefaults() *FeatureEligibilityResponseContent {
	this := FeatureEligibilityResponseContent{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *FeatureEligibilityResponseContent) GetSuccess() []FeatureEligibilityItem {
	if o == nil || IsNil(o.Success) {
		var ret []FeatureEligibilityItem
		return ret
	}
	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureEligibilityResponseContent) GetSuccessOk() ([]FeatureEligibilityItem, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *FeatureEligibilityResponseContent) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given []FeatureEligibilityItem and assigns it to the Success field.
func (o *FeatureEligibilityResponseContent) SetSuccess(v []FeatureEligibilityItem) {
	o.Success = v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *FeatureEligibilityResponseContent) GetError() []FeatureEligibilityError {
	if o == nil || IsNil(o.Error) {
		var ret []FeatureEligibilityError
		return ret
	}
	return o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureEligibilityResponseContent) GetErrorOk() ([]FeatureEligibilityError, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *FeatureEligibilityResponseContent) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given []FeatureEligibilityError and assigns it to the Error field.
func (o *FeatureEligibilityResponseContent) SetError(v []FeatureEligibilityError) {
	o.Error = v
}

func (o FeatureEligibilityResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

type NullableFeatureEligibilityResponseContent struct {
	value *FeatureEligibilityResponseContent
	isSet bool
}

func (v NullableFeatureEligibilityResponseContent) Get() *FeatureEligibilityResponseContent {
	return v.value
}

func (v *NullableFeatureEligibilityResponseContent) Set(val *FeatureEligibilityResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableFeatureEligibilityResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableFeatureEligibilityResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeatureEligibilityResponseContent(val *FeatureEligibilityResponseContent) *NullableFeatureEligibilityResponseContent {
	return &NullableFeatureEligibilityResponseContent{value: val, isSet: true}
}

func (v NullableFeatureEligibilityResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableFeatureEligibilityResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
