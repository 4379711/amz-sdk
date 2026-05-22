package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the BidRecommendationPerTargetingExpressionV4 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BidRecommendationPerTargetingExpressionV4{}

// BidRecommendationPerTargetingExpressionV4 struct for BidRecommendationPerTargetingExpressionV4
type BidRecommendationPerTargetingExpressionV4 struct {
	BidValues           []BidValue            `json:"bidValues"`
	TargetingExpression TargetingExpressionV4 `json:"targetingExpression"`
}

type _BidRecommendationPerTargetingExpressionV4 BidRecommendationPerTargetingExpressionV4

// NewBidRecommendationPerTargetingExpressionV4 instantiates a new BidRecommendationPerTargetingExpressionV4 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBidRecommendationPerTargetingExpressionV4(bidValues []BidValue, targetingExpression TargetingExpressionV4) *BidRecommendationPerTargetingExpressionV4 {
	this := BidRecommendationPerTargetingExpressionV4{}
	this.BidValues = bidValues
	this.TargetingExpression = targetingExpression
	return &this
}

// NewBidRecommendationPerTargetingExpressionV4WithDefaults instantiates a new BidRecommendationPerTargetingExpressionV4 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBidRecommendationPerTargetingExpressionV4WithDefaults() *BidRecommendationPerTargetingExpressionV4 {
	this := BidRecommendationPerTargetingExpressionV4{}
	return &this
}

// GetBidValues returns the BidValues field value
func (o *BidRecommendationPerTargetingExpressionV4) GetBidValues() []BidValue {
	if o == nil {
		var ret []BidValue
		return ret
	}

	return o.BidValues
}

// GetBidValuesOk returns a tuple with the BidValues field value
// and a boolean to check if the value has been set.
func (o *BidRecommendationPerTargetingExpressionV4) GetBidValuesOk() ([]BidValue, bool) {
	if o == nil {
		return nil, false
	}
	return o.BidValues, true
}

// SetBidValues sets field value
func (o *BidRecommendationPerTargetingExpressionV4) SetBidValues(v []BidValue) {
	o.BidValues = v
}

// GetTargetingExpression returns the TargetingExpression field value
func (o *BidRecommendationPerTargetingExpressionV4) GetTargetingExpression() TargetingExpressionV4 {
	if o == nil {
		var ret TargetingExpressionV4
		return ret
	}

	return o.TargetingExpression
}

// GetTargetingExpressionOk returns a tuple with the TargetingExpression field value
// and a boolean to check if the value has been set.
func (o *BidRecommendationPerTargetingExpressionV4) GetTargetingExpressionOk() (*TargetingExpressionV4, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetingExpression, true
}

// SetTargetingExpression sets field value
func (o *BidRecommendationPerTargetingExpressionV4) SetTargetingExpression(v TargetingExpressionV4) {
	o.TargetingExpression = v
}

func (o BidRecommendationPerTargetingExpressionV4) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["bidValues"] = o.BidValues
	toSerialize["targetingExpression"] = o.TargetingExpression
	return toSerialize, nil
}

type NullableBidRecommendationPerTargetingExpressionV4 struct {
	value *BidRecommendationPerTargetingExpressionV4
	isSet bool
}

func (v NullableBidRecommendationPerTargetingExpressionV4) Get() *BidRecommendationPerTargetingExpressionV4 {
	return v.value
}

func (v *NullableBidRecommendationPerTargetingExpressionV4) Set(val *BidRecommendationPerTargetingExpressionV4) {
	v.value = val
	v.isSet = true
}

func (v NullableBidRecommendationPerTargetingExpressionV4) IsSet() bool {
	return v.isSet
}

func (v *NullableBidRecommendationPerTargetingExpressionV4) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBidRecommendationPerTargetingExpressionV4(val *BidRecommendationPerTargetingExpressionV4) *NullableBidRecommendationPerTargetingExpressionV4 {
	return &NullableBidRecommendationPerTargetingExpressionV4{value: val, isSet: true}
}

func (v NullableBidRecommendationPerTargetingExpressionV4) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBidRecommendationPerTargetingExpressionV4) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
