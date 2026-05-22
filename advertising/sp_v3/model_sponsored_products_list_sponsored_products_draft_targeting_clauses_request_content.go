package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsListSponsoredProductsDraftTargetingClausesRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsListSponsoredProductsDraftTargetingClausesRequestContent{}

// SponsoredProductsListSponsoredProductsDraftTargetingClausesRequestContent struct for SponsoredProductsListSponsoredProductsDraftTargetingClausesRequestContent
type SponsoredProductsListSponsoredProductsDraftTargetingClausesRequestContent struct {
	CampaignIdFilter     *SponsoredProductsReducedObjectIdFilter `json:"campaignIdFilter,omitempty"`
	ExpressionTypeFilter *SponsoredProductsExpressionTypeFilter  `json:"expressionTypeFilter,omitempty"`
	// Number of records to include in the paginated response. Defaults to max page size for given API
	MaxResults *int32 `json:"maxResults,omitempty"`
	// token value allowing to navigate to the next response page
	NextToken       *string                                 `json:"nextToken,omitempty"`
	TargetIdFilter  *SponsoredProductsObjectIdFilter        `json:"targetIdFilter,omitempty"`
	AsinFilter      *SponsoredProductsAsinFilter            `json:"asinFilter,omitempty"`
	AdGroupIdFilter *SponsoredProductsReducedObjectIdFilter `json:"adGroupIdFilter,omitempty"`
	// Whether to get entity with extended data fields such as creationDate, lastUpdateDate, servingStatus
	IncludeExtendedDataFields *bool `json:"includeExtendedDataFields,omitempty"`
}

// NewSponsoredProductsListSponsoredProductsDraftTargetingClausesRequestContent instantiates a new SponsoredProductsListSponsoredProductsDraftTargetingClausesRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsListSponsoredProductsDraftTargetingClausesRequestContent() *SponsoredProductsListSponsoredProductsDraftTargetingClausesRequestContent {
	this := SponsoredProductsListSponsoredProductsDraftTargetingClausesRequestContent{}
	return &this
}

// NewSponsoredProductsListSponsoredProductsDraftTargetingClausesRequestContentWithDefaults instantiates a new SponsoredProductsListSponsoredProductsDraftTargetingClausesRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsListSponsoredProductsDraftTargetingClausesRequestContentWithDefaults() *SponsoredProductsListSponsoredProductsDraftTargetingClausesRequestContent {
	this := SponsoredProductsListSponsoredProductsDraftTargetingClausesRequestContent{}
	return &this
}

// GetCampaignIdFilter returns the CampaignIdFilter field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsDraftTargetingClausesRequestContent) GetCampaignIdFilter() SponsoredProductsReducedObjectIdFilter {
	if o == nil || IsNil(o.CampaignIdFilter) {
		var ret SponsoredProductsReducedObjectIdFilter
		return ret
	}
	return *o.CampaignIdFilter
}

// GetCampaignIdFilterOk returns a tuple with the CampaignIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsDraftTargetingClausesRequestContent) GetCampaignIdFilterOk() (*SponsoredProductsReducedObjectIdFilter, bool) {
	if o == nil || IsNil(o.CampaignIdFilter) {
		return nil, false
	}
	return o.CampaignIdFilter, true
}

// HasCampaignIdFilter returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsDraftTargetingClausesRequestContent) HasCampaignIdFilter() bool {
	if o != nil && !IsNil(o.CampaignIdFilter) {
		return true
	}

	return false
}

// SetCampaignIdFilter gets a reference to the given SponsoredProductsReducedObjectIdFilter and assigns it to the CampaignIdFilter field.
func (o *SponsoredProductsListSponsoredProductsDraftTargetingClausesRequestContent) SetCampaignIdFilter(v SponsoredProductsReducedObjectIdFilter) {
	o.CampaignIdFilter = &v
}

// GetExpressionTypeFilter returns the ExpressionTypeFilter field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsDraftTargetingClausesRequestContent) GetExpressionTypeFilter() SponsoredProductsExpressionTypeFilter {
	if o == nil || IsNil(o.ExpressionTypeFilter) {
		var ret SponsoredProductsExpressionTypeFilter
		return ret
	}
	return *o.ExpressionTypeFilter
}

// GetExpressionTypeFilterOk returns a tuple with the ExpressionTypeFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsDraftTargetingClausesRequestContent) GetExpressionTypeFilterOk() (*SponsoredProductsExpressionTypeFilter, bool) {
	if o == nil || IsNil(o.ExpressionTypeFilter) {
		return nil, false
	}
	return o.ExpressionTypeFilter, true
}

