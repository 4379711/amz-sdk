package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the BidRecommendationPerTargetingExpression type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BidRecommendationPerTargetingExpression{}

// BidRecommendationPerTargetingExpression struct for BidRecommendationPerTargetingExpression
type BidRecommendationPerTargetingExpression struct {
	BidValues           []BidValue          `json:"bidValues"`
	TargetingExpression TargetingExpression `json:"targetingExpression"`
}

type _BidRecommendationPerTargetingExpression BidRecommendationPerTargetingExpression

// NewBidRecommendationPerTargetingExpression instantiates a new BidRecommendationPerTargetingExpression object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBidRecommendationPerTargetingExpression(bidValues []BidValue, targetingExpression TargetingExpression) *BidRecommendationPerTargetingExpression {
	this := BidRecommendationPerTargetingExpression{}
	this.BidValues = bidValues
	this.TargetingExpression = targetingExpression
	return &this
}

// NewBidRecommendationPerTargetingExpressionWithDefaults instantiates a new BidRecommendationPerTargetingExpression object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBidRecommendationPerTargetingExpressionWithDefaults() *BidRecommendationPerTargetingExpression {
	this := BidRecommendationPerTargetingExpression{}
	return &this
}

// GetBidValues returns the BidValues field value
func (o *BidRecommendationPerTargetingExpression) GetBidValues() []BidValue {
	if o == nil {
		var ret []BidValue
		return ret
	}

	return o.BidValues
}

// GetBidValuesOk returns a tuple with the BidValues field value
// and a boolean to check if the value has been set.
func (o *BidRecommendationPerTargetingExpression) GetBidValuesOk() ([]BidValue, bool) {
	if o == nil {
		return nil, false
	}
	return o.BidValues, true
}

// SetBidValues sets field value
func (o *BidRecommendationPerTargetingExpression) SetBidValues(v []BidValue) {
	o.BidValues = v
}

// GetTargetingExpression returns the TargetingExpression field value
func (o *BidRecommendationPerTargetingExpression) GetTargetingExpression() TargetingExpression {
	if o == nil {
		var ret TargetingExpression
		return ret
	}

	return o.TargetingExpression
}

// GetTargetingExpressionOk returns a tuple with the TargetingExpression field value
// and a boolean to check if the value has been set.
func (o *BidRecommendationPerTargetingExpression) GetTargetingExpressionOk() (*TargetingExpression, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetingExpression, true
}

// SetTargetingExpression sets field value
func (o *BidRecommendationPerTargetingExpression) SetTargetingExpression(v TargetingExpression) {
	o.TargetingExpression = v
}

func (o BidRecommendationPerTargetingExpression) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["bidValues"] = o.BidValues
	toSerialize["targetingExpression"] = o.TargetingExpression
	return toSerialize, nil
}

type NullableBidRecommendationPerTargetingExpression struct {
	value *BidRecommendationPerTargetingExpression
	isSet bool
}

func (v NullableBidRecommendationPerTargetingExpression) Get() *BidRecommendationPerTargetingExpression {
	return v.value
}

func (v *NullableBidRecommendationPerTargetingExpression) Set(val *BidRecommendationPerTargetingExpression) {
	v.value = val
	v.isSet = true
}

func (v NullableBidRecommendationPerTargetingExpression) IsSet() bool {
	return v.isSet
}

func (v *NullableBidRecommendationPerTargetingExpression) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBidRecommendationPerTargetingExpression(val *BidRecommendationPerTargetingExpression) *NullableBidRecommendationPerTargetingExpression {
	return &NullableBidRecommendationPerTargetingExpression{value: val, isSet: true}
}

func (v NullableBidRecommendationPerTargetingExpression) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBidRecommendationPerTargetingExpression) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
