package fba_inventory

import (
	"github.com/bytedance/sonic"
)

// checks if the Granularity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Granularity{}

// Granularity Describes a granularity at which inventory data can be aggregated. For example, if you use Marketplace granularity, the fulfillable quantity will reflect inventory that could be fulfilled in the given marketplace.
type Granularity struct {
	// The granularity type for the inventory aggregation level.
	GranularityType *string `json:"granularityType,omitempty"`
	// The granularity ID for the specified granularity type. When granularityType is Marketplace, specify the marketplaceId.
	GranularityId *string `json:"granularityId,omitempty"`
}

// NewGranularity instantiates a new Granularity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGranularity() *Granularity {
	this := Granularity{}
	return &this
}

// NewGranularityWithDefaults instantiates a new Granularity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGranularityWithDefaults() *Granularity {
	this := Granularity{}
	return &this
}

// GetGranularityType returns the GranularityType field value if set, zero value otherwise.
func (o *Granularity) GetGranularityType() string {
	if o == nil || IsNil(o.GranularityType) {
		var ret string
		return ret
	}
	return *o.GranularityType
}

// GetGranularityTypeOk returns a tuple with the GranularityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Granularity) GetGranularityTypeOk() (*string, bool) {
	if o == nil || IsNil(o.GranularityType) {
		return nil, false
	}
	return o.GranularityType, true
}

// HasGranularityType returns a boolean if a field has been set.
func (o *Granularity) HasGranularityType() bool {
	if o != nil && !IsNil(o.GranularityType) {
		return true
	}

	return false
}

// SetGranularityType gets a reference to the given string and assigns it to the GranularityType field.
func (o *Granularity) SetGranularityType(v string) {
	o.GranularityType = &v
}

// GetGranularityId returns the GranularityId field value if set, zero value otherwise.
func (o *Granularity) GetGranularityId() string {
	if o == nil || IsNil(o.GranularityId) {
		var ret string
		return ret
	}
	return *o.GranularityId
}

// GetGranularityIdOk returns a tuple with the GranularityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Granularity) GetGranularityIdOk() (*string, bool) {
	if o == nil || IsNil(o.GranularityId) {
		return nil, false
	}
	return o.GranularityId, true
}

// HasGranularityId returns a boolean if a field has been set.
func (o *Granularity) HasGranularityId() bool {
	if o != nil && !IsNil(o.GranularityId) {
		return true
	}

	return false
}

// SetGranularityId gets a reference to the given string and assigns it to the GranularityId field.
func (o *Granularity) SetGranularityId(v string) {
	o.GranularityId = &v
}

func (o Granularity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GranularityType) {
		toSerialize["granularityType"] = o.GranularityType
	}
	if !IsNil(o.GranularityId) {
		toSerialize["granularityId"] = o.GranularityId
	}
	return toSerialize, nil
}

type NullableGranularity struct {
	value *Granularity
	isSet bool
}

func (v NullableGranularity) Get() *Granularity {
	return v.value
}

func (v *NullableGranularity) Set(val *Granularity) {
	v.value = val
	v.isSet = true
}

func (v NullableGranularity) IsSet() bool {
	return v.isSet
}

func (v *NullableGranularity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGranularity(val *Granularity) *NullableGranularity {
	return &NullableGranularity{value: val, isSet: true}
}

func (v NullableGranularity) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGranularity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
