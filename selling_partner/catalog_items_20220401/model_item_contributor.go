package catalog_items_20220401

import (
	"github.com/bytedance/sonic"
)

// checks if the ItemContributor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ItemContributor{}

// ItemContributor Individual contributor to the creation of an item, such as an author or actor.
type ItemContributor struct {
	Role ItemContributorRole `json:"role"`
	// Name of the contributor, such as `Jane Austen`.
	Value string `json:"value"`
}

type _ItemContributor ItemContributor

// NewItemContributor instantiates a new ItemContributor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemContributor(role ItemContributorRole, value string) *ItemContributor {
	this := ItemContributor{}
	this.Role = role
	this.Value = value
	return &this
}

// NewItemContributorWithDefaults instantiates a new ItemContributor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemContributorWithDefaults() *ItemContributor {
	this := ItemContributor{}
	return &this
}

// GetRole returns the Role field value
func (o *ItemContributor) GetRole() ItemContributorRole {
	if o == nil {
		var ret ItemContributorRole
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *ItemContributor) GetRoleOk() (*ItemContributorRole, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *ItemContributor) SetRole(v ItemContributorRole) {
	o.Role = v
}

// GetValue returns the Value field value
func (o *ItemContributor) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ItemContributor) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *ItemContributor) SetValue(v string) {
	o.Value = v
}

func (o ItemContributor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["role"] = o.Role
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

type NullableItemContributor struct {
	value *ItemContributor
	isSet bool
}

func (v NullableItemContributor) Get() *ItemContributor {
	return v.value
}

func (v *NullableItemContributor) Set(val *ItemContributor) {
	v.value = val
	v.isSet = true
}

func (v NullableItemContributor) IsSet() bool {
	return v.isSet
}

func (v *NullableItemContributor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemContributor(val *ItemContributor) *NullableItemContributor {
	return &NullableItemContributor{value: val, isSet: true}
}

func (v NullableItemContributor) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableItemContributor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
