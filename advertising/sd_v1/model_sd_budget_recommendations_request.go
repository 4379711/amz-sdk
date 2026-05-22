package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the SDBudgetRecommendationsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SDBudgetRecommendationsRequest{}

// SDBudgetRecommendationsRequest Request for budget recommendations.
type SDBudgetRecommendationsRequest struct {
	// A list of campaign ids for which to get budget recommendations and missed opportunities.
	CampaignIds []string `json:"campaignIds"`
}

type _SDBudgetRecommendationsRequest SDBudgetRecommendationsRequest

// NewSDBudgetRecommendationsRequest instantiates a new SDBudgetRecommendationsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSDBudgetRecommendationsRequest(campaignIds []string) *SDBudgetRecommendationsRequest {
	this := SDBudgetRecommendationsRequest{}
	this.CampaignIds = campaignIds
	return &this
}

// NewSDBudgetRecommendationsRequestWithDefaults instantiates a new SDBudgetRecommendationsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSDBudgetRecommendationsRequestWithDefaults() *SDBudgetRecommendationsRequest {
	this := SDBudgetRecommendationsRequest{}
	return &this
}

// GetCampaignIds returns the CampaignIds field value
func (o *SDBudgetRecommendationsRequest) GetCampaignIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.CampaignIds
}

// GetCampaignIdsOk returns a tuple with the CampaignIds field value
// and a boolean to check if the value has been set.
func (o *SDBudgetRecommendationsRequest) GetCampaignIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CampaignIds, true
}

// SetCampaignIds sets field value
func (o *SDBudgetRecommendationsRequest) SetCampaignIds(v []string) {
	o.CampaignIds = v
}

func (o SDBudgetRecommendationsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["campaignIds"] = o.CampaignIds
	return toSerialize, nil
}

type NullableSDBudgetRecommendationsRequest struct {
	value *SDBudgetRecommendationsRequest
	isSet bool
}

func (v NullableSDBudgetRecommendationsRequest) Get() *SDBudgetRecommendationsRequest {
	return v.value
}

func (v *NullableSDBudgetRecommendationsRequest) Set(val *SDBudgetRecommendationsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSDBudgetRecommendationsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSDBudgetRecommendationsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDBudgetRecommendationsRequest(val *SDBudgetRecommendationsRequest) *NullableSDBudgetRecommendationsRequest {
	return &NullableSDBudgetRecommendationsRequest{value: val, isSet: true}
}

func (v NullableSDBudgetRecommendationsRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDBudgetRecommendationsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