// HasExpressionTypeFilter returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsDraftTargetingClausesRequestContent) HasExpressionTypeFilter() bool {
	if o != nil && !IsNil(o.ExpressionTypeFilter) {
		return true
	}

	return false
}

// SetExpressionTypeFilter gets a reference to the given SponsoredProductsExpressionTypeFilter and assigns it to the ExpressionTypeFilter field.
func (o *SponsoredProductsListSponsoredProductsDraftTargetingClausesRequestContent) SetExpressionTypeFilter(v SponsoredProductsExpressionTypeFilter) {
	o.ExpressionTypeFilter = &v
}

// GetMaxResults returns the MaxResults field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsDraftTargetingClausesRequestContent) GetMaxResults() int32 {
	if o == nil || IsNil(o.MaxResults) {
		var ret int32
		return ret
	}
	return *o.MaxResults
}

// GetMaxResultsOk returns a tuple with the MaxResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsDraftTargetingClausesRequestContent) GetMaxResultsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxResults) {
		return nil, false
	}
	return o.MaxResults, true
}

// HasMaxResults returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsDraftTargetingClausesRequestContent) HasMaxResults() bool {
	if o != nil && !IsNil(o.MaxResults) {
		return true
	}

	return false
}

// SetMaxResults gets a reference to the given int32 and assigns it to the MaxResults field.
func (o *SponsoredProductsListSponsoredProductsDraftTargetingClausesRequestContent) SetMaxResults(v int32) {
	o.MaxResults = &v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsDraftTargetingClausesRequestContent) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsDraftTargetingClausesRequestContent) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsDraftTargetingClausesRequestContent) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *SponsoredProductsListSponsoredProductsDraftTargetingClausesRequestContent) SetNextToken(v string) {
	o.NextToken = &v
}

// GetTargetIdFilter returns the TargetIdFilter field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsDraftTargetingClausesRequestContent) GetTargetIdFilter() SponsoredProductsObjectIdFilter {
	if o == nil || IsNil(o.TargetIdFilter) {
		var ret SponsoredProductsObjectIdFilter
		return ret
	}
	return *o.TargetIdFilter
}

// GetTargetIdFilterOk returns a tuple with the TargetIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsDraftTargetingClausesRequestContent) GetTargetIdFilterOk() (*SponsoredProductsObjectIdFilter, bool) {
	if o == nil || IsNil(o.TargetIdFilter) {
		return nil, false
	}
	return o.TargetIdFilter, true
}

// HasTargetIdFilter returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsDraftTargetingClausesRequestContent) HasTargetIdFilter() bool {
	if o != nil && !IsNil(o.TargetIdFilter) {
		return true
	}

	return false
}

// SetTargetIdFilter gets a reference to the given SponsoredProductsObjectIdFilter and assigns it to the TargetIdFilter field.
func (o *SponsoredProductsListSponsoredProductsDraftTargetingClausesRequestContent) SetTargetIdFilter(v SponsoredProductsObjectIdFilter) {
	o.TargetIdFilter = &v
}

// GetAsinFilter returns the AsinFilter field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsDraftTargetingClausesRequestContent) GetAsinFilter() SponsoredProductsAsinFilter {
	if o == nil || IsNil(o.AsinFilter) {
		var ret SponsoredProductsAsinFilter
		return ret
	}
	return *o.AsinFilter
}

// GetAsinFilterOk returns a tuple with the AsinFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsDraftTargetingClausesRequestContent) GetAsinFilterOk() (*SponsoredProductsAsinFilter, bool) {
	if o == nil || IsNil(o.AsinFilter) {
		return nil, false
	}
	return o.AsinFilter, true
}

// HasAsinFilter returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsDraftTargetingClausesRequestContent) HasAsinFilter() bool {
	if o != nil && !IsNil(o.AsinFilter) {
		return true
	}

	return false
}

// SetAsinFilter gets a reference to the given SponsoredProductsAsinFilter and assigns it to the AsinFilter field.
func (o *SponsoredProductsListSponsoredProductsDraftTargetingClausesRequestContent) SetAsinFilter(v SponsoredProductsAsinFilter) {
	o.AsinFilter = &v
}

// GetAdGroupIdFilter returns the AdGroupIdFilter field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsDraftTargetingClausesRequestContent) GetAdGroupIdFilter() SponsoredProductsReducedObjectIdFilter {
	if o == nil || IsNil(o.AdGroupIdFilter) {
		var ret SponsoredProductsReducedObjectIdFilter
		return ret
	}
	return *o.AdGroupIdFilter
}

