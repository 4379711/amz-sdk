package product_pricing_20220501

import (
	"github.com/bytedance/sonic"
)

// checks if the LowestPricedOffersInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LowestPricedOffersInput{}

// LowestPricedOffersInput The input required for building `LowestPricedOffers` data in the response.
type LowestPricedOffersInput struct {
	ItemCondition Condition `json:"itemCondition"`
	// The input parameter specifies the type of offers requested for `LowestPricedOffers`. This applies to `Consumer` and `Business` offers. `Consumer` is the default `offerType`.
	OfferType string `json:"offerType"`
}

type _LowestPricedOffersInput LowestPricedOffersInput

// NewLowestPricedOffersInput instantiates a new LowestPricedOffersInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLowestPricedOffersInput(itemCondition Condition, offerType string) *LowestPricedOffersInput {
	this := LowestPricedOffersInput{}
	this.ItemCondition = itemCondition
	this.OfferType = offerType
	return &this
}

// NewLowestPricedOffersInputWithDefaults instantiates a new LowestPricedOffersInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLowestPricedOffersInputWithDefaults() *LowestPricedOffersInput {
	this := LowestPricedOffersInput{}
	return &this
}

// GetItemCondition returns the ItemCondition field value
func (o *LowestPricedOffersInput) GetItemCondition() Condition {
	if o == nil {
		var ret Condition
		return ret
	}

	return o.ItemCondition
}

// GetItemConditionOk returns a tuple with the ItemCondition field value
// and a boolean to check if the value has been set.
func (o *LowestPricedOffersInput) GetItemConditionOk() (*Condition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ItemCondition, true
}

// SetItemCondition sets field value
func (o *LowestPricedOffersInput) SetItemCondition(v Condition) {
	o.ItemCondition = v
}

// GetOfferType returns the OfferType field value
func (o *LowestPricedOffersInput) GetOfferType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OfferType
}

// GetOfferTypeOk returns a tuple with the OfferType field value
// and a boolean to check if the value has been set.
func (o *LowestPricedOffersInput) GetOfferTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OfferType, true
}

// SetOfferType sets field value
func (o *LowestPricedOffersInput) SetOfferType(v string) {
	o.OfferType = v
}

func (o LowestPricedOffersInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["itemCondition"] = o.ItemCondition
	toSerialize["offerType"] = o.OfferType
	return toSerialize, nil
}

type NullableLowestPricedOffersInput struct {
	value *LowestPricedOffersInput
	isSet bool
}

func (v NullableLowestPricedOffersInput) Get() *LowestPricedOffersInput {
	return v.value
}

func (v *NullableLowestPricedOffersInput) Set(val *LowestPricedOffersInput) {
	v.value = val
	v.isSet = true
}

func (v NullableLowestPricedOffersInput) IsSet() bool {
	return v.isSet
}

func (v *NullableLowestPricedOffersInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLowestPricedOffersInput(val *LowestPricedOffersInput) *NullableLowestPricedOffersInput {
	return &NullableLowestPricedOffersInput{value: val, isSet: true}
}

func (v NullableLowestPricedOffersInput) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableLowestPricedOffersInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
