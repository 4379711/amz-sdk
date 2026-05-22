package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the SBForecastingProductTarget type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SBForecastingProductTarget{}

// SBForecastingProductTarget The target associated with the ad group.
type SBForecastingProductTarget struct {
	Expressions []SBForecastingProductExpression `json:"expressions,omitempty"`
	// The associated bid. Note that this value must be less than the budget associated with the Advertiser account.
	Bid *float32 `json:"bid,omitempty"`
}

// NewSBForecastingProductTarget instantiates a new SBForecastingProductTarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSBForecastingProductTarget() *SBForecastingProductTarget {
	this := SBForecastingProductTarget{}
	return &this
}

// NewSBForecastingProductTargetWithDefaults instantiates a new SBForecastingProductTarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSBForecastingProductTargetWithDefaults() *SBForecastingProductTarget {
	this := SBForecastingProductTarget{}
	return &this
}

// GetExpressions returns the Expressions field value if set, zero value otherwise.
func (o *SBForecastingProductTarget) GetExpressions() []SBForecastingProductExpression {
	if o == nil || IsNil(o.Expressions) {
		var ret []SBForecastingProductExpression
		return ret
	}
	return o.Expressions
}

// GetExpressionsOk returns a tuple with the Expressions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBForecastingProductTarget) GetExpressionsOk() ([]SBForecastingProductExpression, bool) {
	if o == nil || IsNil(o.Expressions) {
		return nil, false
	}
	return o.Expressions, true
}

// HasExpressions returns a boolean if a field has been set.
func (o *SBForecastingProductTarget) HasExpressions() bool {
	if o != nil && !IsNil(o.Expressions) {
		return true
	}

	return false
}

// SetExpressions gets a reference to the given []SBForecastingProductExpression and assigns it to the Expressions field.
func (o *SBForecastingProductTarget) SetExpressions(v []SBForecastingProductExpression) {
	o.Expressions = v
}

// GetBid returns the Bid field value if set, zero value otherwise.
func (o *SBForecastingProductTarget) GetBid() float32 {
	if o == nil || IsNil(o.Bid) {
		var ret float32
		return ret
	}
	return *o.Bid
}

// GetBidOk returns a tuple with the Bid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBForecastingProductTarget) GetBidOk() (*float32, bool) {
	if o == nil || IsNil(o.Bid) {
		return nil, false
	}
	return o.Bid, true
}

// HasBid returns a boolean if a field has been set.
func (o *SBForecastingProductTarget) HasBid() bool {
	if o != nil && !IsNil(o.Bid) {
		return true
	}

	return false
}

// SetBid gets a reference to the given float32 and assigns it to the Bid field.
func (o *SBForecastingProductTarget) SetBid(v float32) {
	o.Bid = &v
}

func (o SBForecastingProductTarget) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Expressions) {
		toSerialize["expressions"] = o.Expressions
	}
	if !IsNil(o.Bid) {
		toSerialize["bid"] = o.Bid
	}
	return toSerialize, nil
}

type NullableSBForecastingProductTarget struct {
	value *SBForecastingProductTarget
	isSet bool
}

func (v NullableSBForecastingProductTarget) Get() *SBForecastingProductTarget {
	return v.value
}

func (v *NullableSBForecastingProductTarget) Set(val *SBForecastingProductTarget) {
	v.value = val
	v.isSet = true
}

func (v NullableSBForecastingProductTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableSBForecastingProductTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSBForecastingProductTarget(val *SBForecastingProductTarget) *NullableSBForecastingProductTarget {
	return &NullableSBForecastingProductTarget{value: val, isSet: true}
}

func (v NullableSBForecastingProductTarget) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSBForecastingProductTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
