package orders_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the SubstitutionPreferences type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubstitutionPreferences{}

// SubstitutionPreferences Substitution preferences for an order item.
type SubstitutionPreferences struct {
	// The type of substitution that these preferences represent.
	SubstitutionType string `json:"SubstitutionType"`
	// A collection of substitution options.
	SubstitutionOptions []SubstitutionOption `json:"SubstitutionOptions,omitempty"`
}

type _SubstitutionPreferences SubstitutionPreferences

// NewSubstitutionPreferences instantiates a new SubstitutionPreferences object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubstitutionPreferences(substitutionType string) *SubstitutionPreferences {
	this := SubstitutionPreferences{}
	this.SubstitutionType = substitutionType
	return &this
}

// NewSubstitutionPreferencesWithDefaults instantiates a new SubstitutionPreferences object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubstitutionPreferencesWithDefaults() *SubstitutionPreferences {
	this := SubstitutionPreferences{}
	return &this
}

// GetSubstitutionType returns the SubstitutionType field value
func (o *SubstitutionPreferences) GetSubstitutionType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubstitutionType
}

// GetSubstitutionTypeOk returns a tuple with the SubstitutionType field value
// and a boolean to check if the value has been set.
func (o *SubstitutionPreferences) GetSubstitutionTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubstitutionType, true
}

// SetSubstitutionType sets field value
func (o *SubstitutionPreferences) SetSubstitutionType(v string) {
	o.SubstitutionType = v
}

// GetSubstitutionOptions returns the SubstitutionOptions field value if set, zero value otherwise.
func (o *SubstitutionPreferences) GetSubstitutionOptions() []SubstitutionOption {
	if o == nil || IsNil(o.SubstitutionOptions) {
		var ret []SubstitutionOption
		return ret
	}
	return o.SubstitutionOptions
}

// GetSubstitutionOptionsOk returns a tuple with the SubstitutionOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubstitutionPreferences) GetSubstitutionOptionsOk() ([]SubstitutionOption, bool) {
	if o == nil || IsNil(o.SubstitutionOptions) {
		return nil, false
	}
	return o.SubstitutionOptions, true
}

// HasSubstitutionOptions returns a boolean if a field has been set.
func (o *SubstitutionPreferences) HasSubstitutionOptions() bool {
	if o != nil && !IsNil(o.SubstitutionOptions) {
		return true
	}

	return false
}

// SetSubstitutionOptions gets a reference to the given []SubstitutionOption and assigns it to the SubstitutionOptions field.
func (o *SubstitutionPreferences) SetSubstitutionOptions(v []SubstitutionOption) {
	o.SubstitutionOptions = v
}

func (o SubstitutionPreferences) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["SubstitutionType"] = o.SubstitutionType
	if !IsNil(o.SubstitutionOptions) {
		toSerialize["SubstitutionOptions"] = o.SubstitutionOptions
	}
	return toSerialize, nil
}

type NullableSubstitutionPreferences struct {
	value *SubstitutionPreferences
	isSet bool
}

func (v NullableSubstitutionPreferences) Get() *SubstitutionPreferences {
	return v.value
}

func (v *NullableSubstitutionPreferences) Set(val *SubstitutionPreferences) {
	v.value = val
	v.isSet = true
}

func (v NullableSubstitutionPreferences) IsSet() bool {
	return v.isSet
}

func (v *NullableSubstitutionPreferences) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubstitutionPreferences(val *SubstitutionPreferences) *NullableSubstitutionPreferences {
	return &NullableSubstitutionPreferences{value: val, isSet: true}
}

func (v NullableSubstitutionPreferences) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSubstitutionPreferences) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
