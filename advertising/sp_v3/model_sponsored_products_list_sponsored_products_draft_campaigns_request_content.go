package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsListSponsoredProductsDraftCampaignsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsListSponsoredProductsDraftCampaignsRequestContent{}

// SponsoredProductsListSponsoredProductsDraftCampaignsRequestContent struct for SponsoredProductsListSponsoredProductsDraftCampaignsRequestContent
type SponsoredProductsListSponsoredProductsDraftCampaignsRequestContent struct {
	CampaignIdFilter  *SponsoredProductsObjectIdFilter        `json:"campaignIdFilter,omitempty"`
	PortfolioIdFilter *SponsoredProductsReducedObjectIdFilter `json:"portfolioIdFilter,omitempty"`
	// Number of records to include in the paginated response. Defaults to max page size for given API
	MaxResults *int32 `json:"maxResults,omitempty"`
	// token value allowing to navigate to the next response page
	NextToken *string `json:"nextToken,omitempty"`
	// Whether to get entity with extended data fields such as creationDate, lastUpdateDate, servingStatus
	IncludeExtendedDataFields *bool                        `json:"includeExtendedDataFields,omitempty"`
	NameFilter                *SponsoredProductsNameFilter `json:"nameFilter,omitempty"`
}

// NewSponsoredProductsListSponsoredProductsDraftCampaignsRequestContent instantiates a new SponsoredProductsListSponsoredProductsDraftCampaignsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsListSponsoredProductsDraftCampaignsRequestContent() *SponsoredProductsListSponsoredProductsDraftCampaignsRequestContent {
	this := SponsoredProductsListSponsoredProductsDraftCampaignsRequestContent{}
	return &this
}

// NewSponsoredProductsListSponsoredProductsDraftCampaignsRequestContentWithDefaults instantiates a new SponsoredProductsListSponsoredProductsDraftCampaignsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsListSponsoredProductsDraftCampaignsRequestContentWithDefaults() *SponsoredProductsListSponsoredProductsDraftCampaignsRequestContent {
	this := SponsoredProductsListSponsoredProductsDraftCampaignsRequestContent{}
	return &this
}

// GetCampaignIdFilter returns the CampaignIdFilter field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsDraftCampaignsRequestContent) GetCampaignIdFilter() SponsoredProductsObjectIdFilter {
	if o == nil || IsNil(o.CampaignIdFilter) {
		var ret SponsoredProductsObjectIdFilter
		return ret
	}
	return *o.CampaignIdFilter
}

// GetCampaignIdFilterOk returns a tuple with the CampaignIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsDraftCampaignsRequestContent) GetCampaignIdFilterOk() (*SponsoredProductsObjectIdFilter, bool) {
	if o == nil || IsNil(o.CampaignIdFilter) {
		return nil, false
	}
	return o.CampaignIdFilter, true
}

// HasCampaignIdFilter returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsDraftCampaignsRequestContent) HasCampaignIdFilter() bool {
	if o != nil && !IsNil(o.CampaignIdFilter) {
		return true
	}

	return false
}

// SetCampaignIdFilter gets a reference to the given SponsoredProductsObjectIdFilter and assigns it to the CampaignIdFilter field.
func (o *SponsoredProductsListSponsoredProductsDraftCampaignsRequestContent) SetCampaignIdFilter(v SponsoredProductsObjectIdFilter) {
	o.CampaignIdFilter = &v
}

// GetPortfolioIdFilter returns the PortfolioIdFilter field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsDraftCampaignsRequestContent) GetPortfolioIdFilter() SponsoredProductsReducedObjectIdFilter {
	if o == nil || IsNil(o.PortfolioIdFilter) {
		var ret SponsoredProductsReducedObjectIdFilter
		return ret
	}
	return *o.PortfolioIdFilter
}

// GetPortfolioIdFilterOk returns a tuple with the PortfolioIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsDraftCampaignsRequestContent) GetPortfolioIdFilterOk() (*SponsoredProductsReducedObjectIdFilter, bool) {
	if o == nil || IsNil(o.PortfolioIdFilter) {
		return nil, false
	}
	return o.PortfolioIdFilter, true
}

// HasPortfolioIdFilter returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsDraftCampaignsRequestContent) HasPortfolioIdFilter() bool {
	if o != nil && !IsNil(o.PortfolioIdFilter) {
		return true
	}

	return false
}

// SetPortfolioIdFilter gets a reference to the given SponsoredProductsReducedObjectIdFilter and assigns it to the PortfolioIdFilter field.
func (o *SponsoredProductsListSponsoredProductsDraftCampaignsRequestContent) SetPortfolioIdFilter(v SponsoredProductsReducedObjectIdFilter) {
	o.PortfolioIdFilter = &v
}

