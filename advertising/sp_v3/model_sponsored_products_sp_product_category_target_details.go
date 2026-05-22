package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsSPProductCategoryTargetDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsSPProductCategoryTargetDetails{}

// SponsoredProductsSPProductCategoryTargetDetails struct for SponsoredProductsSPProductCategoryTargetDetails
type SponsoredProductsSPProductCategoryTargetDetails struct {
	// The age range ID to target.
	ProductAgeRangeId *string `json:"productAgeRangeId,omitempty"`
	// Refinement to target products with a price less than the value within the product category.
	ProductPriceLessThan *float64 `json:"productPriceLessThan,omitempty"`
	// Target based on if a product is Prime-shipping eligible.
	ProductPrimeShippingEligible *bool `json:"productPrimeShippingEligible,omitempty"`
	// The resolved product category.
	ProductCategoryIdResolved *string `json:"productCategoryIdResolved,omitempty"`
	// The brand ID to target.
	ProductBrandId *string `json:"productBrandId,omitempty"`
	// The resolved name of the brand.
	ProductBrandIdResolved *string `json:"productBrandIdResolved,omitempty"`
	// Refinement to target products with a rating less than the value within the product category.
	ProductRatingLessThan *float64 `json:"productRatingLessThan,omitempty"`
	// The product genre ID to target.
	ProductGenreId *string `json:"productGenreId,omitempty"`
	// The product category ID to target.
	ProductCategoryId *string `json:"productCategoryId,omitempty"`
	// Refinement to target products with a rating greater than the value within the product category.
	ProductRatingGreaterThan *float64 `json:"productRatingGreaterThan,omitempty"`
	// The resolved age range to target.
	ProductAgeRangeIdResolved *string `json:"productAgeRangeIdResolved,omitempty"`
	// The resolved product genre to target.
	ProductGenreIdResolved *string `json:"productGenreIdResolved,omitempty"`
	// Refinement to target products with a price greater than the value within the product category.
	ProductPriceGreaterThan *float64 `json:"productPriceGreaterThan,omitempty"`
}

// NewSponsoredProductsSPProductCategoryTargetDetails instantiates a new SponsoredProductsSPProductCategoryTargetDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsSPProductCategoryTargetDetails() *SponsoredProductsSPProductCategoryTargetDetails {
	this := SponsoredProductsSPProductCategoryTargetDetails{}
	return &this
}

// NewSponsoredProductsSPProductCategoryTargetDetailsWithDefaults instantiates a new SponsoredProductsSPProductCategoryTargetDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsSPProductCategoryTargetDetailsWithDefaults() *SponsoredProductsSPProductCategoryTargetDetails {
	this := SponsoredProductsSPProductCategoryTargetDetails{}
	return &this
}

// GetProductAgeRangeId returns the ProductAgeRangeId field value if set, zero value otherwise.
func (o *SponsoredProductsSPProductCategoryTargetDetails) GetProductAgeRangeId() string {
	if o == nil || IsNil(o.ProductAgeRangeId) {
		var ret string
		return ret
	}
	return *o.ProductAgeRangeId
}

// GetProductAgeRangeIdOk returns a tuple with the ProductAgeRangeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsSPProductCategoryTargetDetails) GetProductAgeRangeIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProductAgeRangeId) {
		return nil, false
	}
	return o.ProductAgeRangeId, true
}

// HasProductAgeRangeId returns a boolean if a field has been set.
func (o *SponsoredProductsSPProductCategoryTargetDetails) HasProductAgeRangeId() bool {
	if o != nil && !IsNil(o.ProductAgeRangeId) {
		return true
	}

	return false
}

// SetProductAgeRangeId gets a reference to the given string and assigns it to the ProductAgeRangeId field.
func (o *SponsoredProductsSPProductCategoryTargetDetails) SetProductAgeRangeId(v string) {
	o.ProductAgeRangeId = &v
}

// GetProductPriceLessThan returns the ProductPriceLessThan field value if set, zero value otherwise.
func (o *SponsoredProductsSPProductCategoryTargetDetails) GetProductPriceLessThan() float64 {
	if o == nil || IsNil(o.ProductPriceLessThan) {
		var ret float64
		return ret
	}
	return *o.ProductPriceLessThan
}

