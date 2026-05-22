package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCreateTargetPromotionGroupTargetsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateTargetPromotionGroupTargetsResponseContent{}

// SponsoredProductsCreateTargetPromotionGroupTargetsResponseContent Response object for creating target promotion group targets.
type SponsoredProductsCreateTargetPromotionGroupTargetsResponseContent struct {
	// List of successfully created targets.
	Success []SponsoredProductsCreateTargetPromotionGroupTargetsSuccessResponseItem `json:"success,omitempty"`
	// List of targets that failed to create.
	Errors []SponsoredProductsCreateTargetPromotionGroupTargetsFailureResponseItem `json:"errors,omitempty"`
}

// NewSponsoredProductsCreateTargetPromotionGroupTargetsResponseContent instantiates a new SponsoredProductsCreateTargetPromotionGroupTargetsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateTargetPromotionGroupTargetsResponseContent() *SponsoredProductsCreateTargetPromotionGroupTargetsResponseContent {
	this := SponsoredProductsCreateTargetPromotionGroupTargetsResponseContent{}
	return &this
}

// NewSponsoredProductsCreateTargetPromotionGroupTargetsResponseContentWithDefaults instantiates a new SponsoredProductsCreateTargetPromotionGroupTargetsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateTargetPromotionGroupTargetsResponseContentWithDefaults() *SponsoredProductsCreateTargetPromotionGroupTargetsResponseContent {
	this := SponsoredProductsCreateTargetPromotionGroupTargetsResponseContent{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *SponsoredProductsCreateTargetPromotionGroupTargetsResponseContent) GetSuccess() []SponsoredProductsCreateTargetPromotionGroupTargetsSuccessResponseItem {
	if o == nil || IsNil(o.Success) {
		var ret []SponsoredProductsCreateTargetPromotionGroupTargetsSuccessResponseItem
		return ret
	}
	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateTargetPromotionGroupTargetsResponseContent) GetSuccessOk() ([]SponsoredProductsCreateTargetPromotionGroupTargetsSuccessResponseItem, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *SponsoredProductsCreateTargetPromotionGroupTargetsResponseContent) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given []SponsoredProductsCreateTargetPromotionGroupTargetsSuccessResponseItem and assigns it to the Success field.
func (o *SponsoredProductsCreateTargetPromotionGroupTargetsResponseContent) SetSuccess(v []SponsoredProductsCreateTargetPromotionGroupTargetsSuccessResponseItem) {
	o.Success = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *SponsoredProductsCreateTargetPromotionGroupTargetsResponseContent) GetErrors() []SponsoredProductsCreateTargetPromotionGroupTargetsFailureResponseItem {
	if o == nil || IsNil(o.Errors) {
		var ret []SponsoredProductsCreateTargetPromotionGroupTargetsFailureResponseItem
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateTargetPromotionGroupTargetsResponseContent) GetErrorsOk() ([]SponsoredProductsCreateTargetPromotionGroupTargetsFailureResponseItem, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *SponsoredProductsCreateTargetPromotionGroupTargetsResponseContent) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []SponsoredProductsCreateTargetPromotionGroupTargetsFailureResponseItem and assigns it to the Errors field.
func (o *SponsoredProductsCreateTargetPromotionGroupTargetsResponseContent) SetErrors(v []SponsoredProductsCreateTargetPromotionGroupTargetsFailureResponseItem) {
	o.Errors = v
}

func (o SponsoredProductsCreateTargetPromotionGroupTargetsResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableSponsoredProductsCreateTargetPromotionGroupTargetsResponseContent struct {
	value *SponsoredProductsCreateTargetPromotionGroupTargetsResponseContent
	isSet bool
}

func (v NullableSponsoredProductsCreateTargetPromotionGroupTargetsResponseContent) Get() *SponsoredProductsCreateTargetPromotionGroupTargetsResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsCreateTargetPromotionGroupTargetsResponseContent) Set(val *SponsoredProductsCreateTargetPromotionGroupTargetsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateTargetPromotionGroupTargetsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateTargetPromotionGroupTargetsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateTargetPromotionGroupTargetsResponseContent(val *SponsoredProductsCreateTargetPromotionGroupTargetsResponseContent) *NullableSponsoredProductsCreateTargetPromotionGroupTargetsResponseContent {
	return &NullableSponsoredProductsCreateTargetPromotionGroupTargetsResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateTargetPromotionGroupTargetsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateTargetPromotionGroupTargetsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