// GetMaxResults returns the MaxResults field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsDraftCampaignsRequestContent) GetMaxResults() int32 {
	if o == nil || IsNil(o.MaxResults) {
		var ret int32
		return ret
	}
	return *o.MaxResults
}

// GetMaxResultsOk returns a tuple with the MaxResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsDraftCampaignsRequestContent) GetMaxResultsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxResults) {
		return nil, false
	}
	return o.MaxResults, true
}

// HasMaxResults returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsDraftCampaignsRequestContent) HasMaxResults() bool {
	if o != nil && !IsNil(o.MaxResults) {
		return true
	}

	return false
}

// SetMaxResults gets a reference to the given int32 and assigns it to the MaxResults field.
func (o *SponsoredProductsListSponsoredProductsDraftCampaignsRequestContent) SetMaxResults(v int32) {
	o.MaxResults = &v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsDraftCampaignsRequestContent) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsDraftCampaignsRequestContent) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsDraftCampaignsRequestContent) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *SponsoredProductsListSponsoredProductsDraftCampaignsRequestContent) SetNextToken(v string) {
	o.NextToken = &v
}

// GetIncludeExtendedDataFields returns the IncludeExtendedDataFields field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsDraftCampaignsRequestContent) GetIncludeExtendedDataFields() bool {
	if o == nil || IsNil(o.IncludeExtendedDataFields) {
		var ret bool
		return ret
	}
	return *o.IncludeExtendedDataFields
}

// GetIncludeExtendedDataFieldsOk returns a tuple with the IncludeExtendedDataFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsDraftCampaignsRequestContent) GetIncludeExtendedDataFieldsOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeExtendedDataFields) {
		return nil, false
	}
	return o.IncludeExtendedDataFields, true
}

// HasIncludeExtendedDataFields returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsDraftCampaignsRequestContent) HasIncludeExtendedDataFields() bool {
	if o != nil && !IsNil(o.IncludeExtendedDataFields) {
		return true
	}

	return false
}

// SetIncludeExtendedDataFields gets a reference to the given bool and assigns it to the IncludeExtendedDataFields field.
func (o *SponsoredProductsListSponsoredProductsDraftCampaignsRequestContent) SetIncludeExtendedDataFields(v bool) {
	o.IncludeExtendedDataFields = &v
}

// GetNameFilter returns the NameFilter field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsDraftCampaignsRequestContent) GetNameFilter() SponsoredProductsNameFilter {
	if o == nil || IsNil(o.NameFilter) {
		var ret SponsoredProductsNameFilter
		return ret
	}
	return *o.NameFilter
}

// GetNameFilterOk returns a tuple with the NameFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsDraftCampaignsRequestContent) GetNameFilterOk() (*SponsoredProductsNameFilter, bool) {
	if o == nil || IsNil(o.NameFilter) {
		return nil, false
	}
	return o.NameFilter, true
}

// HasNameFilter returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsDraftCampaignsRequestContent) HasNameFilter() bool {
	if o != nil && !IsNil(o.NameFilter) {
		return true
	}

	return false
}

// SetNameFilter gets a reference to the given SponsoredProductsNameFilter and assigns it to the NameFilter field.
func (o *SponsoredProductsListSponsoredProductsDraftCampaignsRequestContent) SetNameFilter(v SponsoredProductsNameFilter) {
	o.NameFilter = &v
}

func (o SponsoredProductsListSponsoredProductsDraftCampaignsRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CampaignIdFilter) {
		toSerialize["campaignIdFilter"] = o.CampaignIdFilter
	}
	if !IsNil(o.PortfolioIdFilter) {
		toSerialize["portfolioIdFilter"] = o.PortfolioIdFilter
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
	if !IsNil(o.NameFilter) {
		toSerialize["nameFilter"] = o.NameFilter
	}
	return toSerialize, nil
}

type NullableSponsoredProductsListSponsoredProductsDraftCampaignsRequestContent struct {
	value *SponsoredProductsListSponsoredProductsDraftCampaignsRequestContent
	isSet bool
}

func (v NullableSponsoredProductsListSponsoredProductsDraftCampaignsRequestContent) Get() *SponsoredProductsListSponsoredProductsDraftCampaignsRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsListSponsoredProductsDraftCampaignsRequestContent) Set(val *SponsoredProductsListSponsoredProductsDraftCampaignsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsListSponsoredProductsDraftCampaignsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsListSponsoredProductsDraftCampaignsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsListSponsoredProductsDraftCampaignsRequestContent(val *SponsoredProductsListSponsoredProductsDraftCampaignsRequestContent) *NullableSponsoredProductsListSponsoredProductsDraftCampaignsRequestContent {
	return &NullableSponsoredProductsListSponsoredProductsDraftCampaignsRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsListSponsoredProductsDraftCampaignsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsListSponsoredProductsDraftCampaignsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
