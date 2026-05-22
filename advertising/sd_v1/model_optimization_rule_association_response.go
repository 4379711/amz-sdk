package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the OptimizationRuleAssociationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptimizationRuleAssociationResponse{}

// OptimizationRuleAssociationResponse struct for OptimizationRuleAssociationResponse
type OptimizationRuleAssociationResponse struct {
	// The HTTP status code of the response.
	Code *string `json:"code,omitempty"`
	// An array of response objects. Each response object has code, details and optimizationRuleId.
	Responses []SingleOptimizationRuleAssociationResponse `json:"responses,omitempty"`
}

// NewOptimizationRuleAssociationResponse instantiates a new OptimizationRuleAssociationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptimizationRuleAssociationResponse() *OptimizationRuleAssociationResponse {
	this := OptimizationRuleAssociationResponse{}
	return &this
}

// NewOptimizationRuleAssociationResponseWithDefaults instantiates a new OptimizationRuleAssociationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptimizationRuleAssociationResponseWithDefaults() *OptimizationRuleAssociationResponse {
	this := OptimizationRuleAssociationResponse{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *OptimizationRuleAssociationResponse) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizationRuleAssociationResponse) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *OptimizationRuleAssociationResponse) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *OptimizationRuleAssociationResponse) SetCode(v string) {
	o.Code = &v
}

// GetResponses returns the Responses field value if set, zero value otherwise.
func (o *OptimizationRuleAssociationResponse) GetResponses() []SingleOptimizationRuleAssociationResponse {
	if o == nil || IsNil(o.Responses) {
		var ret []SingleOptimizationRuleAssociationResponse
		return ret
	}
	return o.Responses
}

// GetResponsesOk returns a tuple with the Responses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizationRuleAssociationResponse) GetResponsesOk() ([]SingleOptimizationRuleAssociationResponse, bool) {
	if o == nil || IsNil(o.Responses) {
		return nil, false
	}
	return o.Responses, true
}

// HasResponses returns a boolean if a field has been set.
func (o *OptimizationRuleAssociationResponse) HasResponses() bool {
	if o != nil && !IsNil(o.Responses) {
		return true
	}

	return false
}

// SetResponses gets a reference to the given []SingleOptimizationRuleAssociationResponse and assigns it to the Responses field.
func (o *OptimizationRuleAssociationResponse) SetResponses(v []SingleOptimizationRuleAssociationResponse) {
	o.Responses = v
}

func (o OptimizationRuleAssociationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Responses) {
		toSerialize["responses"] = o.Responses
	}
	return toSerialize, nil
}

type NullableOptimizationRuleAssociationResponse struct {
	value *OptimizationRuleAssociationResponse
	isSet bool
}

func (v NullableOptimizationRuleAssociationResponse) Get() *OptimizationRuleAssociationResponse {
	return v.value
}

func (v *NullableOptimizationRuleAssociationResponse) Set(val *OptimizationRuleAssociationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOptimizationRuleAssociationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOptimizationRuleAssociationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptimizationRuleAssociationResponse(val *OptimizationRuleAssociationResponse) *NullableOptimizationRuleAssociationResponse {
	return &NullableOptimizationRuleAssociationResponse{value: val, isSet: true}
}

func (v NullableOptimizationRuleAssociationResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOptimizationRuleAssociationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
