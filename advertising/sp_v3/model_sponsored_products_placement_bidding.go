package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsPlacementBidding type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsPlacementBidding{}

// SponsoredProductsPlacementBidding struct for SponsoredProductsPlacementBidding
type SponsoredProductsPlacementBidding struct {
	Percentage *int32                      `json:"percentage,omitempty"`
	Placement  *SponsoredProductsPlacement `json:"placement,omitempty"`
}

// NewSponsoredProductsPlacementBidding instantiates a new SponsoredProductsPlacementBidding object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsPlacementBidding() *SponsoredProductsPlacementBidding {
	this := SponsoredProductsPlacementBidding{}
	return &this
}

// NewSponsoredProductsPlacementBiddingWithDefaults instantiates a new SponsoredProductsPlacementBidding object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsPlacementBiddingWithDefaults() *SponsoredProductsPlacementBidding {
	this := SponsoredProductsPlacementBidding{}
	return &this
}

// GetPercentage returns the Percentage field value if set, zero value otherwise.
func (o *SponsoredProductsPlacementBidding) GetPercentage() int32 {
	if o == nil || IsNil(o.Percentage) {
		var ret int32
		return ret
	}
	return *o.Percentage
}

// GetPercentageOk returns a tuple with the Percentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsPlacementBidding) GetPercentageOk() (*int32, bool) {
	if o == nil || IsNil(o.Percentage) {
		return nil, false
	}
	return o.Percentage, true
}

// HasPercentage returns a boolean if a field has been set.
func (o *SponsoredProductsPlacementBidding) HasPercentage() bool {
	if o != nil && !IsNil(o.Percentage) {
		return true
	}

	return false
}

// SetPercentage gets a reference to the given int32 and assigns it to the Percentage field.
func (o *SponsoredProductsPlacementBidding) SetPercentage(v int32) {
	o.Percentage = &v
}

// GetPlacement returns the Placement field value if set, zero value otherwise.
func (o *SponsoredProductsPlacementBidding) GetPlacement() SponsoredProductsPlacement {
	if o == nil || IsNil(o.Placement) {
		var ret SponsoredProductsPlacement
		return ret
	}
	return *o.Placement
}

// GetPlacementOk returns a tuple with the Placement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsPlacementBidding) GetPlacementOk() (*SponsoredProductsPlacement, bool) {
	if o == nil || IsNil(o.Placement) {
		return nil, false
	}
	return o.Placement, true
}

// HasPlacement returns a boolean if a field has been set.
func (o *SponsoredProductsPlacementBidding) HasPlacement() bool {
	if o != nil && !IsNil(o.Placement) {
		return true
	}

	return false
}

// SetPlacement gets a reference to the given SponsoredProductsPlacement and assigns it to the Placement field.
func (o *SponsoredProductsPlacementBidding) SetPlacement(v SponsoredProductsPlacement) {
	o.Placement = &v
}

func (o SponsoredProductsPlacementBidding) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Percentage) {
		toSerialize["percentage"] = o.Percentage
	}
	if !IsNil(o.Placement) {
		toSerialize["placement"] = o.Placement
	}
	return toSerialize, nil
}

type NullableSponsoredProductsPlacementBidding struct {
	value *SponsoredProductsPlacementBidding
	isSet bool
}

func (v NullableSponsoredProductsPlacementBidding) Get() *SponsoredProductsPlacementBidding {
	return v.value
}

func (v *NullableSponsoredProductsPlacementBidding) Set(val *SponsoredProductsPlacementBidding) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsPlacementBidding) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsPlacementBidding) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsPlacementBidding(val *SponsoredProductsPlacementBidding) *NullableSponsoredProductsPlacementBidding {
	return &NullableSponsoredProductsPlacementBidding{value: val, isSet: true}
}

func (v NullableSponsoredProductsPlacementBidding) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsPlacementBidding) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
