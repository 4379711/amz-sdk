package finances_20240619

import (
	"github.com/bytedance/sonic"
)

// checks if the ItemRelatedIdentifier type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ItemRelatedIdentifier{}

// ItemRelatedIdentifier Related business identifiers of the item.
type ItemRelatedIdentifier struct {
	// Enumerated set of related item identifier names for the item.
	ItemRelatedIdentifierName *string `json:"itemRelatedIdentifierName,omitempty"`
	// Corresponding value of ItemRelatedIdentifierName
	ItemRelatedIdentifierValue *string `json:"itemRelatedIdentifierValue,omitempty"`
}

// NewItemRelatedIdentifier instantiates a new ItemRelatedIdentifier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemRelatedIdentifier() *ItemRelatedIdentifier {
	this := ItemRelatedIdentifier{}
	return &this
}

// NewItemRelatedIdentifierWithDefaults instantiates a new ItemRelatedIdentifier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemRelatedIdentifierWithDefaults() *ItemRelatedIdentifier {
	this := ItemRelatedIdentifier{}
	return &this
}

// GetItemRelatedIdentifierName returns the ItemRelatedIdentifierName field value if set, zero value otherwise.
func (o *ItemRelatedIdentifier) GetItemRelatedIdentifierName() string {
	if o == nil || IsNil(o.ItemRelatedIdentifierName) {
		var ret string
		return ret
	}
	return *o.ItemRelatedIdentifierName
}

// GetItemRelatedIdentifierNameOk returns a tuple with the ItemRelatedIdentifierName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemRelatedIdentifier) GetItemRelatedIdentifierNameOk() (*string, bool) {
	if o == nil || IsNil(o.ItemRelatedIdentifierName) {
		return nil, false
	}
	return o.ItemRelatedIdentifierName, true
}

// HasItemRelatedIdentifierName returns a boolean if a field has been set.
func (o *ItemRelatedIdentifier) HasItemRelatedIdentifierName() bool {
	if o != nil && !IsNil(o.ItemRelatedIdentifierName) {
		return true
	}

	return false
}

// SetItemRelatedIdentifierName gets a reference to the given string and assigns it to the ItemRelatedIdentifierName field.
func (o *ItemRelatedIdentifier) SetItemRelatedIdentifierName(v string) {
	o.ItemRelatedIdentifierName = &v
}

// GetItemRelatedIdentifierValue returns the ItemRelatedIdentifierValue field value if set, zero value otherwise.
func (o *ItemRelatedIdentifier) GetItemRelatedIdentifierValue() string {
	if o == nil || IsNil(o.ItemRelatedIdentifierValue) {
		var ret string
		return ret
	}
	return *o.ItemRelatedIdentifierValue
}

// GetItemRelatedIdentifierValueOk returns a tuple with the ItemRelatedIdentifierValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemRelatedIdentifier) GetItemRelatedIdentifierValueOk() (*string, bool) {
	if o == nil || IsNil(o.ItemRelatedIdentifierValue) {
		return nil, false
	}
	return o.ItemRelatedIdentifierValue, true
}

// HasItemRelatedIdentifierValue returns a boolean if a field has been set.
func (o *ItemRelatedIdentifier) HasItemRelatedIdentifierValue() bool {
	if o != nil && !IsNil(o.ItemRelatedIdentifierValue) {
		return true
	}

	return false
}

// SetItemRelatedIdentifierValue gets a reference to the given string and assigns it to the ItemRelatedIdentifierValue field.
func (o *ItemRelatedIdentifier) SetItemRelatedIdentifierValue(v string) {
	o.ItemRelatedIdentifierValue = &v
}

func (o ItemRelatedIdentifier) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ItemRelatedIdentifierName) {
		toSerialize["itemRelatedIdentifierName"] = o.ItemRelatedIdentifierName
	}
	if !IsNil(o.ItemRelatedIdentifierValue) {
		toSerialize["itemRelatedIdentifierValue"] = o.ItemRelatedIdentifierValue
	}
	return toSerialize, nil
}

type NullableItemRelatedIdentifier struct {
	value *ItemRelatedIdentifier
	isSet bool
}

func (v NullableItemRelatedIdentifier) Get() *ItemRelatedIdentifier {
	return v.value
}

func (v *NullableItemRelatedIdentifier) Set(val *ItemRelatedIdentifier) {
	v.value = val
	v.isSet = true
}

func (v NullableItemRelatedIdentifier) IsSet() bool {
	return v.isSet
}

func (v *NullableItemRelatedIdentifier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemRelatedIdentifier(val *ItemRelatedIdentifier) *NullableItemRelatedIdentifier {
	return &NullableItemRelatedIdentifier{value: val, isSet: true}
}

func (v NullableItemRelatedIdentifier) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableItemRelatedIdentifier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
