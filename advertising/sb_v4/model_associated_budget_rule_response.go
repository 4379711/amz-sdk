package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the AssociatedBudgetRuleResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssociatedBudgetRuleResponse{}

// AssociatedBudgetRuleResponse struct for AssociatedBudgetRuleResponse
type AssociatedBudgetRuleResponse struct {
	// An enumerated success or error code for machine use.
	Code *string `json:"code,omitempty"`
	// A human-readable description of the error, if unsuccessful
	Details *string `json:"details,omitempty"`
	// The budget rule identifier.
	RuleId *string `json:"ruleId,omitempty"`
}

// NewAssociatedBudgetRuleResponse instantiates a new AssociatedBudgetRuleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssociatedBudgetRuleResponse() *AssociatedBudgetRuleResponse {
	this := AssociatedBudgetRuleResponse{}
	return &this
}

// NewAssociatedBudgetRuleResponseWithDefaults instantiates a new AssociatedBudgetRuleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssociatedBudgetRuleResponseWithDefaults() *AssociatedBudgetRuleResponse {
	this := AssociatedBudgetRuleResponse{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *AssociatedBudgetRuleResponse) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssociatedBudgetRuleResponse) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *AssociatedBudgetRuleResponse) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *AssociatedBudgetRuleResponse) SetCode(v string) {
	o.Code = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *AssociatedBudgetRuleResponse) GetDetails() string {
	if o == nil || IsNil(o.Details) {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssociatedBudgetRuleResponse) GetDetailsOk() (*string, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *AssociatedBudgetRuleResponse) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *AssociatedBudgetRuleResponse) SetDetails(v string) {
	o.Details = &v
}

// GetRuleId returns the RuleId field value if set, zero value otherwise.
func (o *AssociatedBudgetRuleResponse) GetRuleId() string {
	if o == nil || IsNil(o.RuleId) {
		var ret string
		return ret
	}
	return *o.RuleId
}

// GetRuleIdOk returns a tuple with the RuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssociatedBudgetRuleResponse) GetRuleIdOk() (*string, bool) {
	if o == nil || IsNil(o.RuleId) {
		return nil, false
	}
	return o.RuleId, true
}

// HasRuleId returns a boolean if a field has been set.
func (o *AssociatedBudgetRuleResponse) HasRuleId() bool {
	if o != nil && !IsNil(o.RuleId) {
		return true
	}

	return false
}

// SetRuleId gets a reference to the given string and assigns it to the RuleId field.
func (o *AssociatedBudgetRuleResponse) SetRuleId(v string) {
	o.RuleId = &v
}

func (o AssociatedBudgetRuleResponse) ToMap() (map[string]interface{}, error) {
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
	return toSerialize, nil
}

type NullableAssociatedBudgetRuleResponse struct {
	value *AssociatedBudgetRuleResponse
	isSet bool
}

func (v NullableAssociatedBudgetRuleResponse) Get() *AssociatedBudgetRuleResponse {
	return v.value
}

func (v *NullableAssociatedBudgetRuleResponse) Set(val *AssociatedBudgetRuleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAssociatedBudgetRuleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAssociatedBudgetRuleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssociatedBudgetRuleResponse(val *AssociatedBudgetRuleResponse) *NullableAssociatedBudgetRuleResponse {
	return &NullableAssociatedBudgetRuleResponse{value: val, isSet: true}
}

func (v NullableAssociatedBudgetRuleResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAssociatedBudgetRuleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
