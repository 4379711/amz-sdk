package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseErrorItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseErrorItem{}

// OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseErrorItem struct for OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseErrorItem
type OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseErrorItem struct {
	// Index of the optimization rule in the request.
	Index int32 `json:"index"`
	// The id that uniquely identifies the optimization rule that was failed to be disassociated.
	OptimizationRuleId string `json:"optimizationRuleId"`
	// The http status error code for machine use.
	HttpStatusCode string                                              `json:"httpStatusCode"`
	SubErrors      []OptimizationRulesAPIOptimizationRuleBatchSubError `json:"subErrors"`
}

type _OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseErrorItem OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseErrorItem

// NewOptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseErrorItem instantiates a new OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseErrorItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseErrorItem(index int32, optimizationRuleId string, httpStatusCode string, subErrors []OptimizationRulesAPIOptimizationRuleBatchSubError) *OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseErrorItem {
	this := OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseErrorItem{}
	this.Index = index
	this.OptimizationRuleId = optimizationRuleId
	this.HttpStatusCode = httpStatusCode
	this.SubErrors = subErrors
	return &this
}

// NewOptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseErrorItemWithDefaults instantiates a new OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseErrorItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseErrorItemWithDefaults() *OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseErrorItem {
	this := OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseErrorItem{}
	return &this
}

// GetIndex returns the Index field value
func (o *OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseErrorItem) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseErrorItem) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseErrorItem) SetIndex(v int32) {
	o.Index = v
}

// GetOptimizationRuleId returns the OptimizationRuleId field value
func (o *OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseErrorItem) GetOptimizationRuleId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OptimizationRuleId
}

// GetOptimizationRuleIdOk returns a tuple with the OptimizationRuleId field value
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseErrorItem) GetOptimizationRuleIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OptimizationRuleId, true
}

// SetOptimizationRuleId sets field value
func (o *OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseErrorItem) SetOptimizationRuleId(v string) {
	o.OptimizationRuleId = v
}

// GetHttpStatusCode returns the HttpStatusCode field value
func (o *OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseErrorItem) GetHttpStatusCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HttpStatusCode
}

// GetHttpStatusCodeOk returns a tuple with the HttpStatusCode field value
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseErrorItem) GetHttpStatusCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HttpStatusCode, true
}

// SetHttpStatusCode sets field value
func (o *OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseErrorItem) SetHttpStatusCode(v string) {
	o.HttpStatusCode = v
}

// GetSubErrors returns the SubErrors field value
func (o *OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseErrorItem) GetSubErrors() []OptimizationRulesAPIOptimizationRuleBatchSubError {
	if o == nil {
		var ret []OptimizationRulesAPIOptimizationRuleBatchSubError
		return ret
	}

	return o.SubErrors
}

// GetSubErrorsOk returns a tuple with the SubErrors field value
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseErrorItem) GetSubErrorsOk() ([]OptimizationRulesAPIOptimizationRuleBatchSubError, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubErrors, true
}

// SetSubErrors sets field value
func (o *OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseErrorItem) SetSubErrors(v []OptimizationRulesAPIOptimizationRuleBatchSubError) {
	o.SubErrors = v
}

func (o OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseErrorItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["index"] = o.Index
	toSerialize["optimizationRuleId"] = o.OptimizationRuleId
	toSerialize["httpStatusCode"] = o.HttpStatusCode
	toSerialize["subErrors"] = o.SubErrors
	return toSerialize, nil
}

type NullableOptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseErrorItem struct {
	value *OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseErrorItem
	isSet bool
}

func (v NullableOptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseErrorItem) Get() *OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseErrorItem {
	return v.value
}

func (v *NullableOptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseErrorItem) Set(val *OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseErrorItem) {
	v.value = val
	v.isSet = true
}

func (v NullableOptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseErrorItem) IsSet() bool {
	return v.isSet
}

func (v *NullableOptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseErrorItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseErrorItem(val *OptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseErrorItem) *NullableOptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseErrorItem {
	return &NullableOptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseErrorItem{value: val, isSet: true}
}

func (v NullableOptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseErrorItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOptimizationRulesAPIDisassociateOptimizationRulesFromCampaignResponseErrorItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
