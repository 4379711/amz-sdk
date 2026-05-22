package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the OptimizationRulesAPIGetRuleRecommendationsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptimizationRulesAPIGetRuleRecommendationsResponse{}

// OptimizationRulesAPIGetRuleRecommendationsResponse Response object for GetRuleRecommendations API
type OptimizationRulesAPIGetRuleRecommendationsResponse struct {
	// An enumerated success or error code for machine use.
	Code *string `json:"code,omitempty"`
	// To retrieve the next page of results, call the same operation and specify this token in the request. If the field is empty, the first page of results will be returned.
	NextToken *string                                  `json:"nextToken,omitempty"`
	Responses []OptimizationRulesAPIRuleRecommendation `json:"responses,omitempty"`
}

// NewOptimizationRulesAPIGetRuleRecommendationsResponse instantiates a new OptimizationRulesAPIGetRuleRecommendationsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptimizationRulesAPIGetRuleRecommendationsResponse() *OptimizationRulesAPIGetRuleRecommendationsResponse {
	this := OptimizationRulesAPIGetRuleRecommendationsResponse{}
	return &this
}

// NewOptimizationRulesAPIGetRuleRecommendationsResponseWithDefaults instantiates a new OptimizationRulesAPIGetRuleRecommendationsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptimizationRulesAPIGetRuleRecommendationsResponseWithDefaults() *OptimizationRulesAPIGetRuleRecommendationsResponse {
	this := OptimizationRulesAPIGetRuleRecommendationsResponse{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *OptimizationRulesAPIGetRuleRecommendationsResponse) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIGetRuleRecommendationsResponse) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *OptimizationRulesAPIGetRuleRecommendationsResponse) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *OptimizationRulesAPIGetRuleRecommendationsResponse) SetCode(v string) {
	o.Code = &v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *OptimizationRulesAPIGetRuleRecommendationsResponse) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIGetRuleRecommendationsResponse) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *OptimizationRulesAPIGetRuleRecommendationsResponse) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *OptimizationRulesAPIGetRuleRecommendationsResponse) SetNextToken(v string) {
	o.NextToken = &v
}

// GetResponses returns the Responses field value if set, zero value otherwise.
func (o *OptimizationRulesAPIGetRuleRecommendationsResponse) GetResponses() []OptimizationRulesAPIRuleRecommendation {
	if o == nil || IsNil(o.Responses) {
		var ret []OptimizationRulesAPIRuleRecommendation
		return ret
	}
	return o.Responses
}

// GetResponsesOk returns a tuple with the Responses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIGetRuleRecommendationsResponse) GetResponsesOk() ([]OptimizationRulesAPIRuleRecommendation, bool) {
	if o == nil || IsNil(o.Responses) {
		return nil, false
	}
	return o.Responses, true
}

// HasResponses returns a boolean if a field has been set.
func (o *OptimizationRulesAPIGetRuleRecommendationsResponse) HasResponses() bool {
	if o != nil && !IsNil(o.Responses) {
		return true
	}

	return false
}

// SetResponses gets a reference to the given []OptimizationRulesAPIRuleRecommendation and assigns it to the Responses field.
func (o *OptimizationRulesAPIGetRuleRecommendationsResponse) SetResponses(v []OptimizationRulesAPIRuleRecommendation) {
	o.Responses = v
}

func (o OptimizationRulesAPIGetRuleRecommendationsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	if !IsNil(o.Responses) {
		toSerialize["responses"] = o.Responses
	}
	return toSerialize, nil
}

type NullableOptimizationRulesAPIGetRuleRecommendationsResponse struct {
	value *OptimizationRulesAPIGetRuleRecommendationsResponse
	isSet bool
}

func (v NullableOptimizationRulesAPIGetRuleRecommendationsResponse) Get() *OptimizationRulesAPIGetRuleRecommendationsResponse {
	return v.value
}

func (v *NullableOptimizationRulesAPIGetRuleRecommendationsResponse) Set(val *OptimizationRulesAPIGetRuleRecommendationsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOptimizationRulesAPIGetRuleRecommendationsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOptimizationRulesAPIGetRuleRecommendationsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptimizationRulesAPIGetRuleRecommendationsResponse(val *OptimizationRulesAPIGetRuleRecommendationsResponse) *NullableOptimizationRulesAPIGetRuleRecommendationsResponse {
	return &NullableOptimizationRulesAPIGetRuleRecommendationsResponse{value: val, isSet: true}
}

func (v NullableOptimizationRulesAPIGetRuleRecommendationsResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOptimizationRulesAPIGetRuleRecommendationsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
