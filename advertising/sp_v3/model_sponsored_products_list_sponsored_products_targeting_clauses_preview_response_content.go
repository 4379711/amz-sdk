package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsListSponsoredProductsTargetingClausesPreviewResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsListSponsoredProductsTargetingClausesPreviewResponseContent{}

// SponsoredProductsListSponsoredProductsTargetingClausesPreviewResponseContent struct for SponsoredProductsListSponsoredProductsTargetingClausesPreviewResponseContent
type SponsoredProductsListSponsoredProductsTargetingClausesPreviewResponseContent struct {
	// The total number of entities
	TotalResults *int64 `json:"totalResults,omitempty"`
	// token value allowing to navigate to the next response page
	NextToken        *string                            `json:"nextToken,omitempty"`
	TargetingClauses []SponsoredProductsTargetingClause `json:"targetingClauses,omitempty"`
}

// NewSponsoredProductsListSponsoredProductsTargetingClausesPreviewResponseContent instantiates a new SponsoredProductsListSponsoredProductsTargetingClausesPreviewResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsListSponsoredProductsTargetingClausesPreviewResponseContent() *SponsoredProductsListSponsoredProductsTargetingClausesPreviewResponseContent {
	this := SponsoredProductsListSponsoredProductsTargetingClausesPreviewResponseContent{}
	return &this
}

// NewSponsoredProductsListSponsoredProductsTargetingClausesPreviewResponseContentWithDefaults instantiates a new SponsoredProductsListSponsoredProductsTargetingClausesPreviewResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsListSponsoredProductsTargetingClausesPreviewResponseContentWithDefaults() *SponsoredProductsListSponsoredProductsTargetingClausesPreviewResponseContent {
	this := SponsoredProductsListSponsoredProductsTargetingClausesPreviewResponseContent{}
	return &this
}

// GetTotalResults returns the TotalResults field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsTargetingClausesPreviewResponseContent) GetTotalResults() int64 {
	if o == nil || IsNil(o.TotalResults) {
		var ret int64
		return ret
	}
	return *o.TotalResults
}

// GetTotalResultsOk returns a tuple with the TotalResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsTargetingClausesPreviewResponseContent) GetTotalResultsOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalResults) {
		return nil, false
	}
	return o.TotalResults, true
}

// HasTotalResults returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsTargetingClausesPreviewResponseContent) HasTotalResults() bool {
	if o != nil && !IsNil(o.TotalResults) {
		return true
	}

	return false
}

// SetTotalResults gets a reference to the given int64 and assigns it to the TotalResults field.
func (o *SponsoredProductsListSponsoredProductsTargetingClausesPreviewResponseContent) SetTotalResults(v int64) {
	o.TotalResults = &v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsTargetingClausesPreviewResponseContent) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsTargetingClausesPreviewResponseContent) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsTargetingClausesPreviewResponseContent) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *SponsoredProductsListSponsoredProductsTargetingClausesPreviewResponseContent) SetNextToken(v string) {
	o.NextToken = &v
}

// GetTargetingClauses returns the TargetingClauses field value if set, zero value otherwise.
func (o *SponsoredProductsListSponsoredProductsTargetingClausesPreviewResponseContent) GetTargetingClauses() []SponsoredProductsTargetingClause {
	if o == nil || IsNil(o.TargetingClauses) {
		var ret []SponsoredProductsTargetingClause
		return ret
	}
	return o.TargetingClauses
}

// GetTargetingClausesOk returns a tuple with the TargetingClauses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsListSponsoredProductsTargetingClausesPreviewResponseContent) GetTargetingClausesOk() ([]SponsoredProductsTargetingClause, bool) {
	if o == nil || IsNil(o.TargetingClauses) {
		return nil, false
	}
	return o.TargetingClauses, true
}

// HasTargetingClauses returns a boolean if a field has been set.
func (o *SponsoredProductsListSponsoredProductsTargetingClausesPreviewResponseContent) HasTargetingClauses() bool {
	if o != nil && !IsNil(o.TargetingClauses) {
		return true
	}

	return false
}

// SetTargetingClauses gets a reference to the given []SponsoredProductsTargetingClause and assigns it to the TargetingClauses field.
func (o *SponsoredProductsListSponsoredProductsTargetingClausesPreviewResponseContent) SetTargetingClauses(v []SponsoredProductsTargetingClause) {
	o.TargetingClauses = v
}

func (o SponsoredProductsListSponsoredProductsTargetingClausesPreviewResponseContent) ToMap() (map[string]interface{}, error) {
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

type NullableSponsoredProductsListSponsoredProductsTargetingClausesPreviewResponseContent struct {
	value *SponsoredProductsListSponsoredProductsTargetingClausesPreviewResponseContent
	isSet bool
}

func (v NullableSponsoredProductsListSponsoredProductsTargetingClausesPreviewResponseContent) Get() *SponsoredProductsListSponsoredProductsTargetingClausesPreviewResponseContent {
	return v.value
}

func (v *NullableSponsoredProductsListSponsoredProductsTargetingClausesPreviewResponseContent) Set(val *SponsoredProductsListSponsoredProductsTargetingClausesPreviewResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsListSponsoredProductsTargetingClausesPreviewResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsListSponsoredProductsTargetingClausesPreviewResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsListSponsoredProductsTargetingClausesPreviewResponseContent(val *SponsoredProductsListSponsoredProductsTargetingClausesPreviewResponseContent) *NullableSponsoredProductsListSponsoredProductsTargetingClausesPreviewResponseContent {
	return &NullableSponsoredProductsListSponsoredProductsTargetingClausesPreviewResponseContent{value: val, isSet: true}
}

func (v NullableSponsoredProductsListSponsoredProductsTargetingClausesPreviewResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsListSponsoredProductsTargetingClausesPreviewResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
