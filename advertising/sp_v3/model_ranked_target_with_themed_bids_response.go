package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the RankedTargetWithThemedBidsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RankedTargetWithThemedBidsResponse{}

// RankedTargetWithThemedBidsResponse struct for RankedTargetWithThemedBidsResponse
type RankedTargetWithThemedBidsResponse struct {
	KeywordTargetList []RankedTargetWithThemedBids `json:"keywordTargetList,omitempty"`
	// A list of impact metrics which anticipates the number of clicks and orders you will receive if you target all targeting expressions using the low, medium, and high suggested bid.
	ImpactMetrics []ImpactMetrics `json:"impactMetrics,omitempty"`
}

// NewRankedTargetWithThemedBidsResponse instantiates a new RankedTargetWithThemedBidsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRankedTargetWithThemedBidsResponse() *RankedTargetWithThemedBidsResponse {
	this := RankedTargetWithThemedBidsResponse{}
	return &this
}

// NewRankedTargetWithThemedBidsResponseWithDefaults instantiates a new RankedTargetWithThemedBidsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRankedTargetWithThemedBidsResponseWithDefaults() *RankedTargetWithThemedBidsResponse {
	this := RankedTargetWithThemedBidsResponse{}
	return &this
}

// GetKeywordTargetList returns the KeywordTargetList field value if set, zero value otherwise.
func (o *RankedTargetWithThemedBidsResponse) GetKeywordTargetList() []RankedTargetWithThemedBids {
	if o == nil || IsNil(o.KeywordTargetList) {
		var ret []RankedTargetWithThemedBids
		return ret
	}
	return o.KeywordTargetList
}

// GetKeywordTargetListOk returns a tuple with the KeywordTargetList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RankedTargetWithThemedBidsResponse) GetKeywordTargetListOk() ([]RankedTargetWithThemedBids, bool) {
	if o == nil || IsNil(o.KeywordTargetList) {
		return nil, false
	}
	return o.KeywordTargetList, true
}

// HasKeywordTargetList returns a boolean if a field has been set.
func (o *RankedTargetWithThemedBidsResponse) HasKeywordTargetList() bool {
	if o != nil && !IsNil(o.KeywordTargetList) {
		return true
	}

	return false
}

// SetKeywordTargetList gets a reference to the given []RankedTargetWithThemedBids and assigns it to the KeywordTargetList field.
func (o *RankedTargetWithThemedBidsResponse) SetKeywordTargetList(v []RankedTargetWithThemedBids) {
	o.KeywordTargetList = v
}

// GetImpactMetrics returns the ImpactMetrics field value if set, zero value otherwise.
func (o *RankedTargetWithThemedBidsResponse) GetImpactMetrics() []ImpactMetrics {
	if o == nil || IsNil(o.ImpactMetrics) {
		var ret []ImpactMetrics
		return ret
	}
	return o.ImpactMetrics
}

// GetImpactMetricsOk returns a tuple with the ImpactMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RankedTargetWithThemedBidsResponse) GetImpactMetricsOk() ([]ImpactMetrics, bool) {
	if o == nil || IsNil(o.ImpactMetrics) {
		return nil, false
	}
	return o.ImpactMetrics, true
}

// HasImpactMetrics returns a boolean if a field has been set.
func (o *RankedTargetWithThemedBidsResponse) HasImpactMetrics() bool {
	if o != nil && !IsNil(o.ImpactMetrics) {
		return true
	}

	return false
}

// SetImpactMetrics gets a reference to the given []ImpactMetrics and assigns it to the ImpactMetrics field.
func (o *RankedTargetWithThemedBidsResponse) SetImpactMetrics(v []ImpactMetrics) {
	o.ImpactMetrics = v
}

func (o RankedTargetWithThemedBidsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.KeywordTargetList) {
		toSerialize["keywordTargetList"] = o.KeywordTargetList
	}
	if !IsNil(o.ImpactMetrics) {
		toSerialize["impactMetrics"] = o.ImpactMetrics
	}
	return toSerialize, nil
}

type NullableRankedTargetWithThemedBidsResponse struct {
	value *RankedTargetWithThemedBidsResponse
	isSet bool
}

func (v NullableRankedTargetWithThemedBidsResponse) Get() *RankedTargetWithThemedBidsResponse {
	return v.value
}

func (v *NullableRankedTargetWithThemedBidsResponse) Set(val *RankedTargetWithThemedBidsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRankedTargetWithThemedBidsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRankedTargetWithThemedBidsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRankedTargetWithThemedBidsResponse(val *RankedTargetWithThemedBidsResponse) *NullableRankedTargetWithThemedBidsResponse {
	return &NullableRankedTargetWithThemedBidsResponse{value: val, isSet: true}
}

func (v NullableRankedTargetWithThemedBidsResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableRankedTargetWithThemedBidsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
