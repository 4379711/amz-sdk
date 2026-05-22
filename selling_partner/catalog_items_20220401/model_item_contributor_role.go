package catalog_items_20220401

import (
	"github.com/bytedance/sonic"
)

// checks if the ItemContributorRole type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ItemContributorRole{}

// ItemContributorRole Role of an individual contributor in the creation of an item, such as author or actor.
type ItemContributorRole struct {
	// Display name of the role in the requested locale, such as `Author` or `Actor`.
	DisplayName *string `json:"displayName,omitempty"`
	// Role value for the Amazon catalog item, such as `author` or `actor`.
	Value string `json:"value"`
}

type _ItemContributorRole ItemContributorRole

// NewItemContributorRole instantiates a new ItemContributorRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemContributorRole(value string) *ItemContributorRole {
	this := ItemContributorRole{}
	this.Value = value
	return &this
}

// NewItemContributorRoleWithDefaults instantiates a new ItemContributorRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemContributorRoleWithDefaults() *ItemContributorRole {
	this := ItemContributorRole{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *ItemContributorRole) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemContributorRole) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ItemContributorRole) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *ItemContributorRole) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetValue returns the Value field value
func (o *ItemContributorRole) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ItemContributorRole) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *ItemContributorRole) SetValue(v string) {
	o.Value = v
}

func (o ItemContributorRole) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

type NullableItemContributorRole struct {
	value *ItemContributorRole
	isSet bool
}

func (v NullableItemContributorRole) Get() *ItemContributorRole {
	return v.value
}

func (v *NullableItemContributorRole) Set(val *ItemContributorRole) {
	v.value = val
	v.isSet = true
}

func (v NullableItemContributorRole) IsSet() bool {
	return v.isSet
}

func (v *NullableItemContributorRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemContributorRole(val *ItemContributorRole) *NullableItemContributorRole {
	return &NullableItemContributorRole{value: val, isSet: true}
}

func (v NullableItemContributorRole) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableItemContributorRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
