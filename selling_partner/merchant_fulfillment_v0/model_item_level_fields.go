package merchant_fulfillment_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the ItemLevelFields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ItemLevelFields{}

// ItemLevelFields A list of item level fields.
type ItemLevelFields struct {
	// The Amazon Standard Identification Number (ASIN) of the item.
	Asin string `json:"Asin"`
	// A list of additional inputs.
	AdditionalInputs []AdditionalInputs `json:"AdditionalInputs"`
}

type _ItemLevelFields ItemLevelFields

// NewItemLevelFields instantiates a new ItemLevelFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemLevelFields(asin string, additionalInputs []AdditionalInputs) *ItemLevelFields {
	this := ItemLevelFields{}
	this.Asin = asin
	this.AdditionalInputs = additionalInputs
	return &this
}

// NewItemLevelFieldsWithDefaults instantiates a new ItemLevelFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemLevelFieldsWithDefaults() *ItemLevelFields {
	this := ItemLevelFields{}
	return &this
}

// GetAsin returns the Asin field value
func (o *ItemLevelFields) GetAsin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Asin
}

// GetAsinOk returns a tuple with the Asin field value
// and a boolean to check if the value has been set.
func (o *ItemLevelFields) GetAsinOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Asin, true
}

// SetAsin sets field value
func (o *ItemLevelFields) SetAsin(v string) {
	o.Asin = v
}

// GetAdditionalInputs returns the AdditionalInputs field value
func (o *ItemLevelFields) GetAdditionalInputs() []AdditionalInputs {
	if o == nil {
		var ret []AdditionalInputs
		return ret
	}

	return o.AdditionalInputs
}

// GetAdditionalInputsOk returns a tuple with the AdditionalInputs field value
// and a boolean to check if the value has been set.
func (o *ItemLevelFields) GetAdditionalInputsOk() ([]AdditionalInputs, bool) {
	if o == nil {
		return nil, false
	}
	return o.AdditionalInputs, true
}

// SetAdditionalInputs sets field value
func (o *ItemLevelFields) SetAdditionalInputs(v []AdditionalInputs) {
	o.AdditionalInputs = v
}

func (o ItemLevelFields) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Asin"] = o.Asin
	toSerialize["AdditionalInputs"] = o.AdditionalInputs
	return toSerialize, nil
}

type NullableItemLevelFields struct {
	value *ItemLevelFields
	isSet bool
}

func (v NullableItemLevelFields) Get() *ItemLevelFields {
	return v.value
}

func (v *NullableItemLevelFields) Set(val *ItemLevelFields) {
	v.value = val
	v.isSet = true
}

func (v NullableItemLevelFields) IsSet() bool {
	return v.isSet
}

func (v *NullableItemLevelFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemLevelFields(val *ItemLevelFields) *NullableItemLevelFields {
	return &NullableItemLevelFields{value: val, isSet: true}
}

func (v NullableItemLevelFields) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableItemLevelFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
