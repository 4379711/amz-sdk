package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsListSponsoredProductsDraftTargetingClausesResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsListSponsoredProductsDraftTargetingClausesResponseContent{}

// SponsoredProductsListSponsoredProductsDraftTargetingClausesResponseContent struct for SponsoredProductsListSponsoredProductsDraftTargetingClausesResponseContent
type SponsoredProductsListSponsoredProductsDraftTargetingClausesResponseContent struct {
	// The total number of entities
	TotalResults *int64 `json:"totalResults,omitempty"`
	// token value allowing to navigate to the next response page
	NextToken        *string                                 `json:"nextToken,omitempty"`
	TargetingClauses []SponsoredProductsDraftTargetingClause `json:"targetingClauses,omitempty"`
}

// NewSponsoredProductsListSponsoredProductsDraftTargetingClausesResponseContent instantiates a new SponsoredProductsListSponsoredProductsDraftTargetingClausesResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsListSponsoredProductsDraftTargetingClausesResponseContent() *SponsoredProductsListSponsoredProductsDraftTargetingClausesResponseContent {
	this := SponsoredProductsListSponsoredProductsDraftTargetingClausesResponseContent{}
	return &this
}

// NewSponsoredProductsListSponsoredProductsDraftTargetingClausesResponseContentWithDefaults instantiates a new SponsoredProductsListSponsoredProductsDraftTargetingClausesResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsListSponsoredProductsDraftTargetingClausesResponseContentWithDefaults() *SponsoredProductsListSponsoredProductsDraftTargetingClausesResponseContent {
	this := SponsoredProductsListSponsoredProductsDraftTargetingClausesResponseContent{}
	return &this
}

// GetTotalResults returns the TotalResults field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsDraftTargetingClausesResponseContent) GetTotalResults() int64 {
	if o == nil || IsNil(o.TotalResults) {
		var ret int64
		return ret
	}
	return *o.TotalResults
}

// GetTotalResultsOk returns a tuple with the TotalResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsDraftTargetingClausesResponseContent) GetTotalResultsOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalResults) {
		return nil, false
	}
	return o.TotalResults, true
}

// HasTotalResults returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsDraftTargetingClausesResponseContent) HasTotalResults() bool {
	if o != nil && !IsNil(o.TotalResults) {
		return true
	}

	return false
}

// SetTotalResults gets a reference to the given int64 and assigns it to the TotalResults field.
func (o *SponsoredProductsListSponsoredProductsDraftTargetingClausesResponseContent) SetTotalResults(v int64) {
	o.TotalResults = &v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsDraftTargetingClausesResponseContent) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsDraftTargetingClausesResponseContent) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsDraftTargetingClausesResponseContent) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *SponsoredProductsListSponsoredProductsDraftTargetingClausesResponseContent) SetNextToken(v string) {
	o.NextToken = &v
}

// GetTargetingClauses returns the TargetingClauses field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsDraftTargetingClausesResponseContent) GetTargetingClauses() []SponsoredProductsDraftTargetingClause {
	if o == nil || IsNil(o.TargetingClauses) {
		var ret []SponsoredProductsDraftTargetingClause
		return ret
	}
	return o.TargetingClauses
}

// GetTargetingClausesOk returns a tuple with the TargetingClauses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsDraftTargetingClausesResponseContent) GetTargetingClausesOk() ([]SponsoredProductsDraftTargetingClause, bool) {
	if o == nil || IsNil(o.TargetingClauses) {
		return nil, false
	}
	return o.TargetingClauses, true
}

// HasTargetingClauses returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsDraftTargetingClausesResponseContent) HasTargetingClauses() bool {
	if o != nil && !IsNil(o.TargetingClauses) {
		return true
	}

	return false
}

// SetTargetingClauses gets a reference to the given []SponsoredProductsDraftTargetingClause and assigns it to the TargetingClauses field.
func (o *SponsoredProductsListSponsoredProductsDraftTargetingClausesResponseContent) SetTargetingClauses(v []SponsoredProductsDraftTargetingClause) {
	o.TargetingClauses = v
}

func (o SponsoredProductsListSponsoredProductsDraftTargetingClausesResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TotalResults) {
		toSerialize["totalResults"] = o.TotalResults
	}
	if !IsNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	if !IsNil(o.TargetingClauses) {
		toSerialize["targetingClauses"] = o.TargetingClauses
	}
	return toSerialize, nil
}

type NullableSponsoredProductsListSponsoredProductsDraftTargetingClausesResponseContent struct {
	value *SponsoredProductsListSponsoredProductsDraftTargetingClausesResponseContent
	isSet bool
}

func (v NullableSponsoredProductsListSponsoredProductsDraftTargetingClausesResponseContent) Get() *SponsoredProductsListSponsoredProductsDraftTargetingClausesResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsListSponsoredProductsDraftTargetingClausesResponseContent) Set(val *SponsoredProductsListSponsoredProductsDraftTargetingClausesResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsListSponsoredProductsDraftTargetingClausesResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsListSponsoredProductsDraftTargetingClausesResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsListSponsoredProductsDraftTargetingClausesResponseContent(val *SponsoredProductsListSponsoredProductsDraftTargetingClausesResponseContent) *NullableSponsoredProductsListSponsoredProductsDraftTargetingClausesResponseContent {
	return &NullableSponsoredProductsListSponsoredProductsDraftTargetingClausesResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsListSponsoredProductsDraftTargetingClausesResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsListSponsoredProductsDraftTargetingClausesResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
