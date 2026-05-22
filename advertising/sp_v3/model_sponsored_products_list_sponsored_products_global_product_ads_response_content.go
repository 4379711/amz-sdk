package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsListSponsoredProductsGlobalProductAdsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsListSponsoredProductsGlobalProductAdsResponseContent{}

// SponsoredProductsListSponsoredProductsGlobalProductAdsResponseContent struct for SponsoredProductsListSponsoredProductsGlobalProductAdsResponseContent
type SponsoredProductsListSponsoredProductsGlobalProductAdsResponseContent struct {
	// The total number of entities
	TotalResults *int64 `json:"totalResults,omitempty"`
	// token value allowing to navigate to the next response page
	NextToken  *string                            `json:"nextToken,omitempty"`
	ProductAds []SponsoredProductsGlobalProductAd `json:"productAds,omitempty"`
}

// NewSponsoredProductsListSponsoredProductsGlobalProductAdsResponseContent instantiates a new SponsoredProductsListSponsoredProductsGlobalProductAdsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsListSponsoredProductsGlobalProductAdsResponseContent() *SponsoredProductsListSponsoredProductsGlobalProductAdsResponseContent {
	this := SponsoredProductsListSponsoredProductsGlobalProductAdsResponseContent{}
	return &this
}

// NewSponsoredProductsListSponsoredProductsGlobalProductAdsResponseContentWithDefaults instantiates a new SponsoredProductsListSponsoredProductsGlobalProductAdsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsListSponsoredProductsGlobalProductAdsResponseContentWithDefaults() *SponsoredProductsListSponsoredProductsGlobalProductAdsResponseContent {
	this := SponsoredProductsListSponsoredProductsGlobalProductAdsResponseContent{}
	return &this
}

// GetTotalResults returns the TotalResults field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsGlobalProductAdsResponseContent) GetTotalResults() int64 {
	if o == nil || IsNil(o.TotalResults) {
		var ret int64
		return ret
	}
	return *o.TotalResults
}

// GetTotalResultsOk returns a tuple with the TotalResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalProductAdsResponseContent) GetTotalResultsOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalResults) {
		return nil, false
	}
	return o.TotalResults, true
}

// HasTotalResults returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalProductAdsResponseContent) HasTotalResults() bool {
	if o != nil && !IsNil(o.TotalResults) {
		return true
	}

	return false
}

// SetTotalResults gets a reference to the given int64 and assigns it to the TotalResults field.
func (o *SponsoredProductsListSponsoredProductsGlobalProductAdsResponseContent) SetTotalResults(v int64) {
	o.TotalResults = &v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsGlobalProductAdsResponseContent) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalProductAdsResponseContent) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalProductAdsResponseContent) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *SponsoredProductsListSponsoredProductsGlobalProductAdsResponseContent) SetNextToken(v string) {
	o.NextToken = &v
}

// GetProductAds returns the ProductAds field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsGlobalProductAdsResponseContent) GetProductAds() []SponsoredProductsGlobalProductAd {
	if o == nil || IsNil(o.ProductAds) {
		var ret []SponsoredProductsGlobalProductAd
		return ret
	}
	return o.ProductAds
}

// GetProductAdsOk returns a tuple with the ProductAds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalProductAdsResponseContent) GetProductAdsOk() ([]SponsoredProductsGlobalProductAd, bool) {
	if o == nil || IsNil(o.ProductAds) {
		return nil, false
	}
	return o.ProductAds, true
}

// HasProductAds returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalProductAdsResponseContent) HasProductAds() bool {
	if o != nil && !IsNil(o.ProductAds) {
		return true
	}

	return false
}

// SetProductAds gets a reference to the given []SponsoredProductsGlobalProductAd and assigns it to the ProductAds field.
func (o *SponsoredProductsListSponsoredProductsGlobalProductAdsResponseContent) SetProductAds(v []SponsoredProductsGlobalProductAd) {
	o.ProductAds = v
}

func (o SponsoredProductsListSponsoredProductsGlobalProductAdsResponseContent) ToMap() (map[string]interface{}, error) {
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

type NullableSponsoredProductsListSponsoredProductsGlobalProductAdsResponseContent struct {
	value *SponsoredProductsListSponsoredProductsGlobalProductAdsResponseContent
	isSet bool
}

func (v NullableSponsoredProductsListSponsoredProductsGlobalProductAdsResponseContent) Get() *SponsoredProductsListSponsoredProductsGlobalProductAdsResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsListSponsoredProductsGlobalProductAdsResponseContent) Set(val *SponsoredProductsListSponsoredProductsGlobalProductAdsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsListSponsoredProductsGlobalProductAdsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsListSponsoredProductsGlobalProductAdsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsListSponsoredProductsGlobalProductAdsResponseContent(val *SponsoredProductsListSponsoredProductsGlobalProductAdsResponseContent) *NullableSponsoredProductsListSponsoredProductsGlobalProductAdsResponseContent {
	return &NullableSponsoredProductsListSponsoredProductsGlobalProductAdsResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsListSponsoredProductsGlobalProductAdsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsListSponsoredProductsGlobalProductAdsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