// GetProductPriceLessThanOk returns a tuple with the ProductPriceLessThan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsSPProductCategoryTargetDetails) GetProductPriceLessThanOk() (*float64, bool) {
	if o == nil || IsNil(o.ProductPriceLessThan) {
		return nil, false
	}
	return o.ProductPriceLessThan, true
}

// HasProductPriceLessThan returns a boolean if a field has been set.
func (o *SponsoredProductsSPProductCategoryTargetDetails) HasProductPriceLessThan() bool {
	if o != nil && !IsNil(o.ProductPriceLessThan) {
		return true
	}

	return false
}

// SetProductPriceLessThan gets a reference to the given float64 and assigns it to the ProductPriceLessThan field.
func (o *SponsoredProductsSPProductCategoryTargetDetails) SetProductPriceLessThan(v float64) {
	o.ProductPriceLessThan = &v
}

// GetProductPrimeShippingEligible returns the ProductPrimeShippingEligible field value if set, zero value otherwise.
func (o *SponsoredProductsSPProductCategoryTargetDetails) GetProductPrimeShippingEligible() bool {
	if o == nil || IsNil(o.ProductPrimeShippingEligible) {
		var ret bool
		return ret
	}
	return *o.ProductPrimeShippingEligible
}

// GetProductPrimeShippingEligibleOk returns a tuple with the ProductPrimeShippingEligible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsSPProductCategoryTargetDetails) GetProductPrimeShippingEligibleOk() (*bool, bool) {
	if o == nil || IsNil(o.ProductPrimeShippingEligible) {
		return nil, false
	}
	return o.ProductPrimeShippingEligible, true
}

// HasProductPrimeShippingEligible returns a boolean if a field has been set.
func (o *SponsoredProductsSPProductCategoryTargetDetails) HasProductPrimeShippingEligible() bool {
	if o != nil && !IsNil(o.ProductPrimeShippingEligible) {
		return true
	}

	return false
}

// SetProductPrimeShippingEligible gets a reference to the given bool and assigns it to the ProductPrimeShippingEligible field.
func (o *SponsoredProductsSPProductCategoryTargetDetails) SetProductPrimeShippingEligible(v bool) {
	o.ProductPrimeShippingEligible = &v
}

// GetProductCategoryIdResolved returns the ProductCategoryIdResolved field value if set, zero value otherwise.
func (o *SponsoredProductsSPProductCategoryTargetDetails) GetProductCategoryIdResolved() string {
	if o == nil || IsNil(o.ProductCategoryIdResolved) {
		var ret string
		return ret
	}
	return *o.ProductCategoryIdResolved
}

// GetProductCategoryIdResolvedOk returns a tuple with the ProductCategoryIdResolved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsSPProductCategoryTargetDetails) GetProductCategoryIdResolvedOk() (*string, bool) {
	if o == nil || IsNil(o.ProductCategoryIdResolved) {
		return nil, false
	}
	return o.ProductCategoryIdResolved, true
}

// HasProductCategoryIdResolved returns a boolean if a field has been set.
func (o *SponsoredProductsSPProductCategoryTargetDetails) HasProductCategoryIdResolved() bool {
	if o != nil && !IsNil(o.ProductCategoryIdResolved) {
		return true
	}

	return false
}

// SetProductCategoryIdResolved gets a reference to the given string and assigns it to the ProductCategoryIdResolved field.
func (o *SponsoredProductsSPProductCategoryTargetDetails) SetProductCategoryIdResolved(v string) {
	o.ProductCategoryIdResolved = &v
}

// GetProductBrandId returns the ProductBrandId field value if set, zero value otherwise.
func (o *SponsoredProductsSPProductCategoryTargetDetails) GetProductBrandId() string {
	if o == nil || IsNil(o.ProductBrandId) {
		var ret string
		return ret
	}
	return *o.ProductBrandId
}

// GetProductBrandIdOk returns a tuple with the ProductBrandId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsSPProductCategoryTargetDetails) GetProductBrandIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProductBrandId) {
		return nil, false
	}
	return o.ProductBrandId, true
}

// HasProductBrandId returns a boolean if a field has been set.
func (o *SponsoredProductsSPProductCategoryTargetDetails) HasProductBrandId() bool {
	if o != nil && !IsNil(o.ProductBrandId) {
		return true
	}

	return false
}

