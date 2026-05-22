package product_pricing_20220501

import (
	"github.com/bytedance/sonic"
)

// checks if the FeaturedOfferSegment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FeaturedOfferSegment{}

// FeaturedOfferSegment Describes the segment in which the offer is featured.
type FeaturedOfferSegment struct {
	// The customer membership type that makes up this segment
	CustomerMembership string         `json:"customerMembership"`
	SegmentDetails     SegmentDetails `json:"segmentDetails"`
}

type _FeaturedOfferSegment FeaturedOfferSegment

// NewFeaturedOfferSegment instantiates a new FeaturedOfferSegment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeaturedOfferSegment(customerMembership string, segmentDetails SegmentDetails) *FeaturedOfferSegment {
	this := FeaturedOfferSegment{}
	this.CustomerMembership = customerMembership
	this.SegmentDetails = segmentDetails
	return &this
}

// NewFeaturedOfferSegmentWithDefaults instantiates a new FeaturedOfferSegment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeaturedOfferSegmentWithDefaults() *FeaturedOfferSegment {
	this := FeaturedOfferSegment{}
	return &this
}

// GetCustomerMembership returns the CustomerMembership field value
func (o *FeaturedOfferSegment) GetCustomerMembership() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerMembership
}

// GetCustomerMembershipOk returns a tuple with the CustomerMembership field value
// and a boolean to check if the value has been set.
func (o *FeaturedOfferSegment) GetCustomerMembershipOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerMembership, true
}

// SetCustomerMembership sets field value
func (o *FeaturedOfferSegment) SetCustomerMembership(v string) {
	o.CustomerMembership = v
}

// GetSegmentDetails returns the SegmentDetails field value
func (o *FeaturedOfferSegment) GetSegmentDetails() SegmentDetails {
	if o == nil {
		var ret SegmentDetails
		return ret
	}

	return o.SegmentDetails
}

// GetSegmentDetailsOk returns a tuple with the SegmentDetails field value
// and a boolean to check if the value has been set.
func (o *FeaturedOfferSegment) GetSegmentDetailsOk() (*SegmentDetails, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SegmentDetails, true
}

// SetSegmentDetails sets field value
func (o *FeaturedOfferSegment) SetSegmentDetails(v SegmentDetails) {
	o.SegmentDetails = v
}

func (o FeaturedOfferSegment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["customerMembership"] = o.CustomerMembership
	toSerialize["segmentDetails"] = o.SegmentDetails
	return toSerialize, nil
}

type NullableFeaturedOfferSegment struct {
	value *FeaturedOfferSegment
	isSet bool
}

func (v NullableFeaturedOfferSegment) Get() *FeaturedOfferSegment {
	return v.value
}

func (v *NullableFeaturedOfferSegment) Set(val *FeaturedOfferSegment) {
	v.value = val
	v.isSet = true
}

func (v NullableFeaturedOfferSegment) IsSet() bool {
	return v.isSet
}

func (v *NullableFeaturedOfferSegment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeaturedOfferSegment(val *FeaturedOfferSegment) *NullableFeaturedOfferSegment {
	return &NullableFeaturedOfferSegment{value: val, isSet: true}
}

func (v NullableFeaturedOfferSegment) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableFeaturedOfferSegment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
