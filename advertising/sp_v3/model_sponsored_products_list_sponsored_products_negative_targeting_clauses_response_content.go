package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsListSponsoredProductsNegativeTargetingClausesResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsListSponsoredProductsNegativeTargetingClausesResponseContent{}

// SponsoredProductsListSponsoredProductsNegativeTargetingClausesResponseContent struct for SponsoredProductsListSponsoredProductsNegativeTargetingClausesResponseContent
type SponsoredProductsListSponsoredProductsNegativeTargetingClausesResponseContent struct {
	// The total number of entities
	TotalResults *int64 `json:"totalResults,omitempty"`
	// token value allowing to navigate to the next response page
	NextToken                *string                                    `json:"nextToken,omitempty"`
	NegativeTargetingClauses []SponsoredProductsNegativeTargetingClause `json:"negativeTargetingClauses,omitempty"`
}

// NewSponsoredProductsListSponsoredProductsNegativeTargetingClausesResponseContent instantiates a new SponsoredProductsListSponsoredProductsNegativeTargetingClausesResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsListSponsoredProductsNegativeTargetingClausesResponseContent() *SponsoredProductsListSponsoredProductsNegativeTargetingClausesResponseContent {
	this := SponsoredProductsListSponsoredProductsNegativeTargetingClausesResponseContent{}
	return &this
}

// NewSponsoredProductsListSponsoredProductsNegativeTargetingClausesResponseContentWithDefaults instantiates a new SponsoredProductsListSponsoredProductsNegativeTargetingClausesResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsListSponsoredProductsNegativeTargetingClausesResponseContentWithDefaults() *SponsoredProductsListSponsoredProductsNegativeTargetingClausesResponseContent {
	this := SponsoredProductsListSponsoredProductsNegativeTargetingClausesResponseContent{}
	return &this
}

// GetTotalResults returns the TotalResults field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsNegativeTargetingClausesResponseContent) GetTotalResults() int64 {
	if o == nil || IsNil(o.TotalResults) {
		var ret int64
		return ret
	}
	return *o.TotalResults
}

// GetTotalResultsOk returns a tuple with the TotalResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsNegativeTargetingClausesResponseContent) GetTotalResultsOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalResults) {
		return nil, false
	}
	return o.TotalResults, true
}

// HasTotalResults returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsNegativeTargetingClausesResponseContent) HasTotalResults() bool {
	if o != nil && !IsNil(o.TotalResults) {
		return true
	}

	return false
}

// SetTotalResults gets a reference to the given int64 and assigns it to the TotalResults field.
func (o *SponsoredProductsListSponsoredProductsNegativeTargetingClausesResponseContent) SetTotalResults(v int64) {
	o.TotalResults = &v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsNegativeTargetingClausesResponseContent) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsNegativeTargetingClausesResponseContent) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsNegativeTargetingClausesResponseContent) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *SponsoredProductsListSponsoredProductsNegativeTargetingClausesResponseContent) SetNextToken(v string) {
	o.NextToken = &v
}

// GetNegativeTargetingClauses returns the NegativeTargetingClauses field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsNegativeTargetingClausesResponseContent) GetNegativeTargetingClauses() []SponsoredProductsNegativeTargetingClause {
	if o == nil || IsNil(o.NegativeTargetingClauses) {
		var ret []SponsoredProductsNegativeTargetingClause
		return ret
	}
	return o.NegativeTargetingClauses
}

// GetNegativeTargetingClausesOk returns a tuple with the NegativeTargetingClauses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsNegativeTargetingClausesResponseContent) GetNegativeTargetingClausesOk() ([]SponsoredProductsNegativeTargetingClause, bool) {
	if o == nil || IsNil(o.NegativeTargetingClauses) {
		return nil, false
	}
	return o.NegativeTargetingClauses, true
}

// HasNegativeTargetingClauses returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsNegativeTargetingClausesResponseContent) HasNegativeTargetingClauses() bool {
	if o != nil && !IsNil(o.NegativeTargetingClauses) {
		return true
	}

	return false
}

// SetNegativeTargetingClauses gets a reference to the given []SponsoredProductsNegativeTargetingClause and assigns it to the NegativeTargetingClauses field.
func (o *SponsoredProductsListSponsoredProductsNegativeTargetingClausesResponseContent) SetNegativeTargetingClauses(v []SponsoredProductsNegativeTargetingClause) {
	o.NegativeTargetingClauses = v
}

func (o SponsoredProductsListSponsoredProductsNegativeTargetingClausesResponseContent) ToMap() (map[string]interface{}, error) {
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

type NullableSponsoredProductsListSponsoredProductsNegativeTargetingClausesResponseContent struct {
	value *SponsoredProductsListSponsoredProductsNegativeTargetingClausesResponseContent
	isSet bool
}

func (v NullableSponsoredProductsListSponsoredProductsNegativeTargetingClausesResponseContent) Get() *SponsoredProductsListSponsoredProductsNegativeTargetingClausesResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsListSponsoredProductsNegativeTargetingClausesResponseContent) Set(val *SponsoredProductsListSponsoredProductsNegativeTargetingClausesResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsListSponsoredProductsNegativeTargetingClausesResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsListSponsoredProductsNegativeTargetingClausesResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsListSponsoredProductsNegativeTargetingClausesResponseContent(val *SponsoredProductsListSponsoredProductsNegativeTargetingClausesResponseContent) *NullableSponsoredProductsListSponsoredProductsNegativeTargetingClausesResponseContent {
	return &NullableSponsoredProductsListSponsoredProductsNegativeTargetingClausesResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsListSponsoredProductsNegativeTargetingClausesResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsListSponsoredProductsNegativeTargetingClausesResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
