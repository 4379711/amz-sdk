package portfolios_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the BudgetUsageCampaignBatchError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BudgetUsageCampaignBatchError{}

// BudgetUsageCampaignBatchError struct for BudgetUsageCampaignBatchError
type BudgetUsageCampaignBatchError struct {
	// An enumerated error code for machine use.
	Code *string `json:"code,omitempty"`
	// ID of requested resource
	CampaignId *string `json:"campaignId,omitempty"`
	// An index to maintain order of the campaignIds
	Index *float32 `json:"index,omitempty"`
	// A human-readable description of the response.
	Details *string `json:"details,omitempty"`
}

// NewBudgetUsageCampaignBatchError instantiates a new BudgetUsageCampaignBatchError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBudgetUsageCampaignBatchError() *BudgetUsageCampaignBatchError {
	this := BudgetUsageCampaignBatchError{}
	return &this
}

// NewBudgetUsageCampaignBatchErrorWithDefaults instantiates a new BudgetUsageCampaignBatchError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBudgetUsageCampaignBatchErrorWithDefaults() *BudgetUsageCampaignBatchError {
	this := BudgetUsageCampaignBatchError{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *BudgetUsageCampaignBatchError) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetUsageCampaignBatchError) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *BudgetUsageCampaignBatchError) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *BudgetUsageCampaignBatchError) SetCode(v string) {
	o.Code = &v
}

// GetCampaignId returns the CampaignId field value if set, zero value otherwise.
func (o *BudgetUsageCampaignBatchError) GetCampaignId() string {
	if o == nil || IsNil(o.CampaignId) {
		var ret string
		return ret
	}
	return *o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetUsageCampaignBatchError) GetCampaignIdOk() (*string, bool) {
	if o == nil || IsNil(o.CampaignId) {
		return nil, false
	}
	return o.CampaignId, true
}

// HasCampaignId returns a boolean if a field has been set.
func (o *BudgetUsageCampaignBatchError) HasCampaignId() bool {
	if o != nil && !IsNil(o.CampaignId) {
		return true
	}

	return false
}

// SetCampaignId gets a reference to the given string and assigns it to the CampaignId field.
func (o *BudgetUsageCampaignBatchError) SetCampaignId(v string) {
	o.CampaignId = &v
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *BudgetUsageCampaignBatchError) GetIndex() float32 {
	if o == nil || IsNil(o.Index) {
		var ret float32
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetUsageCampaignBatchError) GetIndexOk() (*float32, bool) {
	if o == nil || IsNil(o.Index) {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *BudgetUsageCampaignBatchError) HasIndex() bool {
	if o != nil && !IsNil(o.Index) {
		return true
	}

	return false
}

// SetIndex gets a reference to the given float32 and assigns it to the Index field.
func (o *BudgetUsageCampaignBatchError) SetIndex(v float32) {
	o.Index = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *BudgetUsageCampaignBatchError) GetDetails() string {
	if o == nil || IsNil(o.Details) {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetUsageCampaignBatchError) GetDetailsOk() (*string, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *BudgetUsageCampaignBatchError) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *BudgetUsageCampaignBatchError) SetDetails(v string) {
	o.Details = &v
}

func (o BudgetUsageCampaignBatchError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.CampaignId) {
		toSerialize["campaignId"] = o.CampaignId
	}
	if !IsNil(o.Index) {
		toSerialize["index"] = o.Index
	}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	return toSerialize, nil
}

type NullableBudgetUsageCampaignBatchError struct {
	value *BudgetUsageCampaignBatchError
	isSet bool
}

func (v NullableBudgetUsageCampaignBatchError) Get() *BudgetUsageCampaignBatchError {
	return v.value
}

func (v *NullableBudgetUsageCampaignBatchError) Set(val *BudgetUsageCampaignBatchError) {
	v.value = val
	v.isSet = true
}

func (v NullableBudgetUsageCampaignBatchError) IsSet() bool {
	return v.isSet
}

func (v *NullableBudgetUsageCampaignBatchError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBudgetUsageCampaignBatchError(val *BudgetUsageCampaignBatchError) *NullableBudgetUsageCampaignBatchError {
	return &NullableBudgetUsageCampaignBatchError{value: val, isSet: true}
}

func (v NullableBudgetUsageCampaignBatchError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBudgetUsageCampaignBatchError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
