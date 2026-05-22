package sd_v1

import (
	"time"

	"github.com/bytedance/sonic"
)

// checks if the BudgetUsageCampaign type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BudgetUsageCampaign{}

// BudgetUsageCampaign struct for BudgetUsageCampaign
type BudgetUsageCampaign struct {
	// Budget usage percentage (spend / available budget) for the given budget policy.
	BudgetUsagePercent *float32 `json:"budgetUsagePercent,omitempty"`
	// ID of requested resource
	CampaignId *string `json:"campaignId,omitempty"`
	// Last evaluation time for budget usage
	UsageUpdatedTimestamp *time.Time `json:"usageUpdatedTimestamp,omitempty"`
	// An index to maintain order of the campaignIds
	Index *float32 `json:"index,omitempty"`
	// Budget amount of resource requested
	Budget *float32 `json:"budget,omitempty"`
}

// NewBudgetUsageCampaign instantiates a new BudgetUsageCampaign object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBudgetUsageCampaign() *BudgetUsageCampaign {
	this := BudgetUsageCampaign{}
	return &this
}

// NewBudgetUsageCampaignWithDefaults instantiates a new BudgetUsageCampaign object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBudgetUsageCampaignWithDefaults() *BudgetUsageCampaign {
	this := BudgetUsageCampaign{}
	return &this
}

// GetBudgetUsagePercent returns the BudgetUsagePercent field value if set, zero value otherwise.
func (o *BudgetUsageCampaign) GetBudgetUsagePercent() float32 {
	if o == nil || IsNil(o.BudgetUsagePercent) {
		var ret float32
		return ret
	}
	return *o.BudgetUsagePercent
}

// GetBudgetUsagePercentOk returns a tuple with the BudgetUsagePercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetUsageCampaign) GetBudgetUsagePercentOk() (*float32, bool) {
	if o == nil || IsNil(o.BudgetUsagePercent) {
		return nil, false
	}
	return o.BudgetUsagePercent, true
}

// HasBudgetUsagePercent returns a boolean if a field has been set.
func (o *BudgetUsageCampaign) HasBudgetUsagePercent() bool {
	if o != nil && !IsNil(o.BudgetUsagePercent) {
		return true
	}

	return false
}

// SetBudgetUsagePercent gets a reference to the given float32 and assigns it to the BudgetUsagePercent field.
func (o *BudgetUsageCampaign) SetBudgetUsagePercent(v float32) {
	o.BudgetUsagePercent = &v
}

// GetCampaignId returns the CampaignId field value if set, zero value otherwise.
func (o *BudgetUsageCampaign) GetCampaignId() string {
	if o == nil || IsNil(o.CampaignId) {
		var ret string
		return ret
	}
	return *o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetUsageCampaign) GetCampaignIdOk() (*string, bool) {
	if o == nil || IsNil(o.CampaignId) {
		return nil, false
	}
	return o.CampaignId, true
}

// HasCampaignId returns a boolean if a field has been set.
func (o *BudgetUsageCampaign) HasCampaignId() bool {
	if o != nil && !IsNil(o.CampaignId) {
		return true
	}

	return false
}

// SetCampaignId gets a reference to the given string and assigns it to the CampaignId field.
func (o *BudgetUsageCampaign) SetCampaignId(v string) {
	o.CampaignId = &v
}

// GetUsageUpdatedTimestamp returns the UsageUpdatedTimestamp field value if set, zero value otherwise.
func (o *BudgetUsageCampaign) GetUsageUpdatedTimestamp() time.Time {
	if o == nil || IsNil(o.UsageUpdatedTimestamp) {
		var ret time.Time
		return ret
	}
	return *o.UsageUpdatedTimestamp
}

// GetUsageUpdatedTimestampOk returns a tuple with the UsageUpdatedTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetUsageCampaign) GetUsageUpdatedTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UsageUpdatedTimestamp) {
		return nil, false
	}
	return o.UsageUpdatedTimestamp, true
}

// HasUsageUpdatedTimestamp returns a boolean if a field has been set.
func (o *BudgetUsageCampaign) HasUsageUpdatedTimestamp() bool {
	if o != nil && !IsNil(o.UsageUpdatedTimestamp) {
		return true
	}

	return false
}

// SetUsageUpdatedTimestamp gets a reference to the given time.Time and assigns it to the UsageUpdatedTimestamp field.
func (o *BudgetUsageCampaign) SetUsageUpdatedTimestamp(v time.Time) {
	o.UsageUpdatedTimestamp = &v
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *BudgetUsageCampaign) GetIndex() float32 {
	if o == nil || IsNil(o.Index) {
		var ret float32
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetUsageCampaign) GetIndexOk() (*float32, bool) {
	if o == nil || IsNil(o.Index) {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *BudgetUsageCampaign) HasIndex() bool {
	if o != nil && !IsNil(o.Index) {
		return true
	}

	return false
}

// SetIndex gets a reference to the given float32 and assigns it to the Index field.
func (o *BudgetUsageCampaign) SetIndex(v float32) {
	o.Index = &v
}

// GetBudget returns the Budget field value if set, zero value otherwise.
func (o *BudgetUsageCampaign) GetBudget() float32 {
	if o == nil || IsNil(o.Budget) {
		var ret float32
		return ret
	}
	return *o.Budget
}

// GetBudgetOk returns a tuple with the Budget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetUsageCampaign) GetBudgetOk() (*float32, bool) {
	if o == nil || IsNil(o.Budget) {
		return nil, false
	}
	return o.Budget, true
}

// HasBudget returns a boolean if a field has been set.
func (o *BudgetUsageCampaign) HasBudget() bool {
	if o != nil && !IsNil(o.Budget) {
		return true
	}

	return false
}

// SetBudget gets a reference to the given float32 and assigns it to the Budget field.
func (o *BudgetUsageCampaign) SetBudget(v float32) {
	o.Budget = &v
}

func (o BudgetUsageCampaign) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BudgetUsagePercent) {
		toSerialize["budgetUsagePercent"] = o.BudgetUsagePercent
	}
	if !IsNil(o.CampaignId) {
		toSerialize["campaignId"] = o.CampaignId
	}
	if !IsNil(o.UsageUpdatedTimestamp) {
		toSerialize["usageUpdatedTimestamp"] = o.UsageUpdatedTimestamp
	}
	if !IsNil(o.Index) {
		toSerialize["index"] = o.Index
	}
	if !IsNil(o.Budget) {
		toSerialize["budget"] = o.Budget
	}
	return toSerialize, nil
}

type NullableBudgetUsageCampaign struct {
	value *BudgetUsageCampaign
	isSet bool
}

func (v NullableBudgetUsageCampaign) Get() *BudgetUsageCampaign {
	return v.value
}

func (v *NullableBudgetUsageCampaign) Set(val *BudgetUsageCampaign) {
	v.value = val
	v.isSet = true
}

func (v NullableBudgetUsageCampaign) IsSet() bool {
	return v.isSet
}

func (v *NullableBudgetUsageCampaign) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBudgetUsageCampaign(val *BudgetUsageCampaign) *NullableBudgetUsageCampaign {
	return &NullableBudgetUsageCampaign{value: val, isSet: true}
}

func (v NullableBudgetUsageCampaign) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBudgetUsageCampaign) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
