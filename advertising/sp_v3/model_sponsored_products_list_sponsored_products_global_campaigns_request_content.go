package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsListSponsoredProductsGlobalCampaignsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsListSponsoredProductsGlobalCampaignsRequestContent{}

// SponsoredProductsListSponsoredProductsGlobalCampaignsRequestContent struct for SponsoredProductsListSponsoredProductsGlobalCampaignsRequestContent
type SponsoredProductsListSponsoredProductsGlobalCampaignsRequestContent struct {
	CampaignIdFilter *SponsoredProductsObjectIdFilter `json:"campaignIdFilter,omitempty"`
	// Number of records to include in the paginated response. Defaults to max page size for given API
	MaxResults *int32 `json:"maxResults,omitempty"`
	// token value allowing to navigate to the next response page
	NextToken *string `json:"nextToken,omitempty"`
	// Whether to get entity with extended data fields such as creationDate, lastUpdateDate, servingStatus
	IncludeExtendedDataFields *bool `json:"includeExtendedDataFields,omitempty"`
}

// NewSponsoredProductsListSponsoredProductsGlobalCampaignsRequestContent instantiates a new SponsoredProductsListSponsoredProductsGlobalCampaignsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsListSponsoredProductsGlobalCampaignsRequestContent() *SponsoredProductsListSponsoredProductsGlobalCampaignsRequestContent {
	this := SponsoredProductsListSponsoredProductsGlobalCampaignsRequestContent{}
	return &this
}

// NewSponsoredProductsListSponsoredProductsGlobalCampaignsRequestContentWithDefaults instantiates a new SponsoredProductsListSponsoredProductsGlobalCampaignsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsListSponsoredProductsGlobalCampaignsRequestContentWithDefaults() *SponsoredProductsListSponsoredProductsGlobalCampaignsRequestContent {
	this := SponsoredProductsListSponsoredProductsGlobalCampaignsRequestContent{}
	return &this
}

// GetCampaignIdFilter returns the CampaignIdFilter field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsGlobalCampaignsRequestContent) GetCampaignIdFilter() SponsoredProductsObjectIdFilter {
	if o == nil || IsNil(o.CampaignIdFilter) {
		var ret SponsoredProductsObjectIdFilter
		return ret
	}
	return *o.CampaignIdFilter
}

// GetCampaignIdFilterOk returns a tuple with the CampaignIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalCampaignsRequestContent) GetCampaignIdFilterOk() (*SponsoredProductsObjectIdFilter, bool) {
	if o == nil || IsNil(o.CampaignIdFilter) {
		return nil, false
	}
	return o.CampaignIdFilter, true
}

// HasCampaignIdFilter returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalCampaignsRequestContent) HasCampaignIdFilter() bool {
	if o != nil && !IsNil(o.CampaignIdFilter) {
		return true
	}

	return false
}

// SetCampaignIdFilter gets a reference to the given SponsoredProductsObjectIdFilter and assigns it to the CampaignIdFilter field.
func (o *SponsoredProductsListSponsoredProductsGlobalCampaignsRequestContent) SetCampaignIdFilter(v SponsoredProductsObjectIdFilter) {
	o.CampaignIdFilter = &v
}

// GetMaxResults returns the MaxResults field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsGlobalCampaignsRequestContent) GetMaxResults() int32 {
	if o == nil || IsNil(o.MaxResults) {
		var ret int32
		return ret
	}
	return *o.MaxResults
}

// GetMaxResultsOk returns a tuple with the MaxResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalCampaignsRequestContent) GetMaxResultsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxResults) {
		return nil, false
	}
	return o.MaxResults, true
}

// HasMaxResults returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalCampaignsRequestContent) HasMaxResults() bool {
	if o != nil && !IsNil(o.MaxResults) {
		return true
	}

	return false
}

// SetMaxResults gets a reference to the given int32 and assigns it to the MaxResults field.
func (o *SponsoredProductsListSponsoredProductsGlobalCampaignsRequestContent) SetMaxResults(v int32) {
	o.MaxResults = &v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsGlobalCampaignsRequestContent) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalCampaignsRequestContent) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalCampaignsRequestContent) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *SponsoredProductsListSponsoredProductsGlobalCampaignsRequestContent) SetNextToken(v string) {
	o.NextToken = &v
}

// GetIncludeExtendedDataFields returns the IncludeExtendedDataFields field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsGlobalCampaignsRequestContent) GetIncludeExtendedDataFields() bool {
	if o == nil || IsNil(o.IncludeExtendedDataFields) {
		var ret bool
		return ret
	}
	return *o.IncludeExtendedDataFields
}

// GetIncludeExtendedDataFieldsOk returns a tuple with the IncludeExtendedDataFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalCampaignsRequestContent) GetIncludeExtendedDataFieldsOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeExtendedDataFields) {
		return nil, false
	}
	return o.IncludeExtendedDataFields, true
}

// HasIncludeExtendedDataFields returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalCampaignsRequestContent) HasIncludeExtendedDataFields() bool {
	if o != nil && !IsNil(o.IncludeExtendedDataFields) {
		return true
	}

	return false
}

// SetIncludeExtendedDataFields gets a reference to the given bool and assigns it to the IncludeExtendedDataFields field.
func (o *SponsoredProductsListSponsoredProductsGlobalCampaignsRequestContent) SetIncludeExtendedDataFields(v bool) {
	o.IncludeExtendedDataFields = &v
}

func (o SponsoredProductsListSponsoredProductsGlobalCampaignsRequestContent) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.IncludeExtendedDataFields) {
		toSerialize["includeExtendedDataFields"] = o.IncludeExtendedDataFields
	}
	return toSerialize, nil
}

type NullableSponsoredProductsListSponsoredProductsGlobalCampaignsRequestContent struct {
	value *SponsoredProductsListSponsoredProductsGlobalCampaignsRequestContent
	isSet bool
}

func (v NullableSponsoredProductsListSponsoredProductsGlobalCampaignsRequestContent) Get() *SponsoredProductsListSponsoredProductsGlobalCampaignsRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsListSponsoredProductsGlobalCampaignsRequestContent) Set(val *SponsoredProductsListSponsoredProductsGlobalCampaignsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsListSponsoredProductsGlobalCampaignsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsListSponsoredProductsGlobalCampaignsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsListSponsoredProductsGlobalCampaignsRequestContent(val *SponsoredProductsListSponsoredProductsGlobalCampaignsRequestContent) *NullableSponsoredProductsListSponsoredProductsGlobalCampaignsRequestContent {
	return &NullableSponsoredProductsListSponsoredProductsGlobalCampaignsRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsListSponsoredProductsGlobalCampaignsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsListSponsoredProductsGlobalCampaignsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
