package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsTargetDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsTargetDetails{}

// SponsoredProductsTargetDetails struct for SponsoredProductsTargetDetails
type SponsoredProductsTargetDetails struct {
	KeywordTarget         *SponsoredProductsSPKeywordTargetDetails         `json:"keywordTarget,omitempty"`
	ProductCategoryTarget *SponsoredProductsSPProductCategoryTargetDetails `json:"productCategoryTarget,omitempty"`
	AutoTarget            *SponsoredProductsSPAutoTargetDetails            `json:"autoTarget,omitempty"`
	ProductTarget         *SponsoredProductsSPProductTargetDetails         `json:"productTarget,omitempty"`
}

// NewSponsoredProductsTargetDetails instantiates a new SponsoredProductsTargetDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsTargetDetails() *SponsoredProductsTargetDetails {
	this := SponsoredProductsTargetDetails{}
	return &this
}

// NewSponsoredProductsTargetDetailsWithDefaults instantiates a new SponsoredProductsTargetDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsTargetDetailsWithDefaults() *SponsoredProductsTargetDetails {
	this := SponsoredProductsTargetDetails{}
	return &this
}

// GetKeywordTarget returns the KeywordTarget field value if set, zero value otherwise.
func (o *SponsoredProductsTargetDetails) GetKeywordTarget() SponsoredProductsSPKeywordTargetDetails {
	if o == nil || IsNil(o.KeywordTarget) {
		var ret SponsoredProductsSPKeywordTargetDetails
		return ret
	}
	return *o.KeywordTarget
}

// GetKeywordTargetOk returns a tuple with the KeywordTarget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsTargetDetails) GetKeywordTargetOk() (*SponsoredProductsSPKeywordTargetDetails, bool) {
	if o == nil || IsNil(o.KeywordTarget) {
		return nil, false
	}
	return o.KeywordTarget, true
}

// HasKeywordTarget returns a boolean if a field has been set.
func (o *SponsoredProductsTargetDetails) HasKeywordTarget() bool {
	if o != nil && !IsNil(o.KeywordTarget) {
		return true
	}

	return false
}

// SetKeywordTarget gets a reference to the given SponsoredProductsSPKeywordTargetDetails and assigns it to the KeywordTarget field.
func (o *SponsoredProductsTargetDetails) SetKeywordTarget(v SponsoredProductsSPKeywordTargetDetails) {
	o.KeywordTarget = &v
}

// GetProductCategoryTarget returns the ProductCategoryTarget field value if set, zero value otherwise.
func (o *SponsoredProductsTargetDetails) GetProductCategoryTarget() SponsoredProductsSPProductCategoryTargetDetails {
	if o == nil || IsNil(o.ProductCategoryTarget) {
		var ret SponsoredProductsSPProductCategoryTargetDetails
		return ret
	}
	return *o.ProductCategoryTarget
}

// GetProductCategoryTargetOk returns a tuple with the ProductCategoryTarget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsTargetDetails) GetProductCategoryTargetOk() (*SponsoredProductsSPProductCategoryTargetDetails, bool) {
	if o == nil || IsNil(o.ProductCategoryTarget) {
		return nil, false
	}
	return o.ProductCategoryTarget, true
}

// HasProductCategoryTarget returns a boolean if a field has been set.
func (o *SponsoredProductsTargetDetails) HasProductCategoryTarget() bool {
	if o != nil && !IsNil(o.ProductCategoryTarget) {
		return true
	}

	return false
}

// SetProductCategoryTarget gets a reference to the given SponsoredProductsSPProductCategoryTargetDetails and assigns it to the ProductCategoryTarget field.
func (o *SponsoredProductsTargetDetails) SetProductCategoryTarget(v SponsoredProductsSPProductCategoryTargetDetails) {
	o.ProductCategoryTarget = &v
}

// GetAutoTarget returns the AutoTarget field value if set, zero value otherwise.
func (o *SponsoredProductsTargetDetails) GetAutoTarget() SponsoredProductsSPAutoTargetDetails {
	if o == nil || IsNil(o.AutoTarget) {
		var ret SponsoredProductsSPAutoTargetDetails
		return ret
	}
	return *o.AutoTarget
}

// GetAutoTargetOk returns a tuple with the AutoTarget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsTargetDetails) GetAutoTargetOk() (*SponsoredProductsSPAutoTargetDetails, bool) {
	if o == nil || IsNil(o.AutoTarget) {
		return nil, false
	}
	return o.AutoTarget, true
}

// HasAutoTarget returns a boolean if a field has been set.
func (o *SponsoredProductsTargetDetails) HasAutoTarget() bool {
	if o != nil && !IsNil(o.AutoTarget) {
		return true
	}

	return false
}

// SetAutoTarget gets a reference to the given SponsoredProductsSPAutoTargetDetails and assigns it to the AutoTarget field.
func (o *SponsoredProductsTargetDetails) SetAutoTarget(v SponsoredProductsSPAutoTargetDetails) {
	o.AutoTarget = &v
}

// GetProductTarget returns the ProductTarget field value if set, zero value otherwise.
func (o *SponsoredProductsTargetDetails) GetProductTarget() SponsoredProductsSPProductTargetDetails {
	if o == nil || IsNil(o.ProductTarget) {
		var ret SponsoredProductsSPProductTargetDetails
		return ret
	}
	return *o.ProductTarget
}

// GetProductTargetOk returns a tuple with the ProductTarget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsTargetDetails) GetProductTargetOk() (*SponsoredProductsSPProductTargetDetails, bool) {
	if o == nil || IsNil(o.ProductTarget) {
		return nil, false
	}
	return o.ProductTarget, true
}

// HasProductTarget returns a boolean if a field has been set.
func (o *SponsoredProductsTargetDetails) HasProductTarget() bool {
	if o != nil && !IsNil(o.ProductTarget) {
		return true
	}

	return false
}

// SetProductTarget gets a reference to the given SponsoredProductsSPProductTargetDetails and assigns it to the ProductTarget field.
func (o *SponsoredProductsTargetDetails) SetProductTarget(v SponsoredProductsSPProductTargetDetails) {
	o.ProductTarget = &v
}

func (o SponsoredProductsTargetDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.KeywordTarget) {
		toSerialize["keywordTarget"] = o.KeywordTarget
	}
	if !IsNil(o.ProductCategoryTarget) {
		toSerialize["productCategoryTarget"] = o.ProductCategoryTarget
	}
	if !IsNil(o.AutoTarget) {
		toSerialize["autoTarget"] = o.AutoTarget
	}
	if !IsNil(o.ProductTarget) {
		toSerialize["productTarget"] = o.ProductTarget
	}
	return toSerialize, nil
}

type NullableSponsoredProductsTargetDetails struct {
	value *SponsoredProductsTargetDetails
	isSet bool
}

func (v NullableSponsoredProductsTargetDetails) Get() *SponsoredProductsTargetDetails {
	return v.value
}

func (v *NullableSponsoredProductsTargetDetails) Set(val *SponsoredProductsTargetDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsTargetDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsTargetDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsTargetDetails(val *SponsoredProductsTargetDetails) *NullableSponsoredProductsTargetDetails {
	return &NullableSponsoredProductsTargetDetails{value: val, isSet: true}
}

func (v NullableSponsoredProductsTargetDetails) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsTargetDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
