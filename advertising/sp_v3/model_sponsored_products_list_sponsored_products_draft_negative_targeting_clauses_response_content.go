package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsListSponsoredProductsDraftNegativeTargetingClausesResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsListSponsoredProductsDraftNegativeTargetingClausesResponseContent{}

// SponsoredProductsListSponsoredProductsDraftNegativeTargetingClausesResponseContent struct for SponsoredProductsListSponsoredProductsDraftNegativeTargetingClausesResponseContent
type SponsoredProductsListSponsoredProductsDraftNegativeTargetingClausesResponseContent struct {
	// The total number of entities
	TotalResults *int64 `json:"totalResults,omitempty"`
	// token value allowing to navigate to the next response page
	NextToken                *string                                         `json:"nextToken,omitempty"`
	NegativeTargetingClauses []SponsoredProductsDraftNegativeTargetingClause `json:"negativeTargetingClauses,omitempty"`
}

// NewSponsoredProductsListSponsoredProductsDraftNegativeTargetingClausesResponseContent instantiates a new SponsoredProductsListSponsoredProductsDraftNegativeTargetingClausesResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsListSponsoredProductsDraftNegativeTargetingClausesResponseContent() *SponsoredProductsListSponsoredProductsDraftNegativeTargetingClausesResponseContent {
	this := SponsoredProductsListSponsoredProductsDraftNegativeTargetingClausesResponseContent{}
	return &this
}

// NewSponsoredProductsListSponsoredProductsDraftNegativeTargetingClausesResponseContentWithDefaults instantiates a new SponsoredProductsListSponsoredProductsDraftNegativeTargetingClausesResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsListSponsoredProductsDraftNegativeTargetingClausesResponseContentWithDefaults() *SponsoredProductsListSponsoredProductsDraftNegativeTargetingClausesResponseContent {
	this := SponsoredProductsListSponsoredProductsDraftNegativeTargetingClausesResponseContent{}
	return &this
}

// GetTotalResults returns the TotalResults field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsDraftNegativeTargetingClausesResponseContent) GetTotalResults() int64 {
	if o == nil || IsNil(o.TotalResults) {
		var ret int64
		return ret
	}
	return *o.TotalResults
}

// GetTotalResultsOk returns a tuple with the TotalResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsDraftNegativeTargetingClausesResponseContent) GetTotalResultsOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalResults) {
		return nil, false
	}
	return o.TotalResults, true
}

// HasTotalResults returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsDraftNegativeTargetingClausesResponseContent) HasTotalResults() bool {
	if o != nil && !IsNil(o.TotalResults) {
		return true
	}

	return false
}

// SetTotalResults gets a reference to the given int64 and assigns it to the TotalResults field.
func (o *SponsoredProductsListSponsoredProductsDraftNegativeTargetingClausesResponseContent) SetTotalResults(v int64) {
	o.TotalResults = &v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsDraftNegativeTargetingClausesResponseContent) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsDraftNegativeTargetingClausesResponseContent) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsDraftNegativeTargetingClausesResponseContent) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *SponsoredProductsListSponsoredProductsDraftNegativeTargetingClausesResponseContent) SetNextToken(v string) {
	o.NextToken = &v
}

// GetNegativeTargetingClauses returns the NegativeTargetingClauses field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsDraftNegativeTargetingClausesResponseContent) GetNegativeTargetingClauses() []SponsoredProductsDraftNegativeTargetingClause {
	if o == nil || IsNil(o.NegativeTargetingClauses) {
		var ret []SponsoredProductsDraftNegativeTargetingClause
		return ret
	}
	return o.NegativeTargetingClauses
}

// GetNegativeTargetingClausesOk returns a tuple with the NegativeTargetingClauses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsDraftNegativeTargetingClausesResponseContent) GetNegativeTargetingClausesOk() ([]SponsoredProductsDraftNegativeTargetingClause, bool) {
	if o == nil || IsNil(o.NegativeTargetingClauses) {
		return nil, false
	}
	return o.NegativeTargetingClauses, true
}

// HasNegativeTargetingClauses returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsDraftNegativeTargetingClausesResponseContent) HasNegativeTargetingClauses() bool {
	if o != nil && !IsNil(o.NegativeTargetingClauses) {
		return true
	}

	return false
}

// SetNegativeTargetingClauses gets a reference to the given []SponsoredProductsDraftNegativeTargetingClause and assigns it to the NegativeTargetingClauses field.
func (o *SponsoredProductsListSponsoredProductsDraftNegativeTargetingClausesResponseContent) SetNegativeTargetingClauses(v []SponsoredProductsDraftNegativeTargetingClause) {
	o.NegativeTargetingClauses = v
}

func (o SponsoredProductsListSponsoredProductsDraftNegativeTargetingClausesResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TotalResults) {
		toSerialize["totalResults"] = o.TotalResults
	}
	if !IsNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	if !IsNil(o.NegativeTargetingClauses) {
		toSerialize["negativeTargetingClauses"] = o.NegativeTargetingClauses
	}
	return toSerialize, nil
}

type NullableSponsoredProductsListSponsoredProductsDraftNegativeTargetingClausesResponseContent struct {
	value *SponsoredProductsListSponsoredProductsDraftNegativeTargetingClausesResponseContent
	isSet bool
}

func (v NullableSponsoredProductsListSponsoredProductsDraftNegativeTargetingClausesResponseContent) Get() *SponsoredProductsListSponsoredProductsDraftNegativeTargetingClausesResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsListSponsoredProductsDraftNegativeTargetingClausesResponseContent) Set(val *SponsoredProductsListSponsoredProductsDraftNegativeTargetingClausesResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsListSponsoredProductsDraftNegativeTargetingClausesResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsListSponsoredProductsDraftNegativeTargetingClausesResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsListSponsoredProductsDraftNegativeTargetingClausesResponseContent(val *SponsoredProductsListSponsoredProductsDraftNegativeTargetingClausesResponseContent) *NullableSponsoredProductsListSponsoredProductsDraftNegativeTargetingClausesResponseContent {
	return &NullableSponsoredProductsListSponsoredProductsDraftNegativeTargetingClausesResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsListSponsoredProductsDraftNegativeTargetingClausesResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsListSponsoredProductsDraftNegativeTargetingClausesResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
