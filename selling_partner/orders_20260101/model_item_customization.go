package orders_20260101

import (
	"github.com/bytedance/sonic"
)

// checks if the ItemCustomization type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ItemCustomization{}

// ItemCustomization Information about any personalization, customization, or special modifications applied to this order item.
type ItemCustomization struct {
	// The URL of the customized data for custom orders from the Amazon Custom program.
	CustomizedUrl *string `json:"customizedUrl,omitempty"`
}

// NewItemCustomization instantiates a new ItemCustomization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemCustomization() *ItemCustomization {
	this := ItemCustomization{}
	return &this
}

// NewItemCustomizationWithDefaults instantiates a new ItemCustomization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemCustomizationWithDefaults() *ItemCustomization {
	this := ItemCustomization{}
	return &this
}

// GetCustomizedUrl returns the CustomizedUrl field value if set, zero value otherwise.
func (o *ItemCustomization) GetCustomizedUrl() string {
	if o == nil || IsNil(o.CustomizedUrl) {
		var ret string
		return ret
	}
	return *o.CustomizedUrl
}

// GetCustomizedUrlOk returns a tuple with the CustomizedUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemCustomization) GetCustomizedUrlOk() (*string, bool) {
	if o == nil || IsNil(o.CustomizedUrl) {
		return nil, false
	}
	return o.CustomizedUrl, true
}

// HasCustomizedUrl returns a boolean if a field has been set.
func (o *ItemCustomization) HasCustomizedUrl() bool {
	if o != nil && !IsNil(o.CustomizedUrl) {
		return true
	}

	return false
}

// SetCustomizedUrl gets a reference to the given string and assigns it to the CustomizedUrl field.
func (o *ItemCustomization) SetCustomizedUrl(v string) {
	o.CustomizedUrl = &v
}

func (o ItemCustomization) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CustomizedUrl) {
		toSerialize["customizedUrl"] = o.CustomizedUrl
	}
	return toSerialize, nil
}

type NullableItemCustomization struct {
	value *ItemCustomization
	isSet bool
}

func (v NullableItemCustomization) Get() *ItemCustomization {
	return v.value
}

func (v *NullableItemCustomization) Set(val *ItemCustomization) {
	v.value = val
	v.isSet = true
}

func (v NullableItemCustomization) IsSet() bool {
	return v.isSet
}

func (v *NullableItemCustomization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemCustomization(val *ItemCustomization) *NullableItemCustomization {
	return &NullableItemCustomization{value: val, isSet: true}
}

func (v NullableItemCustomization) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableItemCustomization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
