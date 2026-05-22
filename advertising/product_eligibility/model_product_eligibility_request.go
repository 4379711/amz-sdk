package product_eligibility

import (
	"github.com/bytedance/sonic"
)

// checks if the ProductEligibilityRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductEligibilityRequest{}

// ProductEligibilityRequest A product advertising eligibility request object.
type ProductEligibilityRequest struct {
	// Set to 'sp' to check product eligibility for Sponsored Products advertisements. Set to 'sb' to check product eligibility for Sponsored Brands advertisements. Set to 'sd' to check product eligibility for Sponsored Displays advertisements. Set to 'dsp' to check product eligibility for Demand Side Platform advertisements.
	AdType *string `json:"adType,omitempty"`
	// A list of product identifier objects.
	ProductDetailsList []ProductDetails `json:"productDetailsList"`
	// Set to the locale string in the table below to specify the language in which the response is returned
	Locale *string `json:"locale,omitempty"`
}

type _ProductEligibilityRequest ProductEligibilityRequest

// NewProductEligibilityRequest instantiates a new ProductEligibilityRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductEligibilityRequest(productDetailsList []ProductDetails) *ProductEligibilityRequest {
	this := ProductEligibilityRequest{}
	var adType string = "sp"
	this.AdType = &adType
	this.ProductDetailsList = productDetailsList
	return &this
}

// NewProductEligibilityRequestWithDefaults instantiates a new ProductEligibilityRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductEligibilityRequestWithDefaults() *ProductEligibilityRequest {
	this := ProductEligibilityRequest{}
	var adType string = "sp"
	this.AdType = &adType
	return &this
}

// GetAdType returns the AdType field value if set, zero value otherwise.
func (o *ProductEligibilityRequest) GetAdType() string {
	if o == nil || IsNil(o.AdType) {
		var ret string
		return ret
	}
	return *o.AdType
}

// GetAdTypeOk returns a tuple with the AdType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductEligibilityRequest) GetAdTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AdType) {
		return nil, false
	}
	return o.AdType, true
}

// HasAdType returns a boolean if a field has been set.
func (o *ProductEligibilityRequest) HasAdType() bool {
	if o != nil && !IsNil(o.AdType) {
		return true
	}

	return false
}

// SetAdType gets a reference to the given string and assigns it to the AdType field.
func (o *ProductEligibilityRequest) SetAdType(v string) {
	o.AdType = &v
}

// GetProductDetailsList returns the ProductDetailsList field value
func (o *ProductEligibilityRequest) GetProductDetailsList() []ProductDetails {
	if o == nil {
		var ret []ProductDetails
		return ret
	}

	return o.ProductDetailsList
}

// GetProductDetailsListOk returns a tuple with the ProductDetailsList field value
// and a boolean to check if the value has been set.
func (o *ProductEligibilityRequest) GetProductDetailsListOk() ([]ProductDetails, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProductDetailsList, true
}

// SetProductDetailsList sets field value
func (o *ProductEligibilityRequest) SetProductDetailsList(v []ProductDetails) {
	o.ProductDetailsList = v
}

// GetLocale returns the Locale field value if set, zero value otherwise.
func (o *ProductEligibilityRequest) GetLocale() string {
	if o == nil || IsNil(o.Locale) {
		var ret string
		return ret
	}
	return *o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductEligibilityRequest) GetLocaleOk() (*string, bool) {
	if o == nil || IsNil(o.Locale) {
		return nil, false
	}
	return o.Locale, true
}

// HasLocale returns a boolean if a field has been set.
func (o *ProductEligibilityRequest) HasLocale() bool {
	if o != nil && !IsNil(o.Locale) {
		return true
	}

	return false
}

// SetLocale gets a reference to the given string and assigns it to the Locale field.
func (o *ProductEligibilityRequest) SetLocale(v string) {
	o.Locale = &v
}

func (o ProductEligibilityRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdType) {
		toSerialize["adType"] = o.AdType
	}
	toSerialize["productDetailsList"] = o.ProductDetailsList
	if !IsNil(o.Locale) {
		toSerialize["locale"] = o.Locale
	}
	return toSerialize, nil
}

type NullableProductEligibilityRequest struct {
	value *ProductEligibilityRequest
	isSet bool
}

func (v NullableProductEligibilityRequest) Get() *ProductEligibilityRequest {
	return v.value
}

func (v *NullableProductEligibilityRequest) Set(val *ProductEligibilityRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProductEligibilityRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProductEligibilityRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductEligibilityRequest(val *ProductEligibilityRequest) *NullableProductEligibilityRequest {
	return &NullableProductEligibilityRequest{value: val, isSet: true}
}

func (v NullableProductEligibilityRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableProductEligibilityRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
