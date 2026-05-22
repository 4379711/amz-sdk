package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsListSponsoredProductsTargetingClausesPreviewRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsListSponsoredProductsTargetingClausesPreviewRequestContent{}

// SponsoredProductsListSponsoredProductsTargetingClausesPreviewRequestContent struct for SponsoredProductsListSponsoredProductsTargetingClausesPreviewRequestContent
type SponsoredProductsListSponsoredProductsTargetingClausesPreviewRequestContent struct {
	CampaignIdFilter     *SponsoredProductsReducedObjectIdFilter `json:"campaignIdFilter,omitempty"`
	StateFilter          *SponsoredProductsEntityStateFilter     `json:"stateFilter,omitempty"`
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

// NewSponsoredProductsListSponsoredProductsTargetingClausesPreviewRequestContent instantiates a new SponsoredProductsListSponsoredProductsTargetingClausesPreviewRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsListSponsoredProductsTargetingClausesPreviewRequestContent() *SponsoredProductsListSponsoredProductsTargetingClausesPreviewRequestContent {
	this := SponsoredProductsListSponsoredProductsTargetingClausesPreviewRequestContent{}
	return &this
}

// NewSponsoredProductsListSponsoredProductsTargetingClausesPreviewRequestContentWithDefaults instantiates a new SponsoredProductsListSponsoredProductsTargetingClausesPreviewRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsListSponsoredProductsTargetingClausesPreviewRequestContentWithDefaults() *SponsoredProductsListSponsoredProductsTargetingClausesPreviewRequestContent {
	this := SponsoredProductsListSponsoredProductsTargetingClausesPreviewRequestContent{}
	return &this
}

// GetCampaignIdFilter returns the CampaignIdFilter field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsTargetingClausesPreviewRequestContent) GetCampaignIdFilter() SponsoredProductsReducedObjectIdFilter {
	if o == nil || IsNil(o.CampaignIdFilter) {
		var ret SponsoredProductsReducedObjectIdFilter
		return ret
	}
	return *o.CampaignIdFilter
}

// GetCampaignIdFilterOk returns a tuple with the CampaignIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsTargetingClausesPreviewRequestContent) GetCampaignIdFilterOk() (*SponsoredProductsReducedObjectIdFilter, bool) {
	if o == nil || IsNil(o.CampaignIdFilter) {
		return nil, false
	}
	return o.CampaignIdFilter, true
}

// HasCampaignIdFilter returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsTargetingClausesPreviewRequestContent) HasCampaignIdFilter() bool {
	if o != nil && !IsNil(o.CampaignIdFilter) {
		return true
	}

	return false
}

// SetCampaignIdFilter gets a reference to the given SponsoredProductsReducedObjectIdFilter and assigns it to the CampaignIdFilter field.
func (o *SponsoredProductsListSponsoredProductsTargetingClausesPreviewRequestContent) SetCampaignIdFilter(v SponsoredProductsReducedObjectIdFilter) {
	o.CampaignIdFilter = &v
}

// GetStateFilter returns the StateFilter field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsTargetingClausesPreviewRequestContent) GetStateFilter() SponsoredProductsEntityStateFilter {
	if o == nil || IsNil(o.StateFilter) {
		var ret SponsoredProductsEntityStateFilter
		return ret
	}
	return *o.StateFilter
}

// GetStateFilterOk returns a tuple with the StateFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsTargetingClausesPreviewRequestContent) GetStateFilterOk() (*SponsoredProductsEntityStateFilter, bool) {
	if o == nil || IsNil(o.StateFilter) {
		return nil, false
	}
	return o.StateFilter, true
}

// HasStateFilter returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsTargetingClausesPreviewRequestContent) HasStateFilter() bool {
	if o != nil && !IsNil(o.StateFilter) {
		return true
	}

	return false
}

// SetStateFilter gets a reference to the given SponsoredProductsEntityStateFilter and assigns it to the StateFilter field.
func (o *SponsoredProductsListSponsoredProductsTargetingClausesPreviewRequestContent) SetStateFilter(v SponsoredProductsEntityStateFilter) {
	o.StateFilter = &v
}

// GetExpressionTypeFilter returns the ExpressionTypeFilter field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsTargetingClausesPreviewRequestContent) GetExpressionTypeFilter() SponsoredProductsExpressionTypeFilter {
	if o == nil || IsNil(o.ExpressionTypeFilter) {
		var ret SponsoredProductsExpressionTypeFilter
		return ret
	}
	return *o.ExpressionTypeFilter
}

