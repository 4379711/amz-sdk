package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsListSponsoredProductsGlobalTargetingClausesResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsListSponsoredProductsGlobalTargetingClausesResponseContent{}

// SponsoredProductsListSponsoredProductsGlobalTargetingClausesResponseContent struct for SponsoredProductsListSponsoredProductsGlobalTargetingClausesResponseContent
type SponsoredProductsListSponsoredProductsGlobalTargetingClausesResponseContent struct {
	// The total number of entities
	TotalResults *int64 `json:"totalResults,omitempty"`
	// token value allowing to navigate to the next response page
	NextToken        *string                                  `json:"nextToken,omitempty"`
	TargetingClauses []SponsoredProductsGlobalTargetingClause `json:"targetingClauses,omitempty"`
}

// NewSponsoredProductsListSponsoredProductsGlobalTargetingClausesResponseContent instantiates a new SponsoredProductsListSponsoredProductsGlobalTargetingClausesResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsListSponsoredProductsGlobalTargetingClausesResponseContent() *SponsoredProductsListSponsoredProductsGlobalTargetingClausesResponseContent {
	this := SponsoredProductsListSponsoredProductsGlobalTargetingClausesResponseContent{}
	return &this
}

// NewSponsoredProductsListSponsoredProductsGlobalTargetingClausesResponseContentWithDefaults instantiates a new SponsoredProductsListSponsoredProductsGlobalTargetingClausesResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsListSponsoredProductsGlobalTargetingClausesResponseContentWithDefaults() *SponsoredProductsListSponsoredProductsGlobalTargetingClausesResponseContent {
	this := SponsoredProductsListSponsoredProductsGlobalTargetingClausesResponseContent{}
	return &this
}

// GetTotalResults returns the TotalResults field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsGlobalTargetingClausesResponseContent) GetTotalResults() int64 {
	if o == nil || IsNil(o.TotalResults) {
		var ret int64
		return ret
	}
	return *o.TotalResults
}

// GetTotalResultsOk returns a tuple with the TotalResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalTargetingClausesResponseContent) GetTotalResultsOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalResults) {
		return nil, false
	}
	return o.TotalResults, true
}

// HasTotalResults returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalTargetingClausesResponseContent) HasTotalResults() bool {
	if o != nil && !IsNil(o.TotalResults) {
		return true
	}

	return false
}

// SetTotalResults gets a reference to the given int64 and assigns it to the TotalResults field.
func (o *SponsoredProductsListSponsoredProductsGlobalTargetingClausesResponseContent) SetTotalResults(v int64) {
	o.TotalResults = &v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsGlobalTargetingClausesResponseContent) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalTargetingClausesResponseContent) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalTargetingClausesResponseContent) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *SponsoredProductsListSponsoredProductsGlobalTargetingClausesResponseContent) SetNextToken(v string) {
	o.NextToken = &v
}

// GetTargetingClauses returns the TargetingClauses field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsGlobalTargetingClausesResponseContent) GetTargetingClauses() []SponsoredProductsGlobalTargetingClause {
	if o == nil || IsNil(o.TargetingClauses) {
		var ret []SponsoredProductsGlobalTargetingClause
		return ret
	}
	return o.TargetingClauses
}

// GetTargetingClausesOk returns a tuple with the TargetingClauses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalTargetingClausesResponseContent) GetTargetingClausesOk() ([]SponsoredProductsGlobalTargetingClause, bool) {
	if o == nil || IsNil(o.TargetingClauses) {
		return nil, false
	}
	return o.TargetingClauses, true
}

// HasTargetingClauses returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsGlobalTargetingClausesResponseContent) HasTargetingClauses() bool {
	if o != nil && !IsNil(o.TargetingClauses) {
		return true
	}

	return false
}

// SetTargetingClauses gets a reference to the given []SponsoredProductsGlobalTargetingClause and assigns it to the TargetingClauses field.
func (o *SponsoredProductsListSponsoredProductsGlobalTargetingClausesResponseContent) SetTargetingClauses(v []SponsoredProductsGlobalTargetingClause) {
	o.TargetingClauses = v
}

func (o SponsoredProductsListSponsoredProductsGlobalTargetingClausesResponseContent) ToMap() (map[string]interface{}, error) {
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

type NullableSponsoredProductsListSponsoredProductsGlobalTargetingClausesResponseContent struct {
	value *SponsoredProductsListSponsoredProductsGlobalTargetingClausesResponseContent
	isSet bool
}

func (v NullableSponsoredProductsListSponsoredProductsGlobalTargetingClausesResponseContent) Get() *SponsoredProductsListSponsoredProductsGlobalTargetingClausesResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsListSponsoredProductsGlobalTargetingClausesResponseContent) Set(val *SponsoredProductsListSponsoredProductsGlobalTargetingClausesResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsListSponsoredProductsGlobalTargetingClausesResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsListSponsoredProductsGlobalTargetingClausesResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsListSponsoredProductsGlobalTargetingClausesResponseContent(val *SponsoredProductsListSponsoredProductsGlobalTargetingClausesResponseContent) *NullableSponsoredProductsListSponsoredProductsGlobalTargetingClausesResponseContent {
	return &NullableSponsoredProductsListSponsoredProductsGlobalTargetingClausesResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsListSponsoredProductsGlobalTargetingClausesResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsListSponsoredProductsGlobalTargetingClausesResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
