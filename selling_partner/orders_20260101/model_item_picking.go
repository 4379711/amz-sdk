package orders_20260101

import (
	"github.com/bytedance/sonic"
)

// checks if the ItemPicking type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ItemPicking{}

// ItemPicking Information related to the warehouse picking process for an order item.
type ItemPicking struct {
	SubstitutionPreference *ItemSubstitutionPreference `json:"substitutionPreference,omitempty"`
}

// NewItemPicking instantiates a new ItemPicking object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemPicking() *ItemPicking {
	this := ItemPicking{}
	return &this
}

// NewItemPickingWithDefaults instantiates a new ItemPicking object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemPickingWithDefaults() *ItemPicking {
	this := ItemPicking{}
	return &this
}

// GetSubstitutionPreference returns the SubstitutionPreference field value if set, zero value otherwise.
func (o *ItemPicking) GetSubstitutionPreference() ItemSubstitutionPreference {
	if o == nil || IsNil(o.SubstitutionPreference) {
		var ret ItemSubstitutionPreference
		return ret
	}
	return *o.SubstitutionPreference
}

// GetSubstitutionPreferenceOk returns a tuple with the SubstitutionPreference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemPicking) GetSubstitutionPreferenceOk() (*ItemSubstitutionPreference, bool) {
	if o == nil || IsNil(o.SubstitutionPreference) {
		return nil, false
	}
	return o.SubstitutionPreference, true
}

// HasSubstitutionPreference returns a boolean if a field has been set.
func (o *ItemPicking) HasSubstitutionPreference() bool {
	if o != nil && !IsNil(o.SubstitutionPreference) {
		return true
	}

	return false
}

// SetSubstitutionPreference gets a reference to the given ItemSubstitutionPreference and assigns it to the SubstitutionPreference field.
func (o *ItemPicking) SetSubstitutionPreference(v ItemSubstitutionPreference) {
	o.SubstitutionPreference = &v
}

func (o ItemPicking) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SubstitutionPreference) {
		toSerialize["substitutionPreference"] = o.SubstitutionPreference
	}
	return toSerialize, nil
}

type NullableItemPicking struct {
	value *ItemPicking
	isSet bool
}

func (v NullableItemPicking) Get() *ItemPicking {
	return v.value
}

func (v *NullableItemPicking) Set(val *ItemPicking) {
	v.value = val
	v.isSet = true
}

func (v NullableItemPicking) IsSet() bool {
	return v.isSet
}

func (v *NullableItemPicking) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemPicking(val *ItemPicking) *NullableItemPicking {
	return &NullableItemPicking{value: val, isSet: true}
}

func (v NullableItemPicking) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableItemPicking) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
