package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the OptimizationRulesAPIGetAssociatedCampaignsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptimizationRulesAPIGetAssociatedCampaignsResponse{}

// OptimizationRulesAPIGetAssociatedCampaignsResponse Response object for getting campaigns associated with a rule.
type OptimizationRulesAPIGetAssociatedCampaignsResponse struct {
	AssociatedCampaigns []OptimizationRulesAPISingleCampaignRuleAssociationStatus `json:"associatedCampaigns,omitempty"`
	// An enumerated error code for machine use.
	Code *string `json:"code,omitempty"`
	// To retrieve the next page of results, call the same operation and specify this token in the request. If the nextToken field is empty, there are no further results.
	NextToken *string `json:"nextToken,omitempty"`
}

// NewOptimizationRulesAPIGetAssociatedCampaignsResponse instantiates a new OptimizationRulesAPIGetAssociatedCampaignsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptimizationRulesAPIGetAssociatedCampaignsResponse() *OptimizationRulesAPIGetAssociatedCampaignsResponse {
	this := OptimizationRulesAPIGetAssociatedCampaignsResponse{}
	return &this
}

// NewOptimizationRulesAPIGetAssociatedCampaignsResponseWithDefaults instantiates a new OptimizationRulesAPIGetAssociatedCampaignsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptimizationRulesAPIGetAssociatedCampaignsResponseWithDefaults() *OptimizationRulesAPIGetAssociatedCampaignsResponse {
	this := OptimizationRulesAPIGetAssociatedCampaignsResponse{}
	return &this
}

// GetAssociatedCampaigns returns the AssociatedCampaigns field value if set, zero value otherwise.
func (o *OptimizationRulesAPIGetAssociatedCampaignsResponse) GetAssociatedCampaigns() []OptimizationRulesAPISingleCampaignRuleAssociationStatus {
	if o == nil || IsNil(o.AssociatedCampaigns) {
		var ret []OptimizationRulesAPISingleCampaignRuleAssociationStatus
		return ret
	}
	return o.AssociatedCampaigns
}

// GetAssociatedCampaignsOk returns a tuple with the AssociatedCampaigns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIGetAssociatedCampaignsResponse) GetAssociatedCampaignsOk() ([]OptimizationRulesAPISingleCampaignRuleAssociationStatus, bool) {
	if o == nil || IsNil(o.AssociatedCampaigns) {
		return nil, false
	}
	return o.AssociatedCampaigns, true
}

// HasAssociatedCampaigns returns a boolean if a field has been set.
func (o *OptimizationRulesAPIGetAssociatedCampaignsResponse) HasAssociatedCampaigns() bool {
	if o != nil && !IsNil(o.AssociatedCampaigns) {
		return true
	}

	return false
}

// SetAssociatedCampaigns gets a reference to the given []OptimizationRulesAPISingleCampaignRuleAssociationStatus and assigns it to the AssociatedCampaigns field.
func (o *OptimizationRulesAPIGetAssociatedCampaignsResponse) SetAssociatedCampaigns(v []OptimizationRulesAPISingleCampaignRuleAssociationStatus) {
	o.AssociatedCampaigns = v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *OptimizationRulesAPIGetAssociatedCampaignsResponse) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIGetAssociatedCampaignsResponse) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *OptimizationRulesAPIGetAssociatedCampaignsResponse) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *OptimizationRulesAPIGetAssociatedCampaignsResponse) SetCode(v string) {
	o.Code = &v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *OptimizationRulesAPIGetAssociatedCampaignsResponse) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIGetAssociatedCampaignsResponse) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *OptimizationRulesAPIGetAssociatedCampaignsResponse) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *OptimizationRulesAPIGetAssociatedCampaignsResponse) SetNextToken(v string) {
	o.NextToken = &v
}

func (o OptimizationRulesAPIGetAssociatedCampaignsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AssociatedCampaigns) {
		toSerialize["associatedCampaigns"] = o.AssociatedCampaigns
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	return toSerialize, nil
}

type NullableOptimizationRulesAPIGetAssociatedCampaignsResponse struct {
	value *OptimizationRulesAPIGetAssociatedCampaignsResponse
	isSet bool
}

func (v NullableOptimizationRulesAPIGetAssociatedCampaignsResponse) Get() *OptimizationRulesAPIGetAssociatedCampaignsResponse {
	return v.value
}

func (v *NullableOptimizationRulesAPIGetAssociatedCampaignsResponse) Set(val *OptimizationRulesAPIGetAssociatedCampaignsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOptimizationRulesAPIGetAssociatedCampaignsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOptimizationRulesAPIGetAssociatedCampaignsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptimizationRulesAPIGetAssociatedCampaignsResponse(val *OptimizationRulesAPIGetAssociatedCampaignsResponse) *NullableOptimizationRulesAPIGetAssociatedCampaignsResponse {
	return &NullableOptimizationRulesAPIGetAssociatedCampaignsResponse{value: val, isSet: true}
}

func (v NullableOptimizationRulesAPIGetAssociatedCampaignsResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOptimizationRulesAPIGetAssociatedCampaignsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