// SetProductBrandId gets a reference to the given string and assigns it to the ProductBrandId field.
func (o *SponsoredProductsSPProductCategoryTargetDetails) SetProductBrandId(v string) {
	o.ProductBrandId = &v
}

// GetProductBrandIdResolved returns the ProductBrandIdResolved field value if set, zero value otherwise.
func (o *SponsoredProductsSPProductCategoryTargetDetails) GetProductBrandIdResolved() string {
	if o == nil || IsNil(o.ProductBrandIdResolved) {
		var ret string
		return ret
	}
	return *o.ProductBrandIdResolved
}

// GetProductBrandIdResolvedOk returns a tuple with the ProductBrandIdResolved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsSPProductCategoryTargetDetails) GetProductBrandIdResolvedOk() (*string, bool) {
	if o == nil || IsNil(o.ProductBrandIdResolved) {
		return nil, false
	}
	return o.ProductBrandIdResolved, true
}

// HasProductBrandIdResolved returns a boolean if a field has been set.
func (o *SponsoredProductsSPProductCategoryTargetDetails) HasProductBrandIdResolved() bool {
	if o != nil && !IsNil(o.ProductBrandIdResolved) {
		return true
	}

	return false
}

// SetProductBrandIdResolved gets a reference to the given string and assigns it to the ProductBrandIdResolved field.
func (o *SponsoredProductsSPProductCategoryTargetDetails) SetProductBrandIdResolved(v string) {
	o.ProductBrandIdResolved = &v
}

// GetProductRatingLessThan returns the ProductRatingLessThan field value if set, zero value otherwise.
func (o *SponsoredProductsSPProductCategoryTargetDetails) GetProductRatingLessThan() float64 {
	if o == nil || IsNil(o.ProductRatingLessThan) {
		var ret float64
		return ret
	}
	return *o.ProductRatingLessThan
}

// GetProductRatingLessThanOk returns a tuple with the ProductRatingLessThan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsSPProductCategoryTargetDetails) GetProductRatingLessThanOk() (*float64, bool) {
	if o == nil || IsNil(o.ProductRatingLessThan) {
		return nil, false
	}
	return o.ProductRatingLessThan, true
}

// HasProductRatingLessThan returns a boolean if a field has been set.
func (o *SponsoredProductsSPProductCategoryTargetDetails) HasProductRatingLessThan() bool {
	if o != nil && !IsNil(o.ProductRatingLessThan) {
		return true
	}

	return false
}

// SetProductRatingLessThan gets a reference to the given float64 and assigns it to the ProductRatingLessThan field.
func (o *SponsoredProductsSPProductCategoryTargetDetails) SetProductRatingLessThan(v float64) {
	o.ProductRatingLessThan = &v
}

// GetProductGenreId returns the ProductGenreId field value if set, zero value otherwise.
func (o *SponsoredProductsSPProductCategoryTargetDetails) GetProductGenreId() string {
	if o == nil || IsNil(o.ProductGenreId) {
		var ret string
		return ret
	}
	return *o.ProductGenreId
}

// GetProductGenreIdOk returns a tuple with the ProductGenreId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsSPProductCategoryTargetDetails) GetProductGenreIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProductGenreId) {
		return nil, false
	}
	return o.ProductGenreId, true
}

// HasProductGenreId returns a boolean if a field has been set.
func (o *SponsoredProductsSPProductCategoryTargetDetails) HasProductGenreId() bool {
	if o != nil && !IsNil(o.ProductGenreId) {
		return true
	}

	return false
}

// SetProductGenreId gets a reference to the given string and assigns it to the ProductGenreId field.
func (o *SponsoredProductsSPProductCategoryTargetDetails) SetProductGenreId(v string) {
	o.ProductGenreId = &v
}

// GetProductCategoryId returns the ProductCategoryId field value if set, zero value otherwise.
func (o *SponsoredProductsSPProductCategoryTargetDetails) GetProductCategoryId() string {
	if o == nil || IsNil(o.ProductCategoryId) {
		var ret string
		return ret
	}
	return *o.ProductCategoryId
}

// GetProductCategoryIdOk returns a tuple with the ProductCategoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsSPProductCategoryTargetDetails) GetProductCategoryIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProductCategoryId) {
		return nil, false
	}
	return o.ProductCategoryId, true
}

