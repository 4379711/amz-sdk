package notifications

import (
	"github.com/bytedance/sonic"
)

// checks if the AggregationSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AggregationSettings{}

// AggregationSettings A container that holds all of the necessary properties to configure the aggregation of notifications.
type AggregationSettings struct {
	AggregationTimePeriod AggregationTimePeriod `json:"aggregationTimePeriod"`
}

type _AggregationSettings AggregationSettings

// NewAggregationSettings instantiates a new AggregationSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAggregationSettings(aggregationTimePeriod AggregationTimePeriod) *AggregationSettings {
	this := AggregationSettings{}
	this.AggregationTimePeriod = aggregationTimePeriod
	return &this
}

// NewAggregationSettingsWithDefaults instantiates a new AggregationSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAggregationSettingsWithDefaults() *AggregationSettings {
	this := AggregationSettings{}
	return &this
}

// GetAggregationTimePeriod returns the AggregationTimePeriod field value
func (o *AggregationSettings) GetAggregationTimePeriod() AggregationTimePeriod {
	if o == nil {
		var ret AggregationTimePeriod
		return ret
	}

	return o.AggregationTimePeriod
}

// GetAggregationTimePeriodOk returns a tuple with the AggregationTimePeriod field value
// and a boolean to check if the value has been set.
func (o *AggregationSettings) GetAggregationTimePeriodOk() (*AggregationTimePeriod, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AggregationTimePeriod, true
}

// SetAggregationTimePeriod sets field value
func (o *AggregationSettings) SetAggregationTimePeriod(v AggregationTimePeriod) {
	o.AggregationTimePeriod = v
}

func (o AggregationSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["aggregationTimePeriod"] = o.AggregationTimePeriod
	return toSerialize, nil
}

type NullableAggregationSettings struct {
	value *AggregationSettings
	isSet bool
}

func (v NullableAggregationSettings) Get() *AggregationSettings {
	return v.value
}

func (v *NullableAggregationSettings) Set(val *AggregationSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableAggregationSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableAggregationSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAggregationSettings(val *AggregationSettings) *NullableAggregationSettings {
	return &NullableAggregationSettings{value: val, isSet: true}
}

func (v NullableAggregationSettings) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAggregationSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
