package merchant_fulfillment_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the AdditionalInputs type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdditionalInputs{}

// AdditionalInputs Maps the additional seller input to the definition. The key to the map is the field name.
type AdditionalInputs struct {
	// The field name.
	AdditionalInputFieldName *string                `json:"AdditionalInputFieldName,omitempty"`
	SellerInputDefinition    *SellerInputDefinition `json:"SellerInputDefinition,omitempty"`
}

// NewAdditionalInputs instantiates a new AdditionalInputs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdditionalInputs() *AdditionalInputs {
	this := AdditionalInputs{}
	return &this
}

// NewAdditionalInputsWithDefaults instantiates a new AdditionalInputs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdditionalInputsWithDefaults() *AdditionalInputs {
	this := AdditionalInputs{}
	return &this
}

// GetAdditionalInputFieldName returns the AdditionalInputFieldName field value if set, zero value otherwise.
func (o *AdditionalInputs) GetAdditionalInputFieldName() string {
	if o == nil || IsNil(o.AdditionalInputFieldName) {
		var ret string
		return ret
	}
	return *o.AdditionalInputFieldName
}

// GetAdditionalInputFieldNameOk returns a tuple with the AdditionalInputFieldName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalInputs) GetAdditionalInputFieldNameOk() (*string, bool) {
	if o == nil || IsNil(o.AdditionalInputFieldName) {
		return nil, false
	}
	return o.AdditionalInputFieldName, true
}

// HasAdditionalInputFieldName returns a boolean if a field has been set.
func (o *AdditionalInputs) HasAdditionalInputFieldName() bool {
	if o != nil && !IsNil(o.AdditionalInputFieldName) {
		return true
	}

	return false
}

// SetAdditionalInputFieldName gets a reference to the given string and assigns it to the AdditionalInputFieldName field.
func (o *AdditionalInputs) SetAdditionalInputFieldName(v string) {
	o.AdditionalInputFieldName = &v
}

// GetSellerInputDefinition returns the SellerInputDefinition field value if set, zero value otherwise.
func (o *AdditionalInputs) GetSellerInputDefinition() SellerInputDefinition {
	if o == nil || IsNil(o.SellerInputDefinition) {
		var ret SellerInputDefinition
		return ret
	}
	return *o.SellerInputDefinition
}

// GetSellerInputDefinitionOk returns a tuple with the SellerInputDefinition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalInputs) GetSellerInputDefinitionOk() (*SellerInputDefinition, bool) {
	if o == nil || IsNil(o.SellerInputDefinition) {
		return nil, false
	}
	return o.SellerInputDefinition, true
}

// HasSellerInputDefinition returns a boolean if a field has been set.
func (o *AdditionalInputs) HasSellerInputDefinition() bool {
	if o != nil && !IsNil(o.SellerInputDefinition) {
		return true
	}

	return false
}

// SetSellerInputDefinition gets a reference to the given SellerInputDefinition and assigns it to the SellerInputDefinition field.
func (o *AdditionalInputs) SetSellerInputDefinition(v SellerInputDefinition) {
	o.SellerInputDefinition = &v
}

func (o AdditionalInputs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdditionalInputFieldName) {
		toSerialize["AdditionalInputFieldName"] = o.AdditionalInputFieldName
	}
	if !IsNil(o.SellerInputDefinition) {
		toSerialize["SellerInputDefinition"] = o.SellerInputDefinition
	}
	return toSerialize, nil
}

type NullableAdditionalInputs struct {
	value *AdditionalInputs
	isSet bool
}

func (v NullableAdditionalInputs) Get() *AdditionalInputs {
	return v.value
}

func (v *NullableAdditionalInputs) Set(val *AdditionalInputs) {
	v.value = val
	v.isSet = true
}

func (v NullableAdditionalInputs) IsSet() bool {
	return v.isSet
}

func (v *NullableAdditionalInputs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdditionalInputs(val *AdditionalInputs) *NullableAdditionalInputs {
	return &NullableAdditionalInputs{value: val, isSet: true}
}

func (v NullableAdditionalInputs) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAdditionalInputs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
