package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the OptimizationRulesAPISearchOptimizationRulesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptimizationRulesAPISearchOptimizationRulesRequest{}

// OptimizationRulesAPISearchOptimizationRulesRequest Request object for searching or getting optimization rules.
type OptimizationRulesAPISearchOptimizationRulesRequest struct {
	NextToken              *string                                     `json:"nextToken,omitempty"`
	OptimizationRuleFilter *OptimizationRulesAPIOptimizationRuleFilter `json:"optimizationRuleFilter,omitempty"`
	PageSize               *float32                                    `json:"pageSize,omitempty"`
	CampaignFilter         *OptimizationRulesAPICampaignFilter         `json:"campaignFilter,omitempty"`
}

// NewOptimizationRulesAPISearchOptimizationRulesRequest instantiates a new OptimizationRulesAPISearchOptimizationRulesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptimizationRulesAPISearchOptimizationRulesRequest() *OptimizationRulesAPISearchOptimizationRulesRequest {
	this := OptimizationRulesAPISearchOptimizationRulesRequest{}
	return &this
}

// NewOptimizationRulesAPISearchOptimizationRulesRequestWithDefaults instantiates a new OptimizationRulesAPISearchOptimizationRulesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptimizationRulesAPISearchOptimizationRulesRequestWithDefaults() *OptimizationRulesAPISearchOptimizationRulesRequest {
	this := OptimizationRulesAPISearchOptimizationRulesRequest{}
	return &this
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *OptimizationRulesAPISearchOptimizationRulesRequest) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPISearchOptimizationRulesRequest) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *OptimizationRulesAPISearchOptimizationRulesRequest) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *OptimizationRulesAPISearchOptimizationRulesRequest) SetNextToken(v string) {
	o.NextToken = &v
}

// GetOptimizationRuleFilter returns the OptimizationRuleFilter field value if set, zero value otherwise.
func (o *OptimizationRulesAPISearchOptimizationRulesRequest) GetOptimizationRuleFilter() OptimizationRulesAPIOptimizationRuleFilter {
	if o == nil || IsNil(o.OptimizationRuleFilter) {
		var ret OptimizationRulesAPIOptimizationRuleFilter
		return ret
	}
	return *o.OptimizationRuleFilter
}

// GetOptimizationRuleFilterOk returns a tuple with the OptimizationRuleFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPISearchOptimizationRulesRequest) GetOptimizationRuleFilterOk() (*OptimizationRulesAPIOptimizationRuleFilter, bool) {
	if o == nil || IsNil(o.OptimizationRuleFilter) {
		return nil, false
	}
	return o.OptimizationRuleFilter, true
}

// HasOptimizationRuleFilter returns a boolean if a field has been set.
func (o *OptimizationRulesAPISearchOptimizationRulesRequest) HasOptimizationRuleFilter() bool {
	if o != nil && !IsNil(o.OptimizationRuleFilter) {
		return true
	}

	return false
}

// SetOptimizationRuleFilter gets a reference to the given OptimizationRulesAPIOptimizationRuleFilter and assigns it to the OptimizationRuleFilter field.
func (o *OptimizationRulesAPISearchOptimizationRulesRequest) SetOptimizationRuleFilter(v OptimizationRulesAPIOptimizationRuleFilter) {
	o.OptimizationRuleFilter = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *OptimizationRulesAPISearchOptimizationRulesRequest) GetPageSize() float32 {
	if o == nil || IsNil(o.PageSize) {
		var ret float32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPISearchOptimizationRulesRequest) GetPageSizeOk() (*float32, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *OptimizationRulesAPISearchOptimizationRulesRequest) HasPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given float32 and assigns it to the PageSize field.
func (o *OptimizationRulesAPISearchOptimizationRulesRequest) SetPageSize(v float32) {
	o.PageSize = &v
}

// GetCampaignFilter returns the CampaignFilter field value if set, zero value otherwise.
func (o *OptimizationRulesAPISearchOptimizationRulesRequest) GetCampaignFilter() OptimizationRulesAPICampaignFilter {
	if o == nil || IsNil(o.CampaignFilter) {
		var ret OptimizationRulesAPICampaignFilter
		return ret
	}
	return *o.CampaignFilter
}

// GetCampaignFilterOk returns a tuple with the CampaignFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPISearchOptimizationRulesRequest) GetCampaignFilterOk() (*OptimizationRulesAPICampaignFilter, bool) {
	if o == nil || IsNil(o.CampaignFilter) {
		return nil, false
	}
	return o.CampaignFilter, true
}

// HasCampaignFilter returns a boolean if a field has been set.
func (o *OptimizationRulesAPISearchOptimizationRulesRequest) HasCampaignFilter() bool {
	if o != nil && !IsNil(o.CampaignFilter) {
		return true
	}

	return false
}

// SetCampaignFilter gets a reference to the given OptimizationRulesAPICampaignFilter and assigns it to the CampaignFilter field.
func (o *OptimizationRulesAPISearchOptimizationRulesRequest) SetCampaignFilter(v OptimizationRulesAPICampaignFilter) {
	o.CampaignFilter = &v
}

func (o OptimizationRulesAPISearchOptimizationRulesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	if !IsNil(o.OptimizationRuleFilter) {
		toSerialize["optimizationRuleFilter"] = o.OptimizationRuleFilter
	}
	if !IsNil(o.PageSize) {
		toSerialize["pageSize"] = o.PageSize
	}
	if !IsNil(o.CampaignFilter) {
		toSerialize["campaignFilter"] = o.CampaignFilter
	}
	return toSerialize, nil
}

type NullableOptimizationRulesAPISearchOptimizationRulesRequest struct {
	value *OptimizationRulesAPISearchOptimizationRulesRequest
	isSet bool
}

func (v NullableOptimizationRulesAPISearchOptimizationRulesRequest) Get() *OptimizationRulesAPISearchOptimizationRulesRequest {
	return v.value
}

func (v *NullableOptimizationRulesAPISearchOptimizationRulesRequest) Set(val *OptimizationRulesAPISearchOptimizationRulesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOptimizationRulesAPISearchOptimizationRulesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOptimizationRulesAPISearchOptimizationRulesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptimizationRulesAPISearchOptimizationRulesRequest(val *OptimizationRulesAPISearchOptimizationRulesRequest) *NullableOptimizationRulesAPISearchOptimizationRulesRequest {
	return &NullableOptimizationRulesAPISearchOptimizationRulesRequest{value: val, isSet: true}
}

func (v NullableOptimizationRulesAPISearchOptimizationRulesRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOptimizationRulesAPISearchOptimizationRulesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