// HasProductCategoryId returns a boolean if a field has been set.
func (o *SponsoredProductsSPProductCategoryTargetDetails) HasProductCategoryId() bool {
	if o != nil && !IsNil(o.ProductCategoryId) {
		return true
	}

	return false
}

// SetProductCategoryId gets a reference to the given string and assigns it to the ProductCategoryId field.
func (o *SponsoredProductsSPProductCategoryTargetDetails) SetProductCategoryId(v string) {
	o.ProductCategoryId = &v
}

// GetProductRatingGreaterThan returns the ProductRatingGreaterThan field value if set, zero value otherwise.
func (o *SponsoredProductsSPProductCategoryTargetDetails) GetProductRatingGreaterThan() float64 {
	if o == nil || IsNil(o.ProductRatingGreaterThan) {
		var ret float64
		return ret
	}
	return *o.ProductRatingGreaterThan
}

// GetProductRatingGreaterThanOk returns a tuple with the ProductRatingGreaterThan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsSPProductCategoryTargetDetails) GetProductRatingGreaterThanOk() (*float64, bool) {
	if o == nil || IsNil(o.ProductRatingGreaterThan) {
		return nil, false
	}
	return o.ProductRatingGreaterThan, true
}

// HasProductRatingGreaterThan returns a boolean if a field has been set.
func (o *SponsoredProductsSPProductCategoryTargetDetails) HasProductRatingGreaterThan() bool {
	if o != nil && !IsNil(o.ProductRatingGreaterThan) {
		return true
	}

	return false
}

// SetProductRatingGreaterThan gets a reference to the given float64 and assigns it to the ProductRatingGreaterThan field.
func (o *SponsoredProductsSPProductCategoryTargetDetails) SetProductRatingGreaterThan(v float64) {
	o.ProductRatingGreaterThan = &v
}

// GetProductAgeRangeIdResolved returns the ProductAgeRangeIdResolved field value if set, zero value otherwise.
func (o *SponsoredProductsSPProductCategoryTargetDetails) GetProductAgeRangeIdResolved() string {
	if o == nil || IsNil(o.ProductAgeRangeIdResolved) {
		var ret string
		return ret
	}
	return *o.ProductAgeRangeIdResolved
}

// GetProductAgeRangeIdResolvedOk returns a tuple with the ProductAgeRangeIdResolved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsSPProductCategoryTargetDetails) GetProductAgeRangeIdResolvedOk() (*string, bool) {
	if o == nil || IsNil(o.ProductAgeRangeIdResolved) {
		return nil, false
	}
	return o.ProductAgeRangeIdResolved, true
}

// HasProductAgeRangeIdResolved returns a boolean if a field has been set.
func (o *SponsoredProductsSPProductCategoryTargetDetails) HasProductAgeRangeIdResolved() bool {
	if o != nil && !IsNil(o.ProductAgeRangeIdResolved) {
		return true
	}

	return false
}

// SetProductAgeRangeIdResolved gets a reference to the given string and assigns it to the ProductAgeRangeIdResolved field.
func (o *SponsoredProductsSPProductCategoryTargetDetails) SetProductAgeRangeIdResolved(v string) {
	o.ProductAgeRangeIdResolved = &v
}

// GetProductGenreIdResolved returns the ProductGenreIdResolved field value if set, zero value otherwise.
func (o *SponsoredProductsSPProductCategoryTargetDetails) GetProductGenreIdResolved() string {
	if o == nil || IsNil(o.ProductGenreIdResolved) {
		var ret string
		return ret
	}
	return *o.ProductGenreIdResolved
}

// GetProductGenreIdResolvedOk returns a tuple with the ProductGenreIdResolved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsSPProductCategoryTargetDetails) GetProductGenreIdResolvedOk() (*string, bool) {
	if o == nil || IsNil(o.ProductGenreIdResolved) {
		return nil, false
	}
	return o.ProductGenreIdResolved, true
}

// HasProductGenreIdResolved returns a boolean if a field has been set.
func (o *SponsoredProductsSPProductCategoryTargetDetails) HasProductGenreIdResolved() bool {
	if o != nil && !IsNil(o.ProductGenreIdResolved) {
		return true
	}

	return false
}

// SetProductGenreIdResolved gets a reference to the given string and assigns it to the ProductGenreIdResolved field.
func (o *SponsoredProductsSPProductCategoryTargetDetails) SetProductGenreIdResolved(v string) {
	o.ProductGenreIdResolved = &v
}

