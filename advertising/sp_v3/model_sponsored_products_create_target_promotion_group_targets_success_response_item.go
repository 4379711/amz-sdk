package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCreateTargetPromotionGroupTargetsSuccessResponseItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateTargetPromotionGroupTargetsSuccessResponseItem{}

// SponsoredProductsCreateTargetPromotionGroupTargetsSuccessResponseItem Response object of successfully created target promotion group target.
type SponsoredProductsCreateTargetPromotionGroupTargetsSuccessResponseItem struct {
	// The expression type of the target that was requested to be created.
	ExpressionType *string                        `json:"expressionType,omitempty"`
	TargetDetails  *SponsoredProductsCreateTarget `json:"targetDetails,omitempty"`
	// The target that was requested to be created.
	Target *string `json:"target,omitempty"`
}

// NewSponsoredProductsCreateTargetPromotionGroupTargetsSuccessResponseItem instantiates a new SponsoredProductsCreateTargetPromotionGroupTargetsSuccessResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateTargetPromotionGroupTargetsSuccessResponseItem() *SponsoredProductsCreateTargetPromotionGroupTargetsSuccessResponseItem {
	this := SponsoredProductsCreateTargetPromotionGroupTargetsSuccessResponseItem{}
	return &this
}

// NewSponsoredProductsCreateTargetPromotionGroupTargetsSuccessResponseItemWithDefaults instantiates a new SponsoredProductsCreateTargetPromotionGroupTargetsSuccessResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateTargetPromotionGroupTargetsSuccessResponseItemWithDefaults() *SponsoredProductsCreateTargetPromotionGroupTargetsSuccessResponseItem {
	this := SponsoredProductsCreateTargetPromotionGroupTargetsSuccessResponseItem{}
	return &this
}

// GetExpressionType returns the ExpressionType field value if set, zero value otherwise.
func (o *SponsoredProductsCreateTargetPromotionGroupTargetsSuccessResponseItem) GetExpressionType() string {
	if o == nil || IsNil(o.ExpressionType) {
		var ret string
		return ret
	}
	return *o.ExpressionType
}

// GetExpressionTypeOk returns a tuple with the ExpressionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateTargetPromotionGroupTargetsSuccessResponseItem) GetExpressionTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ExpressionType) {
		return nil, false
	}
	return o.ExpressionType, true
}

// HasExpressionType returns a boolean if a field has been set.
func (o *SponsoredProductsCreateTargetPromotionGroupTargetsSuccessResponseItem) HasExpressionType() bool {
	if o != nil && !IsNil(o.ExpressionType) {
		return true
	}

	return false
}

// SetExpressionType gets a reference to the given string and assigns it to the ExpressionType field.
func (o *SponsoredProductsCreateTargetPromotionGroupTargetsSuccessResponseItem) SetExpressionType(v string) {
	o.ExpressionType = &v
}

// GetTargetDetails returns the TargetDetails field value if set, zero value otherwise.
func (o *SponsoredProductsCreateTargetPromotionGroupTargetsSuccessResponseItem) GetTargetDetails() SponsoredProductsCreateTarget {
	if o == nil || IsNil(o.TargetDetails) {
		var ret SponsoredProductsCreateTarget
		return ret
	}
	return *o.TargetDetails
}

// GetTargetDetailsOk returns a tuple with the TargetDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateTargetPromotionGroupTargetsSuccessResponseItem) GetTargetDetailsOk() (*SponsoredProductsCreateTarget, bool) {
	if o == nil || IsNil(o.TargetDetails) {
		return nil, false
	}
	return o.TargetDetails, true
}

// HasTargetDetails returns a boolean if a field has been set.
func (o *SponsoredProductsCreateTargetPromotionGroupTargetsSuccessResponseItem) HasTargetDetails() bool {
	if o != nil && !IsNil(o.TargetDetails) {
		return true
	}

	return false
}

// SetTargetDetails gets a reference to the given SponsoredProductsCreateTarget and assigns it to the TargetDetails field.
func (o *SponsoredProductsCreateTargetPromotionGroupTargetsSuccessResponseItem) SetTargetDetails(v SponsoredProductsCreateTarget) {
	o.TargetDetails = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *SponsoredProductsCreateTargetPromotionGroupTargetsSuccessResponseItem) GetTarget() string {
	if o == nil || IsNil(o.Target) {
		var ret string
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateTargetPromotionGroupTargetsSuccessResponseItem) GetTargetOk() (*string, bool) {
	if o == nil || IsNil(o.Target) {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *SponsoredProductsCreateTargetPromotionGroupTargetsSuccessResponseItem) HasTarget() bool {
	if o != nil && !IsNil(o.Target) {
		return true
	}

	return false
}

// SetTarget gets a reference to the given string and assigns it to the Target field.
func (o *SponsoredProductsCreateTargetPromotionGroupTargetsSuccessResponseItem) SetTarget(v string) {
	o.Target = &v
}

func (o SponsoredProductsCreateTargetPromotionGroupTargetsSuccessResponseItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExpressionType) {
		toSerialize["expressionType"] = o.ExpressionType
	}
	if !IsNil(o.TargetDetails) {
		toSerialize["targetDetails"] = o.TargetDetails
	}
	if !IsNil(o.Target) {
		toSerialize["target"] = o.Target
	}
	return toSerialize, nil
}

type NullableSponsoredProductsCreateTargetPromotionGroupTargetsSuccessResponseItem struct {
	value *SponsoredProductsCreateTargetPromotionGroupTargetsSuccessResponseItem
	isSet bool
}

func (v NullableSponsoredProductsCreateTargetPromotionGroupTargetsSuccessResponseItem) Get() *SponsoredProductsCreateTargetPromotionGroupTargetsSuccessResponseItem {
	return v.value
}

func (v *NullableSponsoredProductsCreateTargetPromotionGroupTargetsSuccessResponseItem) Set(val *SponsoredProductsCreateTargetPromotionGroupTargetsSuccessResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateTargetPromotionGroupTargetsSuccessResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateTargetPromotionGroupTargetsSuccessResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateTargetPromotionGroupTargetsSuccessResponseItem(val *SponsoredProductsCreateTargetPromotionGroupTargetsSuccessResponseItem) *NullableSponsoredProductsCreateTargetPromotionGroupTargetsSuccessResponseItem {
	return &NullableSponsoredProductsCreateTargetPromotionGroupTargetsSuccessResponseItem{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateTargetPromotionGroupTargetsSuccessResponseItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateTargetPromotionGroupTargetsSuccessResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
