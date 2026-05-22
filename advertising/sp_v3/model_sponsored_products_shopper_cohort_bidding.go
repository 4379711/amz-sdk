package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsShopperCohortBidding type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsShopperCohortBidding{}

// SponsoredProductsShopperCohortBidding struct for SponsoredProductsShopperCohortBidding
type SponsoredProductsShopperCohortBidding struct {
	ShopperCohortType SponsoredProductsShopperCohortType `json:"shopperCohortType"`
	Percentage        *int32                             `json:"percentage,omitempty"`
	// A list of Audience Segments. Shoppers belonging to these segments will be selected for applying the bid adjustments. This is a required field when using \"AUDIENCE_SEGMENT\" option for `shopperCohortType`.
	AudienceSegments []SponsoredProductsAudienceSegment `json:"audienceSegments,omitempty"`
}

type _SponsoredProductsShopperCohortBidding SponsoredProductsShopperCohortBidding

// NewSponsoredProductsShopperCohortBidding instantiates a new SponsoredProductsShopperCohortBidding object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsShopperCohortBidding(shopperCohortType SponsoredProductsShopperCohortType) *SponsoredProductsShopperCohortBidding {
	this := SponsoredProductsShopperCohortBidding{}
	this.ShopperCohortType = shopperCohortType
	return &this
}

// NewSponsoredProductsShopperCohortBiddingWithDefaults instantiates a new SponsoredProductsShopperCohortBidding object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsShopperCohortBiddingWithDefaults() *SponsoredProductsShopperCohortBidding {
	this := SponsoredProductsShopperCohortBidding{}
	return &this
}

// GetShopperCohortType returns the ShopperCohortType field value
func (o *SponsoredProductsShopperCohortBidding) GetShopperCohortType() SponsoredProductsShopperCohortType {
	if o == nil {
		var ret SponsoredProductsShopperCohortType
		return ret
	}

	return o.ShopperCohortType
}

// GetShopperCohortTypeOk returns a tuple with the ShopperCohortType field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsShopperCohortBidding) GetShopperCohortTypeOk() (*SponsoredProductsShopperCohortType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShopperCohortType, true
}

// SetShopperCohortType sets field value
func (o *SponsoredProductsShopperCohortBidding) SetShopperCohortType(v SponsoredProductsShopperCohortType) {
	o.ShopperCohortType = v
}

// GetPercentage returns the Percentage field value if set, zero value otherwise.
func (o *SponsoredProductsShopperCohortBidding) GetPercentage() int32 {
	if o == nil || IsNil(o.Percentage) {
		var ret int32
		return ret
	}
	return *o.Percentage
}

// GetPercentageOk returns a tuple with the Percentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsShopperCohortBidding) GetPercentageOk() (*int32, bool) {
	if o == nil || IsNil(o.Percentage) {
		return nil, false
	}
	return o.Percentage, true
}

// HasPercentage returns a boolean if a field has been set.
func (o *SponsoredProductsShopperCohortBidding) HasPercentage() bool {
	if o != nil && !IsNil(o.Percentage) {
		return true
	}

	return false
}

// SetPercentage gets a reference to the given int32 and assigns it to the Percentage field.
func (o *SponsoredProductsShopperCohortBidding) SetPercentage(v int32) {
	o.Percentage = &v
}

// GetAudienceSegments returns the AudienceSegments field value if set, zero value otherwise.
func (o *SponsoredProductsShopperCohortBidding) GetAudienceSegments() []SponsoredProductsAudienceSegment {
	if o == nil || IsNil(o.AudienceSegments) {
		var ret []SponsoredProductsAudienceSegment
		return ret
	}
	return o.AudienceSegments
}

// GetAudienceSegmentsOk returns a tuple with the AudienceSegments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsShopperCohortBidding) GetAudienceSegmentsOk() ([]SponsoredProductsAudienceSegment, bool) {
	if o == nil || IsNil(o.AudienceSegments) {
		return nil, false
	}
	return o.AudienceSegments, true
}

// HasAudienceSegments returns a boolean if a field has been set.
func (o *SponsoredProductsShopperCohortBidding) HasAudienceSegments() bool {
	if o != nil && !IsNil(o.AudienceSegments) {
		return true
	}

	return false
}

// SetAudienceSegments gets a reference to the given []SponsoredProductsAudienceSegment and assigns it to the AudienceSegments field.
func (o *SponsoredProductsShopperCohortBidding) SetAudienceSegments(v []SponsoredProductsAudienceSegment) {
	o.AudienceSegments = v
}

func (o SponsoredProductsShopperCohortBidding) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["shopperCohortType"] = o.ShopperCohortType
	if !IsNil(o.Percentage) {
		toSerialize["percentage"] = o.Percentage
	}
	if !IsNil(o.AudienceSegments) {
		toSerialize["audienceSegments"] = o.AudienceSegments
	}
	return toSerialize, nil
}

type NullableSponsoredProductsShopperCohortBidding struct {
	value *SponsoredProductsShopperCohortBidding
	isSet bool
}

func (v NullableSponsoredProductsShopperCohortBidding) Get() *SponsoredProductsShopperCohortBidding {
	return v.value
}

func (v *NullableSponsoredProductsShopperCohortBidding) Set(val *SponsoredProductsShopperCohortBidding) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsShopperCohortBidding) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsShopperCohortBidding) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsShopperCohortBidding(val *SponsoredProductsShopperCohortBidding) *NullableSponsoredProductsShopperCohortBidding {
	return &NullableSponsoredProductsShopperCohortBidding{value: val, isSet: true}
}

func (v NullableSponsoredProductsShopperCohortBidding) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsShopperCohortBidding) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
