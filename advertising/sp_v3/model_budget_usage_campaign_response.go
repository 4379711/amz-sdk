package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the BudgetUsageCampaignResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BudgetUsageCampaignResponse{}

// BudgetUsageCampaignResponse struct for BudgetUsageCampaignResponse
type BudgetUsageCampaignResponse struct {
	// List of budget usage percentages that were successfully pulled
	Success []BudgetUsageCampaign `json:"success,omitempty"`
	// List of budget usage percentages that failed to pull
	Error []BudgetUsageCampaignBatchError `json:"error,omitempty"`
}

// NewBudgetUsageCampaignResponse instantiates a new BudgetUsageCampaignResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBudgetUsageCampaignResponse() *BudgetUsageCampaignResponse {
	this := BudgetUsageCampaignResponse{}
	return &this
}

// NewBudgetUsageCampaignResponseWithDefaults instantiates a new BudgetUsageCampaignResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBudgetUsageCampaignResponseWithDefaults() *BudgetUsageCampaignResponse {
	this := BudgetUsageCampaignResponse{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *BudgetUsageCampaignResponse) GetSuccess() []BudgetUsageCampaign {
	if o == nil || IsNil(o.Success) {
		var ret []BudgetUsageCampaign
		return ret
	}
	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetUsageCampaignResponse) GetSuccessOk() ([]BudgetUsageCampaign, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *BudgetUsageCampaignResponse) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given []BudgetUsageCampaign and assigns it to the Success field.
func (o *BudgetUsageCampaignResponse) SetSuccess(v []BudgetUsageCampaign) {
	o.Success = v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *BudgetUsageCampaignResponse) GetError() []BudgetUsageCampaignBatchError {
	if o == nil || IsNil(o.Error) {
		var ret []BudgetUsageCampaignBatchError
		return ret
	}
	return o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetUsageCampaignResponse) GetErrorOk() ([]BudgetUsageCampaignBatchError, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *BudgetUsageCampaignResponse) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given []BudgetUsageCampaignBatchError and assigns it to the Error field.
func (o *BudgetUsageCampaignResponse) SetError(v []BudgetUsageCampaignBatchError) {
	o.Error = v
}

func (o BudgetUsageCampaignResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

type NullableBudgetUsageCampaignResponse struct {
	value *BudgetUsageCampaignResponse
	isSet bool
}

func (v NullableBudgetUsageCampaignResponse) Get() *BudgetUsageCampaignResponse {
	return v.value
}

func (v *NullableBudgetUsageCampaignResponse) Set(val *BudgetUsageCampaignResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBudgetUsageCampaignResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBudgetUsageCampaignResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBudgetUsageCampaignResponse(val *BudgetUsageCampaignResponse) *NullableBudgetUsageCampaignResponse {
	return &NullableBudgetUsageCampaignResponse{value: val, isSet: true}
}

func (v NullableBudgetUsageCampaignResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBudgetUsageCampaignResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
