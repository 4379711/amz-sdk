package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsListSponsoredProductsDraftProductAdsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsListSponsoredProductsDraftProductAdsRequestContent{}

// SponsoredProductsListSponsoredProductsDraftProductAdsRequestContent struct for SponsoredProductsListSponsoredProductsDraftProductAdsRequestContent
type SponsoredProductsListSponsoredProductsDraftProductAdsRequestContent struct {
	CampaignIdFilter *SponsoredProductsReducedObjectIdFilter `json:"campaignIdFilter,omitempty"`
	// Number of records to include in the paginated response. Defaults to max page size for given API
	MaxResults *int32 `json:"maxResults,omitempty"`
	// token value allowing to navigate to the next response page
	NextToken       *string                                 `json:"nextToken,omitempty"`
	AdIdFilter      *SponsoredProductsObjectIdFilter        `json:"adIdFilter,omitempty"`
	AdGroupIdFilter *SponsoredProductsReducedObjectIdFilter `json:"adGroupIdFilter,omitempty"`
	// Whether to get entity with extended data fields such as creationDate, lastUpdateDate, servingStatus
	IncludeExtendedDataFields *bool `json:"includeExtendedDataFields,omitempty"`
}

// NewSponsoredProductsListSponsoredProductsDraftProductAdsRequestContent instantiates a new SponsoredProductsListSponsoredProductsDraftProductAdsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsListSponsoredProductsDraftProductAdsRequestContent() *SponsoredProductsListSponsoredProductsDraftProductAdsRequestContent {
	this := SponsoredProductsListSponsoredProductsDraftProductAdsRequestContent{}
	return &this
}

// NewSponsoredProductsListSponsoredProductsDraftProductAdsRequestContentWithDefaults instantiates a new SponsoredProductsListSponsoredProductsDraftProductAdsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsListSponsoredProductsDraftProductAdsRequestContentWithDefaults() *SponsoredProductsListSponsoredProductsDraftProductAdsRequestContent {
	this := SponsoredProductsListSponsoredProductsDraftProductAdsRequestContent{}
	return &this
}

// GetCampaignIdFilter returns the CampaignIdFilter field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsDraftProductAdsRequestContent) GetCampaignIdFilter() SponsoredProductsReducedObjectIdFilter {
	if o == nil || IsNil(o.CampaignIdFilter) {
		var ret SponsoredProductsReducedObjectIdFilter
		return ret
	}
	return *o.CampaignIdFilter
}

// GetCampaignIdFilterOk returns a tuple with the CampaignIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsDraftProductAdsRequestContent) GetCampaignIdFilterOk() (*SponsoredProductsReducedObjectIdFilter, bool) {
	if o == nil || IsNil(o.CampaignIdFilter) {
		return nil, false
	}
	return o.CampaignIdFilter, true
}

// HasCampaignIdFilter returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsDraftProductAdsRequestContent) HasCampaignIdFilter() bool {
	if o != nil && !IsNil(o.CampaignIdFilter) {
		return true
	}

	return false
}

// SetCampaignIdFilter gets a reference to the given SponsoredProductsReducedObjectIdFilter and assigns it to the CampaignIdFilter field.
func (o *SponsoredProductsListSponsoredProductsDraftProductAdsRequestContent) SetCampaignIdFilter(v SponsoredProductsReducedObjectIdFilter) {
	o.CampaignIdFilter = &v
}

// GetMaxResults returns the MaxResults field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsDraftProductAdsRequestContent) GetMaxResults() int32 {
	if o == nil || IsNil(o.MaxResults) {
		var ret int32
		return ret
	}
	return *o.MaxResults
}

// GetMaxResultsOk returns a tuple with the MaxResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsDraftProductAdsRequestContent) GetMaxResultsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxResults) {
		return nil, false
	}
	return o.MaxResults, true
}

// HasMaxResults returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsDraftProductAdsRequestContent) HasMaxResults() bool {
	if o != nil && !IsNil(o.MaxResults) {
		return true
	}

	return false
}

// SetMaxResults gets a reference to the given int32 and assigns it to the MaxResults field.
func (o *SponsoredProductsListSponsoredProductsDraftProductAdsRequestContent) SetMaxResults(v int32) {
	o.MaxResults = &v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsDraftProductAdsRequestContent) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsDraftProductAdsRequestContent) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsDraftProductAdsRequestContent) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *SponsoredProductsListSponsoredProductsDraftProductAdsRequestContent) SetNextToken(v string) {
	o.NextToken = &v
}

// GetAdIdFilter returns the AdIdFilter field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsDraftProductAdsRequestContent) GetAdIdFilter() SponsoredProductsObjectIdFilter {
	if o == nil || IsNil(o.AdIdFilter) {
		var ret SponsoredProductsObjectIdFilter
		return ret
	}
	return *o.AdIdFilter
}