// GetExpressionTypeFilterOk returns a tuple with the ExpressionTypeFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsTargetingClausesPreviewRequestContent) GetExpressionTypeFilterOk() (*SponsoredProductsExpressionTypeFilter, bool) {
	if o == nil || IsNil(o.ExpressionTypeFilter) {
		return nil, false
	}
	return o.ExpressionTypeFilter, true
}

// HasExpressionTypeFilter returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsTargetingClausesPreviewRequestContent) HasExpressionTypeFilter() bool {
	if o != nil && !IsNil(o.ExpressionTypeFilter) {
		return true
	}

	return false
}

// SetExpressionTypeFilter gets a reference to the given SponsoredProductsExpressionTypeFilter and assigns it to the ExpressionTypeFilter field.
func (o *SponsoredProductsListSponsoredProductsTargetingClausesPreviewRequestContent) SetExpressionTypeFilter(v SponsoredProductsExpressionTypeFilter) {
	o.ExpressionTypeFilter = &v
}

// GetMaxResults returns the MaxResults field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsTargetingClausesPreviewRequestContent) GetMaxResults() int32 {
	if o == nil || IsNil(o.MaxResults) {
		var ret int32
		return ret
	}
	return *o.MaxResults
}

// GetMaxResultsOk returns a tuple with the MaxResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsTargetingClausesPreviewRequestContent) GetMaxResultsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxResults) {
		return nil, false
	}
	return o.MaxResults, true
}

// HasMaxResults returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsTargetingClausesPreviewRequestContent) HasMaxResults() bool {
	if o != nil && !IsNil(o.MaxResults) {
		return true
	}

	return false
}

// SetMaxResults gets a reference to the given int32 and assigns it to the MaxResults field.
func (o *SponsoredProductsListSponsoredProductsTargetingClausesPreviewRequestContent) SetMaxResults(v int32) {
	o.MaxResults = &v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsTargetingClausesPreviewRequestContent) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsTargetingClausesPreviewRequestContent) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsTargetingClausesPreviewRequestContent) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *SponsoredProductsListSponsoredProductsTargetingClausesPreviewRequestContent) SetNextToken(v string) {
	o.NextToken = &v
}

// GetTargetIdFilter returns the TargetIdFilter field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsTargetingClausesPreviewRequestContent) GetTargetIdFilter() SponsoredProductsObjectIdFilter {
	if o == nil || IsNil(o.TargetIdFilter) {
		var ret SponsoredProductsObjectIdFilter
		return ret
	}
	return *o.TargetIdFilter
}

// GetTargetIdFilterOk returns a tuple with the TargetIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsTargetingClausesPreviewRequestContent) GetTargetIdFilterOk() (*SponsoredProductsObjectIdFilter, bool) {
	if o == nil || IsNil(o.TargetIdFilter) {
		return nil, false
	}
	return o.TargetIdFilter, true
}

// HasTargetIdFilter returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsTargetingClausesPreviewRequestContent) HasTargetIdFilter() bool {
	if o != nil && !IsNil(o.TargetIdFilter) {
		return true
	}

	return false
}

// SetTargetIdFilter gets a reference to the given SponsoredProductsObjectIdFilter and assigns it to the TargetIdFilter field.
func (o *SponsoredProductsListSponsoredProductsTargetingClausesPreviewRequestContent) SetTargetIdFilter(v SponsoredProductsObjectIdFilter) {
	o.TargetIdFilter = &v
}

// GetAsinFilter returns the AsinFilter field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsTargetingClausesPreviewRequestContent) GetAsinFilter() SponsoredProductsAsinFilter {
	if o == nil || IsNil(o.AsinFilter) {
		var ret SponsoredProductsAsinFilter
		return ret
	}
	return *o.AsinFilter
}

// GetAsinFilterOk returns a tuple with the AsinFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsTargetingClausesPreviewRequestContent) GetAsinFilterOk() (*SponsoredProductsAsinFilter, bool) {
	if o == nil || IsNil(o.AsinFilter) {
		return nil, false
	}
	return o.AsinFilter, true
}

// HasAsinFilter returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsTargetingClausesPreviewRequestContent) HasAsinFilter() bool {
	if o != nil && !IsNil(o.AsinFilter) {
		return true
	}

	return false
}

