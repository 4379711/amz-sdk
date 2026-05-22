package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the OptimizationRulesAPIDeleteOptimizationRulesResponseErrorItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptimizationRulesAPIDeleteOptimizationRulesResponseErrorItem{}

// OptimizationRulesAPIDeleteOptimizationRulesResponseErrorItem struct for OptimizationRulesAPIDeleteOptimizationRulesResponseErrorItem
type OptimizationRulesAPIDeleteOptimizationRulesResponseErrorItem struct {
	// Index of the optimization rule in the request.
	Index int32 `json:"index"`
	// The id that uniquely identifies the optimization rule that failed deletion.
	OptimizationRuleId string `json:"optimizationRuleId"`
	// The http status error code for machine use.
	HttpStatusCode string                                              `json:"httpStatusCode"`
	SubErrors      []OptimizationRulesAPIOptimizationRuleBatchSubError `json:"subErrors"`
}

type _OptimizationRulesAPIDeleteOptimizationRulesResponseErrorItem OptimizationRulesAPIDeleteOptimizationRulesResponseErrorItem

// NewOptimizationRulesAPIDeleteOptimizationRulesResponseErrorItem instantiates a new OptimizationRulesAPIDeleteOptimizationRulesResponseErrorItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptimizationRulesAPIDeleteOptimizationRulesResponseErrorItem(index int32, optimizationRuleId string, httpStatusCode string, subErrors []OptimizationRulesAPIOptimizationRuleBatchSubError) *OptimizationRulesAPIDeleteOptimizationRulesResponseErrorItem {
	this := OptimizationRulesAPIDeleteOptimizationRulesResponseErrorItem{}
	this.Index = index
	this.OptimizationRuleId = optimizationRuleId
	this.HttpStatusCode = httpStatusCode
	this.SubErrors = subErrors
	return &this
}

// NewOptimizationRulesAPIDeleteOptimizationRulesResponseErrorItemWithDefaults instantiates a new OptimizationRulesAPIDeleteOptimizationRulesResponseErrorItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptimizationRulesAPIDeleteOptimizationRulesResponseErrorItemWithDefaults() *OptimizationRulesAPIDeleteOptimizationRulesResponseErrorItem {
	this := OptimizationRulesAPIDeleteOptimizationRulesResponseErrorItem{}
	return &this
}

// GetIndex returns the Index field value
func (o *OptimizationRulesAPIDeleteOptimizationRulesResponseErrorItem) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIDeleteOptimizationRulesResponseErrorItem) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *OptimizationRulesAPIDeleteOptimizationRulesResponseErrorItem) SetIndex(v int32) {
	o.Index = v
}

// GetOptimizationRuleId returns the OptimizationRuleId field value
func (o *OptimizationRulesAPIDeleteOptimizationRulesResponseErrorItem) GetOptimizationRuleId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OptimizationRuleId
}

// GetOptimizationRuleIdOk returns a tuple with the OptimizationRuleId field value
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIDeleteOptimizationRulesResponseErrorItem) GetOptimizationRuleIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OptimizationRuleId, true
}

// SetOptimizationRuleId sets field value
func (o *OptimizationRulesAPIDeleteOptimizationRulesResponseErrorItem) SetOptimizationRuleId(v string) {
	o.OptimizationRuleId = v
}

// GetHttpStatusCode returns the HttpStatusCode field value
func (o *OptimizationRulesAPIDeleteOptimizationRulesResponseErrorItem) GetHttpStatusCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HttpStatusCode
}

// GetHttpStatusCodeOk returns a tuple with the HttpStatusCode field value
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIDeleteOptimizationRulesResponseErrorItem) GetHttpStatusCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HttpStatusCode, true
}

// SetHttpStatusCode sets field value
func (o *OptimizationRulesAPIDeleteOptimizationRulesResponseErrorItem) SetHttpStatusCode(v string) {
	o.HttpStatusCode = v
}

// GetSubErrors returns the SubErrors field value
func (o *OptimizationRulesAPIDeleteOptimizationRulesResponseErrorItem) GetSubErrors() []OptimizationRulesAPIOptimizationRuleBatchSubError {
	if o == nil {
		var ret []OptimizationRulesAPIOptimizationRuleBatchSubError
		return ret
	}

	return o.SubErrors
}

// GetSubErrorsOk returns a tuple with the SubErrors field value
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIDeleteOptimizationRulesResponseErrorItem) GetSubErrorsOk() ([]OptimizationRulesAPIOptimizationRuleBatchSubError, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubErrors, true
}

// SetSubErrors sets field value
func (o *OptimizationRulesAPIDeleteOptimizationRulesResponseErrorItem) SetSubErrors(v []OptimizationRulesAPIOptimizationRuleBatchSubError) {
	o.SubErrors = v
}

func (o OptimizationRulesAPIDeleteOptimizationRulesResponseErrorItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["index"] = o.Index
	toSerialize["optimizationRuleId"] = o.OptimizationRuleId
	toSerialize["httpStatusCode"] = o.HttpStatusCode
	toSerialize["subErrors"] = o.SubErrors
	return toSerialize, nil
}

type NullableOptimizationRulesAPIDeleteOptimizationRulesResponseErrorItem struct {
	value *OptimizationRulesAPIDeleteOptimizationRulesResponseErrorItem
	isSet bool
}

func (v NullableOptimizationRulesAPIDeleteOptimizationRulesResponseErrorItem) Get() *OptimizationRulesAPIDeleteOptimizationRulesResponseErrorItem {
	return v.value
}

func (v *NullableOptimizationRulesAPIDeleteOptimizationRulesResponseErrorItem) Set(val *OptimizationRulesAPIDeleteOptimizationRulesResponseErrorItem) {
	v.value = val
	v.isSet = true
}

func (v NullableOptimizationRulesAPIDeleteOptimizationRulesResponseErrorItem) IsSet() bool {
	return v.isSet
}

func (v *NullableOptimizationRulesAPIDeleteOptimizationRulesResponseErrorItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptimizationRulesAPIDeleteOptimizationRulesResponseErrorItem(val *OptimizationRulesAPIDeleteOptimizationRulesResponseErrorItem) *NullableOptimizationRulesAPIDeleteOptimizationRulesResponseErrorItem {
	return &NullableOptimizationRulesAPIDeleteOptimizationRulesResponseErrorItem{value: val, isSet: true}
}

func (v NullableOptimizationRulesAPIDeleteOptimizationRulesResponseErrorItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOptimizationRulesAPIDeleteOptimizationRulesResponseErrorItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
