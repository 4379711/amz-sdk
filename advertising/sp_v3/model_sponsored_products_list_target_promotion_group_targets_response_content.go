package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsListTargetPromotionGroupTargetsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsListTargetPromotionGroupTargetsResponseContent{}

// SponsoredProductsListTargetPromotionGroupTargetsResponseContent Response object for querying target promotion group targets.
type SponsoredProductsListTargetPromotionGroupTargetsResponseContent struct {
	// The total number of results available.
	TotalResults *int64 `json:"totalResults,omitempty"`
	// To retrieve the next page of results, call the same operation and specify this token in the     request. If the nextToken field is empty, there are no further results.
	NextToken *string                   `json:"nextToken,omitempty"`
	Targets   []SponsoredProductsTarget `json:"targets,omitempty"`
}

// NewSponsoredProductsListTargetPromotionGroupTargetsResponseContent instantiates a new SponsoredProductsListTargetPromotionGroupTargetsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsListTargetPromotionGroupTargetsResponseContent() *SponsoredProductsListTargetPromotionGroupTargetsResponseContent {
	this := SponsoredProductsListTargetPromotionGroupTargetsResponseContent{}
	return &this
}

// NewSponsoredProductsListTargetPromotionGroupTargetsResponseContentWithDefaults instantiates a new SponsoredProductsListTargetPromotionGroupTargetsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsListTargetPromotionGroupTargetsResponseContentWithDefaults() *SponsoredProductsListTargetPromotionGroupTargetsResponseContent {
	this := SponsoredProductsListTargetPromotionGroupTargetsResponseContent{}
	return &this
}

// GetTotalResults returns the TotalResults field value if set, zero value otherwise.
func (o *SponsoredProductsListTargetPromotionGroupTargetsResponseContent) GetTotalResults() int64 {
	if o == nil || IsNil(o.TotalResults) {
		var ret int64
		return ret
	}
	return *o.TotalResults
}

// GetTotalResultsOk returns a tuple with the TotalResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListTargetPromotionGroupTargetsResponseContent) GetTotalResultsOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalResults) {
		return nil, false
	}
	return o.TotalResults, true
}

// HasTotalResults returns a boolean if a field has been set.
func (o *SponsoredProductsListTargetPromotionGroupTargetsResponseContent) HasTotalResults() bool {
	if o != nil && !IsNil(o.TotalResults) {
		return true
	}

	return false
}

// SetTotalResults gets a reference to the given int64 and assigns it to the TotalResults field.
func (o *SponsoredProductsListTargetPromotionGroupTargetsResponseContent) SetTotalResults(v int64) {
	o.TotalResults = &v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *SponsoredProductsListTargetPromotionGroupTargetsResponseContent) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListTargetPromotionGroupTargetsResponseContent) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *SponsoredProductsListTargetPromotionGroupTargetsResponseContent) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *SponsoredProductsListTargetPromotionGroupTargetsResponseContent) SetNextToken(v string) {
	o.NextToken = &v
}

// GetTargets returns the Targets field value if set, zero value otherwise.
func (o *SponsoredProductsListTargetPromotionGroupTargetsResponseContent) GetTargets() []SponsoredProductsTarget {
	if o == nil || IsNil(o.Targets) {
		var ret []SponsoredProductsTarget
		return ret
	}
	return o.Targets
}

// GetTargetsOk returns a tuple with the Targets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListTargetPromotionGroupTargetsResponseContent) GetTargetsOk() ([]SponsoredProductsTarget, bool) {
	if o == nil || IsNil(o.Targets) {
		return nil, false
	}
	return o.Targets, true
}

// HasTargets returns a boolean if a field has been set.
func (o *SponsoredProductsListTargetPromotionGroupTargetsResponseContent) HasTargets() bool {
	if o != nil && !IsNil(o.Targets) {
		return true
	}

	return false
}

// SetTargets gets a reference to the given []SponsoredProductsTarget and assigns it to the Targets field.
func (o *SponsoredProductsListTargetPromotionGroupTargetsResponseContent) SetTargets(v []SponsoredProductsTarget) {
	o.Targets = v
}

func (o SponsoredProductsListTargetPromotionGroupTargetsResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TotalResults) {
		toSerialize["totalResults"] = o.TotalResults
	}
	if !IsNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	if !IsNil(o.Targets) {
		toSerialize["targets"] = o.Targets
	}
	return toSerialize, nil
}

type NullableSponsoredProductsListTargetPromotionGroupTargetsResponseContent struct {
	value *SponsoredProductsListTargetPromotionGroupTargetsResponseContent
	isSet bool
}

func (v NullableSponsoredProductsListTargetPromotionGroupTargetsResponseContent) Get() *SponsoredProductsListTargetPromotionGroupTargetsResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsListTargetPromotionGroupTargetsResponseContent) Set(val *SponsoredProductsListTargetPromotionGroupTargetsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsListTargetPromotionGroupTargetsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsListTargetPromotionGroupTargetsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsListTargetPromotionGroupTargetsResponseContent(val *SponsoredProductsListTargetPromotionGroupTargetsResponseContent) *NullableSponsoredProductsListTargetPromotionGroupTargetsResponseContent {
	return &NullableSponsoredProductsListTargetPromotionGroupTargetsResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsListTargetPromotionGroupTargetsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsListTargetPromotionGroupTargetsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
