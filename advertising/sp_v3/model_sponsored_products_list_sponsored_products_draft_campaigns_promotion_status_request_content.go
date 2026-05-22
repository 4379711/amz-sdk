package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsListSponsoredProductsDraftCampaignsPromotionStatusRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsListSponsoredProductsDraftCampaignsPromotionStatusRequestContent{}

// SponsoredProductsListSponsoredProductsDraftCampaignsPromotionStatusRequestContent struct for SponsoredProductsListSponsoredProductsDraftCampaignsPromotionStatusRequestContent
type SponsoredProductsListSponsoredProductsDraftCampaignsPromotionStatusRequestContent struct {
	CampaignIdFilter SponsoredProductsObjectIdFilter `json:"campaignIdFilter"`
	// Number of records to include in the paginated response. Defaults to max page size for given API
	MaxResults *int32 `json:"maxResults,omitempty"`
	// token value allowing to navigate to the next response page
	NextToken *string `json:"nextToken,omitempty"`
}

type _SponsoredProductsListSponsoredProductsDraftCampaignsPromotionStatusRequestContent SponsoredProductsListSponsoredProductsDraftCampaignsPromotionStatusRequestContent

// NewSponsoredProductsListSponsoredProductsDraftCampaignsPromotionStatusRequestContent instantiates a new SponsoredProductsListSponsoredProductsDraftCampaignsPromotionStatusRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsListSponsoredProductsDraftCampaignsPromotionStatusRequestContent(campaignIdFilter SponsoredProductsObjectIdFilter) *SponsoredProductsListSponsoredProductsDraftCampaignsPromotionStatusRequestContent {
	this := SponsoredProductsListSponsoredProductsDraftCampaignsPromotionStatusRequestContent{}
	this.CampaignIdFilter = campaignIdFilter
	return &this
}

// NewSponsoredProductsListSponsoredProductsDraftCampaignsPromotionStatusRequestContentWithDefaults instantiates a new SponsoredProductsListSponsoredProductsDraftCampaignsPromotionStatusRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsListSponsoredProductsDraftCampaignsPromotionStatusRequestContentWithDefaults() *SponsoredProductsListSponsoredProductsDraftCampaignsPromotionStatusRequestContent {
	this := SponsoredProductsListSponsoredProductsDraftCampaignsPromotionStatusRequestContent{}
	return &this
}

// GetCampaignIdFilter returns the CampaignIdFilter field value
func (o *SponsoredProductsListSponsoredProductsDraftCampaignsPromotionStatusRequestContent) GetCampaignIdFilter() SponsoredProductsObjectIdFilter {
	if o == nil {
		var ret SponsoredProductsObjectIdFilter
		return ret
	}

	return o.CampaignIdFilter
}

// GetCampaignIdFilterOk returns a tuple with the CampaignIdFilter field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsDraftCampaignsPromotionStatusRequestContent) GetCampaignIdFilterOk() (*SponsoredProductsObjectIdFilter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignIdFilter, true
}

// SetCampaignIdFilter sets field value
func (o *SponsoredProductsListSponsoredProductsDraftCampaignsPromotionStatusRequestContent) SetCampaignIdFilter(v SponsoredProductsObjectIdFilter) {
	o.CampaignIdFilter = v
}

// GetMaxResults returns the MaxResults field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsDraftCampaignsPromotionStatusRequestContent) GetMaxResults() int32 {
	if o == nil || IsNil(o.MaxResults) {
		var ret int32
		return ret
	}
	return *o.MaxResults
}

// GetMaxResultsOk returns a tuple with the MaxResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsDraftCampaignsPromotionStatusRequestContent) GetMaxResultsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxResults) {
		return nil, false
	}
	return o.MaxResults, true
}

// HasMaxResults returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsDraftCampaignsPromotionStatusRequestContent) HasMaxResults() bool {
	if o != nil && !IsNil(o.MaxResults) {
		return true
	}

	return false
}

// SetMaxResults gets a reference to the given int32 and assigns it to the MaxResults field.
func (o *SponsoredProductsListSponsoredProductsDraftCampaignsPromotionStatusRequestContent) SetMaxResults(v int32) {
	o.MaxResults = &v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsDraftCampaignsPromotionStatusRequestContent) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsDraftCampaignsPromotionStatusRequestContent) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsDraftCampaignsPromotionStatusRequestContent) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *SponsoredProductsListSponsoredProductsDraftCampaignsPromotionStatusRequestContent) SetNextToken(v string) {
	o.NextToken = &v
}

func (o SponsoredProductsListSponsoredProductsDraftCampaignsPromotionStatusRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["campaignIdFilter"] = o.CampaignIdFilter
	if !IsNil(o.MaxResults) {
		toSerialize["maxResults"] = o.MaxResults
	}
	if !IsNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	return toSerialize, nil
}

type NullableSponsoredProductsListSponsoredProductsDraftCampaignsPromotionStatusRequestContent struct {
	value *SponsoredProductsListSponsoredProductsDraftCampaignsPromotionStatusRequestContent
	isSet bool
}

func (v NullableSponsoredProductsListSponsoredProductsDraftCampaignsPromotionStatusRequestContent) Get() *SponsoredProductsListSponsoredProductsDraftCampaignsPromotionStatusRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsListSponsoredProductsDraftCampaignsPromotionStatusRequestContent) Set(val *SponsoredProductsListSponsoredProductsDraftCampaignsPromotionStatusRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsListSponsoredProductsDraftCampaignsPromotionStatusRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsListSponsoredProductsDraftCampaignsPromotionStatusRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsListSponsoredProductsDraftCampaignsPromotionStatusRequestContent(val *SponsoredProductsListSponsoredProductsDraftCampaignsPromotionStatusRequestContent) *NullableSponsoredProductsListSponsoredProductsDraftCampaignsPromotionStatusRequestContent {
	return &NullableSponsoredProductsListSponsoredProductsDraftCampaignsPromotionStatusRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsListSponsoredProductsDraftCampaignsPromotionStatusRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsListSponsoredProductsDraftCampaignsPromotionStatusRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
