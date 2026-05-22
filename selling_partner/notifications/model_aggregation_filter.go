package notifications

import (
	"github.com/bytedance/sonic"
)

// checks if the AggregationFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AggregationFilter{}

// AggregationFilter A filter used to select the aggregation time period at which to send notifications (for example: limit to one notification every five minutes for high frequency notifications).
type AggregationFilter struct {
	AggregationSettings *AggregationSettings `json:"aggregationSettings,omitempty"`
}

// NewAggregationFilter instantiates a new AggregationFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAggregationFilter() *AggregationFilter {
	this := AggregationFilter{}
	return &this
}

// NewAggregationFilterWithDefaults instantiates a new AggregationFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAggregationFilterWithDefaults() *AggregationFilter {
	this := AggregationFilter{}
	return &this
}

// GetAggregationSettings returns the AggregationSettings field value if set, zero value otherwise.
func (o *AggregationFilter) GetAggregationSettings() AggregationSettings {
	if o == nil || IsNil(o.AggregationSettings) {
		var ret AggregationSettings
		return ret
	}
	return *o.AggregationSettings
}

// GetAggregationSettingsOk returns a tuple with the AggregationSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregationFilter) GetAggregationSettingsOk() (*AggregationSettings, bool) {
	if o == nil || IsNil(o.AggregationSettings) {
		return nil, false
	}
	return o.AggregationSettings, true
}

// HasAggregationSettings returns a boolean if a field has been set.
func (o *AggregationFilter) HasAggregationSettings() bool {
	if o != nil && !IsNil(o.AggregationSettings) {
		return true
	}

	return false
}

// SetAggregationSettings gets a reference to the given AggregationSettings and assigns it to the AggregationSettings field.
func (o *AggregationFilter) SetAggregationSettings(v AggregationSettings) {
	o.AggregationSettings = &v
}

func (o AggregationFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AggregationSettings) {
		toSerialize["aggregationSettings"] = o.AggregationSettings
	}
	return toSerialize, nil
}

type NullableAggregationFilter struct {
	value *AggregationFilter
	isSet bool
}

func (v NullableAggregationFilter) Get() *AggregationFilter {
	return v.value
}

func (v *NullableAggregationFilter) Set(val *AggregationFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableAggregationFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableAggregationFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAggregationFilter(val *AggregationFilter) *NullableAggregationFilter {
	return &NullableAggregationFilter{value: val, isSet: true}
}

func (v NullableAggregationFilter) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAggregationFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
