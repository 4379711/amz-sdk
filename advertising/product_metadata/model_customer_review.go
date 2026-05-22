package product_metadata

import (
	"github.com/bytedance/sonic"
)

// checks if the CustomerReview type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerReview{}

// CustomerReview struct for CustomerReview
type CustomerReview struct {
	ReviewCount *ReviewCount `json:"reviewCount,omitempty"`
	ReviewUrl   *string      `json:"reviewUrl,omitempty"`
	Rating      *Rating      `json:"rating,omitempty"`
}

// NewCustomerReview instantiates a new CustomerReview object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerReview() *CustomerReview {
	this := CustomerReview{}
	return &this
}

// NewCustomerReviewWithDefaults instantiates a new CustomerReview object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerReviewWithDefaults() *CustomerReview {
	this := CustomerReview{}
	return &this
}

// GetReviewCount returns the ReviewCount field value if set, zero value otherwise.
func (o *CustomerReview) GetReviewCount() ReviewCount {
	if o == nil || IsNil(o.ReviewCount) {
		var ret ReviewCount
		return ret
	}
	return *o.ReviewCount
}

// GetReviewCountOk returns a tuple with the ReviewCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerReview) GetReviewCountOk() (*ReviewCount, bool) {
	if o == nil || IsNil(o.ReviewCount) {
		return nil, false
	}
	return o.ReviewCount, true
}

// HasReviewCount returns a boolean if a field has been set.
func (o *CustomerReview) HasReviewCount() bool {
	if o != nil && !IsNil(o.ReviewCount) {
		return true
	}

	return false
}

// SetReviewCount gets a reference to the given ReviewCount and assigns it to the ReviewCount field.
func (o *CustomerReview) SetReviewCount(v ReviewCount) {
	o.ReviewCount = &v
}

// GetReviewUrl returns the ReviewUrl field value if set, zero value otherwise.
func (o *CustomerReview) GetReviewUrl() string {
	if o == nil || IsNil(o.ReviewUrl) {
		var ret string
		return ret
	}
	return *o.ReviewUrl
}

// GetReviewUrlOk returns a tuple with the ReviewUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerReview) GetReviewUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ReviewUrl) {
		return nil, false
	}
	return o.ReviewUrl, true
}

// HasReviewUrl returns a boolean if a field has been set.
func (o *CustomerReview) HasReviewUrl() bool {
	if o != nil && !IsNil(o.ReviewUrl) {
		return true
	}

	return false
}

// SetReviewUrl gets a reference to the given string and assigns it to the ReviewUrl field.
func (o *CustomerReview) SetReviewUrl(v string) {
	o.ReviewUrl = &v
}

// GetRating returns the Rating field value if set, zero value otherwise.
func (o *CustomerReview) GetRating() Rating {
	if o == nil || IsNil(o.Rating) {
		var ret Rating
		return ret
	}
	return *o.Rating
}

// GetRatingOk returns a tuple with the Rating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerReview) GetRatingOk() (*Rating, bool) {
	if o == nil || IsNil(o.Rating) {
		return nil, false
	}
	return o.Rating, true
}

// HasRating returns a boolean if a field has been set.
func (o *CustomerReview) HasRating() bool {
	if o != nil && !IsNil(o.Rating) {
		return true
	}

	return false
}

// SetRating gets a reference to the given Rating and assigns it to the Rating field.
func (o *CustomerReview) SetRating(v Rating) {
	o.Rating = &v
}

func (o CustomerReview) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ReviewCount) {
		toSerialize["reviewCount"] = o.ReviewCount
	}
	if !IsNil(o.ReviewUrl) {
		toSerialize["reviewUrl"] = o.ReviewUrl
	}
	if !IsNil(o.Rating) {
		toSerialize["rating"] = o.Rating
	}
	return toSerialize, nil
}

type NullableCustomerReview struct {
	value *CustomerReview
	isSet bool
}

func (v NullableCustomerReview) Get() *CustomerReview {
	return v.value
}

func (v *NullableCustomerReview) Set(val *CustomerReview) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerReview) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerReview) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerReview(val *CustomerReview) *NullableCustomerReview {
	return &NullableCustomerReview{value: val, isSet: true}
}

func (v NullableCustomerReview) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCustomerReview) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
