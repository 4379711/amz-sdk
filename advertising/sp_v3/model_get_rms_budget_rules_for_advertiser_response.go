package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the GetRMSBudgetRulesForAdvertiserResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetRMSBudgetRulesForAdvertiserResponse{}

// GetRMSBudgetRulesForAdvertiserResponse struct for GetRMSBudgetRulesForAdvertiserResponse
type GetRMSBudgetRulesForAdvertiserResponse struct {
	// A list of rules created by the advertiser.
	BudgetRulesForAdvertiserResponse []RMSBudgetRule `json:"budgetRulesForAdvertiserResponse,omitempty"`
	// To retrieve the next page of results, call the same operation and specify this token in the request. If the `nextToken` field is empty, there are no further results.
	NextToken *string `json:"nextToken,omitempty"`
}

// NewGetRMSBudgetRulesForAdvertiserResponse instantiates a new GetRMSBudgetRulesForAdvertiserResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRMSBudgetRulesForAdvertiserResponse() *GetRMSBudgetRulesForAdvertiserResponse {
	this := GetRMSBudgetRulesForAdvertiserResponse{}
	return &this
}

// NewGetRMSBudgetRulesForAdvertiserResponseWithDefaults instantiates a new GetRMSBudgetRulesForAdvertiserResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRMSBudgetRulesForAdvertiserResponseWithDefaults() *GetRMSBudgetRulesForAdvertiserResponse {
	this := GetRMSBudgetRulesForAdvertiserResponse{}
	return &this
}

// GetBudgetRulesForAdvertiserResponse returns the BudgetRulesForAdvertiserResponse field value if set, zero value otherwise.
func (o *GetRMSBudgetRulesForAdvertiserResponse) GetBudgetRulesForAdvertiserResponse() []RMSBudgetRule {
	if o == nil || IsNil(o.BudgetRulesForAdvertiserResponse) {
		var ret []RMSBudgetRule
		return ret
	}
	return o.BudgetRulesForAdvertiserResponse
}

// GetBudgetRulesForAdvertiserResponseOk returns a tuple with the BudgetRulesForAdvertiserResponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRMSBudgetRulesForAdvertiserResponse) GetBudgetRulesForAdvertiserResponseOk() ([]RMSBudgetRule, bool) {
	if o == nil || IsNil(o.BudgetRulesForAdvertiserResponse) {
		return nil, false
	}
	return o.BudgetRulesForAdvertiserResponse, true
}

// HasBudgetRulesForAdvertiserResponse returns a boolean if a field has been set.
func (o *GetRMSBudgetRulesForAdvertiserResponse) HasBudgetRulesForAdvertiserResponse() bool {
	if o != nil && !IsNil(o.BudgetRulesForAdvertiserResponse) {
		return true
	}

	return false
}

// SetBudgetRulesForAdvertiserResponse gets a reference to the given []RMSBudgetRule and assigns it to the BudgetRulesForAdvertiserResponse field.
func (o *GetRMSBudgetRulesForAdvertiserResponse) SetBudgetRulesForAdvertiserResponse(v []RMSBudgetRule) {
	o.BudgetRulesForAdvertiserResponse = v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *GetRMSBudgetRulesForAdvertiserResponse) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRMSBudgetRulesForAdvertiserResponse) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *GetRMSBudgetRulesForAdvertiserResponse) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *GetRMSBudgetRulesForAdvertiserResponse) SetNextToken(v string) {
	o.NextToken = &v
}

func (o GetRMSBudgetRulesForAdvertiserResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BudgetRulesForAdvertiserResponse) {
		toSerialize["budgetRulesForAdvertiserResponse"] = o.BudgetRulesForAdvertiserResponse
	}
	if !IsNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	return toSerialize, nil
}

type NullableGetRMSBudgetRulesForAdvertiserResponse struct {
	value *GetRMSBudgetRulesForAdvertiserResponse
	isSet bool
}

func (v NullableGetRMSBudgetRulesForAdvertiserResponse) Get() *GetRMSBudgetRulesForAdvertiserResponse {
	return v.value
}

func (v *NullableGetRMSBudgetRulesForAdvertiserResponse) Set(val *GetRMSBudgetRulesForAdvertiserResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRMSBudgetRulesForAdvertiserResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRMSBudgetRulesForAdvertiserResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRMSBudgetRulesForAdvertiserResponse(val *GetRMSBudgetRulesForAdvertiserResponse) *NullableGetRMSBudgetRulesForAdvertiserResponse {
	return &NullableGetRMSBudgetRulesForAdvertiserResponse{value: val, isSet: true}
}

func (v NullableGetRMSBudgetRulesForAdvertiserResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetRMSBudgetRulesForAdvertiserResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
