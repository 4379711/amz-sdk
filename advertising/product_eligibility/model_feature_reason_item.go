package product_eligibility

import (
	"github.com/bytedance/sonic"
)

// checks if the FeatureReasonItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FeatureReasonItem{}

// FeatureReasonItem struct for FeatureReasonItem
type FeatureReasonItem struct {
	Code *ReasonCode `json:"code,omitempty"`
	// Message explaining what the status means. Example: Payment preference not found for associated billing account. Please add a new payment method
	Description *string `json:"description,omitempty"`
}

// NewFeatureReasonItem instantiates a new FeatureReasonItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeatureReasonItem() *FeatureReasonItem {
	this := FeatureReasonItem{}
	return &this
}

// NewFeatureReasonItemWithDefaults instantiates a new FeatureReasonItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeatureReasonItemWithDefaults() *FeatureReasonItem {
	this := FeatureReasonItem{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *FeatureReasonItem) GetCode() ReasonCode {
	if o == nil || IsNil(o.Code) {
		var ret ReasonCode
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureReasonItem) GetCodeOk() (*ReasonCode, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *FeatureReasonItem) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given ReasonCode and assigns it to the Code field.
func (o *FeatureReasonItem) SetCode(v ReasonCode) {
	o.Code = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *FeatureReasonItem) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureReasonItem) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *FeatureReasonItem) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *FeatureReasonItem) SetDescription(v string) {
	o.Description = &v
}

func (o FeatureReasonItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullableFeatureReasonItem struct {
	value *FeatureReasonItem
	isSet bool
}

func (v NullableFeatureReasonItem) Get() *FeatureReasonItem {
	return v.value
}

func (v *NullableFeatureReasonItem) Set(val *FeatureReasonItem) {
	v.value = val
	v.isSet = true
}

func (v NullableFeatureReasonItem) IsSet() bool {
	return v.isSet
}

func (v *NullableFeatureReasonItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeatureReasonItem(val *FeatureReasonItem) *NullableFeatureReasonItem {
	return &NullableFeatureReasonItem{value: val, isSet: true}
}

func (v NullableFeatureReasonItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableFeatureReasonItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