// GetProductPriceGreaterThan returns the ProductPriceGreaterThan field value if set, zero value otherwise.
func (o *SponsoredProductsSPProductCategoryTargetDetails) GetProductPriceGreaterThan() float64 {
	if o == nil || IsNil(o.ProductPriceGreaterThan) {
		var ret float64
		return ret
	}
	return *o.ProductPriceGreaterThan
}

// GetProductPriceGreaterThanOk returns a tuple with the ProductPriceGreaterThan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsSPProductCategoryTargetDetails) GetProductPriceGreaterThanOk() (*float64, bool) {
	if o == nil || IsNil(o.ProductPriceGreaterThan) {
		return nil, false
	}
	return o.ProductPriceGreaterThan, true
}

// HasProductPriceGreaterThan returns a boolean if a field has been set.
func (o *SponsoredProductsSPProductCategoryTargetDetails) HasProductPriceGreaterThan() bool {
	if o != nil && !IsNil(o.ProductPriceGreaterThan) {
		return true
	}

	return false
}

// SetProductPriceGreaterThan gets a reference to the given float64 and assigns it to the ProductPriceGreaterThan field.
func (o *SponsoredProductsSPProductCategoryTargetDetails) SetProductPriceGreaterThan(v float64) {
	o.ProductPriceGreaterThan = &v
}

func (o SponsoredProductsSPProductCategoryTargetDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProductAgeRangeId) {
		toSerialize["productAgeRangeId"] = o.ProductAgeRangeId
	}
	if !IsNil(o.ProductPriceLessThan) {
		toSerialize["productPriceLessThan"] = o.ProductPriceLessThan
	}
	if !IsNil(o.ProductPrimeShippingEligible) {
		toSerialize["productPrimeShippingEligible"] = o.ProductPrimeShippingEligible
	}
	if !IsNil(o.ProductCategoryIdResolved) {
		toSerialize["productCategoryIdResolved"] = o.ProductCategoryIdResolved
	}
	if !IsNil(o.ProductBrandId) {
		toSerialize["productBrandId"] = o.ProductBrandId
	}
	if !IsNil(o.ProductBrandIdResolved) {
		toSerialize["productBrandIdResolved"] = o.ProductBrandIdResolved
	}
	if !IsNil(o.ProductRatingLessThan) {
		toSerialize["productRatingLessThan"] = o.ProductRatingLessThan
	}
	if !IsNil(o.ProductGenreId) {
		toSerialize["productGenreId"] = o.ProductGenreId
	}
	if !IsNil(o.ProductCategoryId) {
		toSerialize["productCategoryId"] = o.ProductCategoryId
	}
	if !IsNil(o.ProductRatingGreaterThan) {
		toSerialize["productRatingGreaterThan"] = o.ProductRatingGreaterThan
	}
	if !IsNil(o.ProductAgeRangeIdResolved) {
		toSerialize["productAgeRangeIdResolved"] = o.ProductAgeRangeIdResolved
	}
	if !IsNil(o.ProductGenreIdResolved) {
		toSerialize["productGenreIdResolved"] = o.ProductGenreIdResolved
	}
	if !IsNil(o.ProductPriceGreaterThan) {
		toSerialize["productPriceGreaterThan"] = o.ProductPriceGreaterThan
	}
	return toSerialize, nil
}

type NullableSponsoredProductsSPProductCategoryTargetDetails struct {
	value *SponsoredProductsSPProductCategoryTargetDetails
	isSet bool
}

func (v NullableSponsoredProductsSPProductCategoryTargetDetails) Get() *SponsoredProductsSPProductCategoryTargetDetails {
	return v.value
}

func (v *NullableSponsoredProductsSPProductCategoryTargetDetails) Set(val *SponsoredProductsSPProductCategoryTargetDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsSPProductCategoryTargetDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsSPProductCategoryTargetDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsSPProductCategoryTargetDetails(val *SponsoredProductsSPProductCategoryTargetDetails) *NullableSponsoredProductsSPProductCategoryTargetDetails {
	return &NullableSponsoredProductsSPProductCategoryTargetDetails{value: val, isSet: true}
}

func (v NullableSponsoredProductsSPProductCategoryTargetDetails) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsSPProductCategoryTargetDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
