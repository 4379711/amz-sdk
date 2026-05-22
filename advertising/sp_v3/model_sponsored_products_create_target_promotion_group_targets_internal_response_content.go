package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCreateTargetPromotionGroupTargetsInternalResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateTargetPromotionGroupTargetsInternalResponseContent{}

// SponsoredProductsCreateTargetPromotionGroupTargetsInternalResponseContent Response object for creating target promotion group targets.
type SponsoredProductsCreateTargetPromotionGroupTargetsInternalResponseContent struct {
	// List of successfully created targets.
	Success []SponsoredProductsCreateTargetPromotionGroupTargetsSuccessResponseItem `json:"success,omitempty"`
	// List of targets that failed to create.
	Errors []SponsoredProductsCreateTargetPromotionGroupTargetsFailureResponseItem `json:"errors,omitempty"`
}

// NewSponsoredProductsCreateTargetPromotionGroupTargetsInternalResponseContent instantiates a new SponsoredProductsCreateTargetPromotionGroupTargetsInternalResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateTargetPromotionGroupTargetsInternalResponseContent() *SponsoredProductsCreateTargetPromotionGroupTargetsInternalResponseContent {
	this := SponsoredProductsCreateTargetPromotionGroupTargetsInternalResponseContent{}
	return &this
}

// NewSponsoredProductsCreateTargetPromotionGroupTargetsInternalResponseContentWithDefaults instantiates a new SponsoredProductsCreateTargetPromotionGroupTargetsInternalResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateTargetPromotionGroupTargetsInternalResponseContentWithDefaults() *SponsoredProductsCreateTargetPromotionGroupTargetsInternalResponseContent {
	this := SponsoredProductsCreateTargetPromotionGroupTargetsInternalResponseContent{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *SponsoredProductsCreateTargetPromotionGroupTargetsInternalResponseContent) GetSuccess() []SponsoredProductsCreateTargetPromotionGroupTargetsSuccessResponseItem {
	if o == nil || IsNil(o.Success) {
		var ret []SponsoredProductsCreateTargetPromotionGroupTargetsSuccessResponseItem
		return ret
	}
	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateTargetPromotionGroupTargetsInternalResponseContent) GetSuccessOk() ([]SponsoredProductsCreateTargetPromotionGroupTargetsSuccessResponseItem, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *SponsoredProductsCreateTargetPromotionGroupTargetsInternalResponseContent) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given []SponsoredProductsCreateTargetPromotionGroupTargetsSuccessResponseItem and assigns it to the Success field.
func (o *SponsoredProductsCreateTargetPromotionGroupTargetsInternalResponseContent) SetSuccess(v []SponsoredProductsCreateTargetPromotionGroupTargetsSuccessResponseItem) {
	o.Success = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *SponsoredProductsCreateTargetPromotionGroupTargetsInternalResponseContent) GetErrors() []SponsoredProductsCreateTargetPromotionGroupTargetsFailureResponseItem {
	if o == nil || IsNil(o.Errors) {
		var ret []SponsoredProductsCreateTargetPromotionGroupTargetsFailureResponseItem
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateTargetPromotionGroupTargetsInternalResponseContent) GetErrorsOk() ([]SponsoredProductsCreateTargetPromotionGroupTargetsFailureResponseItem, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *SponsoredProductsCreateTargetPromotionGroupTargetsInternalResponseContent) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []SponsoredProductsCreateTargetPromotionGroupTargetsFailureResponseItem and assigns it to the Errors field.
func (o *SponsoredProductsCreateTargetPromotionGroupTargetsInternalResponseContent) SetErrors(v []SponsoredProductsCreateTargetPromotionGroupTargetsFailureResponseItem) {
	o.Errors = v
}

func (o SponsoredProductsCreateTargetPromotionGroupTargetsInternalResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableSponsoredProductsCreateTargetPromotionGroupTargetsInternalResponseContent struct {
	value *SponsoredProductsCreateTargetPromotionGroupTargetsInternalResponseContent
	isSet bool
}

func (v NullableSponsoredProductsCreateTargetPromotionGroupTargetsInternalResponseContent) Get() *SponsoredProductsCreateTargetPromotionGroupTargetsInternalResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsCreateTargetPromotionGroupTargetsInternalResponseContent) Set(val *SponsoredProductsCreateTargetPromotionGroupTargetsInternalResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateTargetPromotionGroupTargetsInternalResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateTargetPromotionGroupTargetsInternalResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateTargetPromotionGroupTargetsInternalResponseContent(val *SponsoredProductsCreateTargetPromotionGroupTargetsInternalResponseContent) *NullableSponsoredProductsCreateTargetPromotionGroupTargetsInternalResponseContent {
	return &NullableSponsoredProductsCreateTargetPromotionGroupTargetsInternalResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateTargetPromotionGroupTargetsInternalResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateTargetPromotionGroupTargetsInternalResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
