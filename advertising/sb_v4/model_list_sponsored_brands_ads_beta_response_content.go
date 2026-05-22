package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the ListSponsoredBrandsAdsBetaResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListSponsoredBrandsAdsBetaResponseContent{}

// ListSponsoredBrandsAdsBetaResponseContent struct for ListSponsoredBrandsAdsBetaResponseContent
type ListSponsoredBrandsAdsBetaResponseContent struct {
	Ads []Ad `json:"ads,omitempty"`
	// The total number of entities.
	TotalResults *float32 `json:"totalResults,omitempty"`
	// Token value allowing to navigate to the next response page.
	NextToken *string `json:"nextToken,omitempty"`
}

// NewListSponsoredBrandsAdsBetaResponseContent instantiates a new ListSponsoredBrandsAdsBetaResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListSponsoredBrandsAdsBetaResponseContent() *ListSponsoredBrandsAdsBetaResponseContent {
	this := ListSponsoredBrandsAdsBetaResponseContent{}
	return &this
}

// NewListSponsoredBrandsAdsBetaResponseContentWithDefaults instantiates a new ListSponsoredBrandsAdsBetaResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListSponsoredBrandsAdsBetaResponseContentWithDefaults() *ListSponsoredBrandsAdsBetaResponseContent {
	this := ListSponsoredBrandsAdsBetaResponseContent{}
	return &this
}

// GetAds returns the Ads field value if set, zero value otherwise.
func (o *ListSponsoredBrandsAdsBetaResponseContent) GetAds() []Ad {
	if o == nil || IsNil(o.Ads) {
		var ret []Ad
		return ret
	}
	return o.Ads
}

// GetAdsOk returns a tuple with the Ads field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSponsoredBrandsAdsBetaResponseContent) GetAdsOk() ([]Ad, bool) {
	if o == nil || IsNil(o.Ads) {
		return nil, false
	}
	return o.Ads, true
}

// HasAds returns a boolean if a field has been set.
func (o *ListSponsoredBrandsAdsBetaResponseContent) HasAds() bool {
	if o != nil && !IsNil(o.Ads) {
		return true
	}

	return false
}

// SetAds gets a reference to the given []Ad and assigns it to the Ads field.
func (o *ListSponsoredBrandsAdsBetaResponseContent) SetAds(v []Ad) {
	o.Ads = v
}

// GetTotalResults returns the TotalResults field value if set, zero value otherwise.
func (o *ListSponsoredBrandsAdsBetaResponseContent) GetTotalResults() float32 {
	if o == nil || IsNil(o.TotalResults) {
		var ret float32
		return ret
	}
	return *o.TotalResults
}

// GetTotalResultsOk returns a tuple with the TotalResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSponsoredBrandsAdsBetaResponseContent) GetTotalResultsOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalResults) {
		return nil, false
	}
	return o.TotalResults, true
}

// HasTotalResults returns a boolean if a field has been set.
func (o *ListSponsoredBrandsAdsBetaResponseContent) HasTotalResults() bool {
	if o != nil && !IsNil(o.TotalResults) {
		return true
	}

	return false
}

// SetTotalResults gets a reference to the given float32 and assigns it to the TotalResults field.
func (o *ListSponsoredBrandsAdsBetaResponseContent) SetTotalResults(v float32) {
	o.TotalResults = &v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *ListSponsoredBrandsAdsBetaResponseContent) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSponsoredBrandsAdsBetaResponseContent) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *ListSponsoredBrandsAdsBetaResponseContent) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *ListSponsoredBrandsAdsBetaResponseContent) SetNextToken(v string) {
	o.NextToken = &v
}

func (o ListSponsoredBrandsAdsBetaResponseContent) ToMap() (map[string]interface{}, error) {
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

type NullableListSponsoredBrandsAdsBetaResponseContent struct {
	value *ListSponsoredBrandsAdsBetaResponseContent
	isSet bool
}

func (v NullableListSponsoredBrandsAdsBetaResponseContent) Get() *ListSponsoredBrandsAdsBetaResponseContent {
	return v.value
}

func (v *NullableListSponsoredBrandsAdsBetaResponseContent) Set(val *ListSponsoredBrandsAdsBetaResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableListSponsoredBrandsAdsBetaResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableListSponsoredBrandsAdsBetaResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListSponsoredBrandsAdsBetaResponseContent(val *ListSponsoredBrandsAdsBetaResponseContent) *NullableListSponsoredBrandsAdsBetaResponseContent {
	return &NullableListSponsoredBrandsAdsBetaResponseContent{value: val, isSet: true}
}

func (v NullableListSponsoredBrandsAdsBetaResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableListSponsoredBrandsAdsBetaResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
