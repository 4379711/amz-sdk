package merchant_fulfillment_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the AdditionalSellerInputs type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdditionalSellerInputs{}

// AdditionalSellerInputs An additional set of seller inputs required to purchase shipping.
type AdditionalSellerInputs struct {
	// The name of the additional input field.
	AdditionalInputFieldName string                `json:"AdditionalInputFieldName"`
	AdditionalSellerInput    AdditionalSellerInput `json:"AdditionalSellerInput"`
}

type _AdditionalSellerInputs AdditionalSellerInputs

// NewAdditionalSellerInputs instantiates a new AdditionalSellerInputs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdditionalSellerInputs(additionalInputFieldName string, additionalSellerInput AdditionalSellerInput) *AdditionalSellerInputs {
	this := AdditionalSellerInputs{}
	this.AdditionalInputFieldName = additionalInputFieldName
	this.AdditionalSellerInput = additionalSellerInput
	return &this
}

// NewAdditionalSellerInputsWithDefaults instantiates a new AdditionalSellerInputs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdditionalSellerInputsWithDefaults() *AdditionalSellerInputs {
	this := AdditionalSellerInputs{}
	return &this
}

// GetAdditionalInputFieldName returns the AdditionalInputFieldName field value
func (o *AdditionalSellerInputs) GetAdditionalInputFieldName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdditionalInputFieldName
}

// GetAdditionalInputFieldNameOk returns a tuple with the AdditionalInputFieldName field value
// and a boolean to check if the value has been set.
func (o *AdditionalSellerInputs) GetAdditionalInputFieldNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdditionalInputFieldName, true
}

// SetAdditionalInputFieldName sets field value
func (o *AdditionalSellerInputs) SetAdditionalInputFieldName(v string) {
	o.AdditionalInputFieldName = v
}

// GetAdditionalSellerInput returns the AdditionalSellerInput field value
func (o *AdditionalSellerInputs) GetAdditionalSellerInput() AdditionalSellerInput {
	if o == nil {
		var ret AdditionalSellerInput
		return ret
	}

	return o.AdditionalSellerInput
}

// GetAdditionalSellerInputOk returns a tuple with the AdditionalSellerInput field value
// and a boolean to check if the value has been set.
func (o *AdditionalSellerInputs) GetAdditionalSellerInputOk() (*AdditionalSellerInput, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdditionalSellerInput, true
}

// SetAdditionalSellerInput sets field value
func (o *AdditionalSellerInputs) SetAdditionalSellerInput(v AdditionalSellerInput) {
	o.AdditionalSellerInput = v
}

func (o AdditionalSellerInputs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["AdditionalInputFieldName"] = o.AdditionalInputFieldName
	toSerialize["AdditionalSellerInput"] = o.AdditionalSellerInput
	return toSerialize, nil
}

type NullableAdditionalSellerInputs struct {
	value *AdditionalSellerInputs
	isSet bool
}

func (v NullableAdditionalSellerInputs) Get() *AdditionalSellerInputs {
	return v.value
}

func (v *NullableAdditionalSellerInputs) Set(val *AdditionalSellerInputs) {
	v.value = val
	v.isSet = true
}

func (v NullableAdditionalSellerInputs) IsSet() bool {
	return v.isSet
}

func (v *NullableAdditionalSellerInputs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdditionalSellerInputs(val *AdditionalSellerInputs) *NullableAdditionalSellerInputs {
	return &NullableAdditionalSellerInputs{value: val, isSet: true}
}

func (v NullableAdditionalSellerInputs) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAdditionalSellerInputs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