// SetAsinFilter gets a reference to the given SponsoredProductsAsinFilter and assigns it to the AsinFilter field.
func (o *SponsoredProductsListSponsoredProductsTargetingClausesPreviewRequestContent) SetAsinFilter(v SponsoredProductsAsinFilter) {
	o.AsinFilter = &v
}

// GetAdGroupIdFilter returns the AdGroupIdFilter field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsTargetingClausesPreviewRequestContent) GetAdGroupIdFilter() SponsoredProductsReducedObjectIdFilter {
	if o == nil || IsNil(o.AdGroupIdFilter) {
		var ret SponsoredProductsReducedObjectIdFilter
		return ret
	}
	return *o.AdGroupIdFilter
}

// GetAdGroupIdFilterOk returns a tuple with the AdGroupIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsTargetingClausesPreviewRequestContent) GetAdGroupIdFilterOk() (*SponsoredProductsReducedObjectIdFilter, bool) {
	if o == nil || IsNil(o.AdGroupIdFilter) {
		return nil, false
	}
	return o.AdGroupIdFilter, true
}

// HasAdGroupIdFilter returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsTargetingClausesPreviewRequestContent) HasAdGroupIdFilter() bool {
	if o != nil && !IsNil(o.AdGroupIdFilter) {
		return true
	}

	return false
}

// SetAdGroupIdFilter gets a reference to the given SponsoredProductsReducedObjectIdFilter and assigns it to the AdGroupIdFilter field.
func (o *SponsoredProductsListSponsoredProductsTargetingClausesPreviewRequestContent) SetAdGroupIdFilter(v SponsoredProductsReducedObjectIdFilter) {
	o.AdGroupIdFilter = &v
}

// GetIncludeExtendedDataFields returns the IncludeExtendedDataFields field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsTargetingClausesPreviewRequestContent) GetIncludeExtendedDataFields() bool {
	if o == nil || IsNil(o.IncludeExtendedDataFields) {
		var ret bool
		return ret
	}
	return *o.IncludeExtendedDataFields
}

// GetIncludeExtendedDataFieldsOk returns a tuple with the IncludeExtendedDataFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsTargetingClausesPreviewRequestContent) GetIncludeExtendedDataFieldsOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeExtendedDataFields) {
		return nil, false
	}
	return o.IncludeExtendedDataFields, true
}

// HasIncludeExtendedDataFields returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsTargetingClausesPreviewRequestContent) HasIncludeExtendedDataFields() bool {
	if o != nil && !IsNil(o.IncludeExtendedDataFields) {
		return true
	}

	return false
}

// SetIncludeExtendedDataFields gets a reference to the given bool and assigns it to the IncludeExtendedDataFields field.
func (o *SponsoredProductsListSponsoredProductsTargetingClausesPreviewRequestContent) SetIncludeExtendedDataFields(v bool) {
	o.IncludeExtendedDataFields = &v
}

func (o SponsoredProductsListSponsoredProductsTargetingClausesPreviewRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CampaignIdFilter) {
		toSerialize["campaignIdFilter"] = o.CampaignIdFilter
	}
	if !IsNil(o.StateFilter) {
		toSerialize["stateFilter"] = o.StateFilter
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

type NullableSponsoredProductsListSponsoredProductsTargetingClausesPreviewRequestContent struct {
	value *SponsoredProductsListSponsoredProductsTargetingClausesPreviewRequestContent
	isSet bool
}

func (v NullableSponsoredProductsListSponsoredProductsTargetingClausesPreviewRequestContent) Get() *SponsoredProductsListSponsoredProductsTargetingClausesPreviewRequestContent {
	return v.value
}

func (v *NullableSponsoredProductsListSponsoredProductsTargetingClausesPreviewRequestContent) Set(val *SponsoredProductsListSponsoredProductsTargetingClausesPreviewRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsListSponsoredProductsTargetingClausesPreviewRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsListSponsoredProductsTargetingClausesPreviewRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsListSponsoredProductsTargetingClausesPreviewRequestContent(val *SponsoredProductsListSponsoredProductsTargetingClausesPreviewRequestContent) *NullableSponsoredProductsListSponsoredProductsTargetingClausesPreviewRequestContent {
	return &NullableSponsoredProductsListSponsoredProductsTargetingClausesPreviewRequestContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsListSponsoredProductsTargetingClausesPreviewRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsListSponsoredProductsTargetingClausesPreviewRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
