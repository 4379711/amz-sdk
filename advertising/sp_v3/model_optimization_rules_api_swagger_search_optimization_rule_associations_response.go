package sp_v3

import "github.com/bytedance/sonic"

// checks if the OptimizationRulesAPISearchOptimizationRuleAssociationsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptimizationRulesAPISearchOptimizationRuleAssociationsResponse{}

// OptimizationRulesAPISearchOptimizationRuleAssociationsResponse Response object for searching optimization rules associations.
type OptimizationRulesAPISearchOptimizationRuleAssociationsResponse struct {
	Associations []OptimizationRulesAPISearchOptimizationRuleAssociationsResponseItem `json:"associations"`
	// To retrieve the next page of results, call the same operation and specify this token in the request. If the field is empty, the first page of results will be returned.
	NextToken *string `json:"nextToken,omitempty"`
}

type _OptimizationRulesAPISearchOptimizationRuleAssociationsResponse OptimizationRulesAPISearchOptimizationRuleAssociationsResponse

// NewOptimizationRulesAPISearchOptimizationRuleAssociationsResponse instantiates a new OptimizationRulesAPISearchOptimizationRuleAssociationsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptimizationRulesAPISearchOptimizationRuleAssociationsResponse(associations []OptimizationRulesAPISearchOptimizationRuleAssociationsResponseItem) *OptimizationRulesAPISearchOptimizationRuleAssociationsResponse {
	this := OptimizationRulesAPISearchOptimizationRuleAssociationsResponse{}
	this.Associations = associations
	return &this
}

// NewOptimizationRulesAPISearchOptimizationRuleAssociationsResponseWithDefaults instantiates a new OptimizationRulesAPISearchOptimizationRuleAssociationsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptimizationRulesAPISearchOptimizationRuleAssociationsResponseWithDefaults() *OptimizationRulesAPISearchOptimizationRuleAssociationsResponse {
	this := OptimizationRulesAPISearchOptimizationRuleAssociationsResponse{}
	return &this
}

// GetAssociations returns the Associations field value
func (o *OptimizationRulesAPISearchOptimizationRuleAssociationsResponse) GetAssociations() []OptimizationRulesAPISearchOptimizationRuleAssociationsResponseItem {
	if o == nil {
		var ret []OptimizationRulesAPISearchOptimizationRuleAssociationsResponseItem
		return ret
	}

	return o.Associations
}

// GetAssociationsOk returns a tuple with the Associations field value
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPISearchOptimizationRuleAssociationsResponse) GetAssociationsOk() ([]OptimizationRulesAPISearchOptimizationRuleAssociationsResponseItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.Associations, true
}

// SetAssociations sets field value
func (o *OptimizationRulesAPISearchOptimizationRuleAssociationsResponse) SetAssociations(v []OptimizationRulesAPISearchOptimizationRuleAssociationsResponseItem) {
	o.Associations = v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *OptimizationRulesAPISearchOptimizationRuleAssociationsResponse) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPISearchOptimizationRuleAssociationsResponse) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *OptimizationRulesAPISearchOptimizationRuleAssociationsResponse) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *OptimizationRulesAPISearchOptimizationRuleAssociationsResponse) SetNextToken(v string) {
	o.NextToken = &v
}

func (o OptimizationRulesAPISearchOptimizationRuleAssociationsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["associations"] = o.Associations
	if !IsNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	return toSerialize, nil
}

type NullableOptimizationRulesAPISearchOptimizationRuleAssociationsResponse struct {
	value *OptimizationRulesAPISearchOptimizationRuleAssociationsResponse
	isSet bool
}

func (v NullableOptimizationRulesAPISearchOptimizationRuleAssociationsResponse) Get() *OptimizationRulesAPISearchOptimizationRuleAssociationsResponse {
	return v.value
}

func (v *NullableOptimizationRulesAPISearchOptimizationRuleAssociationsResponse) Set(val *OptimizationRulesAPISearchOptimizationRuleAssociationsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOptimizationRulesAPISearchOptimizationRuleAssociationsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOptimizationRulesAPISearchOptimizationRuleAssociationsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptimizationRulesAPISearchOptimizationRuleAssociationsResponse(val *OptimizationRulesAPISearchOptimizationRuleAssociationsResponse) *NullableOptimizationRulesAPISearchOptimizationRuleAssociationsResponse {
	return &NullableOptimizationRulesAPISearchOptimizationRuleAssociationsResponse{value: val, isSet: true}
}

func (v NullableOptimizationRulesAPISearchOptimizationRuleAssociationsResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOptimizationRulesAPISearchOptimizationRuleAssociationsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
