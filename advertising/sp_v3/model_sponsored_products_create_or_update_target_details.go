package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCreateOrUpdateTargetDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateOrUpdateTargetDetails{}

// SponsoredProductsCreateOrUpdateTargetDetails struct for SponsoredProductsCreateOrUpdateTargetDetails
type SponsoredProductsCreateOrUpdateTargetDetails struct {
	KeywordTarget         *SponsoredProductsSPKeywordTargetDetails         `json:"keywordTarget,omitempty"`
	ProductCategoryTarget *SponsoredProductsSPProductCategoryTargetDetails `json:"productCategoryTarget,omitempty"`
	ProductTarget         *SponsoredProductsSPProductTargetDetails         `json:"productTarget,omitempty"`
}

// NewSponsoredProductsCreateOrUpdateTargetDetails instantiates a new SponsoredProductsCreateOrUpdateTargetDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateOrUpdateTargetDetails() *SponsoredProductsCreateOrUpdateTargetDetails {
	this := SponsoredProductsCreateOrUpdateTargetDetails{}
	return &this
}

// NewSponsoredProductsCreateOrUpdateTargetDetailsWithDefaults instantiates a new SponsoredProductsCreateOrUpdateTargetDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateOrUpdateTargetDetailsWithDefaults() *SponsoredProductsCreateOrUpdateTargetDetails {
	this := SponsoredProductsCreateOrUpdateTargetDetails{}
	return &this
}

// GetKeywordTarget returns the KeywordTarget field value if set, zero value otherwise.
func (o *SponsoredProductsCreateOrUpdateTargetDetails) GetKeywordTarget() SponsoredProductsSPKeywordTargetDetails {
	if o == nil || IsNil(o.KeywordTarget) {
		var ret SponsoredProductsSPKeywordTargetDetails
		return ret
	}
	return *o.KeywordTarget
}

// GetKeywordTargetOk returns a tuple with the KeywordTarget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateOrUpdateTargetDetails) GetKeywordTargetOk() (*SponsoredProductsSPKeywordTargetDetails, bool) {
	if o == nil || IsNil(o.KeywordTarget) {
		return nil, false
	}
	return o.KeywordTarget, true
}

// HasKeywordTarget returns a boolean if a field has been set.
func (o *SponsoredProductsCreateOrUpdateTargetDetails) HasKeywordTarget() bool {
	if o != nil && !IsNil(o.KeywordTarget) {
		return true
	}

	return false
}

// SetKeywordTarget gets a reference to the given SponsoredProductsSPKeywordTargetDetails and assigns it to the KeywordTarget field.
func (o *SponsoredProductsCreateOrUpdateTargetDetails) SetKeywordTarget(v SponsoredProductsSPKeywordTargetDetails) {
	o.KeywordTarget = &v
}

// GetProductCategoryTarget returns the ProductCategoryTarget field value if set, zero value otherwise.
func (o *SponsoredProductsCreateOrUpdateTargetDetails) GetProductCategoryTarget() SponsoredProductsSPProductCategoryTargetDetails {
	if o == nil || IsNil(o.ProductCategoryTarget) {
		var ret SponsoredProductsSPProductCategoryTargetDetails
		return ret
	}
	return *o.ProductCategoryTarget
}

// GetProductCategoryTargetOk returns a tuple with the ProductCategoryTarget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateOrUpdateTargetDetails) GetProductCategoryTargetOk() (*SponsoredProductsSPProductCategoryTargetDetails, bool) {
	if o == nil || IsNil(o.ProductCategoryTarget) {
		return nil, false
	}
	return o.ProductCategoryTarget, true
}

// HasProductCategoryTarget returns a boolean if a field has been set.
func (o *SponsoredProductsCreateOrUpdateTargetDetails) HasProductCategoryTarget() bool {
	if o != nil && !IsNil(o.ProductCategoryTarget) {
		return true
	}

	return false
}

// SetProductCategoryTarget gets a reference to the given SponsoredProductsSPProductCategoryTargetDetails and assigns it to the ProductCategoryTarget field.
func (o *SponsoredProductsCreateOrUpdateTargetDetails) SetProductCategoryTarget(v SponsoredProductsSPProductCategoryTargetDetails) {
	o.ProductCategoryTarget = &v
}

// GetProductTarget returns the ProductTarget field value if set, zero value otherwise.
func (o *SponsoredProductsCreateOrUpdateTargetDetails) GetProductTarget() SponsoredProductsSPProductTargetDetails {
	if o == nil || IsNil(o.ProductTarget) {
		var ret SponsoredProductsSPProductTargetDetails
		return ret
	}
	return *o.ProductTarget
}

// GetProductTargetOk returns a tuple with the ProductTarget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateOrUpdateTargetDetails) GetProductTargetOk() (*SponsoredProductsSPProductTargetDetails, bool) {
	if o == nil || IsNil(o.ProductTarget) {
		return nil, false
	}
	return o.ProductTarget, true
}

// HasProductTarget returns a boolean if a field has been set.
func (o *SponsoredProductsCreateOrUpdateTargetDetails) HasProductTarget() bool {
	if o != nil && !IsNil(o.ProductTarget) {
		return true
	}

	return false
}

// SetProductTarget gets a reference to the given SponsoredProductsSPProductTargetDetails and assigns it to the ProductTarget field.
func (o *SponsoredProductsCreateOrUpdateTargetDetails) SetProductTarget(v SponsoredProductsSPProductTargetDetails) {
	o.ProductTarget = &v
}

func (o SponsoredProductsCreateOrUpdateTargetDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.KeywordTarget) {
		toSerialize["keywordTarget"] = o.KeywordTarget
	}
	if !IsNil(o.ProductCategoryTarget) {
		toSerialize["productCategoryTarget"] = o.ProductCategoryTarget
	}
	if !IsNil(o.ProductTarget) {
		toSerialize["productTarget"] = o.ProductTarget
	}
	return toSerialize, nil
}

type NullableSponsoredProductsCreateOrUpdateTargetDetails struct {
	value *SponsoredProductsCreateOrUpdateTargetDetails
	isSet bool
}

func (v NullableSponsoredProductsCreateOrUpdateTargetDetails) Get() *SponsoredProductsCreateOrUpdateTargetDetails {
	return v.value
}

func (v *NullableSponsoredProductsCreateOrUpdateTargetDetails) Set(val *SponsoredProductsCreateOrUpdateTargetDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateOrUpdateTargetDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateOrUpdateTargetDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateOrUpdateTargetDetails(val *SponsoredProductsCreateOrUpdateTargetDetails) *NullableSponsoredProductsCreateOrUpdateTargetDetails {
	return &NullableSponsoredProductsCreateOrUpdateTargetDetails{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateOrUpdateTargetDetails) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateOrUpdateTargetDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
