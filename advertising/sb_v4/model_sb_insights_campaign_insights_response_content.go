package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the SBInsightsCampaignInsightsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SBInsightsCampaignInsightsResponseContent{}

// SBInsightsCampaignInsightsResponseContent Response object for /sb/campaigns/insights containing a list of insights for the campaign.
type SBInsightsCampaignInsightsResponseContent struct {
	Insights []SBInsightsObject `json:"insights,omitempty"`
	// Operations that return paginated results include a pagination token in this field. To retrieve the next page of results, call the same operation and specify this token in the request. If the `NextToken` field is empty, there are no further results.
	NextToken *string `json:"nextToken,omitempty"`
}

// NewSBInsightsCampaignInsightsResponseContent instantiates a new SBInsightsCampaignInsightsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSBInsightsCampaignInsightsResponseContent() *SBInsightsCampaignInsightsResponseContent {
	this := SBInsightsCampaignInsightsResponseContent{}
	return &this
}

// NewSBInsightsCampaignInsightsResponseContentWithDefaults instantiates a new SBInsightsCampaignInsightsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSBInsightsCampaignInsightsResponseContentWithDefaults() *SBInsightsCampaignInsightsResponseContent {
	this := SBInsightsCampaignInsightsResponseContent{}
	return &this
}

// GetInsights returns the Insights field value if set, zero value otherwise.
func (o *SBInsightsCampaignInsightsResponseContent) GetInsights() []SBInsightsObject {
	if o == nil || IsNil(o.Insights) {
		var ret []SBInsightsObject
		return ret
	}
	return o.Insights
}

// GetInsightsOk returns a tuple with the Insights field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBInsightsCampaignInsightsResponseContent) GetInsightsOk() ([]SBInsightsObject, bool) {
	if o == nil || IsNil(o.Insights) {
		return nil, false
	}
	return o.Insights, true
}

// HasInsights returns a boolean if a field has been set.
func (o *SBInsightsCampaignInsightsResponseContent) HasInsights() bool {
	if o != nil && !IsNil(o.Insights) {
		return true
	}

	return false
}

// SetInsights gets a reference to the given []SBInsightsObject and assigns it to the Insights field.
func (o *SBInsightsCampaignInsightsResponseContent) SetInsights(v []SBInsightsObject) {
	o.Insights = v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *SBInsightsCampaignInsightsResponseContent) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBInsightsCampaignInsightsResponseContent) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *SBInsightsCampaignInsightsResponseContent) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *SBInsightsCampaignInsightsResponseContent) SetNextToken(v string) {
	o.NextToken = &v
}

func (o SBInsightsCampaignInsightsResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Insights) {
		toSerialize["insights"] = o.Insights
	}
	if !IsNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	return toSerialize, nil
}

type NullableSBInsightsCampaignInsightsResponseContent struct {
	value *SBInsightsCampaignInsightsResponseContent
	isSet bool
}

func (v NullableSBInsightsCampaignInsightsResponseContent) Get() *SBInsightsCampaignInsightsResponseContent {
	return v.value
}

func (v *NullableSBInsightsCampaignInsightsResponseContent) Set(val *SBInsightsCampaignInsightsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSBInsightsCampaignInsightsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSBInsightsCampaignInsightsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSBInsightsCampaignInsightsResponseContent(val *SBInsightsCampaignInsightsResponseContent) *NullableSBInsightsCampaignInsightsResponseContent {
	return &NullableSBInsightsCampaignInsightsResponseContent{value: val, isSet: true}
}

func (v NullableSBInsightsCampaignInsightsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSBInsightsCampaignInsightsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
