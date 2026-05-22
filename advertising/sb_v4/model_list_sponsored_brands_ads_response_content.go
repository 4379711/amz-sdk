package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the ListSponsoredBrandsAdsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListSponsoredBrandsAdsResponseContent{}

// ListSponsoredBrandsAdsResponseContent struct for ListSponsoredBrandsAdsResponseContent
type ListSponsoredBrandsAdsResponseContent struct {
	Ads []Ad `json:"ads,omitempty"`
	// The total number of entities.
	TotalResults *float32 `json:"totalResults,omitempty"`
	// Token value allowing to navigate to the next response page.
	NextToken *string `json:"nextToken,omitempty"`
}

// NewListSponsoredBrandsAdsResponseContent instantiates a new ListSponsoredBrandsAdsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListSponsoredBrandsAdsResponseContent() *ListSponsoredBrandsAdsResponseContent {
	this := ListSponsoredBrandsAdsResponseContent{}
	return &this
}

// NewListSponsoredBrandsAdsResponseContentWithDefaults instantiates a new ListSponsoredBrandsAdsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListSponsoredBrandsAdsResponseContentWithDefaults() *ListSponsoredBrandsAdsResponseContent {
	this := ListSponsoredBrandsAdsResponseContent{}
	return &this
}

// GetAds returns the Ads field value if set, zero value otherwise.
func (o *ListSponsoredBrandsAdsResponseContent) GetAds() []Ad {
	if o == nil || IsNil(o.Ads) {
		var ret []Ad
		return ret
	}
	return o.Ads
}

// GetAdsOk returns a tuple with the Ads field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSponsoredBrandsAdsResponseContent) GetAdsOk() ([]Ad, bool) {
	if o == nil || IsNil(o.Ads) {
		return nil, false
	}
	return o.Ads, true
}

// HasAds returns a boolean if a field has been set.
func (o *ListSponsoredBrandsAdsResponseContent) HasAds() bool {
	if o != nil && !IsNil(o.Ads) {
		return true
	}

	return false
}

// SetAds gets a reference to the given []Ad and assigns it to the Ads field.
func (o *ListSponsoredBrandsAdsResponseContent) SetAds(v []Ad) {
	o.Ads = v
}

// GetTotalResults returns the TotalResults field value if set, zero value otherwise.
func (o *ListSponsoredBrandsAdsResponseContent) GetTotalResults() float32 {
	if o == nil || IsNil(o.TotalResults) {
		var ret float32
		return ret
	}
	return *o.TotalResults
}

// GetTotalResultsOk returns a tuple with the TotalResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSponsoredBrandsAdsResponseContent) GetTotalResultsOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalResults) {
		return nil, false
	}
	return o.TotalResults, true
}

// HasTotalResults returns a boolean if a field has been set.
func (o *ListSponsoredBrandsAdsResponseContent) HasTotalResults() bool {
	if o != nil && !IsNil(o.TotalResults) {
		return true
	}

	return false
}

// SetTotalResults gets a reference to the given float32 and assigns it to the TotalResults field.
func (o *ListSponsoredBrandsAdsResponseContent) SetTotalResults(v float32) {
	o.TotalResults = &v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *ListSponsoredBrandsAdsResponseContent) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSponsoredBrandsAdsResponseContent) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *ListSponsoredBrandsAdsResponseContent) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *ListSponsoredBrandsAdsResponseContent) SetNextToken(v string) {
	o.NextToken = &v
}

func (o ListSponsoredBrandsAdsResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ads) {
		toSerialize["ads"] = o.Ads
	}
	if !IsNil(o.TotalResults) {
		toSerialize["totalResults"] = o.TotalResults
	}
	if !IsNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	return toSerialize, nil
}

type NullableListSponsoredBrandsAdsResponseContent struct {
	value *ListSponsoredBrandsAdsResponseContent
	isSet bool
}

func (v NullableListSponsoredBrandsAdsResponseContent) Get() *ListSponsoredBrandsAdsResponseContent {
	return v.value
}

func (v *NullableListSponsoredBrandsAdsResponseContent) Set(val *ListSponsoredBrandsAdsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableListSponsoredBrandsAdsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableListSponsoredBrandsAdsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListSponsoredBrandsAdsResponseContent(val *ListSponsoredBrandsAdsResponseContent) *NullableListSponsoredBrandsAdsResponseContent {
	return &NullableListSponsoredBrandsAdsResponseContent{value: val, isSet: true}
}

func (v NullableListSponsoredBrandsAdsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableListSponsoredBrandsAdsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