// GetAdIdFilterOk returns a tuple with the AdIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsDraftProductAdsRequestContent) GetAdIdFilterOk() (*SponsoredProductsObjectIdFilter, bool) {
	if o == nil || IsNil(o.AdIdFilter) {
		return nil, false
	}
	return o.AdIdFilter, true
}

// HasAdIdFilter returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsDraftProductAdsRequestContent) HasAdIdFilter() bool {
	if o != nil && !IsNil(o.AdIdFilter) {
		return true
	}

	return false
}

// SetAdIdFilter gets a reference to the given SponsoredProductsObjectIdFilter and assigns it to the AdIdFilter field.
func (o *SponsoredProductsListSponsoredProductsDraftProductAdsRequestContent) SetAdIdFilter(v SponsoredProductsObjectIdFilter) {
	o.AdIdFilter = &v
}

// GetAdGroupIdFilter returns the AdGroupIdFilter field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsDraftProductAdsRequestContent) GetAdGroupIdFilter() SponsoredProductsReducedObjectIdFilter {
	if o == nil || IsNil(o.AdGroupIdFilter) {
		var ret SponsoredProductsReducedObjectIdFilter
		return ret
	}
	return *o.AdGroupIdFilter
}

// GetAdGroupIdFilterOk returns a tuple with the AdGroupIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsDraftProductAdsRequestContent) GetAdGroupIdFilterOk() (*SponsoredProductsReducedObjectIdFilter, bool) {
	if o == nil || IsNil(o.AdGroupIdFilter) {
		return nil, false
	}
	return o.AdGroupIdFilter, true
}

// HasAdGroupIdFilter returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsDraftProductAdsRequestContent) HasAdGroupIdFilter() bool {
	if o != nil && !IsNil(o.AdGroupIdFilter) {
		return true
	}

	return false
}

// SetAdGroupIdFilter gets a reference to the given SponsoredProductsReducedObjectIdFilter and assigns it to the AdGroupIdFilter field.
func (o *SponsoredProductsListSponsoredProductsDraftProductAdsRequestContent) SetAdGroupIdFilter(v SponsoredProductsReducedObjectIdFilter) {
	o.AdGroupIdFilter = &v
}

// GetIncludeExtendedDataFields returns the IncludeExtendedDataFields field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsDraftProductAdsRequestContent) GetIncludeExtendedDataFields() bool {
	if o == nil || IsNil(o.IncludeExtendedDataFields) {
		var ret bool
		return ret
	}
	return *o.IncludeExtendedDataFields
}

// GetIncludeExtendedDataFieldsOk returns a tuple with the IncludeExtendedDataFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsDraftProductAdsRequestContent) GetIncludeExtendedDataFieldsOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeExtendedDataFields) {
		return nil, false
	}
	return o.IncludeExtendedDataFields, true
}

// HasIncludeExtendedDataFields returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsDraftProductAdsRequestContent) HasIncludeExtendedDataFields() bool {
	if o != nil && !IsNil(o.IncludeExtendedDataFields) {
		return true
	}

	return false
}

// SetIncludeExtendedDataFields gets a reference to the given bool and assigns it to the IncludeExtendedDataFields field.
func (o *SponsoredProductsListSponsoredProductsDraftProductAdsRequestContent) SetIncludeExtendedDataFields(v bool) {
	o.IncludeExtendedDataFields = &v
}

func (o SponsoredProductsListSponsoredProductsDraftProductAdsRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CampaignIdFilter) {
		toSerialize["campaignIdFilter"] = o.CampaignIdFilter
	}
	if !IsNil(o.MaxResults) {
		toSerialize["maxResults"] = o.MaxResults
	}
	if !IsNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	if !IsNil(o.AdIdFilter) {
		toSerialize["adIdFilter"] = o.AdIdFilter
	}
	if !IsNil(o.AdGroupIdFilter) {
		toSerialize["adGroupIdFilter"] = o.AdGroupIdFilter
	}
	if !IsNil(o.IncludeExtendedDataFields) {
		toSerialize["includeExtendedDataFields"] = o.IncludeExtendedDataFields
	}
	return toSerialize, nil
}

type NullableSponsoredProductsListSponsoredProductsDraftProductAdsRequestContent struct {
	value *SponsoredProductsListSponsoredProductsDraftProductAdsRequestContent
	isSet bool
}

func (v NullableSponsoredProductsListSponsoredProductsDraftProductAdsRequestContent) Get() *SponsoredProductsListSponsoredProductsDraftProductAdsRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsListSponsoredProductsDraftProductAdsRequestContent) Set(val *SponsoredProductsListSponsoredProductsDraftProductAdsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsListSponsoredProductsDraftProductAdsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsListSponsoredProductsDraftProductAdsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsListSponsoredProductsDraftProductAdsRequestContent(val *SponsoredProductsListSponsoredProductsDraftProductAdsRequestContent) *NullableSponsoredProductsListSponsoredProductsDraftProductAdsRequestContent {
	return &NullableSponsoredProductsListSponsoredProductsDraftProductAdsRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsListSponsoredProductsDraftProductAdsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsListSponsoredProductsDraftProductAdsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
