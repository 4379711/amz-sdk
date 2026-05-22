package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the RankedTargetWithThemedBidsResponseWithError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RankedTargetWithThemedBidsResponseWithError{}

// RankedTargetWithThemedBidsResponseWithError struct for RankedTargetWithThemedBidsResponseWithError
type RankedTargetWithThemedBidsResponseWithError struct {
	Error             *SPTargetingError            `json:"error,omitempty"`
	KeywordTargetList []RankedTargetWithThemedBids `json:"keywordTargetList,omitempty"`
	// A list of impact metrics which anticipates the number of clicks and orders you will receive if you target all targeting expressions using the low, medium, and high suggested bid.
	ImpactMetrics []ImpactMetrics `json:"impactMetrics,omitempty"`
}

// NewRankedTargetWithThemedBidsResponseWithError instantiates a new RankedTargetWithThemedBidsResponseWithError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRankedTargetWithThemedBidsResponseWithError() *RankedTargetWithThemedBidsResponseWithError {
	this := RankedTargetWithThemedBidsResponseWithError{}
	return &this
}

// NewRankedTargetWithThemedBidsResponseWithErrorWithDefaults instantiates a new RankedTargetWithThemedBidsResponseWithError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRankedTargetWithThemedBidsResponseWithErrorWithDefaults() *RankedTargetWithThemedBidsResponseWithError {
	this := RankedTargetWithThemedBidsResponseWithError{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *RankedTargetWithThemedBidsResponseWithError) GetError() SPTargetingError {
	if o == nil || IsNil(o.Error) {
		var ret SPTargetingError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RankedTargetWithThemedBidsResponseWithError) GetErrorOk() (*SPTargetingError, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *RankedTargetWithThemedBidsResponseWithError) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given SPTargetingError and assigns it to the Error field.
func (o *RankedTargetWithThemedBidsResponseWithError) SetError(v SPTargetingError) {
	o.Error = &v
}

// GetKeywordTargetList returns the KeywordTargetList field value if set, zero value otherwise.
func (o *RankedTargetWithThemedBidsResponseWithError) GetKeywordTargetList() []RankedTargetWithThemedBids {
	if o == nil || IsNil(o.KeywordTargetList) {
		var ret []RankedTargetWithThemedBids
		return ret
	}
	return o.KeywordTargetList
}

// GetKeywordTargetListOk returns a tuple with the KeywordTargetList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RankedTargetWithThemedBidsResponseWithError) GetKeywordTargetListOk() ([]RankedTargetWithThemedBids, bool) {
	if o == nil || IsNil(o.KeywordTargetList) {
		return nil, false
	}
	return o.KeywordTargetList, true
}

// HasKeywordTargetList returns a boolean if a field has been set.
func (o *RankedTargetWithThemedBidsResponseWithError) HasKeywordTargetList() bool {
	if o != nil && !IsNil(o.KeywordTargetList) {
		return true
	}

	return false
}

// SetKeywordTargetList gets a reference to the given []RankedTargetWithThemedBids and assigns it to the KeywordTargetList field.
func (o *RankedTargetWithThemedBidsResponseWithError) SetKeywordTargetList(v []RankedTargetWithThemedBids) {
	o.KeywordTargetList = v
}

// GetImpactMetrics returns the ImpactMetrics field value if set, zero value otherwise.
func (o *RankedTargetWithThemedBidsResponseWithError) GetImpactMetrics() []ImpactMetrics {
	if o == nil || IsNil(o.ImpactMetrics) {
		var ret []ImpactMetrics
		return ret
	}
	return o.ImpactMetrics
}

// GetImpactMetricsOk returns a tuple with the ImpactMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RankedTargetWithThemedBidsResponseWithError) GetImpactMetricsOk() ([]ImpactMetrics, bool) {
	if o == nil || IsNil(o.ImpactMetrics) {
		return nil, false
	}
	return o.ImpactMetrics, true
}

// HasImpactMetrics returns a boolean if a field has been set.
func (o *RankedTargetWithThemedBidsResponseWithError) HasImpactMetrics() bool {
	if o != nil && !IsNil(o.ImpactMetrics) {
		return true
	}

	return false
}

// SetImpactMetrics gets a reference to the given []ImpactMetrics and assigns it to the ImpactMetrics field.
func (o *RankedTargetWithThemedBidsResponseWithError) SetImpactMetrics(v []ImpactMetrics) {
	o.ImpactMetrics = v
}

func (o RankedTargetWithThemedBidsResponseWithError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	if !IsNil(o.KeywordTargetList) {
		toSerialize["keywordTargetList"] = o.KeywordTargetList
	}
	if !IsNil(o.ImpactMetrics) {
		toSerialize["impactMetrics"] = o.ImpactMetrics
	}
	return toSerialize, nil
}

type NullableRankedTargetWithThemedBidsResponseWithError struct {
	value *RankedTargetWithThemedBidsResponseWithError
	isSet bool
}

func (v NullableRankedTargetWithThemedBidsResponseWithError) Get() *RankedTargetWithThemedBidsResponseWithError {
	return v.value
}

func (v *NullableRankedTargetWithThemedBidsResponseWithError) Set(val *RankedTargetWithThemedBidsResponseWithError) {
	v.value = val
	v.isSet = true
}

func (v NullableRankedTargetWithThemedBidsResponseWithError) IsSet() bool {
	return v.isSet
}

func (v *NullableRankedTargetWithThemedBidsResponseWithError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRankedTargetWithThemedBidsResponseWithError(val *RankedTargetWithThemedBidsResponseWithError) *NullableRankedTargetWithThemedBidsResponseWithError {
	return &NullableRankedTargetWithThemedBidsResponseWithError{value: val, isSet: true}
}

func (v NullableRankedTargetWithThemedBidsResponseWithError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableRankedTargetWithThemedBidsResponseWithError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
