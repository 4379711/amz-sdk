package product_eligibility

import (
	"github.com/bytedance/sonic"
)

// checks if the ReasonItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReasonItem{}

// ReasonItem struct for ReasonItem
type ReasonItem struct {
	Code  *ReasonCode      `json:"code,omitempty"`
	Level *IneligibleLevel `json:"level,omitempty"`
	// Message explaining what the status means. Example: Payment preference not found for associated billing account. Please add a new payment method
	Description *string `json:"description,omitempty"`
}

// NewReasonItem instantiates a new ReasonItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReasonItem() *ReasonItem {
	this := ReasonItem{}
	return &this
}

// NewReasonItemWithDefaults instantiates a new ReasonItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReasonItemWithDefaults() *ReasonItem {
	this := ReasonItem{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ReasonItem) GetCode() ReasonCode {
	if o == nil || IsNil(o.Code) {
		var ret ReasonCode
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReasonItem) GetCodeOk() (*ReasonCode, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ReasonItem) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given ReasonCode and assigns it to the Code field.
func (o *ReasonItem) SetCode(v ReasonCode) {
	o.Code = &v
}

// GetLevel returns the Level field value if set, zero value otherwise.
func (o *ReasonItem) GetLevel() IneligibleLevel {
	if o == nil || IsNil(o.Level) {
		var ret IneligibleLevel
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReasonItem) GetLevelOk() (*IneligibleLevel, bool) {
	if o == nil || IsNil(o.Level) {
		return nil, false
	}
	return o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *ReasonItem) HasLevel() bool {
	if o != nil && !IsNil(o.Level) {
		return true
	}

	return false
}

// SetLevel gets a reference to the given IneligibleLevel and assigns it to the Level field.
func (o *ReasonItem) SetLevel(v IneligibleLevel) {
	o.Level = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ReasonItem) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReasonItem) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ReasonItem) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ReasonItem) SetDescription(v string) {
	o.Description = &v
}

func (o ReasonItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Level) {
		toSerialize["level"] = o.Level
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullableReasonItem struct {
	value *ReasonItem
	isSet bool
}

func (v NullableReasonItem) Get() *ReasonItem {
	return v.value
}

func (v *NullableReasonItem) Set(val *ReasonItem) {
	v.value = val
	v.isSet = true
}

func (v NullableReasonItem) IsSet() bool {
	return v.isSet
}

func (v *NullableReasonItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReasonItem(val *ReasonItem) *NullableReasonItem {
	return &NullableReasonItem{value: val, isSet: true}
}

func (v NullableReasonItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableReasonItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
