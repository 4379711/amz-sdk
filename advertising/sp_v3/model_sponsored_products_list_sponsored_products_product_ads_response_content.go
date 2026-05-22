package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsListSponsoredProductsProductAdsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsListSponsoredProductsProductAdsResponseContent{}

// SponsoredProductsListSponsoredProductsProductAdsResponseContent struct for SponsoredProductsListSponsoredProductsProductAdsResponseContent
type SponsoredProductsListSponsoredProductsProductAdsResponseContent struct {
	// The total number of entities
	TotalResults *int64 `json:"totalResults,omitempty"`
	// token value allowing to navigate to the next response page
	NextToken  *string                      `json:"nextToken,omitempty"`
	ProductAds []SponsoredProductsProductAd `json:"productAds,omitempty"`
}

// NewSponsoredProductsListSponsoredProductsProductAdsResponseContent instantiates a new SponsoredProductsListSponsoredProductsProductAdsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsListSponsoredProductsProductAdsResponseContent() *SponsoredProductsListSponsoredProductsProductAdsResponseContent {
	this := SponsoredProductsListSponsoredProductsProductAdsResponseContent{}
	return &this
}

// NewSponsoredProductsListSponsoredProductsProductAdsResponseContentWithDefaults instantiates a new SponsoredProductsListSponsoredProductsProductAdsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsListSponsoredProductsProductAdsResponseContentWithDefaults() *SponsoredProductsListSponsoredProductsProductAdsResponseContent {
	this := SponsoredProductsListSponsoredProductsProductAdsResponseContent{}
	return &this
}

// GetTotalResults returns the TotalResults field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsProductAdsResponseContent) GetTotalResults() int64 {
	if o == nil || IsNil(o.TotalResults) {
		var ret int64
		return ret
	}
	return *o.TotalResults
}

// GetTotalResultsOk returns a tuple with the TotalResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsProductAdsResponseContent) GetTotalResultsOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalResults) {
		return nil, false
	}
	return o.TotalResults, true
}

// HasTotalResults returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsProductAdsResponseContent) HasTotalResults() bool {
	if o != nil && !IsNil(o.TotalResults) {
		return true
	}

	return false
}

// SetTotalResults gets a reference to the given int64 and assigns it to the TotalResults field.
func (o *SponsoredProductsListSponsoredProductsProductAdsResponseContent) SetTotalResults(v int64) {
	o.TotalResults = &v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsProductAdsResponseContent) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsProductAdsResponseContent) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsProductAdsResponseContent) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *SponsoredProductsListSponsoredProductsProductAdsResponseContent) SetNextToken(v string) {
	o.NextToken = &v
}

// GetProductAds returns the ProductAds field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsProductAdsResponseContent) GetProductAds() []SponsoredProductsProductAd {
	if o == nil || IsNil(o.ProductAds) {
		var ret []SponsoredProductsProductAd
		return ret
	}
	return o.ProductAds
}

// GetProductAdsOk returns a tuple with the ProductAds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsProductAdsResponseContent) GetProductAdsOk() ([]SponsoredProductsProductAd, bool) {
	if o == nil || IsNil(o.ProductAds) {
		return nil, false
	}
	return o.ProductAds, true
}

// HasProductAds returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsProductAdsResponseContent) HasProductAds() bool {
	if o != nil && !IsNil(o.ProductAds) {
		return true
	}

	return false
}

// SetProductAds gets a reference to the given []SponsoredProductsProductAd and assigns it to the ProductAds field.
func (o *SponsoredProductsListSponsoredProductsProductAdsResponseContent) SetProductAds(v []SponsoredProductsProductAd) {
	o.ProductAds = v
}

func (o SponsoredProductsListSponsoredProductsProductAdsResponseContent) ToMap() (map[string]interface{}, error) {
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

type NullableSponsoredProductsListSponsoredProductsProductAdsResponseContent struct {
	value *SponsoredProductsListSponsoredProductsProductAdsResponseContent
	isSet bool
}

func (v NullableSponsoredProductsListSponsoredProductsProductAdsResponseContent) Get() *SponsoredProductsListSponsoredProductsProductAdsResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsListSponsoredProductsProductAdsResponseContent) Set(val *SponsoredProductsListSponsoredProductsProductAdsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsListSponsoredProductsProductAdsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsListSponsoredProductsProductAdsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsListSponsoredProductsProductAdsResponseContent(val *SponsoredProductsListSponsoredProductsProductAdsResponseContent) *NullableSponsoredProductsListSponsoredProductsProductAdsResponseContent {
	return &NullableSponsoredProductsListSponsoredProductsProductAdsResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsListSponsoredProductsProductAdsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsListSponsoredProductsProductAdsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
