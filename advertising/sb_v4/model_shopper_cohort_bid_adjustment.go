package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the ShopperCohortBidAdjustment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShopperCohortBidAdjustment{}

// ShopperCohortBidAdjustment struct for ShopperCohortBidAdjustment
type ShopperCohortBidAdjustment struct {
	ShopperCohortType *ShopperCohortType `json:"shopperCohortType,omitempty"`
	Percentage        *float32           `json:"percentage,omitempty"`
	// Required when \"AUDIENCE_SEGMENT\" is used for shopperCohortType.
	AudienceSegments []AudienceSegment `json:"audienceSegments,omitempty"`
}

// NewShopperCohortBidAdjustment instantiates a new ShopperCohortBidAdjustment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShopperCohortBidAdjustment() *ShopperCohortBidAdjustment {
	this := ShopperCohortBidAdjustment{}
	return &this
}

// NewShopperCohortBidAdjustmentWithDefaults instantiates a new ShopperCohortBidAdjustment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShopperCohortBidAdjustmentWithDefaults() *ShopperCohortBidAdjustment {
	this := ShopperCohortBidAdjustment{}
	return &this
}

// GetShopperCohortType returns the ShopperCohortType field value if set, zero value otherwise.
func (o *ShopperCohortBidAdjustment) GetShopperCohortType() ShopperCohortType {
	if o == nil || IsNil(o.ShopperCohortType) {
		var ret ShopperCohortType
		return ret
	}
	return *o.ShopperCohortType
}

// GetShopperCohortTypeOk returns a tuple with the ShopperCohortType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShopperCohortBidAdjustment) GetShopperCohortTypeOk() (*ShopperCohortType, bool) {
	if o == nil || IsNil(o.ShopperCohortType) {
		return nil, false
	}
	return o.ShopperCohortType, true
}

// HasShopperCohortType returns a boolean if a field has been set.
func (o *ShopperCohortBidAdjustment) HasShopperCohortType() bool {
	if o != nil && !IsNil(o.ShopperCohortType) {
		return true
	}

	return false
}

// SetShopperCohortType gets a reference to the given ShopperCohortType and assigns it to the ShopperCohortType field.
func (o *ShopperCohortBidAdjustment) SetShopperCohortType(v ShopperCohortType) {
	o.ShopperCohortType = &v
}

// GetPercentage returns the Percentage field value if set, zero value otherwise.
func (o *ShopperCohortBidAdjustment) GetPercentage() float32 {
	if o == nil || IsNil(o.Percentage) {
		var ret float32
		return ret
	}
	return *o.Percentage
}

// GetPercentageOk returns a tuple with the Percentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShopperCohortBidAdjustment) GetPercentageOk() (*float32, bool) {
	if o == nil || IsNil(o.Percentage) {
		return nil, false
	}
	return o.Percentage, true
}

// HasPercentage returns a boolean if a field has been set.
func (o *ShopperCohortBidAdjustment) HasPercentage() bool {
	if o != nil && !IsNil(o.Percentage) {
		return true
	}

	return false
}

// SetPercentage gets a reference to the given float32 and assigns it to the Percentage field.
func (o *ShopperCohortBidAdjustment) SetPercentage(v float32) {
	o.Percentage = &v
}

// GetAudienceSegments returns the AudienceSegments field value if set, zero value otherwise.
func (o *ShopperCohortBidAdjustment) GetAudienceSegments() []AudienceSegment {
	if o == nil || IsNil(o.AudienceSegments) {
		var ret []AudienceSegment
		return ret
	}
	return o.AudienceSegments
}

// GetAudienceSegmentsOk returns a tuple with the AudienceSegments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShopperCohortBidAdjustment) GetAudienceSegmentsOk() ([]AudienceSegment, bool) {
	if o == nil || IsNil(o.AudienceSegments) {
		return nil, false
	}
	return o.AudienceSegments, true
}

// HasAudienceSegments returns a boolean if a field has been set.
func (o *ShopperCohortBidAdjustment) HasAudienceSegments() bool {
	if o != nil && !IsNil(o.AudienceSegments) {
		return true
	}

	return false
}

// SetAudienceSegments gets a reference to the given []AudienceSegment and assigns it to the AudienceSegments field.
func (o *ShopperCohortBidAdjustment) SetAudienceSegments(v []AudienceSegment) {
	o.AudienceSegments = v
}

func (o ShopperCohortBidAdjustment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ShopperCohortType) {
		toSerialize["shopperCohortType"] = o.ShopperCohortType
	}
	if !IsNil(o.Percentage) {
		toSerialize["percentage"] = o.Percentage
	}
	if !IsNil(o.AudienceSegments) {
		toSerialize["audienceSegments"] = o.AudienceSegments
	}
	return toSerialize, nil
}

type NullableShopperCohortBidAdjustment struct {
	value *ShopperCohortBidAdjustment
	isSet bool
}

func (v NullableShopperCohortBidAdjustment) Get() *ShopperCohortBidAdjustment {
	return v.value
}

func (v *NullableShopperCohortBidAdjustment) Set(val *ShopperCohortBidAdjustment) {
	v.value = val
	v.isSet = true
}

func (v NullableShopperCohortBidAdjustment) IsSet() bool {
	return v.isSet
}

func (v *NullableShopperCohortBidAdjustment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShopperCohortBidAdjustment(val *ShopperCohortBidAdjustment) *NullableShopperCohortBidAdjustment {
	return &NullableShopperCohortBidAdjustment{value: val, isSet: true}
}

func (v NullableShopperCohortBidAdjustment) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableShopperCohortBidAdjustment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
