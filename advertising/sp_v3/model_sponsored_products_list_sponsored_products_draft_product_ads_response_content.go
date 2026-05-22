package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsListSponsoredProductsDraftProductAdsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsListSponsoredProductsDraftProductAdsResponseContent{}

// SponsoredProductsListSponsoredProductsDraftProductAdsResponseContent struct for SponsoredProductsListSponsoredProductsDraftProductAdsResponseContent
type SponsoredProductsListSponsoredProductsDraftProductAdsResponseContent struct {
	// The total number of entities
	TotalResults *int64 `json:"totalResults,omitempty"`
	// token value allowing to navigate to the next response page
	NextToken  *string                           `json:"nextToken,omitempty"`
	ProductAds []SponsoredProductsDraftProductAd `json:"productAds,omitempty"`
}

// NewSponsoredProductsListSponsoredProductsDraftProductAdsResponseContent instantiates a new SponsoredProductsListSponsoredProductsDraftProductAdsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsListSponsoredProductsDraftProductAdsResponseContent() *SponsoredProductsListSponsoredProductsDraftProductAdsResponseContent {
	this := SponsoredProductsListSponsoredProductsDraftProductAdsResponseContent{}
	return &this
}

// NewSponsoredProductsListSponsoredProductsDraftProductAdsResponseContentWithDefaults instantiates a new SponsoredProductsListSponsoredProductsDraftProductAdsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsListSponsoredProductsDraftProductAdsResponseContentWithDefaults() *SponsoredProductsListSponsoredProductsDraftProductAdsResponseContent {
	this := SponsoredProductsListSponsoredProductsDraftProductAdsResponseContent{}
	return &this
}

// GetTotalResults returns the TotalResults field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsDraftProductAdsResponseContent) GetTotalResults() int64 {
	if o == nil || IsNil(o.TotalResults) {
		var ret int64
		return ret
	}
	return *o.TotalResults
}

// GetTotalResultsOk returns a tuple with the TotalResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsDraftProductAdsResponseContent) GetTotalResultsOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalResults) {
		return nil, false
	}
	return o.TotalResults, true
}

// HasTotalResults returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsDraftProductAdsResponseContent) HasTotalResults() bool {
	if o != nil && !IsNil(o.TotalResults) {
		return true
	}

	return false
}

// SetTotalResults gets a reference to the given int64 and assigns it to the TotalResults field.
func (o *SponsoredProductsListSponsoredProductsDraftProductAdsResponseContent) SetTotalResults(v int64) {
	o.TotalResults = &v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsDraftProductAdsResponseContent) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsDraftProductAdsResponseContent) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsDraftProductAdsResponseContent) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *SponsoredProductsListSponsoredProductsDraftProductAdsResponseContent) SetNextToken(v string) {
	o.NextToken = &v
}

// GetProductAds returns the ProductAds field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsDraftProductAdsResponseContent) GetProductAds() []SponsoredProductsDraftProductAd {
	if o == nil || IsNil(o.ProductAds) {
		var ret []SponsoredProductsDraftProductAd
		return ret
	}
	return o.ProductAds
}

// GetProductAdsOk returns a tuple with the ProductAds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsDraftProductAdsResponseContent) GetProductAdsOk() ([]SponsoredProductsDraftProductAd, bool) {
	if o == nil || IsNil(o.ProductAds) {
		return nil, false
	}
	return o.ProductAds, true
}

// HasProductAds returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsDraftProductAdsResponseContent) HasProductAds() bool {
	if o != nil && !IsNil(o.ProductAds) {
		return true
	}

	return false
}

// SetProductAds gets a reference to the given []SponsoredProductsDraftProductAd and assigns it to the ProductAds field.
func (o *SponsoredProductsListSponsoredProductsDraftProductAdsResponseContent) SetProductAds(v []SponsoredProductsDraftProductAd) {
	o.ProductAds = v
}

func (o SponsoredProductsListSponsoredProductsDraftProductAdsResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TotalResults) {
		toSerialize["totalResults"] = o.TotalResults
	}
	if !IsNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	if !IsNil(o.ProductAds) {
		toSerialize["productAds"] = o.ProductAds
	}
	return toSerialize, nil
}

type NullableSponsoredProductsListSponsoredProductsDraftProductAdsResponseContent struct {
	value *SponsoredProductsListSponsoredProductsDraftProductAdsResponseContent
	isSet bool
}

func (v NullableSponsoredProductsListSponsoredProductsDraftProductAdsResponseContent) Get() *SponsoredProductsListSponsoredProductsDraftProductAdsResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsListSponsoredProductsDraftProductAdsResponseContent) Set(val *SponsoredProductsListSponsoredProductsDraftProductAdsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsListSponsoredProductsDraftProductAdsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsListSponsoredProductsDraftProductAdsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsListSponsoredProductsDraftProductAdsResponseContent(val *SponsoredProductsListSponsoredProductsDraftProductAdsResponseContent) *NullableSponsoredProductsListSponsoredProductsDraftProductAdsResponseContent {
	return &NullableSponsoredProductsListSponsoredProductsDraftProductAdsResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsListSponsoredProductsDraftProductAdsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsListSponsoredProductsDraftProductAdsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
