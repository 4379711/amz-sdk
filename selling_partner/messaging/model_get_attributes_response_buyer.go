package messaging

import (
	"github.com/bytedance/sonic"
)

// checks if the GetAttributesResponseBuyer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAttributesResponseBuyer{}

// GetAttributesResponseBuyer The list of attributes related to the buyer.
type GetAttributesResponseBuyer struct {
	// The buyer's language of preference, indicated with a locale-specific language tag. Examples: \"en-US\", \"zh-CN\", and \"en-GB\".
	Locale *string `json:"locale,omitempty"`
}

// NewGetAttributesResponseBuyer instantiates a new GetAttributesResponseBuyer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAttributesResponseBuyer() *GetAttributesResponseBuyer {
	this := GetAttributesResponseBuyer{}
	return &this
}

// NewGetAttributesResponseBuyerWithDefaults instantiates a new GetAttributesResponseBuyer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAttributesResponseBuyerWithDefaults() *GetAttributesResponseBuyer {
	this := GetAttributesResponseBuyer{}
	return &this
}

// GetLocale returns the Locale field value if set, zero value otherwise.
func (o *GetAttributesResponseBuyer) GetLocale() string {
	if o == nil || IsNil(o.Locale) {
		var ret string
		return ret
	}
	return *o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAttributesResponseBuyer) GetLocaleOk() (*string, bool) {
	if o == nil || IsNil(o.Locale) {
		return nil, false
	}
	return o.Locale, true
}

// HasLocale returns a boolean if a field has been set.
func (o *GetAttributesResponseBuyer) HasLocale() bool {
	if o != nil && !IsNil(o.Locale) {
		return true
	}

	return false
}

// SetLocale gets a reference to the given string and assigns it to the Locale field.
func (o *GetAttributesResponseBuyer) SetLocale(v string) {
	o.Locale = &v
}

func (o GetAttributesResponseBuyer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Locale) {
		toSerialize["locale"] = o.Locale
	}
	return toSerialize, nil
}

type NullableGetAttributesResponseBuyer struct {
	value *GetAttributesResponseBuyer
	isSet bool
}

func (v NullableGetAttributesResponseBuyer) Get() *GetAttributesResponseBuyer {
	return v.value
}

func (v *NullableGetAttributesResponseBuyer) Set(val *GetAttributesResponseBuyer) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAttributesResponseBuyer) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAttributesResponseBuyer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAttributesResponseBuyer(val *GetAttributesResponseBuyer) *NullableGetAttributesResponseBuyer {
	return &NullableGetAttributesResponseBuyer{value: val, isSet: true}
}

func (v NullableGetAttributesResponseBuyer) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetAttributesResponseBuyer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
