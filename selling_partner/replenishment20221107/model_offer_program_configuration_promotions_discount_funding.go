package replenishment20221107

import (
	"github.com/bytedance/sonic"
)

// checks if the OfferProgramConfigurationPromotionsDiscountFunding type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OfferProgramConfigurationPromotionsDiscountFunding{}

// OfferProgramConfigurationPromotionsDiscountFunding A promotional percentage discount applied to the offer.
type OfferProgramConfigurationPromotionsDiscountFunding struct {
	// The percentage discount on the offer.
	Percentage *float32 `json:"percentage,omitempty"`
}

// NewOfferProgramConfigurationPromotionsDiscountFunding instantiates a new OfferProgramConfigurationPromotionsDiscountFunding object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOfferProgramConfigurationPromotionsDiscountFunding() *OfferProgramConfigurationPromotionsDiscountFunding {
	this := OfferProgramConfigurationPromotionsDiscountFunding{}
	return &this
}

// NewOfferProgramConfigurationPromotionsDiscountFundingWithDefaults instantiates a new OfferProgramConfigurationPromotionsDiscountFunding object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOfferProgramConfigurationPromotionsDiscountFundingWithDefaults() *OfferProgramConfigurationPromotionsDiscountFunding {
	this := OfferProgramConfigurationPromotionsDiscountFunding{}
	return &this
}

// GetPercentage returns the Percentage field value if set, zero value otherwise.
func (o *OfferProgramConfigurationPromotionsDiscountFunding) GetPercentage() float32 {
	if o == nil || IsNil(o.Percentage) {
		var ret float32
		return ret
	}
	return *o.Percentage
}

// GetPercentageOk returns a tuple with the Percentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferProgramConfigurationPromotionsDiscountFunding) GetPercentageOk() (*float32, bool) {
	if o == nil || IsNil(o.Percentage) {
		return nil, false
	}
	return o.Percentage, true
}

// HasPercentage returns a boolean if a field has been set.
func (o *OfferProgramConfigurationPromotionsDiscountFunding) HasPercentage() bool {
	if o != nil && !IsNil(o.Percentage) {
		return true
	}

	return false
}

// SetPercentage gets a reference to the given float32 and assigns it to the Percentage field.
func (o *OfferProgramConfigurationPromotionsDiscountFunding) SetPercentage(v float32) {
	o.Percentage = &v
}

func (o OfferProgramConfigurationPromotionsDiscountFunding) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Percentage) {
		toSerialize["percentage"] = o.Percentage
	}
	return toSerialize, nil
}

type NullableOfferProgramConfigurationPromotionsDiscountFunding struct {
	value *OfferProgramConfigurationPromotionsDiscountFunding
	isSet bool
}

func (v NullableOfferProgramConfigurationPromotionsDiscountFunding) Get() *OfferProgramConfigurationPromotionsDiscountFunding {
	return v.value
}

func (v *NullableOfferProgramConfigurationPromotionsDiscountFunding) Set(val *OfferProgramConfigurationPromotionsDiscountFunding) {
	v.value = val
	v.isSet = true
}

func (v NullableOfferProgramConfigurationPromotionsDiscountFunding) IsSet() bool {
	return v.isSet
}

func (v *NullableOfferProgramConfigurationPromotionsDiscountFunding) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOfferProgramConfigurationPromotionsDiscountFunding(val *OfferProgramConfigurationPromotionsDiscountFunding) *NullableOfferProgramConfigurationPromotionsDiscountFunding {
	return &NullableOfferProgramConfigurationPromotionsDiscountFunding{value: val, isSet: true}
}

func (v NullableOfferProgramConfigurationPromotionsDiscountFunding) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOfferProgramConfigurationPromotionsDiscountFunding) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
