package promotions_20251201

import (
	"github.com/bytedance/sonic"
)

// checks if the ClaimCode type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClaimCode{}

// ClaimCode Claim code configuration for accessing the promotion.
type ClaimCode struct {
	// The type of claim code.
	Type string `json:"type"`
	// The claim code (6-12 alphanumeric uppercase). Required when `type` is `GROUP`.
	Value *string `json:"value,omitempty" validate:"regexp=^[0-9A-Z]+$"`
}

// NewClaimCode instantiates a new ClaimCode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClaimCode(type_ string) *ClaimCode {
	this := ClaimCode{}
	this.Type = type_
	return &this
}

// NewClaimCodeWithDefaults instantiates a new ClaimCode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClaimCodeWithDefaults() *ClaimCode {
	this := ClaimCode{}
	return &this
}

// GetType returns the Type field value
func (o *ClaimCode) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ClaimCode) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ClaimCode) SetType(v string) {
	o.Type = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ClaimCode) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClaimCode) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ClaimCode) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ClaimCode) SetValue(v string) {
	o.Value = &v
}

func (o ClaimCode) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableClaimCode struct {
	value *ClaimCode
	isSet bool
}

func (v NullableClaimCode) Get() *ClaimCode {
	return v.value
}

func (v *NullableClaimCode) Set(val *ClaimCode) {
	v.value = val
	v.isSet = true
}

func (v NullableClaimCode) IsSet() bool {
	return v.isSet
}

func (v *NullableClaimCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClaimCode(val *ClaimCode) *NullableClaimCode {
	return &NullableClaimCode{value: val, isSet: true}
}

func (v NullableClaimCode) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableClaimCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
