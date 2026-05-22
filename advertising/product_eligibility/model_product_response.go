package product_eligibility

import (
	"github.com/bytedance/sonic"
)

// checks if the ProductResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductResponse{}

// ProductResponse An product advertising eligibility response.
type ProductResponse struct {
	ProductDetails        ProductDetails      `json:"productDetails"`
	EligibilityStatusList []EligibilityStatus `json:"eligibilityStatusList"`
	// A human-readable description of the product's advertising eligibility status. Inherits highest severity from eligibilityStatusList.
	OverallStatus string `json:"overallStatus"`
}

type _ProductResponse ProductResponse

// NewProductResponse instantiates a new ProductResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductResponse(productDetails ProductDetails, eligibilityStatusList []EligibilityStatus, overallStatus string) *ProductResponse {
	this := ProductResponse{}
	this.ProductDetails = productDetails
	this.EligibilityStatusList = eligibilityStatusList
	this.OverallStatus = overallStatus
	return &this
}

// NewProductResponseWithDefaults instantiates a new ProductResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductResponseWithDefaults() *ProductResponse {
	this := ProductResponse{}
	return &this
}

// GetProductDetails returns the ProductDetails field value
func (o *ProductResponse) GetProductDetails() ProductDetails {
	if o == nil {
		var ret ProductDetails
		return ret
	}

	return o.ProductDetails
}

// GetProductDetailsOk returns a tuple with the ProductDetails field value
// and a boolean to check if the value has been set.
func (o *ProductResponse) GetProductDetailsOk() (*ProductDetails, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductDetails, true
}

// SetProductDetails sets field value
func (o *ProductResponse) SetProductDetails(v ProductDetails) {
	o.ProductDetails = v
}

// GetEligibilityStatusList returns the EligibilityStatusList field value
func (o *ProductResponse) GetEligibilityStatusList() []EligibilityStatus {
	if o == nil {
		var ret []EligibilityStatus
		return ret
	}

	return o.EligibilityStatusList
}

// GetEligibilityStatusListOk returns a tuple with the EligibilityStatusList field value
// and a boolean to check if the value has been set.
func (o *ProductResponse) GetEligibilityStatusListOk() ([]EligibilityStatus, bool) {
	if o == nil {
		return nil, false
	}
	return o.EligibilityStatusList, true
}

// SetEligibilityStatusList sets field value
func (o *ProductResponse) SetEligibilityStatusList(v []EligibilityStatus) {
	o.EligibilityStatusList = v
}

// GetOverallStatus returns the OverallStatus field value
func (o *ProductResponse) GetOverallStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OverallStatus
}

// GetOverallStatusOk returns a tuple with the OverallStatus field value
// and a boolean to check if the value has been set.
func (o *ProductResponse) GetOverallStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OverallStatus, true
}

// SetOverallStatus sets field value
func (o *ProductResponse) SetOverallStatus(v string) {
	o.OverallStatus = v
}

func (o ProductResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["productDetails"] = o.ProductDetails
	toSerialize["eligibilityStatusList"] = o.EligibilityStatusList
	toSerialize["overallStatus"] = o.OverallStatus
	return toSerialize, nil
}

type NullableProductResponse struct {
	value *ProductResponse
	isSet bool
}

func (v NullableProductResponse) Get() *ProductResponse {
	return v.value
}

func (v *NullableProductResponse) Set(val *ProductResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProductResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProductResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductResponse(val *ProductResponse) *NullableProductResponse {
	return &NullableProductResponse{value: val, isSet: true}
}

func (v NullableProductResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableProductResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
