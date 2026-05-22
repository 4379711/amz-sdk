package aplus_content_20201101

import (
	"github.com/bytedance/sonic"
)

// checks if the StandardComparisonTableModule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StandardComparisonTableModule{}

// StandardComparisonTableModule The standard product comparison table.
type StandardComparisonTableModule struct {
	ProductColumns  []StandardComparisonProductBlock `json:"productColumns,omitempty"`
	MetricRowLabels []PlainTextItem                  `json:"metricRowLabels,omitempty"`
}

// NewStandardComparisonTableModule instantiates a new StandardComparisonTableModule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStandardComparisonTableModule() *StandardComparisonTableModule {
	this := StandardComparisonTableModule{}
	return &this
}

// NewStandardComparisonTableModuleWithDefaults instantiates a new StandardComparisonTableModule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStandardComparisonTableModuleWithDefaults() *StandardComparisonTableModule {
	this := StandardComparisonTableModule{}
	return &this
}

// GetProductColumns returns the ProductColumns field value if set, zero value otherwise.
func (o *StandardComparisonTableModule) GetProductColumns() []StandardComparisonProductBlock {
	if o == nil || IsNil(o.ProductColumns) {
		var ret []StandardComparisonProductBlock
		return ret
	}
	return o.ProductColumns
}

// GetProductColumnsOk returns a tuple with the ProductColumns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandardComparisonTableModule) GetProductColumnsOk() ([]StandardComparisonProductBlock, bool) {
	if o == nil || IsNil(o.ProductColumns) {
		return nil, false
	}
	return o.ProductColumns, true
}

// HasProductColumns returns a boolean if a field has been set.
func (o *StandardComparisonTableModule) HasProductColumns() bool {
	if o != nil && !IsNil(o.ProductColumns) {
		return true
	}

	return false
}

// SetProductColumns gets a reference to the given []StandardComparisonProductBlock and assigns it to the ProductColumns field.
func (o *StandardComparisonTableModule) SetProductColumns(v []StandardComparisonProductBlock) {
	o.ProductColumns = v
}

// GetMetricRowLabels returns the MetricRowLabels field value if set, zero value otherwise.
func (o *StandardComparisonTableModule) GetMetricRowLabels() []PlainTextItem {
	if o == nil || IsNil(o.MetricRowLabels) {
		var ret []PlainTextItem
		return ret
	}
	return o.MetricRowLabels
}

// GetMetricRowLabelsOk returns a tuple with the MetricRowLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandardComparisonTableModule) GetMetricRowLabelsOk() ([]PlainTextItem, bool) {
	if o == nil || IsNil(o.MetricRowLabels) {
		return nil, false
	}
	return o.MetricRowLabels, true
}

// HasMetricRowLabels returns a boolean if a field has been set.
func (o *StandardComparisonTableModule) HasMetricRowLabels() bool {
	if o != nil && !IsNil(o.MetricRowLabels) {
		return true
	}

	return false
}

// SetMetricRowLabels gets a reference to the given []PlainTextItem and assigns it to the MetricRowLabels field.
func (o *StandardComparisonTableModule) SetMetricRowLabels(v []PlainTextItem) {
	o.MetricRowLabels = v
}

func (o StandardComparisonTableModule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProductColumns) {
		toSerialize["productColumns"] = o.ProductColumns
	}
	if !IsNil(o.MetricRowLabels) {
		toSerialize["metricRowLabels"] = o.MetricRowLabels
	}
	return toSerialize, nil
}

type NullableStandardComparisonTableModule struct {
	value *StandardComparisonTableModule
	isSet bool
}

func (v NullableStandardComparisonTableModule) Get() *StandardComparisonTableModule {
	return v.value
}

func (v *NullableStandardComparisonTableModule) Set(val *StandardComparisonTableModule) {
	v.value = val
	v.isSet = true
}

func (v NullableStandardComparisonTableModule) IsSet() bool {
	return v.isSet
}

func (v *NullableStandardComparisonTableModule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStandardComparisonTableModule(val *StandardComparisonTableModule) *NullableStandardComparisonTableModule {
	return &NullableStandardComparisonTableModule{value: val, isSet: true}
}

func (v NullableStandardComparisonTableModule) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableStandardComparisonTableModule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
