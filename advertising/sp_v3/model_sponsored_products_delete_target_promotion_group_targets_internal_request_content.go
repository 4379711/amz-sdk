package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDeleteTargetPromotionGroupTargetsInternalRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDeleteTargetPromotionGroupTargetsInternalRequestContent{}

// SponsoredProductsDeleteTargetPromotionGroupTargetsInternalRequestContent struct for SponsoredProductsDeleteTargetPromotionGroupTargetsInternalRequestContent
type SponsoredProductsDeleteTargetPromotionGroupTargetsInternalRequestContent struct {
	ApiGatewayContext           SponsoredProductsApiGatewayContext     `json:"apiGatewayContext"`
	TargetPromotionGroupTargets []SponsoredProductsDeleteTargetRequest `json:"targetPromotionGroupTargets,omitempty"`
}

type _SponsoredProductsDeleteTargetPromotionGroupTargetsInternalRequestContent SponsoredProductsDeleteTargetPromotionGroupTargetsInternalRequestContent

// NewSponsoredProductsDeleteTargetPromotionGroupTargetsInternalRequestContent instantiates a new SponsoredProductsDeleteTargetPromotionGroupTargetsInternalRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDeleteTargetPromotionGroupTargetsInternalRequestContent(apiGatewayContext SponsoredProductsApiGatewayContext) *SponsoredProductsDeleteTargetPromotionGroupTargetsInternalRequestContent {
	this := SponsoredProductsDeleteTargetPromotionGroupTargetsInternalRequestContent{}
	this.ApiGatewayContext = apiGatewayContext
	return &this
}

// NewSponsoredProductsDeleteTargetPromotionGroupTargetsInternalRequestContentWithDefaults instantiates a new SponsoredProductsDeleteTargetPromotionGroupTargetsInternalRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDeleteTargetPromotionGroupTargetsInternalRequestContentWithDefaults() *SponsoredProductsDeleteTargetPromotionGroupTargetsInternalRequestContent {
	this := SponsoredProductsDeleteTargetPromotionGroupTargetsInternalRequestContent{}
	return &this
}

// GetApiGatewayContext returns the ApiGatewayContext field value
func (o *SponsoredProductsDeleteTargetPromotionGroupTargetsInternalRequestContent) GetApiGatewayContext() SponsoredProductsApiGatewayContext {
	if o == nil {
		var ret SponsoredProductsApiGatewayContext
		return ret
	}

	return o.ApiGatewayContext
}

// GetApiGatewayContextOk returns a tuple with the ApiGatewayContext field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDeleteTargetPromotionGroupTargetsInternalRequestContent) GetApiGatewayContextOk() (*SponsoredProductsApiGatewayContext, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiGatewayContext, true
}

// SetApiGatewayContext sets field value
func (o *SponsoredProductsDeleteTargetPromotionGroupTargetsInternalRequestContent) SetApiGatewayContext(v SponsoredProductsApiGatewayContext) {
	o.ApiGatewayContext = v
}

// GetTargetPromotionGroupTargets returns the TargetPromotionGroupTargets field value if set, zero value otherwise.
func (o *SponsoredProductsDeleteTargetPromotionGroupTargetsInternalRequestContent) GetTargetPromotionGroupTargets() []SponsoredProductsDeleteTargetRequest {
	if o == nil || IsNil(o.TargetPromotionGroupTargets) {
		var ret []SponsoredProductsDeleteTargetRequest
		return ret
	}
	return o.TargetPromotionGroupTargets
}

// GetTargetPromotionGroupTargetsOk returns a tuple with the TargetPromotionGroupTargets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDeleteTargetPromotionGroupTargetsInternalRequestContent) GetTargetPromotionGroupTargetsOk() ([]SponsoredProductsDeleteTargetRequest, bool) {
	if o == nil || IsNil(o.TargetPromotionGroupTargets) {
		return nil, false
	}
	return o.TargetPromotionGroupTargets, true
}

// HasTargetPromotionGroupTargets returns a boolean if a field has been set.
func (o *SponsoredProductsDeleteTargetPromotionGroupTargetsInternalRequestContent) HasTargetPromotionGroupTargets() bool {
	if o != nil && !IsNil(o.TargetPromotionGroupTargets) {
		return true
	}

	return false
}

// SetTargetPromotionGroupTargets gets a reference to the given []SponsoredProductsDeleteTargetRequest and assigns it to the TargetPromotionGroupTargets field.
func (o *SponsoredProductsDeleteTargetPromotionGroupTargetsInternalRequestContent) SetTargetPromotionGroupTargets(v []SponsoredProductsDeleteTargetRequest) {
	o.TargetPromotionGroupTargets = v
}

func (o SponsoredProductsDeleteTargetPromotionGroupTargetsInternalRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["apiGatewayContext"] = o.ApiGatewayContext
	if !IsNil(o.TargetPromotionGroupTargets) {
		toSerialize["targetPromotionGroupTargets"] = o.TargetPromotionGroupTargets
	}
	return toSerialize, nil
}

type NullableSponsoredProductsDeleteTargetPromotionGroupTargetsInternalRequestContent struct {
	value *SponsoredProductsDeleteTargetPromotionGroupTargetsInternalRequestContent
	isSet bool
}

func (v NullableSponsoredProductsDeleteTargetPromotionGroupTargetsInternalRequestContent) Get() *SponsoredProductsDeleteTargetPromotionGroupTargetsInternalRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsDeleteTargetPromotionGroupTargetsInternalRequestContent) Set(val *SponsoredProductsDeleteTargetPromotionGroupTargetsInternalRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDeleteTargetPromotionGroupTargetsInternalRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDeleteTargetPromotionGroupTargetsInternalRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDeleteTargetPromotionGroupTargetsInternalRequestContent(val *SponsoredProductsDeleteTargetPromotionGroupTargetsInternalRequestContent) *NullableSponsoredProductsDeleteTargetPromotionGroupTargetsInternalRequestContent {
	return &NullableSponsoredProductsDeleteTargetPromotionGroupTargetsInternalRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsDeleteTargetPromotionGroupTargetsInternalRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDeleteTargetPromotionGroupTargetsInternalRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
