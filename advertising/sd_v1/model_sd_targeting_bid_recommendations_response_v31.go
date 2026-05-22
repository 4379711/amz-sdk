package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the SDTargetingBidRecommendationsResponseV31 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SDTargetingBidRecommendationsResponseV31{}

// SDTargetingBidRecommendationsResponseV31 Response to a request for targeting bid recommendations.
type SDTargetingBidRecommendationsResponseV31 struct {
	CostType           SDCostTypeV31                                                     `json:"costType"`
	BidRecommendations []SDTargetingBidRecommendationsResponseV31BidRecommendationsInner `json:"bidRecommendations"`
}

type _SDTargetingBidRecommendationsResponseV31 SDTargetingBidRecommendationsResponseV31

// NewSDTargetingBidRecommendationsResponseV31 instantiates a new SDTargetingBidRecommendationsResponseV31 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSDTargetingBidRecommendationsResponseV31(costType SDCostTypeV31, bidRecommendations []SDTargetingBidRecommendationsResponseV31BidRecommendationsInner) *SDTargetingBidRecommendationsResponseV31 {
	this := SDTargetingBidRecommendationsResponseV31{}
	this.CostType = costType
	this.BidRecommendations = bidRecommendations
	return &this
}

// NewSDTargetingBidRecommendationsResponseV31WithDefaults instantiates a new SDTargetingBidRecommendationsResponseV31 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSDTargetingBidRecommendationsResponseV31WithDefaults() *SDTargetingBidRecommendationsResponseV31 {
	this := SDTargetingBidRecommendationsResponseV31{}
	return &this
}

// GetCostType returns the CostType field value
func (o *SDTargetingBidRecommendationsResponseV31) GetCostType() SDCostTypeV31 {
	if o == nil {
		var ret SDCostTypeV31
		return ret
	}

	return o.CostType
}

// GetCostTypeOk returns a tuple with the CostType field value
// and a boolean to check if the value has been set.
func (o *SDTargetingBidRecommendationsResponseV31) GetCostTypeOk() (*SDCostTypeV31, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CostType, true
}

// SetCostType sets field value
func (o *SDTargetingBidRecommendationsResponseV31) SetCostType(v SDCostTypeV31) {
	o.CostType = v
}

// GetBidRecommendations returns the BidRecommendations field value
func (o *SDTargetingBidRecommendationsResponseV31) GetBidRecommendations() []SDTargetingBidRecommendationsResponseV31BidRecommendationsInner {
	if o == nil {
		var ret []SDTargetingBidRecommendationsResponseV31BidRecommendationsInner
		return ret
	}

	return o.BidRecommendations
}

// GetBidRecommendationsOk returns a tuple with the BidRecommendations field value
// and a boolean to check if the value has been set.
func (o *SDTargetingBidRecommendationsResponseV31) GetBidRecommendationsOk() ([]SDTargetingBidRecommendationsResponseV31BidRecommendationsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.BidRecommendations, true
}

// SetBidRecommendations sets field value
func (o *SDTargetingBidRecommendationsResponseV31) SetBidRecommendations(v []SDTargetingBidRecommendationsResponseV31BidRecommendationsInner) {
	o.BidRecommendations = v
}

func (o SDTargetingBidRecommendationsResponseV31) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["costType"] = o.CostType
	toSerialize["bidRecommendations"] = o.BidRecommendations
	return toSerialize, nil
}

type NullableSDTargetingBidRecommendationsResponseV31 struct {
	value *SDTargetingBidRecommendationsResponseV31
	isSet bool
}

func (v NullableSDTargetingBidRecommendationsResponseV31) Get() *SDTargetingBidRecommendationsResponseV31 {
	return v.value
}

func (v *NullableSDTargetingBidRecommendationsResponseV31) Set(val *SDTargetingBidRecommendationsResponseV31) {
	v.value = val
	v.isSet = true
}

func (v NullableSDTargetingBidRecommendationsResponseV31) IsSet() bool {
	return v.isSet
}

func (v *NullableSDTargetingBidRecommendationsResponseV31) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDTargetingBidRecommendationsResponseV31(val *SDTargetingBidRecommendationsResponseV31) *NullableSDTargetingBidRecommendationsResponseV31 {
	return &NullableSDTargetingBidRecommendationsResponseV31{value: val, isSet: true}
}

func (v NullableSDTargetingBidRecommendationsResponseV31) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDTargetingBidRecommendationsResponseV31) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
