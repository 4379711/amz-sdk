package fba_inventory

import (
	"github.com/bytedance/sonic"
)

// checks if the GetInventorySummariesResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetInventorySummariesResult{}

// GetInventorySummariesResult The payload schema for the getInventorySummaries operation.
type GetInventorySummariesResult struct {
	Granularity Granularity `json:"granularity"`
	// A list of inventory summaries.
	InventorySummaries []InventorySummary `json:"inventorySummaries"`
}

type _GetInventorySummariesResult GetInventorySummariesResult

// NewGetInventorySummariesResult instantiates a new GetInventorySummariesResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetInventorySummariesResult(granularity Granularity, inventorySummaries []InventorySummary) *GetInventorySummariesResult {
	this := GetInventorySummariesResult{}
	this.Granularity = granularity
	this.InventorySummaries = inventorySummaries
	return &this
}

// NewGetInventorySummariesResultWithDefaults instantiates a new GetInventorySummariesResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetInventorySummariesResultWithDefaults() *GetInventorySummariesResult {
	this := GetInventorySummariesResult{}
	return &this
}

// GetGranularity returns the Granularity field value
func (o *GetInventorySummariesResult) GetGranularity() Granularity {
	if o == nil {
		var ret Granularity
		return ret
	}

	return o.Granularity
}

// GetGranularityOk returns a tuple with the Granularity field value
// and a boolean to check if the value has been set.
func (o *GetInventorySummariesResult) GetGranularityOk() (*Granularity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Granularity, true
}

// SetGranularity sets field value
func (o *GetInventorySummariesResult) SetGranularity(v Granularity) {
	o.Granularity = v
}

// GetInventorySummaries returns the InventorySummaries field value
func (o *GetInventorySummariesResult) GetInventorySummaries() []InventorySummary {
	if o == nil {
		var ret []InventorySummary
		return ret
	}

	return o.InventorySummaries
}

// GetInventorySummariesOk returns a tuple with the InventorySummaries field value
// and a boolean to check if the value has been set.
func (o *GetInventorySummariesResult) GetInventorySummariesOk() ([]InventorySummary, bool) {
	if o == nil {
		return nil, false
	}
	return o.InventorySummaries, true
}

// SetInventorySummaries sets field value
func (o *GetInventorySummariesResult) SetInventorySummaries(v []InventorySummary) {
	o.InventorySummaries = v
}

func (o GetInventorySummariesResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["granularity"] = o.Granularity
	toSerialize["inventorySummaries"] = o.InventorySummaries
	return toSerialize, nil
}

type NullableGetInventorySummariesResult struct {
	value *GetInventorySummariesResult
	isSet bool
}

func (v NullableGetInventorySummariesResult) Get() *GetInventorySummariesResult {
	return v.value
}

func (v *NullableGetInventorySummariesResult) Set(val *GetInventorySummariesResult) {
	v.value = val
	v.isSet = true
}

func (v NullableGetInventorySummariesResult) IsSet() bool {
	return v.isSet
}

func (v *NullableGetInventorySummariesResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetInventorySummariesResult(val *GetInventorySummariesResult) *NullableGetInventorySummariesResult {
	return &NullableGetInventorySummariesResult{value: val, isSet: true}
}

func (v NullableGetInventorySummariesResult) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetInventorySummariesResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
