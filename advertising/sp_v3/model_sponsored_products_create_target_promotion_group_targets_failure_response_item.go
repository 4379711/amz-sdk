package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCreateTargetPromotionGroupTargetsFailureResponseItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateTargetPromotionGroupTargetsFailureResponseItem{}

// SponsoredProductsCreateTargetPromotionGroupTargetsFailureResponseItem Response object of failed target promotion group target.
type SponsoredProductsCreateTargetPromotionGroupTargetsFailureResponseItem struct {
	// The expression type of the target that was requested to be created.
	ExpressionType *string `json:"expressionType,omitempty"`
	// Response object of failed target promotion group target.
	Errors []SponsoredProductsCreateTargetError `json:"errors,omitempty"`
	// The target that was requested to be created.
	Target *string `json:"target,omitempty"`
}

// NewSponsoredProductsCreateTargetPromotionGroupTargetsFailureResponseItem instantiates a new SponsoredProductsCreateTargetPromotionGroupTargetsFailureResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateTargetPromotionGroupTargetsFailureResponseItem() *SponsoredProductsCreateTargetPromotionGroupTargetsFailureResponseItem {
	this := SponsoredProductsCreateTargetPromotionGroupTargetsFailureResponseItem{}
	return &this
}

// NewSponsoredProductsCreateTargetPromotionGroupTargetsFailureResponseItemWithDefaults instantiates a new SponsoredProductsCreateTargetPromotionGroupTargetsFailureResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateTargetPromotionGroupTargetsFailureResponseItemWithDefaults() *SponsoredProductsCreateTargetPromotionGroupTargetsFailureResponseItem {
	this := SponsoredProductsCreateTargetPromotionGroupTargetsFailureResponseItem{}
	return &this
}

// GetExpressionType returns the ExpressionType field value if set, zero value otherwise.
func (o *SponsoredProductsCreateTargetPromotionGroupTargetsFailureResponseItem) GetExpressionType() string {
	if o == nil || IsNil(o.ExpressionType) {
		var ret string
		return ret
	}
	return *o.ExpressionType
}

// GetExpressionTypeOk returns a tuple with the ExpressionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateTargetPromotionGroupTargetsFailureResponseItem) GetExpressionTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ExpressionType) {
		return nil, false
	}
	return o.ExpressionType, true
}

// HasExpressionType returns a boolean if a field has been set.
func (o *SponsoredProductsCreateTargetPromotionGroupTargetsFailureResponseItem) HasExpressionType() bool {
	if o != nil && !IsNil(o.ExpressionType) {
		return true
	}

	return false
}

// SetExpressionType gets a reference to the given string and assigns it to the ExpressionType field.
func (o *SponsoredProductsCreateTargetPromotionGroupTargetsFailureResponseItem) SetExpressionType(v string) {
	o.ExpressionType = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *SponsoredProductsCreateTargetPromotionGroupTargetsFailureResponseItem) GetErrors() []SponsoredProductsCreateTargetError {
	if o == nil || IsNil(o.Errors) {
		var ret []SponsoredProductsCreateTargetError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateTargetPromotionGroupTargetsFailureResponseItem) GetErrorsOk() ([]SponsoredProductsCreateTargetError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *SponsoredProductsCreateTargetPromotionGroupTargetsFailureResponseItem) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []SponsoredProductsCreateTargetError and assigns it to the Errors field.
func (o *SponsoredProductsCreateTargetPromotionGroupTargetsFailureResponseItem) SetErrors(v []SponsoredProductsCreateTargetError) {
	o.Errors = v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *SponsoredProductsCreateTargetPromotionGroupTargetsFailureResponseItem) GetTarget() string {
	if o == nil || IsNil(o.Target) {
		var ret string
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateTargetPromotionGroupTargetsFailureResponseItem) GetTargetOk() (*string, bool) {
	if o == nil || IsNil(o.Target) {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *SponsoredProductsCreateTargetPromotionGroupTargetsFailureResponseItem) HasTarget() bool {
	if o != nil && !IsNil(o.Target) {
		return true
	}

	return false
}

// SetTarget gets a reference to the given string and assigns it to the Target field.
func (o *SponsoredProductsCreateTargetPromotionGroupTargetsFailureResponseItem) SetTarget(v string) {
	o.Target = &v
}

func (o SponsoredProductsCreateTargetPromotionGroupTargetsFailureResponseItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExpressionType) {
		toSerialize["expressionType"] = o.ExpressionType
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	if !IsNil(o.Target) {
		toSerialize["target"] = o.Target
	}
	return toSerialize, nil
}

type NullableSponsoredProductsCreateTargetPromotionGroupTargetsFailureResponseItem struct {
	value *SponsoredProductsCreateTargetPromotionGroupTargetsFailureResponseItem
	isSet bool
}

func (v NullableSponsoredProductsCreateTargetPromotionGroupTargetsFailureResponseItem) Get() *SponsoredProductsCreateTargetPromotionGroupTargetsFailureResponseItem {
	return v.value
}

func (v *NullableSponsoredProductsCreateTargetPromotionGroupTargetsFailureResponseItem) Set(val *SponsoredProductsCreateTargetPromotionGroupTargetsFailureResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateTargetPromotionGroupTargetsFailureResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateTargetPromotionGroupTargetsFailureResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateTargetPromotionGroupTargetsFailureResponseItem(val *SponsoredProductsCreateTargetPromotionGroupTargetsFailureResponseItem) *NullableSponsoredProductsCreateTargetPromotionGroupTargetsFailureResponseItem {
	return &NullableSponsoredProductsCreateTargetPromotionGroupTargetsFailureResponseItem{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateTargetPromotionGroupTargetsFailureResponseItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateTargetPromotionGroupTargetsFailureResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
