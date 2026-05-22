package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the Values type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Values{}

// Values Metrics benchmark values.
type Values struct {
	Conversions *Conversions `json:"conversions,omitempty"`
	Clicks      *Clicks      `json:"clicks,omitempty"`
	Impressions *Impressions `json:"impressions,omitempty"`
}

// NewValues instantiates a new Values object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValues() *Values {
	this := Values{}
	return &this
}

// NewValuesWithDefaults instantiates a new Values object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValuesWithDefaults() *Values {
	this := Values{}
	return &this
}

// GetConversions returns the Conversions field value if set, zero value otherwise.
func (o *Values) GetConversions() Conversions {
	if o == nil || IsNil(o.Conversions) {
		var ret Conversions
		return ret
	}
	return *o.Conversions
}

// GetConversionsOk returns a tuple with the Conversions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Values) GetConversionsOk() (*Conversions, bool) {
	if o == nil || IsNil(o.Conversions) {
		return nil, false
	}
	return o.Conversions, true
}

// HasConversions returns a boolean if a field has been set.
func (o *Values) HasConversions() bool {
	if o != nil && !IsNil(o.Conversions) {
		return true
	}

	return false
}

// SetConversions gets a reference to the given Conversions and assigns it to the Conversions field.
func (o *Values) SetConversions(v Conversions) {
	o.Conversions = &v
}

// GetClicks returns the Clicks field value if set, zero value otherwise.
func (o *Values) GetClicks() Clicks {
	if o == nil || IsNil(o.Clicks) {
		var ret Clicks
		return ret
	}
	return *o.Clicks
}

// GetClicksOk returns a tuple with the Clicks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Values) GetClicksOk() (*Clicks, bool) {
	if o == nil || IsNil(o.Clicks) {
		return nil, false
	}
	return o.Clicks, true
}

// HasClicks returns a boolean if a field has been set.
func (o *Values) HasClicks() bool {
	if o != nil && !IsNil(o.Clicks) {
		return true
	}

	return false
}

// SetClicks gets a reference to the given Clicks and assigns it to the Clicks field.
func (o *Values) SetClicks(v Clicks) {
	o.Clicks = &v
}

// GetImpressions returns the Impressions field value if set, zero value otherwise.
func (o *Values) GetImpressions() Impressions {
	if o == nil || IsNil(o.Impressions) {
		var ret Impressions
		return ret
	}
	return *o.Impressions
}

// GetImpressionsOk returns a tuple with the Impressions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Values) GetImpressionsOk() (*Impressions, bool) {
	if o == nil || IsNil(o.Impressions) {
		return nil, false
	}
	return o.Impressions, true
}

// HasImpressions returns a boolean if a field has been set.
func (o *Values) HasImpressions() bool {
	if o != nil && !IsNil(o.Impressions) {
		return true
	}

	return false
}

// SetImpressions gets a reference to the given Impressions and assigns it to the Impressions field.
func (o *Values) SetImpressions(v Impressions) {
	o.Impressions = &v
}

func (o Values) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Conversions) {
		toSerialize["conversions"] = o.Conversions
	}
	if !IsNil(o.Clicks) {
		toSerialize["clicks"] = o.Clicks
	}
	if !IsNil(o.Impressions) {
		toSerialize["impressions"] = o.Impressions
	}
	return toSerialize, nil
}

type NullableValues struct {
	value *Values
	isSet bool
}

func (v NullableValues) Get() *Values {
	return v.value
}

func (v *NullableValues) Set(val *Values) {
	v.value = val
	v.isSet = true
}

func (v NullableValues) IsSet() bool {
	return v.isSet
}

func (v *NullableValues) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValues(val *Values) *NullableValues {
	return &NullableValues{value: val, isSet: true}
}

func (v NullableValues) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableValues) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
