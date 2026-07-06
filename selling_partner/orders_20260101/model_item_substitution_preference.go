package orders_20260101

import (
	"github.com/bytedance/sonic"
)

// checks if the ItemSubstitutionPreference type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ItemSubstitutionPreference{}

// ItemSubstitutionPreference Substitution preference for an order item when it becomes unavailable during fulfillment.
type ItemSubstitutionPreference struct {
	// Source and nature of the substitution preferences for this item.
	SubstitutionType string `json:"substitutionType"`
	// List of alternative products that can be substituted for the original item if it becomes unavailable.
	SubstitutionOptions []ItemSubstitutionOption `json:"substitutionOptions,omitempty"`
}

// NewItemSubstitutionPreference instantiates a new ItemSubstitutionPreference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemSubstitutionPreference(substitutionType string) *ItemSubstitutionPreference {
	this := ItemSubstitutionPreference{}
	this.SubstitutionType = substitutionType
	return &this
}

// NewItemSubstitutionPreferenceWithDefaults instantiates a new ItemSubstitutionPreference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemSubstitutionPreferenceWithDefaults() *ItemSubstitutionPreference {
	this := ItemSubstitutionPreference{}
	return &this
}

// GetSubstitutionType returns the SubstitutionType field value
func (o *ItemSubstitutionPreference) GetSubstitutionType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubstitutionType
}

// GetSubstitutionTypeOk returns a tuple with the SubstitutionType field value
// and a boolean to check if the value has been set.
func (o *ItemSubstitutionPreference) GetSubstitutionTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubstitutionType, true
}

// SetSubstitutionType sets field value
func (o *ItemSubstitutionPreference) SetSubstitutionType(v string) {
	o.SubstitutionType = v
}

// GetSubstitutionOptions returns the SubstitutionOptions field value if set, zero value otherwise.
func (o *ItemSubstitutionPreference) GetSubstitutionOptions() []ItemSubstitutionOption {
	if o == nil || IsNil(o.SubstitutionOptions) {
		var ret []ItemSubstitutionOption
		return ret
	}
	return o.SubstitutionOptions
}

// GetSubstitutionOptionsOk returns a tuple with the SubstitutionOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemSubstitutionPreference) GetSubstitutionOptionsOk() ([]ItemSubstitutionOption, bool) {
	if o == nil || IsNil(o.SubstitutionOptions) {
		return nil, false
	}
	return o.SubstitutionOptions, true
}

// HasSubstitutionOptions returns a boolean if a field has been set.
func (o *ItemSubstitutionPreference) HasSubstitutionOptions() bool {
	if o != nil && !IsNil(o.SubstitutionOptions) {
		return true
	}

	return false
}

// SetSubstitutionOptions gets a reference to the given []ItemSubstitutionOption and assigns it to the SubstitutionOptions field.
func (o *ItemSubstitutionPreference) SetSubstitutionOptions(v []ItemSubstitutionOption) {
	o.SubstitutionOptions = v
}

func (o ItemSubstitutionPreference) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["substitutionType"] = o.SubstitutionType
	if !IsNil(o.SubstitutionOptions) {
		toSerialize["substitutionOptions"] = o.SubstitutionOptions
	}
	return toSerialize, nil
}

type NullableItemSubstitutionPreference struct {
	value *ItemSubstitutionPreference
	isSet bool
}

func (v NullableItemSubstitutionPreference) Get() *ItemSubstitutionPreference {
	return v.value
}

func (v *NullableItemSubstitutionPreference) Set(val *ItemSubstitutionPreference) {
	v.value = val
	v.isSet = true
}

func (v NullableItemSubstitutionPreference) IsSet() bool {
	return v.isSet
}

func (v *NullableItemSubstitutionPreference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemSubstitutionPreference(val *ItemSubstitutionPreference) *NullableItemSubstitutionPreference {
	return &NullableItemSubstitutionPreference{value: val, isSet: true}
}

func (v NullableItemSubstitutionPreference) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableItemSubstitutionPreference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
