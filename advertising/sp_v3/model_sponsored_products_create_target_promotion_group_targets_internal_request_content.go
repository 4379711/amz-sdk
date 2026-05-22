package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCreateTargetPromotionGroupTargetsInternalRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateTargetPromotionGroupTargetsInternalRequestContent{}

// SponsoredProductsCreateTargetPromotionGroupTargetsInternalRequestContent struct for SponsoredProductsCreateTargetPromotionGroupTargetsInternalRequestContent
type SponsoredProductsCreateTargetPromotionGroupTargetsInternalRequestContent struct {
	// Entity object identifier
	TargetPromotionGroupId string                                 `json:"targetPromotionGroupId"`
	ApiGatewayContext      SponsoredProductsApiGatewayContext     `json:"apiGatewayContext"`
	Targets                []SponsoredProductsCreateTargetRequest `json:"targets,omitempty"`
}

type _SponsoredProductsCreateTargetPromotionGroupTargetsInternalRequestContent SponsoredProductsCreateTargetPromotionGroupTargetsInternalRequestContent

// NewSponsoredProductsCreateTargetPromotionGroupTargetsInternalRequestContent instantiates a new SponsoredProductsCreateTargetPromotionGroupTargetsInternalRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateTargetPromotionGroupTargetsInternalRequestContent(targetPromotionGroupId string, apiGatewayContext SponsoredProductsApiGatewayContext) *SponsoredProductsCreateTargetPromotionGroupTargetsInternalRequestContent {
	this := SponsoredProductsCreateTargetPromotionGroupTargetsInternalRequestContent{}
	this.TargetPromotionGroupId = targetPromotionGroupId
	this.ApiGatewayContext = apiGatewayContext
	return &this
}

// NewSponsoredProductsCreateTargetPromotionGroupTargetsInternalRequestContentWithDefaults instantiates a new SponsoredProductsCreateTargetPromotionGroupTargetsInternalRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateTargetPromotionGroupTargetsInternalRequestContentWithDefaults() *SponsoredProductsCreateTargetPromotionGroupTargetsInternalRequestContent {
	this := SponsoredProductsCreateTargetPromotionGroupTargetsInternalRequestContent{}
	return &this
}

// GetTargetPromotionGroupId returns the TargetPromotionGroupId field value
func (o *SponsoredProductsCreateTargetPromotionGroupTargetsInternalRequestContent) GetTargetPromotionGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetPromotionGroupId
}

// GetTargetPromotionGroupIdOk returns a tuple with the TargetPromotionGroupId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateTargetPromotionGroupTargetsInternalRequestContent) GetTargetPromotionGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetPromotionGroupId, true
}

// SetTargetPromotionGroupId sets field value
func (o *SponsoredProductsCreateTargetPromotionGroupTargetsInternalRequestContent) SetTargetPromotionGroupId(v string) {
	o.TargetPromotionGroupId = v
}

// GetApiGatewayContext returns the ApiGatewayContext field value
func (o *SponsoredProductsCreateTargetPromotionGroupTargetsInternalRequestContent) GetApiGatewayContext() SponsoredProductsApiGatewayContext {
	if o == nil {
		var ret SponsoredProductsApiGatewayContext
		return ret
	}

	return o.ApiGatewayContext
}

// GetApiGatewayContextOk returns a tuple with the ApiGatewayContext field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateTargetPromotionGroupTargetsInternalRequestContent) GetApiGatewayContextOk() (*SponsoredProductsApiGatewayContext, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiGatewayContext, true
}

// SetApiGatewayContext sets field value
func (o *SponsoredProductsCreateTargetPromotionGroupTargetsInternalRequestContent) SetApiGatewayContext(v SponsoredProductsApiGatewayContext) {
	o.ApiGatewayContext = v
}

// GetTargets returns the Targets field value if set, zero value otherwise.
func (o *SponsoredProductsCreateTargetPromotionGroupTargetsInternalRequestContent) GetTargets() []SponsoredProductsCreateTargetRequest {
	if o == nil || IsNil(o.Targets) {
		var ret []SponsoredProductsCreateTargetRequest
		return ret
	}
	return o.Targets
}

// GetTargetsOk returns a tuple with the Targets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateTargetPromotionGroupTargetsInternalRequestContent) GetTargetsOk() ([]SponsoredProductsCreateTargetRequest, bool) {
	if o == nil || IsNil(o.Targets) {
		return nil, false
	}
	return o.Targets, true
}

// HasTargets returns a boolean if a field has been set.
func (o *SponsoredProductsCreateTargetPromotionGroupTargetsInternalRequestContent) HasTargets() bool {
	if o != nil && !IsNil(o.Targets) {
		return true
	}

	return false
}

// SetTargets gets a reference to the given []SponsoredProductsCreateTargetRequest and assigns it to the Targets field.
func (o *SponsoredProductsCreateTargetPromotionGroupTargetsInternalRequestContent) SetTargets(v []SponsoredProductsCreateTargetRequest) {
	o.Targets = v
}

func (o SponsoredProductsCreateTargetPromotionGroupTargetsInternalRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["targetPromotionGroupId"] = o.TargetPromotionGroupId
	toSerialize["apiGatewayContext"] = o.ApiGatewayContext
	if !IsNil(o.Targets) {
		toSerialize["targets"] = o.Targets
	}
	return toSerialize, nil
}

type NullableSponsoredProductsCreateTargetPromotionGroupTargetsInternalRequestContent struct {
	value *SponsoredProductsCreateTargetPromotionGroupTargetsInternalRequestContent
	isSet bool
}

func (v NullableSponsoredProductsCreateTargetPromotionGroupTargetsInternalRequestContent) Get() *SponsoredProductsCreateTargetPromotionGroupTargetsInternalRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsCreateTargetPromotionGroupTargetsInternalRequestContent) Set(val *SponsoredProductsCreateTargetPromotionGroupTargetsInternalRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateTargetPromotionGroupTargetsInternalRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateTargetPromotionGroupTargetsInternalRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateTargetPromotionGroupTargetsInternalRequestContent(val *SponsoredProductsCreateTargetPromotionGroupTargetsInternalRequestContent) *NullableSponsoredProductsCreateTargetPromotionGroupTargetsInternalRequestContent {
	return &NullableSponsoredProductsCreateTargetPromotionGroupTargetsInternalRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateTargetPromotionGroupTargetsInternalRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateTargetPromotionGroupTargetsInternalRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
