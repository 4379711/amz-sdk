package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the GetSDBudgetRulesForAdvertiserResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSDBudgetRulesForAdvertiserResponse{}

// GetSDBudgetRulesForAdvertiserResponse struct for GetSDBudgetRulesForAdvertiserResponse
type GetSDBudgetRulesForAdvertiserResponse struct {
	// A list of rules created by the advertiser.
	BudgetRulesForAdvertiserResponse []SDBudgetRule `json:"budgetRulesForAdvertiserResponse,omitempty"`
	// To retrieve the next page of results, call the same operation and specify this token in the request. If the `nextToken` field is empty, there are no further results.
	NextToken *string `json:"nextToken,omitempty"`
}

// NewGetSDBudgetRulesForAdvertiserResponse instantiates a new GetSDBudgetRulesForAdvertiserResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSDBudgetRulesForAdvertiserResponse() *GetSDBudgetRulesForAdvertiserResponse {
	this := GetSDBudgetRulesForAdvertiserResponse{}
	return &this
}

// NewGetSDBudgetRulesForAdvertiserResponseWithDefaults instantiates a new GetSDBudgetRulesForAdvertiserResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSDBudgetRulesForAdvertiserResponseWithDefaults() *GetSDBudgetRulesForAdvertiserResponse {
	this := GetSDBudgetRulesForAdvertiserResponse{}
	return &this
}

// GetBudgetRulesForAdvertiserResponse returns the BudgetRulesForAdvertiserResponse field value if set, zero value otherwise.
func (o *GetSDBudgetRulesForAdvertiserResponse) GetBudgetRulesForAdvertiserResponse() []SDBudgetRule {
	if o == nil || IsNil(o.BudgetRulesForAdvertiserResponse) {
		var ret []SDBudgetRule
		return ret
	}
	return o.BudgetRulesForAdvertiserResponse
}

// GetBudgetRulesForAdvertiserResponseOk returns a tuple with the BudgetRulesForAdvertiserResponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSDBudgetRulesForAdvertiserResponse) GetBudgetRulesForAdvertiserResponseOk() ([]SDBudgetRule, bool) {
	if o == nil || IsNil(o.BudgetRulesForAdvertiserResponse) {
		return nil, false
	}
	return o.BudgetRulesForAdvertiserResponse, true
}

// HasBudgetRulesForAdvertiserResponse returns a boolean if a field has been set.
func (o *GetSDBudgetRulesForAdvertiserResponse) HasBudgetRulesForAdvertiserResponse() bool {
	if o != nil && !IsNil(o.BudgetRulesForAdvertiserResponse) {
		return true
	}

	return false
}

// SetBudgetRulesForAdvertiserResponse gets a reference to the given []SDBudgetRule and assigns it to the BudgetRulesForAdvertiserResponse field.
func (o *GetSDBudgetRulesForAdvertiserResponse) SetBudgetRulesForAdvertiserResponse(v []SDBudgetRule) {
	o.BudgetRulesForAdvertiserResponse = v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *GetSDBudgetRulesForAdvertiserResponse) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSDBudgetRulesForAdvertiserResponse) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *GetSDBudgetRulesForAdvertiserResponse) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *GetSDBudgetRulesForAdvertiserResponse) SetNextToken(v string) {
	o.NextToken = &v
}

func (o GetSDBudgetRulesForAdvertiserResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BudgetRulesForAdvertiserResponse) {
		toSerialize["budgetRulesForAdvertiserResponse"] = o.BudgetRulesForAdvertiserResponse
	}
	if !IsNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	return toSerialize, nil
}

type NullableGetSDBudgetRulesForAdvertiserResponse struct {
	value *GetSDBudgetRulesForAdvertiserResponse
	isSet bool
}

func (v NullableGetSDBudgetRulesForAdvertiserResponse) Get() *GetSDBudgetRulesForAdvertiserResponse {
	return v.value
}

func (v *NullableGetSDBudgetRulesForAdvertiserResponse) Set(val *GetSDBudgetRulesForAdvertiserResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSDBudgetRulesForAdvertiserResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSDBudgetRulesForAdvertiserResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSDBudgetRulesForAdvertiserResponse(val *GetSDBudgetRulesForAdvertiserResponse) *NullableGetSDBudgetRulesForAdvertiserResponse {
	return &NullableGetSDBudgetRulesForAdvertiserResponse{value: val, isSet: true}
}

func (v NullableGetSDBudgetRulesForAdvertiserResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetSDBudgetRulesForAdvertiserResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
