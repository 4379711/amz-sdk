package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the MskuPrepDetailInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MskuPrepDetailInput{}

// MskuPrepDetailInput An MSKU and its related prep details.
type MskuPrepDetailInput struct {
	// The merchant SKU, a merchant-supplied identifier for a specific SKU.
	Msku         string       `json:"msku"`
	PrepCategory PrepCategory `json:"prepCategory"`
	// A list of preparation types associated with a preparation category.
	PrepTypes []PrepType `json:"prepTypes"`
}

type _MskuPrepDetailInput MskuPrepDetailInput

// NewMskuPrepDetailInput instantiates a new MskuPrepDetailInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMskuPrepDetailInput(msku string, prepCategory PrepCategory, prepTypes []PrepType) *MskuPrepDetailInput {
	this := MskuPrepDetailInput{}
	this.Msku = msku
	this.PrepCategory = prepCategory
	this.PrepTypes = prepTypes
	return &this
}

// NewMskuPrepDetailInputWithDefaults instantiates a new MskuPrepDetailInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMskuPrepDetailInputWithDefaults() *MskuPrepDetailInput {
	this := MskuPrepDetailInput{}
	return &this
}

// GetMsku returns the Msku field value
func (o *MskuPrepDetailInput) GetMsku() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Msku
}

// GetMskuOk returns a tuple with the Msku field value
// and a boolean to check if the value has been set.
func (o *MskuPrepDetailInput) GetMskuOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Msku, true
}

// SetMsku sets field value
func (o *MskuPrepDetailInput) SetMsku(v string) {
	o.Msku = v
}

// GetPrepCategory returns the PrepCategory field value
func (o *MskuPrepDetailInput) GetPrepCategory() PrepCategory {
	if o == nil {
		var ret PrepCategory
		return ret
	}

	return o.PrepCategory
}

// GetPrepCategoryOk returns a tuple with the PrepCategory field value
// and a boolean to check if the value has been set.
func (o *MskuPrepDetailInput) GetPrepCategoryOk() (*PrepCategory, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrepCategory, true
}

// SetPrepCategory sets field value
func (o *MskuPrepDetailInput) SetPrepCategory(v PrepCategory) {
	o.PrepCategory = v
}

// GetPrepTypes returns the PrepTypes field value
func (o *MskuPrepDetailInput) GetPrepTypes() []PrepType {
	if o == nil {
		var ret []PrepType
		return ret
	}

	return o.PrepTypes
}

// GetPrepTypesOk returns a tuple with the PrepTypes field value
// and a boolean to check if the value has been set.
func (o *MskuPrepDetailInput) GetPrepTypesOk() ([]PrepType, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrepTypes, true
}

// SetPrepTypes sets field value
func (o *MskuPrepDetailInput) SetPrepTypes(v []PrepType) {
	o.PrepTypes = v
}

func (o MskuPrepDetailInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["msku"] = o.Msku
	toSerialize["prepCategory"] = o.PrepCategory
	toSerialize["prepTypes"] = o.PrepTypes
	return toSerialize, nil
}

type NullableMskuPrepDetailInput struct {
	value *MskuPrepDetailInput
	isSet bool
}

func (v NullableMskuPrepDetailInput) Get() *MskuPrepDetailInput {
	return v.value
}

func (v *NullableMskuPrepDetailInput) Set(val *MskuPrepDetailInput) {
	v.value = val
	v.isSet = true
}

func (v NullableMskuPrepDetailInput) IsSet() bool {
	return v.isSet
}

func (v *NullableMskuPrepDetailInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMskuPrepDetailInput(val *MskuPrepDetailInput) *NullableMskuPrepDetailInput {
	return &NullableMskuPrepDetailInput{value: val, isSet: true}
}

func (v NullableMskuPrepDetailInput) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableMskuPrepDetailInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
