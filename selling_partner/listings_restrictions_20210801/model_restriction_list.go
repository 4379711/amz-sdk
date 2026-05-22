package listings_restrictions_20210801

import (
	"github.com/bytedance/sonic"
)

// checks if the RestrictionList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RestrictionList{}

// RestrictionList A list of restrictions for the specified Amazon catalog item.
type RestrictionList struct {
	Restrictions []Restriction `json:"restrictions"`
}

type _RestrictionList RestrictionList

// NewRestrictionList instantiates a new RestrictionList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRestrictionList(restrictions []Restriction) *RestrictionList {
	this := RestrictionList{}
	this.Restrictions = restrictions
	return &this
}

// NewRestrictionListWithDefaults instantiates a new RestrictionList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRestrictionListWithDefaults() *RestrictionList {
	this := RestrictionList{}
	return &this
}

// GetRestrictions returns the Restrictions field value
func (o *RestrictionList) GetRestrictions() []Restriction {
	if o == nil {
		var ret []Restriction
		return ret
	}

	return o.Restrictions
}

// GetRestrictionsOk returns a tuple with the Restrictions field value
// and a boolean to check if the value has been set.
func (o *RestrictionList) GetRestrictionsOk() ([]Restriction, bool) {
	if o == nil {
		return nil, false
	}
	return o.Restrictions, true
}

// SetRestrictions sets field value
func (o *RestrictionList) SetRestrictions(v []Restriction) {
	o.Restrictions = v
}

func (o RestrictionList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["restrictions"] = o.Restrictions
	return toSerialize, nil
}

type NullableRestrictionList struct {
	value *RestrictionList
	isSet bool
}

func (v NullableRestrictionList) Get() *RestrictionList {
	return v.value
}

func (v *NullableRestrictionList) Set(val *RestrictionList) {
	v.value = val
	v.isSet = true
}

func (v NullableRestrictionList) IsSet() bool {
	return v.isSet
}

func (v *NullableRestrictionList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRestrictionList(val *RestrictionList) *NullableRestrictionList {
	return &NullableRestrictionList{value: val, isSet: true}
}

func (v NullableRestrictionList) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableRestrictionList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
