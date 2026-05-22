package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the SingleOptimizationRuleAssociationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SingleOptimizationRuleAssociationResponse{}

// SingleOptimizationRuleAssociationResponse struct for SingleOptimizationRuleAssociationResponse
type SingleOptimizationRuleAssociationResponse struct {
	// The HTTP status code of the response.
	Code *string `json:"code,omitempty"`
	// A human-readable description of the response.
	Details *string `json:"details,omitempty"`
	// The identifier of the optimization rule.
	OptimizationRuleId *string `json:"optimizationRuleId,omitempty"`
}

// NewSingleOptimizationRuleAssociationResponse instantiates a new SingleOptimizationRuleAssociationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSingleOptimizationRuleAssociationResponse() *SingleOptimizationRuleAssociationResponse {
	this := SingleOptimizationRuleAssociationResponse{}
	return &this
}

// NewSingleOptimizationRuleAssociationResponseWithDefaults instantiates a new SingleOptimizationRuleAssociationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSingleOptimizationRuleAssociationResponseWithDefaults() *SingleOptimizationRuleAssociationResponse {
	this := SingleOptimizationRuleAssociationResponse{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *SingleOptimizationRuleAssociationResponse) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleOptimizationRuleAssociationResponse) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *SingleOptimizationRuleAssociationResponse) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *SingleOptimizationRuleAssociationResponse) SetCode(v string) {
	o.Code = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *SingleOptimizationRuleAssociationResponse) GetDetails() string {
	if o == nil || IsNil(o.Details) {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleOptimizationRuleAssociationResponse) GetDetailsOk() (*string, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *SingleOptimizationRuleAssociationResponse) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *SingleOptimizationRuleAssociationResponse) SetDetails(v string) {
	o.Details = &v
}

// GetOptimizationRuleId returns the OptimizationRuleId field value if set, zero value otherwise.
func (o *SingleOptimizationRuleAssociationResponse) GetOptimizationRuleId() string {
	if o == nil || IsNil(o.OptimizationRuleId) {
		var ret string
		return ret
	}
	return *o.OptimizationRuleId
}

// GetOptimizationRuleIdOk returns a tuple with the OptimizationRuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleOptimizationRuleAssociationResponse) GetOptimizationRuleIdOk() (*string, bool) {
	if o == nil || IsNil(o.OptimizationRuleId) {
		return nil, false
	}
	return o.OptimizationRuleId, true
}

// HasOptimizationRuleId returns a boolean if a field has been set.
func (o *SingleOptimizationRuleAssociationResponse) HasOptimizationRuleId() bool {
	if o != nil && !IsNil(o.OptimizationRuleId) {
		return true
	}

	return false
}

// SetOptimizationRuleId gets a reference to the given string and assigns it to the OptimizationRuleId field.
func (o *SingleOptimizationRuleAssociationResponse) SetOptimizationRuleId(v string) {
	o.OptimizationRuleId = &v
}

func (o SingleOptimizationRuleAssociationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	if !IsNil(o.OptimizationRuleId) {
		toSerialize["optimizationRuleId"] = o.OptimizationRuleId
	}
	return toSerialize, nil
}

type NullableSingleOptimizationRuleAssociationResponse struct {
	value *SingleOptimizationRuleAssociationResponse
	isSet bool
}

func (v NullableSingleOptimizationRuleAssociationResponse) Get() *SingleOptimizationRuleAssociationResponse {
	return v.value
}

func (v *NullableSingleOptimizationRuleAssociationResponse) Set(val *SingleOptimizationRuleAssociationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSingleOptimizationRuleAssociationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSingleOptimizationRuleAssociationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSingleOptimizationRuleAssociationResponse(val *SingleOptimizationRuleAssociationResponse) *NullableSingleOptimizationRuleAssociationResponse {
	return &NullableSingleOptimizationRuleAssociationResponse{value: val, isSet: true}
}

func (v NullableSingleOptimizationRuleAssociationResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSingleOptimizationRuleAssociationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
