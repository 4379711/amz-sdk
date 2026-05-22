package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the GetBudgetRecommendationsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetBudgetRecommendationsResponseContent{}

// GetBudgetRecommendationsResponseContent struct for GetBudgetRecommendationsResponseContent
type GetBudgetRecommendationsResponseContent struct {
	// List of successful budget recommendations for campaigns.
	Success []BudgetRecommendation `json:"success"`
	// List of errors that occurred when generating budget recommendations.
	Error []BudgetRecommendationError `json:"error"`
}

type _GetBudgetRecommendationsResponseContent GetBudgetRecommendationsResponseContent

// NewGetBudgetRecommendationsResponseContent instantiates a new GetBudgetRecommendationsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBudgetRecommendationsResponseContent(success []BudgetRecommendation, error_ []BudgetRecommendationError) *GetBudgetRecommendationsResponseContent {
	this := GetBudgetRecommendationsResponseContent{}
	this.Success = success
	this.Error = error_
	return &this
}

// NewGetBudgetRecommendationsResponseContentWithDefaults instantiates a new GetBudgetRecommendationsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBudgetRecommendationsResponseContentWithDefaults() *GetBudgetRecommendationsResponseContent {
	this := GetBudgetRecommendationsResponseContent{}
	return &this
}

// GetSuccess returns the Success field value
func (o *GetBudgetRecommendationsResponseContent) GetSuccess() []BudgetRecommendation {
	if o == nil {
		var ret []BudgetRecommendation
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *GetBudgetRecommendationsResponseContent) GetSuccessOk() ([]BudgetRecommendation, bool) {
	if o == nil {
		return nil, false
	}
	return o.Success, true
}

// SetSuccess sets field value
func (o *GetBudgetRecommendationsResponseContent) SetSuccess(v []BudgetRecommendation) {
	o.Success = v
}

// GetError returns the Error field value
func (o *GetBudgetRecommendationsResponseContent) GetError() []BudgetRecommendationError {
	if o == nil {
		var ret []BudgetRecommendationError
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *GetBudgetRecommendationsResponseContent) GetErrorOk() ([]BudgetRecommendationError, bool) {
	if o == nil {
		return nil, false
	}
	return o.Error, true
}

// SetError sets field value
func (o *GetBudgetRecommendationsResponseContent) SetError(v []BudgetRecommendationError) {
	o.Error = v
}

func (o GetBudgetRecommendationsResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["success"] = o.Success
	toSerialize["error"] = o.Error
	return toSerialize, nil
}

type NullableGetBudgetRecommendationsResponseContent struct {
	value *GetBudgetRecommendationsResponseContent
	isSet bool
}

func (v NullableGetBudgetRecommendationsResponseContent) Get() *GetBudgetRecommendationsResponseContent {
	return v.value
}

func (v *NullableGetBudgetRecommendationsResponseContent) Set(val *GetBudgetRecommendationsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBudgetRecommendationsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBudgetRecommendationsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBudgetRecommendationsResponseContent(val *GetBudgetRecommendationsResponseContent) *NullableGetBudgetRecommendationsResponseContent {
	return &NullableGetBudgetRecommendationsResponseContent{value: val, isSet: true}
}

func (v NullableGetBudgetRecommendationsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetBudgetRecommendationsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
