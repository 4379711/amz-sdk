package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDeleteTargetPromotionGroupsInternalRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDeleteTargetPromotionGroupsInternalRequestContent{}

// SponsoredProductsDeleteTargetPromotionGroupsInternalRequestContent struct for SponsoredProductsDeleteTargetPromotionGroupsInternalRequestContent
type SponsoredProductsDeleteTargetPromotionGroupsInternalRequestContent struct {
	ApiGatewayContext     SponsoredProductsApiGatewayContext      `json:"apiGatewayContext"`
	TargetPromotionGroups []SponsoredProductsTargetPromotionGroup `json:"targetPromotionGroups,omitempty"`
}

type _SponsoredProductsDeleteTargetPromotionGroupsInternalRequestContent SponsoredProductsDeleteTargetPromotionGroupsInternalRequestContent

// NewSponsoredProductsDeleteTargetPromotionGroupsInternalRequestContent instantiates a new SponsoredProductsDeleteTargetPromotionGroupsInternalRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDeleteTargetPromotionGroupsInternalRequestContent(apiGatewayContext SponsoredProductsApiGatewayContext) *SponsoredProductsDeleteTargetPromotionGroupsInternalRequestContent {
	this := SponsoredProductsDeleteTargetPromotionGroupsInternalRequestContent{}
	this.ApiGatewayContext = apiGatewayContext
	return &this
}

// NewSponsoredProductsDeleteTargetPromotionGroupsInternalRequestContentWithDefaults instantiates a new SponsoredProductsDeleteTargetPromotionGroupsInternalRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDeleteTargetPromotionGroupsInternalRequestContentWithDefaults() *SponsoredProductsDeleteTargetPromotionGroupsInternalRequestContent {
	this := SponsoredProductsDeleteTargetPromotionGroupsInternalRequestContent{}
	return &this
}

// GetApiGatewayContext returns the ApiGatewayContext field value
func (o *SponsoredProductsDeleteTargetPromotionGroupsInternalRequestContent) GetApiGatewayContext() SponsoredProductsApiGatewayContext {
	if o == nil {
		var ret SponsoredProductsApiGatewayContext
		return ret
	}

	return o.ApiGatewayContext
}

// GetApiGatewayContextOk returns a tuple with the ApiGatewayContext field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDeleteTargetPromotionGroupsInternalRequestContent) GetApiGatewayContextOk() (*SponsoredProductsApiGatewayContext, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiGatewayContext, true
}

// SetApiGatewayContext sets field value
func (o *SponsoredProductsDeleteTargetPromotionGroupsInternalRequestContent) SetApiGatewayContext(v SponsoredProductsApiGatewayContext) {
	o.ApiGatewayContext = v
}

// GetTargetPromotionGroups returns the TargetPromotionGroups field value if set, zero value otherwise.
func (o *SponsoredProductsDeleteTargetPromotionGroupsInternalRequestContent) GetTargetPromotionGroups() []SponsoredProductsTargetPromotionGroup {
	if o == nil || IsNil(o.TargetPromotionGroups) {
		var ret []SponsoredProductsTargetPromotionGroup
		return ret
	}
	return o.TargetPromotionGroups
}

// GetTargetPromotionGroupsOk returns a tuple with the TargetPromotionGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDeleteTargetPromotionGroupsInternalRequestContent) GetTargetPromotionGroupsOk() ([]SponsoredProductsTargetPromotionGroup, bool) {
	if o == nil || IsNil(o.TargetPromotionGroups) {
		return nil, false
	}
	return o.TargetPromotionGroups, true
}

// HasTargetPromotionGroups returns a boolean if a field has been set.
func (o *SponsoredProductsDeleteTargetPromotionGroupsInternalRequestContent) HasTargetPromotionGroups() bool {
	if o != nil && !IsNil(o.TargetPromotionGroups) {
		return true
	}

	return false
}

// SetTargetPromotionGroups gets a reference to the given []SponsoredProductsTargetPromotionGroup and assigns it to the TargetPromotionGroups field.
func (o *SponsoredProductsDeleteTargetPromotionGroupsInternalRequestContent) SetTargetPromotionGroups(v []SponsoredProductsTargetPromotionGroup) {
	o.TargetPromotionGroups = v
}

func (o SponsoredProductsDeleteTargetPromotionGroupsInternalRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["apiGatewayContext"] = o.ApiGatewayContext
	if !IsNil(o.TargetPromotionGroups) {
		toSerialize["targetPromotionGroups"] = o.TargetPromotionGroups
	}
	return toSerialize, nil
}

type NullableSponsoredProductsDeleteTargetPromotionGroupsInternalRequestContent struct {
	value *SponsoredProductsDeleteTargetPromotionGroupsInternalRequestContent
	isSet bool
}

func (v NullableSponsoredProductsDeleteTargetPromotionGroupsInternalRequestContent) Get() *SponsoredProductsDeleteTargetPromotionGroupsInternalRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsDeleteTargetPromotionGroupsInternalRequestContent) Set(val *SponsoredProductsDeleteTargetPromotionGroupsInternalRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDeleteTargetPromotionGroupsInternalRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDeleteTargetPromotionGroupsInternalRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDeleteTargetPromotionGroupsInternalRequestContent(val *SponsoredProductsDeleteTargetPromotionGroupsInternalRequestContent) *NullableSponsoredProductsDeleteTargetPromotionGroupsInternalRequestContent {
	return &NullableSponsoredProductsDeleteTargetPromotionGroupsInternalRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsDeleteTargetPromotionGroupsInternalRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDeleteTargetPromotionGroupsInternalRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
