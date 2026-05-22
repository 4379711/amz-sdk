package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the BudgetRuleResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BudgetRuleResponse{}

// BudgetRuleResponse struct for BudgetRuleResponse
type BudgetRuleResponse struct {
	// An enumerated success or error code for machine use.
	Code *string `json:"code,omitempty"`
	// A human-readable description of the error, if unsuccessful
	Details *string `json:"details,omitempty"`
	// The rule identifier.
	RuleId                *string  `json:"ruleId,omitempty"`
	AssociatedCampaignIds []string `json:"associatedCampaignIds,omitempty"`
}

// NewBudgetRuleResponse instantiates a new BudgetRuleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBudgetRuleResponse() *BudgetRuleResponse {
	this := BudgetRuleResponse{}
	return &this
}

// NewBudgetRuleResponseWithDefaults instantiates a new BudgetRuleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBudgetRuleResponseWithDefaults() *BudgetRuleResponse {
	this := BudgetRuleResponse{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *BudgetRuleResponse) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetRuleResponse) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *BudgetRuleResponse) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *BudgetRuleResponse) SetCode(v string) {
	o.Code = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *BudgetRuleResponse) GetDetails() string {
	if o == nil || IsNil(o.Details) {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetRuleResponse) GetDetailsOk() (*string, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *BudgetRuleResponse) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *BudgetRuleResponse) SetDetails(v string) {
	o.Details = &v
}

// GetRuleId returns the RuleId field value if set, zero value otherwise.
func (o *BudgetRuleResponse) GetRuleId() string {
	if o == nil || IsNil(o.RuleId) {
		var ret string
		return ret
	}
	return *o.RuleId
}

// GetRuleIdOk returns a tuple with the RuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetRuleResponse) GetRuleIdOk() (*string, bool) {
	if o == nil || IsNil(o.RuleId) {
		return nil, false
	}
	return o.RuleId, true
}

// HasRuleId returns a boolean if a field has been set.
func (o *BudgetRuleResponse) HasRuleId() bool {
	if o != nil && !IsNil(o.RuleId) {
		return true
	}

	return false
}

// SetRuleId gets a reference to the given string and assigns it to the RuleId field.
func (o *BudgetRuleResponse) SetRuleId(v string) {
	o.RuleId = &v
}

// GetAssociatedCampaignIds returns the AssociatedCampaignIds field value if set, zero value otherwise.
func (o *BudgetRuleResponse) GetAssociatedCampaignIds() []string {
	if o == nil || IsNil(o.AssociatedCampaignIds) {
		var ret []string
		return ret
	}
	return o.AssociatedCampaignIds
}

// GetAssociatedCampaignIdsOk returns a tuple with the AssociatedCampaignIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetRuleResponse) GetAssociatedCampaignIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.AssociatedCampaignIds) {
		return nil, false
	}
	return o.AssociatedCampaignIds, true
}

// HasAssociatedCampaignIds returns a boolean if a field has been set.
func (o *BudgetRuleResponse) HasAssociatedCampaignIds() bool {
	if o != nil && !IsNil(o.AssociatedCampaignIds) {
		return true
	}

	return false
}

// SetAssociatedCampaignIds gets a reference to the given []string and assigns it to the AssociatedCampaignIds field.
func (o *BudgetRuleResponse) SetAssociatedCampaignIds(v []string) {
	o.AssociatedCampaignIds = v
}

func (o BudgetRuleResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	if !IsNil(o.RuleId) {
		toSerialize["ruleId"] = o.RuleId
	}
	if !IsNil(o.AssociatedCampaignIds) {
		toSerialize["associatedCampaignIds"] = o.AssociatedCampaignIds
	}
	return toSerialize, nil
}

type NullableBudgetRuleResponse struct {
	value *BudgetRuleResponse
	isSet bool
}

func (v NullableBudgetRuleResponse) Get() *BudgetRuleResponse {
	return v.value
}

func (v *NullableBudgetRuleResponse) Set(val *BudgetRuleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBudgetRuleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBudgetRuleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBudgetRuleResponse(val *BudgetRuleResponse) *NullableBudgetRuleResponse {
	return &NullableBudgetRuleResponse{value: val, isSet: true}
}

func (v NullableBudgetRuleResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBudgetRuleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