// GetAdGroupIdFilterOk returns a tuple with the AdGroupIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsDraftTargetingClausesRequestContent) GetAdGroupIdFilterOk() (*SponsoredProductsReducedObjectIdFilter, bool) {
	if o == nil || IsNil(o.AdGroupIdFilter) {
		return nil, false
	}
	return o.AdGroupIdFilter, true
}

// HasAdGroupIdFilter returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsDraftTargetingClausesRequestContent) HasAdGroupIdFilter() bool {
	if o != nil && !IsNil(o.AdGroupIdFilter) {
		return true
	}

	return false
}

// SetAdGroupIdFilter gets a reference to the given SponsoredProductsReducedObjectIdFilter and assigns it to the AdGroupIdFilter field.
func (o *SponsoredProductsListSponsoredProductsDraftTargetingClausesRequestContent) SetAdGroupIdFilter(v SponsoredProductsReducedObjectIdFilter) {
	o.AdGroupIdFilter = &v
}

// GetIncludeExtendedDataFields returns the IncludeExtendedDataFields field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsDraftTargetingClausesRequestContent) GetIncludeExtendedDataFields() bool {
	if o == nil || IsNil(o.IncludeExtendedDataFields) {
		var ret bool
		return ret
	}
	return *o.IncludeExtendedDataFields
}

// GetIncludeExtendedDataFieldsOk returns a tuple with the IncludeExtendedDataFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsDraftTargetingClausesRequestContent) GetIncludeExtendedDataFieldsOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeExtendedDataFields) {
		return nil, false
	}
	return o.IncludeExtendedDataFields, true
}

// HasIncludeExtendedDataFields returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsDraftTargetingClausesRequestContent) HasIncludeExtendedDataFields() bool {
	if o != nil && !IsNil(o.IncludeExtendedDataFields) {
		return true
	}

	return false
}

// SetIncludeExtendedDataFields gets a reference to the given bool and assigns it to the IncludeExtendedDataFields field.
func (o *SponsoredProductsListSponsoredProductsDraftTargetingClausesRequestContent) SetIncludeExtendedDataFields(v bool) {
	o.IncludeExtendedDataFields = &v
}

func (o SponsoredProductsListSponsoredProductsDraftTargetingClausesRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CampaignIdFilter) {
		toSerialize["campaignIdFilter"] = o.CampaignIdFilter
	}
	if !IsNil(o.ExpressionTypeFilter) {
		toSerialize["expressionTypeFilter"] = o.ExpressionTypeFilter
	}
	if !IsNil(o.MaxResults) {
		toSerialize["maxResults"] = o.MaxResults
	}
	if !IsNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	if !IsNil(o.TargetIdFilter) {
		toSerialize["targetIdFilter"] = o.TargetIdFilter
	}
	if !IsNil(o.AsinFilter) {
		toSerialize["asinFilter"] = o.AsinFilter
	}
	if !IsNil(o.AdGroupIdFilter) {
		toSerialize["adGroupIdFilter"] = o.AdGroupIdFilter
	}
	if !IsNil(o.IncludeExtendedDataFields) {
		toSerialize["includeExtendedDataFields"] = o.IncludeExtendedDataFields
	}
	return toSerialize, nil
}

type NullableSponsoredProductsListSponsoredProductsDraftTargetingClausesRequestContent struct {
	value *SponsoredProductsListSponsoredProductsDraftTargetingClausesRequestContent
	isSet bool
}

func (v NullableSponsoredProductsListSponsoredProductsDraftTargetingClausesRequestContent) Get() *SponsoredProductsListSponsoredProductsDraftTargetingClausesRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsListSponsoredProductsDraftTargetingClausesRequestContent) Set(val *SponsoredProductsListSponsoredProductsDraftTargetingClausesRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsListSponsoredProductsDraftTargetingClausesRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsListSponsoredProductsDraftTargetingClausesRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsListSponsoredProductsDraftTargetingClausesRequestContent(val *SponsoredProductsListSponsoredProductsDraftTargetingClausesRequestContent) *NullableSponsoredProductsListSponsoredProductsDraftTargetingClausesRequestContent {
	return &NullableSponsoredProductsListSponsoredProductsDraftTargetingClausesRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsListSponsoredProductsDraftTargetingClausesRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsListSponsoredProductsDraftTargetingClausesRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
