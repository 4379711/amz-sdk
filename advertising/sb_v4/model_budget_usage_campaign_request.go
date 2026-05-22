package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the BudgetUsageCampaignRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BudgetUsageCampaignRequest{}

// BudgetUsageCampaignRequest struct for BudgetUsageCampaignRequest
type BudgetUsageCampaignRequest struct {
	// A list of campaign IDs
	CampaignIds []string `json:"campaignIds,omitempty"`
}

// NewBudgetUsageCampaignRequest instantiates a new BudgetUsageCampaignRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBudgetUsageCampaignRequest() *BudgetUsageCampaignRequest {
	this := BudgetUsageCampaignRequest{}
	return &this
}

// NewBudgetUsageCampaignRequestWithDefaults instantiates a new BudgetUsageCampaignRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBudgetUsageCampaignRequestWithDefaults() *BudgetUsageCampaignRequest {
	this := BudgetUsageCampaignRequest{}
	return &this
}

// GetCampaignIds returns the CampaignIds field value if set, zero value otherwise.
func (o *BudgetUsageCampaignRequest) GetCampaignIds() []string {
	if o == nil || IsNil(o.CampaignIds) {
		var ret []string
		return ret
	}
	return o.CampaignIds
}

// GetCampaignIdsOk returns a tuple with the CampaignIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetUsageCampaignRequest) GetCampaignIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.CampaignIds) {
		return nil, false
	}
	return o.CampaignIds, true
}

// HasCampaignIds returns a boolean if a field has been set.
func (o *BudgetUsageCampaignRequest) HasCampaignIds() bool {
	if o != nil && !IsNil(o.CampaignIds) {
		return true
	}

	return false
}

// SetCampaignIds gets a reference to the given []string and assigns it to the CampaignIds field.
func (o *BudgetUsageCampaignRequest) SetCampaignIds(v []string) {
	o.CampaignIds = v
}

func (o BudgetUsageCampaignRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CampaignIds) {
		toSerialize["campaignIds"] = o.CampaignIds
	}
	return toSerialize, nil
}

type NullableBudgetUsageCampaignRequest struct {
	value *BudgetUsageCampaignRequest
	isSet bool
}

func (v NullableBudgetUsageCampaignRequest) Get() *BudgetUsageCampaignRequest {
	return v.value
}

func (v *NullableBudgetUsageCampaignRequest) Set(val *BudgetUsageCampaignRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBudgetUsageCampaignRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBudgetUsageCampaignRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBudgetUsageCampaignRequest(val *BudgetUsageCampaignRequest) *NullableBudgetUsageCampaignRequest {
	return &NullableBudgetUsageCampaignRequest{value: val, isSet: true}
}

func (v NullableBudgetUsageCampaignRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBudgetUsageCampaignRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
