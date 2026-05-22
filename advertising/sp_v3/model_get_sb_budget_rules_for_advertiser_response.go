package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the GetSBBudgetRulesForAdvertiserResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSBBudgetRulesForAdvertiserResponse{}

// GetSBBudgetRulesForAdvertiserResponse struct for GetSBBudgetRulesForAdvertiserResponse
type GetSBBudgetRulesForAdvertiserResponse struct {
	// A list of rules created by the advertiser.
	BudgetRulesForAdvertiserResponse []SBBudgetRule `json:"budgetRulesForAdvertiserResponse,omitempty"`
	// To retrieve the next page of results, call the same operation and specify this token in the request. If the `nextToken` field is empty, there are no further results.
	NextToken *string `json:"nextToken,omitempty"`
}

// NewGetSBBudgetRulesForAdvertiserResponse instantiates a new GetSBBudgetRulesForAdvertiserResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSBBudgetRulesForAdvertiserResponse() *GetSBBudgetRulesForAdvertiserResponse {
	this := GetSBBudgetRulesForAdvertiserResponse{}
	return &this
}

// NewGetSBBudgetRulesForAdvertiserResponseWithDefaults instantiates a new GetSBBudgetRulesForAdvertiserResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSBBudgetRulesForAdvertiserResponseWithDefaults() *GetSBBudgetRulesForAdvertiserResponse {
	this := GetSBBudgetRulesForAdvertiserResponse{}
	return &this
}

// GetBudgetRulesForAdvertiserResponse returns the BudgetRulesForAdvertiserResponse field value if set, zero value otherwise.
func (o *GetSBBudgetRulesForAdvertiserResponse) GetBudgetRulesForAdvertiserResponse() []SBBudgetRule {
	if o == nil || IsNil(o.BudgetRulesForAdvertiserResponse) {
		var ret []SBBudgetRule
		return ret
	}
	return o.BudgetRulesForAdvertiserResponse
}

// GetBudgetRulesForAdvertiserResponseOk returns a tuple with the BudgetRulesForAdvertiserResponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSBBudgetRulesForAdvertiserResponse) GetBudgetRulesForAdvertiserResponseOk() ([]SBBudgetRule, bool) {
	if o == nil || IsNil(o.BudgetRulesForAdvertiserResponse) {
		return nil, false
	}
	return o.BudgetRulesForAdvertiserResponse, true
}

// HasBudgetRulesForAdvertiserResponse returns a boolean if a field has been set.
func (o *GetSBBudgetRulesForAdvertiserResponse) HasBudgetRulesForAdvertiserResponse() bool {
	if o != nil && !IsNil(o.BudgetRulesForAdvertiserResponse) {
		return true
	}

	return false
}

// SetBudgetRulesForAdvertiserResponse gets a reference to the given []SBBudgetRule and assigns it to the BudgetRulesForAdvertiserResponse field.
func (o *GetSBBudgetRulesForAdvertiserResponse) SetBudgetRulesForAdvertiserResponse(v []SBBudgetRule) {
	o.BudgetRulesForAdvertiserResponse = v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *GetSBBudgetRulesForAdvertiserResponse) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSBBudgetRulesForAdvertiserResponse) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *GetSBBudgetRulesForAdvertiserResponse) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *GetSBBudgetRulesForAdvertiserResponse) SetNextToken(v string) {
	o.NextToken = &v
}

func (o GetSBBudgetRulesForAdvertiserResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BudgetRulesForAdvertiserResponse) {
		toSerialize["budgetRulesForAdvertiserResponse"] = o.BudgetRulesForAdvertiserResponse
	}
	if !IsNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	return toSerialize, nil
}

type NullableGetSBBudgetRulesForAdvertiserResponse struct {
	value *GetSBBudgetRulesForAdvertiserResponse
	isSet bool
}

func (v NullableGetSBBudgetRulesForAdvertiserResponse) Get() *GetSBBudgetRulesForAdvertiserResponse {
	return v.value
}

func (v *NullableGetSBBudgetRulesForAdvertiserResponse) Set(val *GetSBBudgetRulesForAdvertiserResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSBBudgetRulesForAdvertiserResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSBBudgetRulesForAdvertiserResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSBBudgetRulesForAdvertiserResponse(val *GetSBBudgetRulesForAdvertiserResponse) *NullableGetSBBudgetRulesForAdvertiserResponse {
	return &NullableGetSBBudgetRulesForAdvertiserResponse{value: val, isSet: true}
}

func (v NullableGetSBBudgetRulesForAdvertiserResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetSBBudgetRulesForAdvertiserResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
