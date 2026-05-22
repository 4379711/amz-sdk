package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the GetBudgetRecommendationsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetBudgetRecommendationsRequestContent{}

// GetBudgetRecommendationsRequestContent struct for GetBudgetRecommendationsRequestContent
type GetBudgetRecommendationsRequestContent struct {
	// List of CampaignIds
	CampaignIds []string `json:"campaignIds"`
}

type _GetBudgetRecommendationsRequestContent GetBudgetRecommendationsRequestContent

// NewGetBudgetRecommendationsRequestContent instantiates a new GetBudgetRecommendationsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBudgetRecommendationsRequestContent(campaignIds []string) *GetBudgetRecommendationsRequestContent {
	this := GetBudgetRecommendationsRequestContent{}
	this.CampaignIds = campaignIds
	return &this
}

// NewGetBudgetRecommendationsRequestContentWithDefaults instantiates a new GetBudgetRecommendationsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBudgetRecommendationsRequestContentWithDefaults() *GetBudgetRecommendationsRequestContent {
	this := GetBudgetRecommendationsRequestContent{}
	return &this
}

// GetCampaignIds returns the CampaignIds field value
func (o *GetBudgetRecommendationsRequestContent) GetCampaignIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.CampaignIds
}

// GetCampaignIdsOk returns a tuple with the CampaignIds field value
// and a boolean to check if the value has been set.
func (o *GetBudgetRecommendationsRequestContent) GetCampaignIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CampaignIds, true
}

// SetCampaignIds sets field value
func (o *GetBudgetRecommendationsRequestContent) SetCampaignIds(v []string) {
	o.CampaignIds = v
}

func (o GetBudgetRecommendationsRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["campaignIds"] = o.CampaignIds
	return toSerialize, nil
}

type NullableGetBudgetRecommendationsRequestContent struct {
	value *GetBudgetRecommendationsRequestContent
	isSet bool
}

func (v NullableGetBudgetRecommendationsRequestContent) Get() *GetBudgetRecommendationsRequestContent {
	return v.value
}

func (v *NullableGetBudgetRecommendationsRequestContent) Set(val *GetBudgetRecommendationsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBudgetRecommendationsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBudgetRecommendationsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBudgetRecommendationsRequestContent(val *GetBudgetRecommendationsRequestContent) *NullableGetBudgetRecommendationsRequestContent {
	return &NullableGetBudgetRecommendationsRequestContent{value: val, isSet: true}
}

func (v NullableGetBudgetRecommendationsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetBudgetRecommendationsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
