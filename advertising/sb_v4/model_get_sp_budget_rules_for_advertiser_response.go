package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the GetSPBudgetRulesForAdvertiserResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSPBudgetRulesForAdvertiserResponse{}

// GetSPBudgetRulesForAdvertiserResponse struct for GetSPBudgetRulesForAdvertiserResponse
type GetSPBudgetRulesForAdvertiserResponse struct {
	// A list of rules created by the advertiser.
	BudgetRulesForAdvertiserResponse []SPBudgetRule `json:"budgetRulesForAdvertiserResponse,omitempty"`
	// To retrieve the next page of results, call the same operation and specify this token in the request. If the `nextToken` field is empty, there are no further results.
	NextToken *string `json:"nextToken,omitempty"`
}

// NewGetSPBudgetRulesForAdvertiserResponse instantiates a new GetSPBudgetRulesForAdvertiserResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSPBudgetRulesForAdvertiserResponse() *GetSPBudgetRulesForAdvertiserResponse {
	this := GetSPBudgetRulesForAdvertiserResponse{}
	return &this
}

// NewGetSPBudgetRulesForAdvertiserResponseWithDefaults instantiates a new GetSPBudgetRulesForAdvertiserResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSPBudgetRulesForAdvertiserResponseWithDefaults() *GetSPBudgetRulesForAdvertiserResponse {
	this := GetSPBudgetRulesForAdvertiserResponse{}
	return &this
}

// GetBudgetRulesForAdvertiserResponse returns the BudgetRulesForAdvertiserResponse field value if set, zero value otherwise.
func (o *GetSPBudgetRulesForAdvertiserResponse) GetBudgetRulesForAdvertiserResponse() []SPBudgetRule {
	if o == nil || IsNil(o.BudgetRulesForAdvertiserResponse) {
		var ret []SPBudgetRule
		return ret
	}
	return o.BudgetRulesForAdvertiserResponse
}

// GetBudgetRulesForAdvertiserResponseOk returns a tuple with the BudgetRulesForAdvertiserResponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSPBudgetRulesForAdvertiserResponse) GetBudgetRulesForAdvertiserResponseOk() ([]SPBudgetRule, bool) {
	if o == nil || IsNil(o.BudgetRulesForAdvertiserResponse) {
		return nil, false
	}
	return o.BudgetRulesForAdvertiserResponse, true
}

// HasBudgetRulesForAdvertiserResponse returns a boolean if a field has been set.
func (o *GetSPBudgetRulesForAdvertiserResponse) HasBudgetRulesForAdvertiserResponse() bool {
	if o != nil && !IsNil(o.BudgetRulesForAdvertiserResponse) {
		return true
	}

	return false
}

// SetBudgetRulesForAdvertiserResponse gets a reference to the given []SPBudgetRule and assigns it to the BudgetRulesForAdvertiserResponse field.
func (o *GetSPBudgetRulesForAdvertiserResponse) SetBudgetRulesForAdvertiserResponse(v []SPBudgetRule) {
	o.BudgetRulesForAdvertiserResponse = v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *GetSPBudgetRulesForAdvertiserResponse) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSPBudgetRulesForAdvertiserResponse) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *GetSPBudgetRulesForAdvertiserResponse) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *GetSPBudgetRulesForAdvertiserResponse) SetNextToken(v string) {
	o.NextToken = &v
}

func (o GetSPBudgetRulesForAdvertiserResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BudgetRulesForAdvertiserResponse) {
		toSerialize["budgetRulesForAdvertiserResponse"] = o.BudgetRulesForAdvertiserResponse
	}
	if !IsNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	return toSerialize, nil
}

type NullableGetSPBudgetRulesForAdvertiserResponse struct {
	value *GetSPBudgetRulesForAdvertiserResponse
	isSet bool
}

func (v NullableGetSPBudgetRulesForAdvertiserResponse) Get() *GetSPBudgetRulesForAdvertiserResponse {
	return v.value
}

func (v *NullableGetSPBudgetRulesForAdvertiserResponse) Set(val *GetSPBudgetRulesForAdvertiserResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSPBudgetRulesForAdvertiserResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSPBudgetRulesForAdvertiserResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSPBudgetRulesForAdvertiserResponse(val *GetSPBudgetRulesForAdvertiserResponse) *NullableGetSPBudgetRulesForAdvertiserResponse {
	return &NullableGetSPBudgetRulesForAdvertiserResponse{value: val, isSet: true}
}

func (v NullableGetSPBudgetRulesForAdvertiserResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetSPBudgetRulesForAdvertiserResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
